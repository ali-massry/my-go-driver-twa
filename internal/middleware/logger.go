package middleware

import (
	"time"

	"github.com/ali-massry/my-go-driver/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Logger returns a gin middleware for logging HTTP requests
func Logger(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get status code
		statusCode := c.Writer.Status()

		// Build log fields
		fields := map[string]interface{}{
			"method":     c.Request.Method,
			"path":       path,
			"status":     statusCode,
			"latency_ms": latency.Milliseconds(),
			"ip":         c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}

		if raw != "" {
			fields["query"] = raw
		}

		// Log based on status code
		switch {
		case statusCode >= 500:
			log.WithFields(fields).Error().Msg("Server error")
		case statusCode >= 400:
			log.WithFields(fields).Warn().Msg("Client error")
		default:
			log.WithFields(fields).Info().Msg("Request processed")
		}
	}
}
