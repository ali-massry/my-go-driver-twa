package repository

import (
	"context"
	"fmt"

	"my-go-driver/internal/domain/company"

	"gorm.io/gorm"
)

type companyRepository struct {
	db *gorm.DB
}

// NewCompanyRepository creates a new company repository
func NewCompanyRepository(db *gorm.DB) company.Repository {
	return &companyRepository{db: db}
}

func (r *companyRepository) Create(ctx context.Context, c *company.Company) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *companyRepository) GetByID(ctx context.Context, id uint64) (*company.Company, error) {
	var c company.Company
	err := r.db.WithContext(ctx).First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *companyRepository) Update(ctx context.Context, c *company.Company) error {
	return r.db.WithContext(ctx).Save(c).Error
}

func (r *companyRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&company.Company{}, id).Error
}

func (r *companyRepository) List(ctx context.Context, query company.ListCompaniesQuery) ([]company.Company, int64, error) {
	var companies []company.Company
	var total int64

	db := r.db.WithContext(ctx).Model(&company.Company{})

	// Apply filters
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	if query.Search != "" {
		searchPattern := "%" + query.Search + "%"
		db = db.Where("name LIKE ? OR email LIKE ? OR phone LIKE ?", searchPattern, searchPattern, searchPattern)
	}

	// Get total count
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Set defaults
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Limit <= 0 {
		query.Limit = 10
	}

	// Apply pagination
	offset := (query.Page - 1) * query.Limit
	err := db.Offset(offset).Limit(query.Limit).Order("created_at DESC").Find(&companies).Error

	return companies, total, err
}

func (r *companyRepository) UpdateBranding(ctx context.Context, id uint64, branding company.UpdateBrandingRequest) error {
	updates := make(map[string]interface{})

	if branding.LogoURL != "" {
		updates["logo_url"] = branding.LogoURL
	}
	if branding.ColorPalette != nil {
		updates["color_palette"] = branding.ColorPalette
	}
	if branding.FontFamily != "" {
		updates["font_family"] = branding.FontFamily
	}

	if len(updates) == 0 {
		return fmt.Errorf("no branding fields to update")
	}

	return r.db.WithContext(ctx).Model(&company.Company{}).Where("id = ?", id).Updates(updates).Error
}

func (r *companyRepository) UpdateStatus(ctx context.Context, id uint64, status company.CompanyStatus) error {
	return r.db.WithContext(ctx).Model(&company.Company{}).Where("id = ?", id).Update("status", status).Error
}

// Company Admin methods
func (r *companyRepository) CreateAdmin(ctx context.Context, admin *company.CompanyAdmin) error {
	return r.db.WithContext(ctx).Create(admin).Error
}

func (r *companyRepository) GetAdminByEmail(ctx context.Context, email string) (*company.CompanyAdmin, error) {
	var admin company.CompanyAdmin
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *companyRepository) GetAdminByID(ctx context.Context, id uint64) (*company.CompanyAdmin, error) {
	var admin company.CompanyAdmin
	err := r.db.WithContext(ctx).Preload("Company").First(&admin, id).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *companyRepository) ListAdmins(ctx context.Context, companyID uint64) ([]company.CompanyAdmin, error) {
	var admins []company.CompanyAdmin
	err := r.db.WithContext(ctx).Where("company_id = ?", companyID).Find(&admins).Error
	return admins, err
}
