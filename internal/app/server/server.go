package server

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/text"
	"github.com/sirupsen/logrus"
)

type Server struct {
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
	s.logger.Info("starting server")

	s.config = config

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

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
}
