package app

import (
	"my-go-driver/internal/config"
	"my-go-driver/internal/handler"
	"my-go-driver/internal/repository"
	"my-go-driver/internal/service"
	"my-go-driver/pkg/logger"

	"gorm.io/gorm"
)

// Container holds all application dependencies
type Container struct {
	Config              *config.Config
	DB                  *gorm.DB
	Logger              *logger.Logger
	AdminCompanyHandler *handler.AdminCompanyHandler
	AdminDriverHandler  *handler.AdminDriverHandler
	AdminModuleHandler  *handler.AdminModuleHandler
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

	// Log migration info
	log.Info().Msg("Database migrations ready. All tables created via SQL migrations.")

	// Wire up dependencies using dependency injection

	// Repository layer
	companyRepo := repository.NewCompanyRepository(db)
	driverRepo := repository.NewDriverRepository(db)
	shiftRepo := repository.NewShiftRepository(db)
	moduleRepo := repository.NewModuleRepository(db)

	// Service layer
	companyService := service.NewCompanyService(companyRepo, cfg.JWT.Secret)
	driverService := service.NewDriverService(driverRepo, shiftRepo)
	shiftService := service.NewShiftService(shiftRepo)
	moduleService := service.NewModuleService(moduleRepo)

	// Handler layer
	adminCompanyHandler := handler.NewAdminCompanyHandler(companyService)
	adminDriverHandler := handler.NewAdminDriverHandler(driverService, shiftService)
	adminModuleHandler := handler.NewAdminModuleHandler(moduleService)

	return &Container{
		Config:              cfg,
		DB:                  db,
		Logger:              log,
		AdminCompanyHandler: adminCompanyHandler,
		AdminDriverHandler:  adminDriverHandler,
		AdminModuleHandler:  adminModuleHandler,
	}, nil
}
