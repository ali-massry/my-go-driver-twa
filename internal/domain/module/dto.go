package module

import "time"

// ModuleResponse represents module response
type ModuleResponse struct {
	ID             uint64    `json:"id"`
	ModuleKey      string    `json:"module_key"`
	Name           string    `json:"name"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	DefaultEnabled bool      `json:"default_enabled"`
	CreatedAt      time.Time `json:"created_at"`
}

// CompanyModuleResponse represents company module response
type CompanyModuleResponse struct {
	ID        uint64         `json:"id"`
	CompanyID uint64         `json:"company_id"`
	Module    ModuleResponse `json:"module"`
	IsEnabled bool           `json:"is_enabled"`
	Config    ModuleConfig   `json:"config"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// AssignModuleRequest represents request to assign module
type AssignModuleRequest struct {
	ModuleID  uint64       `json:"module_id" binding:"required"`
	IsEnabled bool         `json:"is_enabled"`
	Config    ModuleConfig `json:"config"`
}
