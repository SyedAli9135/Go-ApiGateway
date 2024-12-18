package main

import (
	"api-gateway/internal/middlewares"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func main() {
	// Step 1: Initialize Consul client
	consulConfig := &api.Config{
		Address: "localhost:8500",
	}
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	// Step 2: Initialize Gin router
	r := gin.Default()

	// Step 3: Register the ServiceRoutingMiddleware with the Consul client
	r.Use(middlewares.ServiceRoutingWithConsulSupportMiddleware(consulClient))

	// Step 4: Set up a simple health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Step 5: Start the server
	r.Run(":8080")
}
