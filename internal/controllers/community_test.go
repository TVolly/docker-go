package controllers_test

import (
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
