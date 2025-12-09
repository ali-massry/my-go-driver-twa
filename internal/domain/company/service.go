package company

import "context"

// Service defines the interface for company business logic
type Service interface {
	// Company operations
	CreateCompany(ctx context.Context, req CreateCompanyRequest) (*CompanyWithOwnerResponse, error)
	GetCompany(ctx context.Context, id uint64) (*CompanyResponse, error)
	UpdateCompany(ctx context.Context, id uint64, req UpdateCompanyRequest) (*CompanyResponse, error)
	DeleteCompany(ctx context.Context, id uint64) error
	ListCompanies(ctx context.Context, query ListCompaniesQuery) (*PaginatedCompaniesResponse, error)
	UpdateBranding(ctx context.Context, id uint64, req UpdateBrandingRequest) (*CompanyResponse, error)
	SuspendCompany(ctx context.Context, id uint64) error
	ActivateCompany(ctx context.Context, id uint64) error

	// Company Admin operations
	LoginAdmin(ctx context.Context, req LoginRequest) (*LoginResponse, error)
	GetAdminProfile(ctx context.Context, adminID uint64) (*CompanyAdminResponse, error)
}
