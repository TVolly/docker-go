package routes

import (
	"github.com/TVolly/goapi-addresses/internal/controllers"
	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/gorilla/mux"
)

func ConfigureCommunityRoutes(r *mux.Router, repoRegistry repositories.RepositoryRegistry) {
	s := r.PathPrefix("/communities").Subrouter()
	c := controllers.NewCommunityController(repoRegistry.Community())

	s.HandleFunc("", c.Index())
}
