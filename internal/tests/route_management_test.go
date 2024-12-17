package tests

import (
	"api-gateway/internal/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestRouteManagement tests the routes management functionality
func TestRouteManagement(t *testing.T) {
	// Initializes the Gin engine for testing
	r := gin.Default()

	// Create a new RoutesConfigWrapper instance
	routeConfig := routes.NewRoutesConfig()

	// Register route management handlers
	routes.RegisterRouteHandlers(r, routeConfig)

	// Step 1: Test Registering a Route
	t.Run("Test Register Route", func(t *testing.T) {
		// Define a new route to register
		requestBody := map[string]string{
			"path":        "/test-path",
			"service_url": "http://localhost:8081", // Mock service URL
		}
		requestBodyBytes, _ := json.Marshal(requestBody)

		// Send a POST request to register the route
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(requestBodyBytes))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assert that the route was registered successfully
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "Route registered successfully!"}`, w.Body.String())
	})

	// Step 2: Test Listing the Routes
	t.Run("Test List Routes", func(t *testing.T) {
		// Send a GET request to list the registered routes
		req, _ := http.NewRequest(http.MethodGet, "/routes", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assert that the route we registered is listed
		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse := `{"routes":{"/test-path":"http://localhost:8081"}}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})

	// Step 3: Test Unregistering a Route
	t.Run("Test Unregister Route", func(t *testing.T) {
		// Define the route to remove
		requestBody := map[string]string{
			"path": "/test-path",
		}
		requestBodyBytes, _ := json.Marshal(requestBody)

		// Send a DELETE request to unregister the route
		req, _ := http.NewRequest(http.MethodDelete, "/unregister", bytes.NewReader(requestBodyBytes))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assert that the route was unregistered successfully
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "Route unregistered successfully"}`, w.Body.String())
	})

	// Step 4: Test Listing Routes After Unregistration
	t.Run("Test List Routes After Unregister", func(t *testing.T) {
		// Send a GET request to list routes again after unregistration
		req, _ := http.NewRequest(http.MethodGet, "/routes", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assert that no routes are registered now
		assert.Equal(t, http.StatusOK, w.Code)
		expectedResponse := `{"routes":{}}`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}
