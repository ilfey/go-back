package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/img"
	"github.com/ilfey/go-back/internal/app/jwt"
	"github.com/ilfey/go-back/internal/app/store/sqlstore"
	"github.com/ilfey/go-back/internal/app/text"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Server struct {
	db     *pgx.Conn
	store  *sqlstore.Store
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New() *Server {
	return &Server{
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start(config *Config) error {

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

func (s *Server) loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}

func (s *Server) configureRouter() {
	s.router.Use(s.loggingMiddleWare)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	textHandler := text.New()
	textHandler.Register(s.router)
	imgHandler := img.New()
	imgHandler.Register(s.router)

	if s.store != nil {
		jwtHandler := jwt.New(s.store, s.config.Key)
		jwtHandler.Register(s.router)
	} else {
		s.logger.Infof("the server is not connected to the database. routes /jwt/** is not available")
	}
}
