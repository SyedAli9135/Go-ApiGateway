package middlewares

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"api-gateway/internal/common"

	"github.com/gin-gonic/gin"
)

// ServiceRoutingMiddleware forwards requests to appropriate services based on the path
func ServiceRoutingMiddleware(routes common.RoutesConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the incoming path matches any configured route
		target, ok := routes.Routes[c.Request.URL.Path]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
			c.Abort()
			return
		}

		// Parse the backend URL
		parsedURL, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
			c.Abort()
			return
		}

		// Create a reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(parsedURL)
		proxy.ErrorHandler = func(rw http.ResponseWriter, r *http.Request, err error) {
			http.Error(rw, "Reverse proxy error: "+err.Error(), http.StatusBadGateway)
		}
		// Forward the request to the backend service
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
