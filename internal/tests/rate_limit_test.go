package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"api-gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func TestRateLimiter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a rate limiter allowing 2 requests per second
	limiter := middlewares.NewRateLimiter(2, time.Second)

	// Set up a Gin router with the rate-limiting middleware
	r := gin.Default()
	r.Use(middlewares.RateLimitMiddleware(limiter))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	// Create a request to the test endpoint
	req, _ := http.NewRequest("GET", "/test", nil)

	// First two requests should pass
	for i := 1; i <= 2; i++ {
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rr.Code)
		}
	}

	// Third request should be rate limited
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusTooManyRequests {
		t.Errorf("expected status 429, got %d", rr.Code)
	}

	// Wait for the rate limiter to refill
	time.Sleep(1 * time.Second)

	// Fourth request should pass after refill
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	// Check that tokens have been refilled
	if limiter.Allow() {
		t.Errorf("Expected rate limit to still be active")
	}
}
