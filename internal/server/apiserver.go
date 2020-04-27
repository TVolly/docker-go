package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("Configure router")
	s.configureRouter()

	s.logger.Info("Server ready to start")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	log_lvl, err := logrus.ParseLevel(s.config.LogLvl)
	if err != nil {
		return err
	}

	s.logger.SetLevel(log_lvl)

	return nil
}

func (s *APIServer) configureRouter() {

	s.router.PathPrefix("/private").Subrouter()
}
