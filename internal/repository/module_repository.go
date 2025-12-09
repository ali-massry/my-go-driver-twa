package repository

import (
	"context"

	"my-go-driver/internal/domain/module"

	"gorm.io/gorm"
)

type moduleRepository struct {
	db *gorm.DB
}

// NewModuleRepository creates a new module repository
func NewModuleRepository(db *gorm.DB) module.Repository {
	return &moduleRepository{db: db}
}

func (r *moduleRepository) ListAllModules(ctx context.Context) ([]module.ModuleMaster, error) {
	var modules []module.ModuleMaster
	err := r.db.WithContext(ctx).Order("category, name").Find(&modules).Error
	return modules, err
}

func (r *moduleRepository) GetModuleByID(ctx context.Context, id uint64) (*module.ModuleMaster, error) {
	var mod module.ModuleMaster
	err := r.db.WithContext(ctx).First(&mod, id).Error
	if err != nil {
		return nil, err
	}
	return &mod, nil
}

func (r *moduleRepository) AssignModule(ctx context.Context, companyModule *module.CompanyModule) error {
	return r.db.WithContext(ctx).Create(companyModule).Error
}

func (r *moduleRepository) GetCompanyModules(ctx context.Context, companyID uint64) ([]module.CompanyModule, error) {
	var modules []module.CompanyModule
	err := r.db.WithContext(ctx).Preload("Module").Where("company_id = ?", companyID).Find(&modules).Error
	return modules, err
}

func (r *moduleRepository) GetCompanyModule(ctx context.Context, companyID, moduleID uint64) (*module.CompanyModule, error) {
	var companyModule module.CompanyModule
	err := r.db.WithContext(ctx).Where("company_id = ? AND module_id = ?", companyID, moduleID).First(&companyModule).Error
	if err != nil {
		return nil, err
	}
	return &companyModule, nil
}

func (r *moduleRepository) UpdateModuleConfig(ctx context.Context, id uint64, config module.ModuleConfig) error {
	return r.db.WithContext(ctx).Model(&module.CompanyModule{}).Where("id = ?", id).Update("config", config).Error
}

func (r *moduleRepository) RemoveModule(ctx context.Context, companyID, moduleID uint64) error {
	return r.db.WithContext(ctx).Where("company_id = ? AND module_id = ?", companyID, moduleID).Delete(&module.CompanyModule{}).Error
}
