package main

import (
	"log"
	"time"

	"api-gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func main() {
	// Step 1: Initialize Consul client
	consulConfig := &api.Config{
		Address: "localhost:8500", // Update to the address of your Consul server
	}
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	// Step 2: Initialize Gin router
	r := gin.Default()

	// Step 3: Create the rate limiter (e.g., 5 requests per minute)
	rateLimiter := middlewares.NewRateLimiter(5, time.Minute)

	// Step 4: Register Middlewares
	r.Use(middlewares.RateLimitMiddleware(rateLimiter))                        // Rate limiting middleware
	r.Use(middlewares.JWTAuthMiddleware())                                     // JWT authentication middleware
	r.Use(middlewares.ServiceRoutingWithConsulSupportMiddleware(consulClient)) // Service routing via Consul

	// Step 5: Register the Health Check endpr.Use(middlewares.JWTAuthMiddleware())    oint
	r.GET("/ping", middlewares.HealthCheckHandler)

	// Step 6: Register a test endpoint for JWT token generation (this is just for testing JWT)
	// r.POST("/login", func(c *gin.Context) {
	// 	userId := c.DefaultPostForm("userID", "")
	// 	token, err := middlewares.GenerateJWT(userId)
	// 	if err != nil {
	// 		c.JSON(500, gin.H{"error": "Failed to generate token"})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{"token": token})
	// })

	// Step 7: Start the server
	log.Println("Starting server on :8080...")
	r.Run(":8080")
}
