package service

import (
	"context"
	"errors"
	"fmt"
	"math"

	"my-go-driver/internal/domain/company"
	"my-go-driver/pkg/hash"
	"my-go-driver/pkg/jwt"

	"gorm.io/gorm"
)

type companyService struct {
	repo      company.Repository
	jwtSecret string
}

// NewCompanyService creates a new company service
func NewCompanyService(repo company.Repository, jwtSecret string) company.Service {
	return &companyService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *companyService) CreateCompany(ctx context.Context, req company.CreateCompanyRequest) (*company.CompanyWithOwnerResponse, error) {
	// Check if admin email already exists
	existingAdmin, err := s.repo.GetAdminByEmail(ctx, req.OwnerEmail)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("error checking admin email: %w", err)
	}
	if existingAdmin != nil {
		return nil, fmt.Errorf("admin email already exists")
	}

	// Create company
	newCompany := &company.Company{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
		Timezone: req.Timezone,
		Status:   company.CompanyStatusActive,
	}

	if req.Timezone == "" {
		newCompany.Timezone = "UTC"
	}

	if err := s.repo.Create(ctx, newCompany); err != nil {
		return nil, fmt.Errorf("failed to create company: %w", err)
	}

	// Hash password for owner
	hashedPassword, err := hash.HashPassword(req.OwnerPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create company owner
	owner := &company.CompanyAdmin{
		CompanyID:    newCompany.ID,
		FullName:     req.OwnerName,
		Email:        req.OwnerEmail,
		Phone:        req.OwnerPhone,
		PasswordHash: hashedPassword,
		Role:         company.AdminRoleOwner,
		IsActive:     true,
	}

	if err := s.repo.CreateAdmin(ctx, owner); err != nil {
		// Rollback: delete the company
		_ = s.repo.Delete(ctx, newCompany.ID)
		return nil, fmt.Errorf("failed to create owner: %w", err)
	}

	return &company.CompanyWithOwnerResponse{
		Company: s.toCompanyResponse(newCompany),
		Owner:   s.toAdminResponse(owner),
	}, nil
}

func (s *companyService) GetCompany(ctx context.Context, id uint64) (*company.CompanyResponse, error) {
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("company not found")
		}
		return nil, err
	}

	response := s.toCompanyResponse(c)
	return &response, nil
}

func (s *companyService) UpdateCompany(ctx context.Context, id uint64, req company.UpdateCompanyRequest) (*company.CompanyResponse, error) {
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("company not found")
		}
		return nil, err
	}

	// Update fields
	if req.Name != "" {
		c.Name = req.Name
	}
	if req.Email != "" {
		c.Email = req.Email
	}
	if req.Phone != "" {
		c.Phone = req.Phone
	}
	if req.Address != "" {
		c.Address = req.Address
	}
	if req.Timezone != "" {
		c.Timezone = req.Timezone
	}

	if err := s.repo.Update(ctx, c); err != nil {
		return nil, fmt.Errorf("failed to update company: %w", err)
	}

	response := s.toCompanyResponse(c)
	return &response, nil
}

func (s *companyService) DeleteCompany(ctx context.Context, id uint64) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("company not found")
		}
		return err
	}

	return s.repo.Delete(ctx, id)
}

func (s *companyService) ListCompanies(ctx context.Context, query company.ListCompaniesQuery) (*company.PaginatedCompaniesResponse, error) {
	companies, total, err := s.repo.List(ctx, query)
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

	responses := make([]company.CompanyResponse, len(companies))
	for i, c := range companies {
		responses[i] = s.toCompanyResponse(&c)
	}

	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))

	return &company.PaginatedCompaniesResponse{
		Companies:  responses,
		TotalCount: total,
		Page:       query.Page,
		Limit:      query.Limit,
		TotalPages: totalPages,
	}, nil
}

func (s *companyService) UpdateBranding(ctx context.Context, id uint64, req company.UpdateBrandingRequest) (*company.CompanyResponse, error) {
	if err := s.repo.UpdateBranding(ctx, id, req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("company not found")
		}
		return nil, fmt.Errorf("failed to update branding: %w", err)
	}

	// Get updated company
	return s.GetCompany(ctx, id)
}

func (s *companyService) SuspendCompany(ctx context.Context, id uint64) error {
	return s.repo.UpdateStatus(ctx, id, company.CompanyStatusSuspended)
}

func (s *companyService) ActivateCompany(ctx context.Context, id uint64) error {
	return s.repo.UpdateStatus(ctx, id, company.CompanyStatusActive)
}

func (s *companyService) LoginAdmin(ctx context.Context, req company.LoginRequest) (*company.LoginResponse, error) {
	admin, err := s.repo.GetAdminByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("invalid email or password")
		}
		return nil, err
	}

	// Check if admin is active
	if !admin.IsActive {
		return nil, fmt.Errorf("account is inactive")
	}

	// Verify password
	if !hash.CheckPasswordHash(req.Password, admin.PasswordHash) {
		return nil, fmt.Errorf("invalid email or password")
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(admin.ID, s.jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &company.LoginResponse{
		Admin: s.toAdminResponse(admin),
		Token: token,
	}, nil
}

func (s *companyService) GetAdminProfile(ctx context.Context, adminID uint64) (*company.CompanyAdminResponse, error) {
	admin, err := s.repo.GetAdminByID(ctx, adminID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("admin not found")
		}
		return nil, err
	}

	response := s.toAdminResponse(admin)
	return &response, nil
}

// Helper methods
func (s *companyService) toCompanyResponse(c *company.Company) company.CompanyResponse {
	return company.CompanyResponse{
		ID:           c.ID,
		Name:         c.Name,
		Email:        c.Email,
		Phone:        c.Phone,
		Address:      c.Address,
		Timezone:     c.Timezone,
		LogoURL:      c.LogoURL,
		ColorPalette: c.ColorPalette,
		FontFamily:   c.FontFamily,
		Status:       c.Status,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

func (s *companyService) toAdminResponse(admin *company.CompanyAdmin) company.CompanyAdminResponse {
	return company.CompanyAdminResponse{
		ID:        admin.ID,
		CompanyID: admin.CompanyID,
		FullName:  admin.FullName,
		Email:     admin.Email,
		Phone:     admin.Phone,
		Role:      admin.Role,
		IsActive:  admin.IsActive,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}
