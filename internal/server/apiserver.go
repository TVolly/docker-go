package server

import (
	"net/http"

	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/TVolly/goapi-addresses/internal/routes"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config       *Config
	logger       *logrus.Logger
	router       *mux.Router
	repoRegistry repositories.RepositoryRegistry
}

func New(config *Config, repoRegistry repositories.RepositoryRegistry) *APIServer {
	return &APIServer{
		config:       config,
		logger:       logrus.New(),
		router:       mux.NewRouter(),
		repoRegistry: repoRegistry,
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("Configure router")
	s.configureRouter()

	s.logger.Info("Run server")
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
	r := routes.NewRouteRegistry(s.router)

	r.ConfigureCommunityRoutes(s.repoRegistry.Community())
}
