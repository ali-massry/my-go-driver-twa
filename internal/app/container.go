package app

import (
	"github.com/ali-massry/my-go-driver/internal/config"
	"github.com/ali-massry/my-go-driver/internal/domain/user"
	"github.com/ali-massry/my-go-driver/internal/handler"
	"github.com/ali-massry/my-go-driver/internal/repository"
	"github.com/ali-massry/my-go-driver/internal/service"
	jwtpkg "github.com/ali-massry/my-go-driver/pkg/jwt"
	"github.com/ali-massry/my-go-driver/pkg/logger"
	"gorm.io/gorm"
)

// Container holds all application dependencies
type Container struct {
	Config      *config.Config
	DB          *gorm.DB
	Logger      *logger.Logger
	JWTManager  *jwtpkg.Manager
	UserHandler *handler.UserHandler
	AuthHandler *handler.AuthHandler
}

// NewContainer creates a new dependency injection container
func NewContainer(cfg *config.Config) (*Container, error) {
	// Initialize logger
	log := logger.New(logger.Config{
		Environment: cfg.Server.Environment,
		Level:       "info",
	})
	logger.SetGlobal(log)

	// Initialize database
	db, err := config.NewDatabase(cfg.Database)
	if err != nil {
		return nil, err
	}

	// Auto-migrate models
	// NOTE: In production, use golang-migrate to run migration files in /migrations
	// For development, we use AutoMigrate for the user table only
	// All other tables are created via SQL migrations
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return nil, err
	}

	// Log migration info
	log.Info().Msg("Database migrations ready. Run SQL migrations from /migrations folder for full schema.")

	// Initialize JWT manager
	jwtManager := jwtpkg.NewManager(cfg.JWT.Secret, cfg.JWT.Expiration)

	// Wire up dependencies using dependency injection
	// Repository layer
	userRepo := repository.NewUserRepository(db)

	// Service layer
	userService := service.NewUserService(userRepo, jwtManager)

	// Handler layer
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(userService)

	return &Container{
		Config:      cfg,
		DB:          db,
		Logger:      log,
		JWTManager:  jwtManager,
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}, nil
}
