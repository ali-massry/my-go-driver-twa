package module

import "context"

// Service defines the interface for module business logic
type Service interface {
	ListAllModules(ctx context.Context) ([]ModuleResponse, error)
	AssignModuleToCompany(ctx context.Context, companyID uint64, req AssignModuleRequest) (*CompanyModuleResponse, error)
	GetCompanyModules(ctx context.Context, companyID uint64) ([]CompanyModuleResponse, error)
	RemoveModuleFromCompany(ctx context.Context, companyID, moduleID uint64) error
}
