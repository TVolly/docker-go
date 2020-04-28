package controllers

import (
	"net/http"

	"github.com/TVolly/goapi-addresses/internal/models"
	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/TVolly/goapi-addresses/internal/responses"
	"github.com/TVolly/goapi-addresses/internal/rules"
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
		rule := &rules.CommunityCreateRule{}
		if err := rule.Load(r.Body); err != nil {
			responses.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		item := &models.Community{}
		rule.Fill(item)

		if err := c.repo.Create(item); err != nil {
			responses.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		body := responses.SerializeData(item)
		responses.Respond(w, r, http.StatusCreated, body)
	}
}
