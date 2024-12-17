package main

import (
	// Import the common package where RouteConfigWrapper is defined
	"api-gateway/internal/middlewares"
	"api-gateway/internal/routes"

	"github.com/gin-gonic/gin"
)

// main initializes the API Gateway application, registers route handlers, and starts the server.
func main() {
	// Initialize the Gin engine
	r := gin.Default()

	// Create a new RouteConfigWrapper instance
	routeConfigWrapper := routes.NewRoutesConfig()

	// Register route management endpoints (e.g., /register, /unregister, /routes)
	routes.RegisterRouteHandlers(r, routeConfigWrapper)

	// Pass the RoutesConfig from the RouteConfigWrapper to the ServiceRoutingMiddleware
	r.Use(middlewares.ServiceRoutingMiddleware(*routeConfigWrapper.RoutesConfig)) // Pass the actual RoutesConfig

	// Start the API Gateway on port 8080
	r.Run(":8080")
}
