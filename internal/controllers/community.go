package controllers

import (
	"net/http"

	"github.com/TVolly/goapi-addresses/internal/models"
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
		items := c.repo.List()
		p := responses.NewPagination(items)

		responses.Respond(w, r, http.StatusOK, p)
	}
}

func (c *communityController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		item := models.TestCommunity()

		if err := c.repo.Create(item); err != nil {
			responses.Error(w, r, http.StatusOK, err)
			return
		}
		body := responses.SerializeData(item)

		responses.Respond(w, r, http.StatusCreated, body)
	}
}
