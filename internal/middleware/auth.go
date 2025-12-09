package middleware

import (
	"net/http"
	"strings"

	"github.com/ali-massry/my-go-driver/pkg/httputil"
	jwtpkg "github.com/ali-massry/my-go-driver/pkg/jwt"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	UserIDKey           = "user_id"
	EmailKey            = "email"
)

// Auth returns a gin middleware for JWT authentication
func Auth(jwtManager *jwtpkg.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get authorization header
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			httputil.RespondError(c, http.StatusUnauthorized, "Authorization header required", nil)
			c.Abort()
			return
		}

		// Check if it starts with Bearer
		if !strings.HasPrefix(authHeader, BearerPrefix) {
			httputil.RespondError(c, http.StatusUnauthorized, "Invalid authorization format", nil)
			c.Abort()
			return
		}

		// Extract token
		token := strings.TrimPrefix(authHeader, BearerPrefix)
		if token == "" {
			httputil.RespondError(c, http.StatusUnauthorized, "Token is required", nil)
			c.Abort()
			return
		}

		// Validate token
		claims, err := jwtManager.Validate(token)
		if err != nil {
			if err == jwtpkg.ErrExpiredToken {
				httputil.RespondError(c, http.StatusUnauthorized, "Token has expired", nil)
			} else {
				httputil.RespondError(c, http.StatusUnauthorized, "Invalid token", nil)
			}
			c.Abort()
			return
		}

		// Set user info in context
		c.Set(UserIDKey, claims.UserID)
		c.Set(EmailKey, claims.Email)

		c.Next()
	}
}

// GetUserID retrieves the user ID from the context
func GetUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get(UserIDKey)
	if !exists {
		return 0, false
	}
	id, ok := userID.(uint)
	return id, ok
}

// GetEmail retrieves the email from the context
func GetEmail(c *gin.Context) (string, bool) {
	email, exists := c.Get(EmailKey)
	if !exists {
		return "", false
	}
	e, ok := email.(string)
	return e, ok
}
