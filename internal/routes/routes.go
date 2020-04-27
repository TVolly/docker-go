package routes

import (
	"github.com/gorilla/mux"
)

type routesRegistry struct {
	router *mux.Router
}

func NewRouteRegistry(r *mux.Router) *routesRegistry {
	return &routesRegistry{
		router: r,
	}
}
