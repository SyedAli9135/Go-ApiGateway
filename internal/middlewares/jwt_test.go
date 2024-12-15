package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestJwtAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a valid token
	token, err := GenerateJWT("user123")
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	// Set up a Gin router with the JWT middleware
	r := gin.Default()
	r.Use(JWTAuthMiddleware())
	r.GET("/secure", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(200, gin.H{"message": "Secure route", "userID": userID})
	})

	// Test with valid token
	req, _ := http.NewRequest("GET", "/secure", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected value 200, got %d", rr.Code)
	}

	// Test with invalid token
	req, _ = http.NewRequest("GET", "/secure", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", rr.Code)
	}

}
