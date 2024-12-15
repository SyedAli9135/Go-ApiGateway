package main

import (
	"fmt"
	"log"

	"api-gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a gin router
	r := gin.Default()

	// Public route
	r.GET("/health", middlewares.HealthCheckHandler)

	// Protected routes (require JWT)
	authorized := r.Group("/")
	authorized.Use(middlewares.JWTAuthMiddleware())
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
