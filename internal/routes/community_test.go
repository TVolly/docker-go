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

func TestConfigureCommunityRoutes(t *testing.T) {
	router := mux.NewRouter()
	repoRegistry := repositories.NewMemoryRegistry()
	routes.ConfigureCommunityRoutes(router, repoRegistry)

	testCases := []struct {
		method       string
		route        string
		expectedCode int
	}{
		{
			method:       http.MethodGet,
			route:        "/communities",
			expectedCode: http.StatusOK,
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
