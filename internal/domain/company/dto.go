package company

import "time"

// CreateCompanyRequest represents request to create a new company
type CreateCompanyRequest struct {
	// Basic Company Info
	Name      string `json:"name" binding:"required,min=2,max=255"`
	LegalName string `json:"legal_name" binding:"omitempty"`
	Email     string `json:"email" binding:"omitempty,email"`
	Phone     string `json:"phone" binding:"omitempty,max=50"`
	WhatsApp  string `json:"whatsapp" binding:"omitempty,max=50"`
	Address   string `json:"address" binding:"omitempty"`
	Country   string `json:"country" binding:"omitempty"`

	// Localization
	Timezone   string `json:"timezone" binding:"omitempty"`
	Locale     string `json:"locale" binding:"omitempty"`
	DateFormat string `json:"date_format" binding:"omitempty"`
	Currency   string `json:"currency" binding:"omitempty"`

	// Billing & Subscription
	Plan               BillingPlan   `json:"plan" binding:"omitempty,oneof=free basic pro enterprise"`
	BillingCycle       BillingCycle  `json:"billing_cycle" binding:"omitempty,oneof=monthly yearly"`
	MaxAllowedDrivers  int           `json:"max_allowed_drivers" binding:"omitempty,min=1"`

	// Owner details
	OwnerName     string `json:"owner_name" binding:"required,min=2,max=255"`
	OwnerEmail    string `json:"owner_email" binding:"required,email"`
	OwnerPhone    string `json:"owner_phone" binding:"omitempty,max=50"`
	OwnerPassword string `json:"owner_password" binding:"required,min=8"`
}

// UpdateCompanyRequest represents request to update company
type UpdateCompanyRequest struct {
	// Basic Company Info
	Name      string `json:"name" binding:"omitempty,min=2,max=255"`
	LegalName string `json:"legal_name" binding:"omitempty"`
	Email     string `json:"email" binding:"omitempty,email"`
	Phone     string `json:"phone" binding:"omitempty,max=50"`
	WhatsApp  string `json:"whatsapp" binding:"omitempty,max=50"`
	Address   string `json:"address" binding:"omitempty"`
	Country   string `json:"country" binding:"omitempty"`

	// Localization
	Timezone   string `json:"timezone" binding:"omitempty"`
	Locale     string `json:"locale" binding:"omitempty"`
	DateFormat string `json:"date_format" binding:"omitempty"`
	Currency   string `json:"currency" binding:"omitempty"`

	// Business Rules
	PODRequired           *bool                  `json:"pod_required" binding:"omitempty"`
	VehicleAssignmentMode VehicleAssignmentMode  `json:"vehicle_assignment_mode" binding:"omitempty,oneof=auto manual"`
	MaxExtraDeliveryQty   *int                   `json:"max_extra_delivery_qty" binding:"omitempty,min=0"`

	// Logistics Settings
	RoutingMode       RoutingMode `json:"routing_mode" binding:"omitempty,oneof=simple optimized AI"`
	GPSAccuracy       GPSAccuracy `json:"gps_accuracy" binding:"omitempty,oneof=low medium high"`
	HasMultipleStores *bool       `json:"has_multiple_stores" binding:"omitempty"`

	// Inventory Settings
	EnableVehicleStock   *bool `json:"enable_vehicle_stock" binding:"omitempty"`
	EnableProductCatalog *bool `json:"enable_product_catalog" binding:"omitempty"`

	// Notifications
	BroadcastEnabled *bool `json:"broadcast_enabled" binding:"omitempty"`

	// Driver Limit
	MaxAllowedDrivers *int `json:"max_allowed_drivers" binding:"omitempty,min=1"`

	// Billing & Subscription
	Plan         BillingPlan   `json:"plan" binding:"omitempty,oneof=free basic pro enterprise"`
	BillingCycle BillingCycle  `json:"billing_cycle" binding:"omitempty,oneof=monthly yearly"`
	SeatsLimit   *int          `json:"seats_limit" binding:"omitempty,min=1"`
	APIRateLimit *int          `json:"api_rate_limit" binding:"omitempty,min=0"`
}

// UpdateBrandingRequest represents request to update company branding
type UpdateBrandingRequest struct {
	LogoURL      string        `json:"logo_url" binding:"omitempty,url"`
	ColorPalette *ColorPalette `json:"color_palette" binding:"omitempty"`
	FontFamily   string        `json:"font_family" binding:"omitempty,max=100"`
	Theme        Theme         `json:"theme" binding:"omitempty,oneof=light dark custom"`
	CustomCSS    string        `json:"custom_css" binding:"omitempty"`
}

// AssignModuleRequest represents request to assign module to company
type AssignModuleRequest struct {
	ModuleID  uint64                 `json:"module_id" binding:"required"`
	IsEnabled bool                   `json:"is_enabled"`
	Config    map[string]interface{} `json:"config"`
}

// CompanyResponse represents company response
type CompanyResponse struct {
	// Basic Company Info
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	LegalName string `json:"legal_name,omitempty"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	WhatsApp  string `json:"whatsapp,omitempty"`
	Address   string `json:"address"`
	Country   string `json:"country,omitempty"`

	// Branding & UI
	LogoURL      string        `json:"logo_url"`
	ColorPalette *ColorPalette `json:"color_palette"`
	FontFamily   string        `json:"font_family"`
	Theme        Theme         `json:"theme"`
	CustomCSS    string        `json:"custom_css,omitempty"`

	// Localization
	Timezone   string `json:"timezone"`
	Locale     string `json:"locale"`
	DateFormat string `json:"date_format"`
	Currency   string `json:"currency"`

	// Business Rules
	PODRequired           bool                  `json:"pod_required"`
	VehicleAssignmentMode VehicleAssignmentMode `json:"vehicle_assignment_mode"`
	MaxExtraDeliveryQty   int                   `json:"max_extra_delivery_qty"`

	// Logistics Settings
	RoutingMode       RoutingMode `json:"routing_mode"`
	GPSAccuracy       GPSAccuracy `json:"gps_accuracy"`
	DepotID           *uint64     `json:"depot_id,omitempty"`
	HasMultipleStores bool        `json:"has_multiple_stores"`

	// Inventory Settings
	EnableVehicleStock   bool `json:"enable_vehicle_stock"`
	EnableProductCatalog bool `json:"enable_product_catalog"`

	// Notifications
	BroadcastEnabled bool `json:"broadcast_enabled"`

	// Driver Limit
	MaxAllowedDrivers int `json:"max_allowed_drivers"`

	// Billing & Subscription
	Plan         BillingPlan   `json:"plan"`
	BillingCycle BillingCycle  `json:"billing_cycle"`
	SeatsLimit   int           `json:"seats_limit"`
	APIRateLimit int           `json:"api_rate_limit"`

	// Status & Audit
	Status    CompanyStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
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
