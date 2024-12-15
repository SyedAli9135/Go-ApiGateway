package main

import (
	"fmt"
	"log"
	"time"

	"api-gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a gin router
	r := gin.Default()

	// Public route
	r.GET("/health", middlewares.HealthCheckHandler)

	// Create a global rate limiter
	rateLimiter := middlewares.NewRateLimiter(5, time.Second) // 5 requests per second

	// Protected routes (require JWT) with rate limiting
	authorized := r.Group("/")
	authorized.Use(middlewares.JWTAuthMiddleware())
	authorized.Use(middlewares.RateLimitMiddleware(rateLimiter))
	{
		authorized.GET("/secure", func(c *gin.Context) {
			userID, _ := c.Get("userID") // Retrieve userID from context
			c.JSON(200, gin.H{"message": "Secure route accessed",
				"userID": userID})
		})
	}

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
