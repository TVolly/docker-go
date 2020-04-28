package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TVolly/goapi-addresses/internal/models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/TVolly/goapi-addresses/internal/handlers"
	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/TVolly/goapi-addresses/internal/responses"
)

func TestHandlersStore_communityIndex(t *testing.T) {
	r := mux.NewRouter()
	h := handlers.NewHandler(r, repositories.NewMemoryStore())
	h.BindCommunityHandlers()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/communities", nil)

	r.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, responses.CONTENT_TYPE_JSON, rec.Header().Get("Content-Type"))
}

func TestHandlersStore_communityCreate(t *testing.T) {
	r := mux.NewRouter()
	h := handlers.NewHandler(r, repositories.NewMemoryStore())
	h.BindCommunityHandlers()

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"name": "Valid name",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      nil,
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid name",
			payload: map[string]string{
				"name": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			body := encodePayload(tc.payload)
			req, _ := http.NewRequest("POST", "/communities", body)

			r.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Equal(t, responses.CONTENT_TYPE_JSON, rec.Header().Get("Content-Type"))
		})
	}
}

func TestHandlersStore_communityShow(t *testing.T) {
	r := mux.NewRouter()
	s := repositories.NewMemoryStore()
	h := handlers.NewHandler(r, s)
	h.BindCommunityHandlers()

	rec1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/communities/999", nil)
	r.ServeHTTP(rec1, req1)
	assert.Equal(t, http.StatusNotFound, rec1.Code)
	assert.Equal(t, responses.CONTENT_TYPE_JSON, rec1.Header().Get("Content-Type"))

	model := models.TestCommunity()
	s.Community().Create(model)
	rec2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", fmt.Sprintf("/communities/%d", model.ID), nil)
	r.ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, responses.CONTENT_TYPE_JSON, rec2.Header().Get("Content-Type"))
}
