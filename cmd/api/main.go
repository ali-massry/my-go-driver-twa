package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"my-go-driver/internal/app"
	"my-go-driver/internal/config"
	"my-go-driver/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set Gin mode based on environment
	if cfg.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize dependency injection container
	container, err := app.NewContainer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	// Create Gin router
	r := gin.New()

	// Setup routes with middleware
	router.SetupRoutes(
		r,
		container.Logger,
		cfg.JWT.Secret,
		container.AdminCompanyHandler,
		container.AdminDriverHandler,
		container.AdminModuleHandler,
	)

	// Create HTTP server
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// Start server in a goroutine
	go func() {
		container.Logger.Info().
			Str("port", cfg.Server.Port).
			Str("environment", cfg.Server.Environment).
			Msg("Server starting")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			container.Logger.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	container.Logger.Info().Msg("Shutting down server...")

	// Graceful shutdown with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		container.Logger.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	// Close database connection
	if container.DB != nil {
		sqlDB, err := container.DB.DB()
		if err == nil {
			sqlDB.Close()
		}
	}

	container.Logger.Info().Msg("Server exited")
}
