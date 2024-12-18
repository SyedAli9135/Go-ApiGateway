package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

// ServiceRoutingWithConsulSupportMiddleware forwards requests to appropriate services based on the path.
// It queries Consul to find the correct backend service URL at runtime.
func ServiceRoutingWithConsulSupportMiddleware(consulClient *api.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Step 1: Query Consul to get the backend service URL for the requested path
		serviceName := c.Request.URL.Path // Use path as the service name for this example

		// Step 2: Look up the service in Consul
		services, _, err := consulClient.Catalog().Service(serviceName, "", nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query Consul: " + err.Error()})
			c.Abort()
			return
		}

		// Step 3: If no service is found, return 404
		if len(services) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service not found in Consul"})
			c.Abort()
			return
		}

		// Step 4: Get the service URL (assuming we want the first service in the list)
		serviceURL := fmt.Sprintf("http://%s:%d", services[0].ServiceAddress, services[0].ServicePort)

		// Step 5: Parse the service URL
		parsedURL, err := url.Parse(serviceURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL from Consul"})
			c.Abort()
			return
		}

		// Step 6: Create a reverse proxy and forward the request to the backend service
		proxy := httputil.NewSingleHostReverseProxy(parsedURL)
		proxy.ErrorHandler = func(rw http.ResponseWriter, r *http.Request, err error) {
			http.Error(rw, "Reverse proxy error: "+err.Error(), http.StatusBadGateway)
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
