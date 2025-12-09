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

// Company represents a company entity
type Company struct {
	ID           uint64        `json:"id" gorm:"primaryKey"`
	Name         string        `json:"name" gorm:"not null"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	Address      string        `json:"address"`
	Timezone     string        `json:"timezone" gorm:"default:UTC"`
	LogoURL      string        `json:"logo_url"`
	ColorPalette *ColorPalette `json:"color_palette" gorm:"type:json"`
	FontFamily   string        `json:"font_family"`
	Status       CompanyStatus `json:"status" gorm:"type:enum('active','suspended');default:active"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
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
