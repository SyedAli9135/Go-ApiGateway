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

	// Register Heath Check route
	r.GET("/health", middlewares.HealthCheckHandler)

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
