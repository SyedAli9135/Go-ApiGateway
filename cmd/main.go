package main

import (
	"net/http"
	"time"

	"api-gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	// Define routes configuration
	routesConfig := middlewares.RoutesConfig{
		Routes: map[string]string{
			"/users": "http://localhost:8081", // Map /users to backend server on port 8081
		},
	}

	// Apply Middlewares
	r.Use(middlewares.RateLimitMiddleware(middlewares.NewRateLimiter(5, time.Second)))
	r.Use(middlewares.ServiceRoutingMiddleware(routesConfig))

	// Health check route for API Gateway
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API Gateway is running"})
	})

	// Start the API Gateway on port 8080
	r.Run(":8080")
}
