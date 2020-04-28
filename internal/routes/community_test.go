package routes_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TVolly/goapi-addresses/internal/repositories"
	"github.com/TVolly/goapi-addresses/internal/routes"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// Проверка что роуты действительны

func TestConfigureCommunityRoutes(t *testing.T) {
	router := mux.NewRouter()
	repo := repositories.NewMemoryRegistry().Community()

	routes.NewRouteRegistry(router).ConfigureCommunityRoutes(repo)

	testCases := []struct {
		method       string
		route        string
		payload      interface{}
		expectedCode int
	}{
		{
			method:       http.MethodGet,
			route:        "/communities",
			expectedCode: http.StatusOK,
		},
		{
			method:       http.MethodPost,
			route:        "/communities",
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("[%s]%s", tc.method, tc.route), func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, tc.route, nil)

			router.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}
