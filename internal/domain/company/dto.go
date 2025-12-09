package company

import "time"

// CreateCompanyRequest represents request to create a new company
type CreateCompanyRequest struct {
	// Company details
	Name     string `json:"name" binding:"required,min=2,max=255"`
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone" binding:"omitempty,max=50"`
	Address  string `json:"address" binding:"omitempty"`
	Timezone string `json:"timezone" binding:"omitempty"`

	// Owner details
	OwnerName     string `json:"owner_name" binding:"required,min=2,max=255"`
	OwnerEmail    string `json:"owner_email" binding:"required,email"`
	OwnerPhone    string `json:"owner_phone" binding:"omitempty,max=50"`
	OwnerPassword string `json:"owner_password" binding:"required,min=8"`
}

// UpdateCompanyRequest represents request to update company
type UpdateCompanyRequest struct {
	Name     string `json:"name" binding:"omitempty,min=2,max=255"`
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone" binding:"omitempty,max=50"`
	Address  string `json:"address" binding:"omitempty"`
	Timezone string `json:"timezone" binding:"omitempty"`
}

// UpdateBrandingRequest represents request to update company branding
type UpdateBrandingRequest struct {
	LogoURL      string        `json:"logo_url" binding:"omitempty,url"`
	ColorPalette *ColorPalette `json:"color_palette" binding:"omitempty"`
	FontFamily   string        `json:"font_family" binding:"omitempty,max=100"`
}

// AssignModuleRequest represents request to assign module to company
type AssignModuleRequest struct {
	ModuleID  uint64                 `json:"module_id" binding:"required"`
	IsEnabled bool                   `json:"is_enabled"`
	Config    map[string]interface{} `json:"config"`
}

// CompanyResponse represents company response
type CompanyResponse struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	Address      string        `json:"address"`
	Timezone     string        `json:"timezone"`
	LogoURL      string        `json:"logo_url"`
	ColorPalette *ColorPalette `json:"color_palette"`
	FontFamily   string        `json:"font_family"`
	Status       CompanyStatus `json:"status"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

// CompanyWithOwnerResponse includes the owner information
type CompanyWithOwnerResponse struct {
	Company CompanyResponse     `json:"company"`
	Owner   CompanyAdminResponse `json:"owner"`
}

// CompanyAdminResponse represents company admin response
type CompanyAdminResponse struct {
	ID        uint64    `json:"id"`
	CompanyID uint64    `json:"company_id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      AdminRole `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ListCompaniesQuery represents query parameters for listing companies
type ListCompaniesQuery struct {
	Page   int           `form:"page" binding:"omitempty,min=1"`
	Limit  int           `form:"limit" binding:"omitempty,min=1,max=100"`
	Status CompanyStatus `form:"status" binding:"omitempty,oneof=active suspended"`
	Search string        `form:"search" binding:"omitempty"`
}

// PaginatedCompaniesResponse represents paginated companies response
type PaginatedCompaniesResponse struct {
	Companies  []CompanyResponse `json:"companies"`
	TotalCount int64             `json:"total_count"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
	TotalPages int               `json:"total_pages"`
}

// LoginRequest represents company admin login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents login response with token
type LoginResponse struct {
	Admin CompanyAdminResponse `json:"admin"`
	Token string               `json:"token"`
}
