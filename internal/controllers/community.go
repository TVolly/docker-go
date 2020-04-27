package controllers

import (
	"net/http"

	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/TVolly/goapi-addresses/internal/responses"
)

type communityController struct {
	repo repositories.CommunityRepository
}

func NewCommunityController(repo repositories.CommunityRepository) *communityController {
	return &communityController{
		repo: repo,
	}
}

func (c *communityController) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.Respond(w, r, http.StatusOK, nil)
	}
}
