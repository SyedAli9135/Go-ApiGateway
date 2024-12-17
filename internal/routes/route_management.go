package routes

import (
	// Importing common package where RouteConfigWrapper is defined
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRouteHandlers registers HTTP handlers for managing routes
func RegisterRouteHandlers(r *gin.Engine, routeConfig *RoutesConfigWrapper) {
	// POST /register: Add a new route
	r.POST("/register", func(c *gin.Context) {
		var body struct {
			Path       string `json:"path" binding:"required"`
			ServiceURL string `json:"service_url" binding:"required"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			// Return 400 Bad request if input is invalid
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add the route to the configuration
		routeConfig.AddRoute(body.Path, body.ServiceURL)
		c.JSON(http.StatusOK, gin.H{"message": "Route registered successfully!"})
	})

	// DELETE /unregister: Remove an existing route
	r.DELETE("/unregister", func(c *gin.Context) {
		var body struct {
			Path string `json:"path" binding:"required"` // Path to remove
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			// Return 400 Bad Request if input is invalid
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Remove the route from the configuration
		routeConfig.RemoveRoute(body.Path)
		c.JSON(http.StatusOK, gin.H{"message": "Route unregistered successfully"})
	})

	// GET /routes: List all registered routes
	r.GET("/routes", func(c *gin.Context) {
		// Retrieve the routes and return as JSON
		routes := routeConfig.GetRoutes()
		c.JSON(http.StatusOK, gin.H{"routes": routes})
	})
}
