package handlers

import (
	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/gorilla/mux"
)

type handlersStore struct {
	router *mux.Router
	store  repositories.RepositoryStore
}

func NewHandler(r *mux.Router, s repositories.RepositoryStore) *handlersStore {
	return &handlersStore{
		router: r,
		store:  s,
	}
}
