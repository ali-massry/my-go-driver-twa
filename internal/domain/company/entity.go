package company

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type CompanyStatus string

const (
	CompanyStatusActive    CompanyStatus = "active"
	CompanyStatusSuspended CompanyStatus = "suspended"
)

type AdminRole string

const (
	AdminRoleOwner   AdminRole = "owner"
	AdminRoleManager AdminRole = "manager"
)

// Additional enums for new fields
type Theme string

const (
	ThemeLight  Theme = "light"
	ThemeDark   Theme = "dark"
	ThemeCustom Theme = "custom"
)

type VehicleAssignmentMode string

const (
	VehicleAssignmentAuto   VehicleAssignmentMode = "auto"
	VehicleAssignmentManual VehicleAssignmentMode = "manual"
)

type RoutingMode string

const (
	RoutingSimple    RoutingMode = "simple"
	RoutingOptimized RoutingMode = "optimized"
	RoutingAI        RoutingMode = "AI"
)

type GPSAccuracy string

const (
	GPSLow    GPSAccuracy = "low"
	GPSMedium GPSAccuracy = "medium"
	GPSHigh   GPSAccuracy = "high"
)

type BillingPlan string

const (
	PlanFree       BillingPlan = "free"
	PlanBasic      BillingPlan = "basic"
	PlanPro        BillingPlan = "pro"
	PlanEnterprise BillingPlan = "enterprise"
)

type BillingCycle string

const (
	CycleMonthly BillingCycle = "monthly"
	CycleYearly  BillingCycle = "yearly"
)

// ColorPalette represents the company's branding colors
type ColorPalette struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
	Accent    string `json:"accent,omitempty"`
}

// Scan implements sql.Scanner interface
func (cp *ColorPalette) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, cp)
}

// Value implements driver.Valuer interface
func (cp ColorPalette) Value() (driver.Value, error) {
	return json.Marshal(cp)
}

// JSONMap is a generic JSON map type
type JSONMap map[string]interface{}

// Scan implements sql.Scanner interface
func (jm *JSONMap) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, jm)
}

// Value implements driver.Valuer interface
func (jm JSONMap) Value() (driver.Value, error) {
	if jm == nil {
		return nil, nil
	}
	return json.Marshal(jm)
}

// JSONArray is a generic JSON array type
type JSONArray []interface{}

// Scan implements sql.Scanner interface
func (ja *JSONArray) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, ja)
}

// Value implements driver.Valuer interface
func (ja JSONArray) Value() (driver.Value, error) {
	if ja == nil {
		return nil, nil
	}
	return json.Marshal(ja)
}

// Company represents a company entity
type Company struct {
	// Basic Company Info
	ID        uint64 `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	LegalName string `json:"legal_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	WhatsApp  string `json:"whatsapp"`
	Address   string `json:"address"`
	Country   string `json:"country"`

	// Branding & UI
	LogoURL      string        `json:"logo_url"`
	ColorPalette *ColorPalette `json:"color_palette" gorm:"type:json"`
	FontFamily   string        `json:"font_family"`
	Theme        Theme         `json:"theme" gorm:"type:enum('light','dark','custom');default:light"`
	CustomCSS    string        `json:"custom_css" gorm:"type:text"`

	// Localization
	Timezone   string `json:"timezone" gorm:"default:UTC"`
	Locale     string `json:"locale" gorm:"default:en"`
	DateFormat string `json:"date_format" gorm:"default:dd/mm/yyyy"`
	Currency   string `json:"currency" gorm:"default:USD"`

	// Business Rules (JSON fields)
	AutoAssignRules       JSONMap               `json:"auto_assign_rules" gorm:"type:json"`
	DeliveryPricingRules  JSONMap               `json:"delivery_pricing_rules" gorm:"type:json"`
	CashHandlingRules     JSONMap               `json:"cash_handling_rules" gorm:"type:json"`
	PODRequired           bool                  `json:"pod_required" gorm:"default:false"`
	DriverShiftRules      JSONMap               `json:"driver_shift_rules" gorm:"type:json"`
	VehicleAssignmentMode VehicleAssignmentMode `json:"vehicle_assignment_mode" gorm:"type:enum('auto','manual');default:manual"`
	MaxExtraDeliveryQty   int                   `json:"max_extra_delivery_qty" gorm:"default:0"`

	// Logistics Settings
	RoutingMode      RoutingMode `json:"routing_mode" gorm:"type:enum('simple','optimized','AI');default:simple"`
	GPSAccuracy      GPSAccuracy `json:"gps_accuracy" gorm:"type:enum('low','medium','high');default:medium"`
	DepotID          *uint64     `json:"depot_id"`
	HasMultipleStores bool       `json:"has_multiple_stores" gorm:"default:false"`

	// Inventory Settings
	EnableVehicleStock    bool      `json:"enable_vehicle_stock" gorm:"default:false"`
	EnableProductCatalog  bool      `json:"enable_product_catalog" gorm:"default:false"`
	AllowedProducts       JSONArray `json:"allowed_products" gorm:"type:json"`

	// Notifications
	NotificationSettings JSONMap `json:"notification_settings" gorm:"type:json"`
	BroadcastEnabled     bool    `json:"broadcast_enabled" gorm:"default:true"`

	// Driver Limit
	MaxAllowedDrivers int `json:"max_allowed_drivers" gorm:"default:10"`

	// Billing & Subscription
	Plan         BillingPlan  `json:"plan" gorm:"type:enum('free','basic','pro','enterprise');default:free"`
	BillingCycle BillingCycle `json:"billing_cycle" gorm:"type:enum('monthly','yearly');default:monthly"`
	SeatsLimit   int          `json:"seats_limit" gorm:"default:10"`
	APIRateLimit int          `json:"api_rate_limit" gorm:"default:1000"`

	// Status & Audit
	Status    CompanyStatus `json:"status" gorm:"type:enum('active','suspended');default:active"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at,omitempty"`
}

func (Company) TableName() string {
	return "companies"
}

// CompanyAdmin represents a company administrator
type CompanyAdmin struct {
	ID           uint64    `json:"id" gorm:"primaryKey"`
	CompanyID    uint64    `json:"company_id" gorm:"not null"`
	FullName     string    `json:"full_name" gorm:"not null"`
	Email        string    `json:"email" gorm:"not null;uniqueIndex"`
	Phone        string    `json:"phone"`
	PasswordHash string    `json:"-" gorm:"not null"`
	Role         AdminRole `json:"role" gorm:"type:enum('owner','manager');default:manager"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Relations
	Company *Company `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
}

func (CompanyAdmin) TableName() string {
	return "company_admins"
}
