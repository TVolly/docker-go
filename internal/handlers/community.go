package handlers

import (
	"net/http"
	"strconv"

	"github.com/TVolly/goapi-addresses/internal/models"
	"github.com/TVolly/goapi-addresses/internal/responses"
	"github.com/TVolly/goapi-addresses/internal/rules"
	"github.com/gorilla/mux"
)

func (h *handlersStore) BindCommunityHandlers() {
	s := h.router.PathPrefix("/communities").Subrouter()

	s.HandleFunc("", h.communityIndex()).Methods(http.MethodGet)
	s.HandleFunc("", h.communityCreate()).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", h.communityShow()).Methods(http.MethodGet)
}

func (h *handlersStore) communityIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := h.store.Community().List()
		p := responses.NewPagination(items)

		responses.Respond(w, r, http.StatusOK, p)
	}
}

func (h *handlersStore) communityCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rule := &rules.CommunityCreateRule{}
		if err := rule.Load(r.Body); err != nil {
			responses.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		item := &models.Community{}
		rule.Fill(item)

		if err := h.store.Community().Create(item); err != nil {
			responses.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		body := responses.SerializeData(item)
		responses.Respond(w, r, http.StatusCreated, body)
	}
}

func (h *handlersStore) communityShow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			responses.Error(w, r, http.StatusNotFound, err)
			return
		}

		item, err := h.store.Community().Find(id)
		if err != nil {
			responses.Error(w, r, http.StatusNotFound, err)
			return
		}

		body := responses.SerializeData(item)
		responses.Respond(w, r, http.StatusOK, body)
	}
}
