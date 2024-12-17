package routes

import (
	"api-gateway/internal/common"
	"sync"
)

// RoutesConfigWrapper is a wrapper around common.RoutesConfig to add methods in the routes package
type RoutesConfigWrapper struct {
	*common.RoutesConfig
	mu sync.RWMutex // Mutex to protect concurrent access to Routes
}

// NewRoutesConfig creates and initializes a new instance of RoutesConfigWrapper
func NewRoutesConfig() *RoutesConfigWrapper {
	return &RoutesConfigWrapper{
		RoutesConfig: &common.RoutesConfig{
			Routes: make(map[string]string),
		},
	}
}

// AddRoute adds a new route mapping to the RoutesConfig with thread safety
func (rc *RoutesConfigWrapper) AddRoute(path string, serviceURL string) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.Routes[path] = serviceURL
}

// RemoveRoute removes an existing route mapping from the RoutesConfig with thread safety
func (rc *RoutesConfigWrapper) RemoveRoute(path string) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	delete(rc.Routes, path)
}

// GetRoutes retrieves all the route mappings in a thread-safe manner
func (rc *RoutesConfigWrapper) GetRoutes() map[string]string {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return rc.Routes
}

// GetRoute retrieves the service URL for a given path in a thread-safe manner
func (rc *RoutesConfigWrapper) GetRoute(path string) (string, bool) {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	target, exists := rc.Routes[path]
	return target, exists
}
