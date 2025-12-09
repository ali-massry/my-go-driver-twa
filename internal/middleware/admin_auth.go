package middleware

import (
	"net/http"
	"strings"

	"my-go-driver/pkg/httputil"
	"my-go-driver/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// AdminAuth returns a gin middleware for admin JWT authentication
func AdminAuth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get authorization header
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			httputil.RespondError(c, http.StatusUnauthorized, "Authorization header required", "")
			c.Abort()
			return
		}

		// Check if it starts with Bearer
		if !strings.HasPrefix(authHeader, BearerPrefix) {
			httputil.RespondError(c, http.StatusUnauthorized, "Invalid authorization format", "")
			c.Abort()
			return
		}

		// Extract token
		token := strings.TrimPrefix(authHeader, BearerPrefix)
		if token == "" {
			httputil.RespondError(c, http.StatusUnauthorized, "Token is required", "")
			c.Abort()
			return
		}

		// Validate token
		userID, err := jwt.ValidateToken(token, jwtSecret)
		if err != nil {
			httputil.RespondError(c, http.StatusUnauthorized, "Invalid or expired token", err.Error())
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", userID)

		c.Next()
	}
}
