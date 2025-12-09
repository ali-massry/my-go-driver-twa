package company

import "context"

// Repository defines the interface for company data access
type Repository interface {
	// Company operations
	Create(ctx context.Context, company *Company) error
	GetByID(ctx context.Context, id uint64) (*Company, error)
	Update(ctx context.Context, company *Company) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, query ListCompaniesQuery) ([]Company, int64, error)
	UpdateBranding(ctx context.Context, id uint64, branding UpdateBrandingRequest) error
	UpdateStatus(ctx context.Context, id uint64, status CompanyStatus) error

	// Company Admin operations
	CreateAdmin(ctx context.Context, admin *CompanyAdmin) error
	GetAdminByEmail(ctx context.Context, email string) (*CompanyAdmin, error)
	GetAdminByID(ctx context.Context, id uint64) (*CompanyAdmin, error)
	ListAdmins(ctx context.Context, companyID uint64) ([]CompanyAdmin, error)
}
