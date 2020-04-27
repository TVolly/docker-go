package routes

import (
	"github.com/TVolly/goapi-addresses/internal/controllers"
	"github.com/TVolly/goapi-addresses/internal/repositories"
)

func (r *routesRegistry) ConfigureCommunityRoutes(repo repositories.CommunityRepository) {
	s := r.router.PathPrefix("/communities").Subrouter()
	c := controllers.NewCommunityController(repo)

	s.HandleFunc("", c.Index())
}
