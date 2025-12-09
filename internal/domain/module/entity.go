package module

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// ModuleMaster represents a master module definition
type ModuleMaster struct {
	ID             uint64    `json:"id" gorm:"primaryKey"`
	ModuleKey      string    `json:"module_key" gorm:"not null;uniqueIndex"`
	Name           string    `json:"name" gorm:"not null"`
	Category       string    `json:"category"`
	Description    string    `json:"description" gorm:"type:text"`
	DefaultEnabled bool      `json:"default_enabled" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at"`
}

func (ModuleMaster) TableName() string {
	return "modules_master"
}

// ModuleConfig represents the configuration JSON
type ModuleConfig map[string]interface{}

// Scan implements sql.Scanner interface
func (mc *ModuleConfig) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, mc)
}

// Value implements driver.Valuer interface
func (mc ModuleConfig) Value() (driver.Value, error) {
	if mc == nil {
		return nil, nil
	}
	return json.Marshal(mc)
}

// CompanyModule represents a module assigned to a company
type CompanyModule struct {
	ID        uint64       `json:"id" gorm:"primaryKey"`
	CompanyID uint64       `json:"company_id" gorm:"not null"`
	ModuleID  uint64       `json:"module_id" gorm:"not null"`
	IsEnabled bool         `json:"is_enabled" gorm:"default:true"`
	Config    ModuleConfig `json:"config" gorm:"type:json"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`

	// Relations
	Module *ModuleMaster `json:"module,omitempty" gorm:"foreignKey:ModuleID"`
}

func (CompanyModule) TableName() string {
	return "company_modules"
}
