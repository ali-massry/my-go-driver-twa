package module

import "context"

// Repository defines the interface for module data access
type Repository interface {
	// Module master operations
	ListAllModules(ctx context.Context) ([]ModuleMaster, error)
	GetModuleByID(ctx context.Context, id uint64) (*ModuleMaster, error)

	// Company module operations
	AssignModule(ctx context.Context, companyModule *CompanyModule) error
	GetCompanyModules(ctx context.Context, companyID uint64) ([]CompanyModule, error)
	GetCompanyModule(ctx context.Context, companyID, moduleID uint64) (*CompanyModule, error)
	UpdateModuleConfig(ctx context.Context, id uint64, config ModuleConfig) error
	RemoveModule(ctx context.Context, companyID, moduleID uint64) error
}
