package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/config"
	"github.com/ilfey/go-back/internal/app/endpoints/img"
	"github.com/ilfey/go-back/internal/app/endpoints/jwt"
	"github.com/ilfey/go-back/internal/app/endpoints/ping"
	"github.com/ilfey/go-back/internal/app/endpoints/text"
	"github.com/ilfey/go-back/internal/pkg/store"
	"github.com/sirupsen/logrus"
)

type Server struct {
	store  *store.Store
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
}

func New() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

func (s *Server) Start(config *config.Config, store *store.Store, logger *logrus.Logger) error {

	s.config = config
	s.store = store
	s.logger = logger

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

func (s *Server) configureRouter() {
	s.router.Use(s.loggingMiddleWare)
	s.router.Use(s.contentJsonMiddleware)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	textHandler := text.New()
	textHandler.Register(s.router)

	imgHandler := img.New(s.logger)
	imgHandler.Register(s.router)

	pingHandler := ping.New(s.config.Key)
	pingHandler.Register(s.router)

	if s.store != nil {
		jwtHandler := jwt.New(s.store, s.config.Key, s.config.LifeSpan)
		jwtHandler.Register(s.router)

		privateRouter := s.router.PathPrefix("/private/").Subrouter()
		privateRouter.Use(s.bearerMiddleware)

		privateTextHandler := text.New()
		privateTextHandler.Register(privateRouter)

		privateImgHandler := img.New(s.logger)
		privateImgHandler.Register(privateRouter)
	} else {
		s.logger.Warnf("the server is not connected to the database. jwt and private endpoints is not available")
	}
}
