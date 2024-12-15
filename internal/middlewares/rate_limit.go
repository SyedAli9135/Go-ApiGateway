package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter defines a rate limiter with a token bucket
type RateLimiter struct {
	tokens     int           // Current number of tokens in the bucket
	maxTokens  int           // Maximum number of tokens in the bucket
	refillRate time.Duration // Refill rate for tokens
	lastRefill time.Time     // Last time tokens were refilled
	mu         sync.RWMutex  // RWMutex for thread-safe access
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(maxTokens int, refillRate time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:     maxTokens,
		maxTokens:  maxTokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow checks if a request can proceed and decrements a token if allowed
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)
	refillTokens := int(elapsed / rl.refillRate)

	if refillTokens > 0 {
		rl.tokens = min(rl.tokens+refillTokens, rl.maxTokens)
		rl.lastRefill = rl.lastRefill.Add(time.Duration(refillTokens) * rl.refillRate)
	}

	// Check if a token is available
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

// RateLimitMiddleware applies rate limiting to incoming requests
func RateLimitMiddleware(limit *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limit.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}
