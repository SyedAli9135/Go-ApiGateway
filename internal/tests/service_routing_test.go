package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"api-gateway/internal/common"
	"api-gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

// CustomResponseWriter wraps httptest.ResponseRecorder and implements gin.ResponseWriter
type CustomResponseWriter struct {
	*httptest.ResponseRecorder
}

func (w *CustomResponseWriter) CloseNotify() <-chan bool {
	// Returning nil here, since we don't need CloseNotify for testing
	return nil
}

// StartMockServer creates and returns a mock backend server using Gin
func StartMockServer() *httptest.Server {
	r := gin.Default()

	// Define mock backend routes
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success from mock server"})
	})

	// Use httptest to create a mock server
	server := httptest.NewServer(r) // Create and return the test server
	return server
}

func TestServiceRoutingMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Step 1: Start the mock server
	mockServer := StartMockServer()
	mockServerURL := mockServer.URL // Capture the mock server URL

	t.Logf("Mock Server URL: %s", mockServerURL)
	defer mockServer.Close() // Now we're calling Close on the httptest.Server

	// Step 2: Set up RoutesConfig with the mock server URL
	// Update RoutesConfig to map the "/test" path to the mock server URL
	routesConfig := common.RoutesConfig{
		Routes: map[string]string{
			"/test": mockServerURL, // Map "/test" to the mock server URL
		},
	}

	// Step 3: Set up the API Gateway with the routing middleware
	r := gin.Default()

	// Register the ServiceRoutingMiddleware, passing the RoutesConfig
	r.Use(middlewares.ServiceRoutingMiddleware(routesConfig))

	// Step 4: Send a request to the API Gateway
	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()
	// Create a custom response writer that wraps the httptest.ResponseRecorder
	customRR := &CustomResponseWriter{ResponseRecorder: rr}

	r.ServeHTTP(customRR, req)

	// Step 5: Log and verify the response
	t.Logf("Response status: %d", rr.Code)
	t.Logf("Response body: %s", rr.Body.String())

	// Verify the response
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	expectedResponse := `{"message":"Success from mock server"}`

	if rr.Body.String() != expectedResponse {
		t.Errorf("unexpected response body: got %s, want %s", rr.Body.String(), expectedResponse)
	}
}
