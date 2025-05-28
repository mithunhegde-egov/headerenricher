// Package autogin provides automatic header enrichment for Gin applications
// Just use: import "localisationgo/internal/middleware/autogin"
// Then call autogin.NewEngine() instead of gin.New()
package autogin

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// generateRequestID creates a unique ID for each request
func generateRequestID() string {
	return uuid.NewString()
}

// generateTimestamp creates a timestamp for the request
func generateTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// HeaderEnricherMiddleware adds headers to all requests
func HeaderEnricherMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate request ID and timestamp
		requestID := generateRequestID()
		timestamp := generateTimestamp()

		// Add custom headers to the request
		c.Request.Header.Set("X-Request-ID", requestID)
		c.Request.Header.Set("X-Request-Timestamp", timestamp)
		c.Request.Header.Set("X-Service", "LocalisationGoService")

		// Continue with the next handlers
		c.Next()

		// Optionally add headers to the response too
		c.Writer.Header().Set("X-Processed-By", "autogin")
	}
}

// NewEngine creates a new gin.Engine with header enrichment middleware already applied
func NewEngine() *gin.Engine {
	// Create a new engine
	engine := gin.New()

	// Add our middleware
	engine.Use(HeaderEnricherMiddleware())

	return engine
}

// DefaultEngine creates a new gin.Engine with default middlewares plus our header enrichment
func DefaultEngine() *gin.Engine {
	// Create a default engine with recovery and logger
	engine := gin.Default()

	// Add our middleware
	engine.Use(HeaderEnricherMiddleware())

	return engine
}
