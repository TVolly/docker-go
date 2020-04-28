package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TVolly/goapi-addresses/internal/controllers"
	"github.com/TVolly/goapi-addresses/internal/repositories"
)

func TestCommunityController_Index(t *testing.T) {
	repo := repositories.NewCommunityMemoryRepository()
	handler := controllers.NewCommunityController(repo).Index()

	testCases := []struct {
		name         string
		expectedCode int
	}{
		{
			name:         "valid",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, nil)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestCommunityController_Create(t *testing.T) {
	repo := repositories.NewCommunityMemoryRepository()
	handler := controllers.NewCommunityController(repo).Create()

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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest("", "", b)
			handler.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}
