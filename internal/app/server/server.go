package server

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

func (s *Server) configureRouter() {
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/index", s.handleIndex())
}

func (s *Server) handleIndex() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello on my index page!!!")
	}
}
