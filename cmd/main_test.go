package main

import (
	"api-gateway/internal/middlewares"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthCheckHandler(t *testing.T) {
	// Set up a new gin router for testing
	r := gin.Default()
	r.GET("/health", middlewares.HealthCheckHandler)

	// Create a new http request to the /health enpdpoint
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response using httptest.NewRecorder
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Check that the status code is 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d, we got %d", http.StatusOK, rr.Code)
	}

	// Define expected response structure
	expected := map[string]string{
		"message": "API Gateway is up and running",
	}

	// Unmarshal response body into a map
	var actual map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Compare the expected and actual values
	if expected["message"] != actual["message"] {
		t.Errorf("expected body %s, got %s", expected["message"], actual["message"])
	}
}
