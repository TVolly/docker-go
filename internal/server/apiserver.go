package server

import (
	"net/http"

	"github.com/TVolly/goapi-addresses/internal/middlewares"

	"github.com/TVolly/goapi-addresses/internal/handlers"

	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  repositories.RepositoryStore
}

func New(config *Config, store repositories.RepositoryStore) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store,
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.store.Init(); err != nil {
		s.logger.Error(err.Error())

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
	s.router.Use(middlewares.LogRequests(s.logger))

	h := handlers.NewHandler(s.router, s.store)
	h.BindCommunityHandlers()
}
