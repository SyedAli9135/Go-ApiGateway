package middlewares

import (
	"github.com/gin-gonic/gin"
)

// HealthCheckHandler responds with simple health check endpoint
func HealthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API Gateway is up and running",
	})
}
