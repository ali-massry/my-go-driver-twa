package service

import (
	"context"
	"errors"
	"fmt"
	"math"

	"my-go-driver/internal/domain/driver"
	"my-go-driver/internal/domain/shift"
	"my-go-driver/pkg/hash"

	"gorm.io/gorm"
)

type driverService struct {
	repo      driver.Repository
	shiftRepo shift.Repository
}

// NewDriverService creates a new driver service
func NewDriverService(repo driver.Repository, shiftRepo shift.Repository) driver.Service {
	return &driverService{
		repo:      repo,
		shiftRepo: shiftRepo,
	}
}

func (s *driverService) CreateDriver(ctx context.Context, req driver.CreateDriverRequest) (*driver.DriverResponse, error) {
	// Check if phone already exists for this company
	existingDriver, err := s.repo.GetByPhone(ctx, req.Phone, req.CompanyID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("error checking phone: %w", err)
	}
	if existingDriver != nil {
		return nil, fmt.Errorf("driver with this phone already exists in this company")
	}

	// Hash password
	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create driver
	newDriver := &driver.Driver{
		CompanyID:    req.CompanyID,
		StoreID:      req.StoreID,
		FullName:     req.FullName,
		Phone:        req.Phone,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Status:       driver.DriverStatusActive,
		OnlineStatus: driver.OnlineStatusOffline,
		ProfilePhoto: req.ProfilePhoto,
		Rating:       0.0,
	}

	if err := s.repo.Create(ctx, newDriver); err != nil {
		return nil, fmt.Errorf("failed to create driver: %w", err)
	}

	response := s.toDriverResponse(newDriver)
	return &response, nil
}

func (s *driverService) GetDriver(ctx context.Context, id uint64) (*driver.DriverResponse, error) {
	d, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("driver not found")
		}
		return nil, err
	}

	response := s.toDriverResponse(d)
	return &response, nil
}

func (s *driverService) UpdateDriver(ctx context.Context, id uint64, req driver.UpdateDriverRequest) (*driver.DriverResponse, error) {
	d, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("driver not found")
		}
		return nil, err
	}

	// Update fields
	if req.FullName != "" {
		d.FullName = req.FullName
	}
	if req.Phone != "" {
		// Check if new phone already exists for this company
		existingDriver, err := s.repo.GetByPhone(ctx, req.Phone, d.CompanyID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("error checking phone: %w", err)
		}
		if existingDriver != nil && existingDriver.ID != id {
			return nil, fmt.Errorf("driver with this phone already exists in this company")
		}
		d.Phone = req.Phone
	}
	if req.Email != "" {
		d.Email = req.Email
	}
	if req.StoreID != nil {
		d.StoreID = req.StoreID
	}
	if req.ProfilePhoto != "" {
		d.ProfilePhoto = req.ProfilePhoto
	}

	if err := s.repo.Update(ctx, d); err != nil {
		return nil, fmt.Errorf("failed to update driver: %w", err)
	}

	response := s.toDriverResponse(d)
	return &response, nil
}

func (s *driverService) DeleteDriver(ctx context.Context, id uint64) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("driver not found")
		}
		return err
	}

	return s.repo.Delete(ctx, id)
}

func (s *driverService) ListDrivers(ctx context.Context, query driver.ListDriversQuery) (*driver.PaginatedDriversResponse, error) {
	drivers, total, err := s.repo.List(ctx, query)
	if err != nil {
		return nil, err
	}

	// Set defaults
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Limit <= 0 {
		query.Limit = 10
	}

	responses := make([]driver.DriverResponse, len(drivers))
	for i, d := range drivers {
		responses[i] = s.toDriverResponse(&d)
	}

	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))

	return &driver.PaginatedDriversResponse{
		Drivers:    responses,
		TotalCount: total,
		Page:       query.Page,
		Limit:      query.Limit,
		TotalPages: totalPages,
	}, nil
}

func (s *driverService) AssignToCompany(ctx context.Context, driverID uint64, req driver.AssignDriverToCompanyRequest) (*driver.DriverResponse, error) {
	d, err := s.repo.GetByID(ctx, driverID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("driver not found")
		}
		return nil, err
	}

	d.CompanyID = req.CompanyID
	d.StoreID = req.StoreID

	if err := s.repo.Update(ctx, d); err != nil {
		return nil, fmt.Errorf("failed to assign driver: %w", err)
	}

	response := s.toDriverResponse(d)
	return &response, nil
}

func (s *driverService) BlockDriver(ctx context.Context, driverID uint64) error {
	return s.repo.UpdateStatus(ctx, driverID, driver.DriverStatusSuspended)
}

func (s *driverService) UnblockDriver(ctx context.Context, driverID uint64) error {
	return s.repo.UpdateStatus(ctx, driverID, driver.DriverStatusActive)
}

func (s *driverService) GetDriverPerformance(ctx context.Context, driverID uint64) (*driver.DriverPerformance, error) {
	// Check if driver exists
	_, err := s.repo.GetByID(ctx, driverID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("driver not found")
		}
		return nil, err
	}

	performance, err := s.repo.GetPerformance(ctx, driverID)
	if err != nil {
		return nil, fmt.Errorf("failed to get driver performance: %w", err)
	}

	return performance, nil
}

// Helper methods
func (s *driverService) toDriverResponse(d *driver.Driver) driver.DriverResponse {
	return driver.DriverResponse{
		ID:           d.ID,
		CompanyID:    d.CompanyID,
		StoreID:      d.StoreID,
		FullName:     d.FullName,
		Phone:        d.Phone,
		Email:        d.Email,
		Status:       d.Status,
		OnlineStatus: d.OnlineStatus,
		Rating:       d.Rating,
		ProfilePhoto: d.ProfilePhoto,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}
