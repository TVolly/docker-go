package routes

import (
	"net/http"

	"github.com/TVolly/goapi-addresses/internal/controllers"
	"github.com/TVolly/goapi-addresses/internal/repositories"
)

func (r *routesRegistry) ConfigureCommunityRoutes(repo repositories.CommunityRepository) {
	s := r.router.PathPrefix("/communities").Subrouter()
	c := controllers.NewCommunityController(repo)

	s.HandleFunc("", c.Index()).Methods(http.MethodGet)
	s.HandleFunc("", c.Create()).Methods(http.MethodPost)
}
