package server

import (
	"context"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/config"
	"github.com/ilfey/go-back/internal/app/endpoints/img"
	"github.com/ilfey/go-back/internal/app/endpoints/jwt"
	"github.com/ilfey/go-back/internal/app/endpoints/text"
	"github.com/ilfey/go-back/internal/app/store/sqlstore"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Server struct {
	db     *pgx.Conn
	store  *sqlstore.Store
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
}

func New() *Server {
	return &Server{
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start(config *config.Config) error {

	s.config = config

	if err := s.configureLogger(); err != nil {
		return err
	}

	db, err := pgx.Connect(context.Background(), s.config.DatabaseUrl)
	if err != nil {
		logrus.Error(err)
	} else {
		s.db = db
		s.store = sqlstore.New(db, s.logger)
		logrus.Info("server connected to db")
	}

	s.configureRouter()

	s.logger.Infof("starting server on http://%s/", config.Address)

	// log routes
	if err := s.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		s.logger.Infof("http://%s%s", config.Address, t)
		return nil
	}); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.Address, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configureRouter() {
	s.router.Use(s.loggingMiddleWare)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	jsonRouter := s.router.PathPrefix("/").Subrouter()
	jsonRouter.Use(s.contentJsonMiddleware)

	textHandler := text.New()
	textHandler.Register(jsonRouter)

	imgHandler := img.New()
	imgHandler.Register(s.router)

	if s.store != nil {
		jwtHandler := jwt.New(s.store, s.config.Key, s.config.LifeSpan)
		jwtHandler.Register(s.router)
	} else {
		s.logger.Infof("the server is not connected to the database. routes /jwt/** is not available")
	}

	privateRouter := s.router.PathPrefix("/private/").Subrouter()
	privateRouter.Use(s.bearerMiddleware)

	privateJsonRouter := privateRouter.PathPrefix("/").Subrouter()
	privateJsonRouter.Use(s.contentJsonMiddleware)

	privateTextHandler := text.New()
	privateTextHandler.Register(privateJsonRouter)
	
	privateImgHandler := img.New()
	privateImgHandler.Register(privateRouter)
}
