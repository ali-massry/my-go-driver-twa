package router

import (
	"github.com/ali-massry/my-go-driver/internal/handler"
	"github.com/ali-massry/my-go-driver/internal/middleware"
	jwtpkg "github.com/ali-massry/my-go-driver/pkg/jwt"
	"github.com/ali-massry/my-go-driver/pkg/logger"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(
	r *gin.Engine,
	log *logger.Logger,
	jwtManager *jwtpkg.Manager,
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
) {
	// Global middleware
	r.Use(middleware.Logger(log))
	r.Use(gin.Recovery())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Public authentication routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes (require authentication)
		protected := v1.Group("")
		protected.Use(middleware.Auth(jwtManager))
		{
			// Auth routes
			protected.GET("/auth/me", authHandler.Me)

			// User management routes
			users := protected.Group("/users")
			{
				users.GET("", userHandler.GetUsers)
				users.POST("", userHandler.CreateUser)
				users.GET("/:id", userHandler.GetUserByID)
				users.PUT("/:id", userHandler.UpdateUser)
				users.DELETE("/:id", userHandler.DeleteUser)
			}
		}
	}
}
