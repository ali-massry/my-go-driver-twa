package router

import (
	"my-go-driver/internal/handler"
	"my-go-driver/internal/middleware"
	"my-go-driver/pkg/logger"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(
	r *gin.Engine,
	log *logger.Logger,
	jwtSecret string,
	adminCompanyHandler *handler.AdminCompanyHandler,
	adminDriverHandler *handler.AdminDriverHandler,
	adminModuleHandler *handler.AdminModuleHandler,
) {
	// Global middleware
	r.Use(middleware.Logger(log))
	r.Use(gin.Recovery())

	// Health check endpoint

	r.Group("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  "healthy",
			"service": "TWA Driver API",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{

		// Admin routes
		admin := v1.Group("/admin")
		{
			// Public admin authentication routes
			adminAuth := admin.Group("/auth")
			{
				adminAuth.POST("/login", adminCompanyHandler.LoginAdmin)
			}

			// Public company creation (for initial setup)
			admin.POST("/companies", adminCompanyHandler.CreateCompany)

			// Protected admin routes (require authentication)
			protected := admin.Group("")
			protected.Use(middleware.AdminAuth(jwtSecret))
			{
				// Admin profile
				protected.GET("/auth/me", adminCompanyHandler.GetAdminProfile)

				// Company management
				companies := protected.Group("/companies")
				{
					companies.GET("", adminCompanyHandler.ListCompanies)
					companies.GET("/:id", adminCompanyHandler.GetCompany)
					companies.PUT("/:id", adminCompanyHandler.UpdateCompany)
					companies.DELETE("/:id", adminCompanyHandler.DeleteCompany)
					companies.PUT("/:id/branding", adminCompanyHandler.UpdateBranding)
					companies.PUT("/:id/suspend", adminCompanyHandler.SuspendCompany)
					companies.PUT("/:id/activate", adminCompanyHandler.ActivateCompany)

					// Company modules
					companies.POST("/:id/modules", adminModuleHandler.AssignModuleToCompany)
					companies.GET("/:id/modules", adminModuleHandler.GetCompanyModules)
					companies.DELETE("/:id/modules/:module_id", adminModuleHandler.RemoveModuleFromCompany)
				}

				// Driver management
				drivers := protected.Group("/drivers")
				{
					drivers.POST("", adminDriverHandler.CreateDriver)
					drivers.GET("", adminDriverHandler.ListDrivers)
					drivers.GET("/:id", adminDriverHandler.GetDriver)
					drivers.PUT("/:id", adminDriverHandler.UpdateDriver)
					drivers.DELETE("/:id", adminDriverHandler.DeleteDriver)
					drivers.PUT("/:id/assign-company", adminDriverHandler.AssignDriverToCompany)
					drivers.PUT("/:id/block", adminDriverHandler.BlockDriver)
					drivers.PUT("/:id/unblock", adminDriverHandler.UnblockDriver)
					drivers.GET("/:id/performance", adminDriverHandler.GetDriverPerformance)
					drivers.GET("/:id/shifts", adminDriverHandler.GetDriverShifts)
				}

				// Modules
				modules := protected.Group("/modules")
				{
					modules.GET("", adminModuleHandler.ListAllModules)
				}
			}
		}
	}
}
