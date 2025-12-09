package service

import (
	"context"
	"errors"
	"fmt"

	"my-go-driver/internal/domain/module"

	"gorm.io/gorm"
)

type moduleService struct {
	repo module.Repository
}

// NewModuleService creates a new module service
func NewModuleService(repo module.Repository) module.Service {
	return &moduleService{repo: repo}
}

func (s *moduleService) ListAllModules(ctx context.Context) ([]module.ModuleResponse, error) {
	modules, err := s.repo.ListAllModules(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]module.ModuleResponse, len(modules))
	for i, m := range modules {
		responses[i] = s.toModuleResponse(&m)
	}

	return responses, nil
}

func (s *moduleService) AssignModuleToCompany(ctx context.Context, companyID uint64, req module.AssignModuleRequest) (*module.CompanyModuleResponse, error) {
	// Check if module exists
	mod, err := s.repo.GetModuleByID(ctx, req.ModuleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("module not found")
		}
		return nil, err
	}

	// Check if already assigned
	existing, err := s.repo.GetCompanyModule(ctx, companyID, req.ModuleID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existing != nil {
		return nil, fmt.Errorf("module already assigned to company")
	}

	// Create company module
	companyModule := &module.CompanyModule{
		CompanyID: companyID,
		ModuleID:  req.ModuleID,
		IsEnabled: req.IsEnabled,
		Config:    req.Config,
	}

	if err := s.repo.AssignModule(ctx, companyModule); err != nil {
		return nil, fmt.Errorf("failed to assign module: %w", err)
	}

	// Reload with module details
	companyModule.Module = mod

	return s.toCompanyModuleResponse(companyModule), nil
}

func (s *moduleService) GetCompanyModules(ctx context.Context, companyID uint64) ([]module.CompanyModuleResponse, error) {
	modules, err := s.repo.GetCompanyModules(ctx, companyID)
	if err != nil {
		return nil, err
	}

	responses := make([]module.CompanyModuleResponse, len(modules))
	for i, m := range modules {
		responses[i] = *s.toCompanyModuleResponse(&m)
	}

	return responses, nil
}

func (s *moduleService) RemoveModuleFromCompany(ctx context.Context, companyID, moduleID uint64) error {
	// Check if assigned
	_, err := s.repo.GetCompanyModule(ctx, companyID, moduleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("module not assigned to company")
		}
		return err
	}

	return s.repo.RemoveModule(ctx, companyID, moduleID)
}

// Helper methods
func (s *moduleService) toModuleResponse(m *module.ModuleMaster) module.ModuleResponse {
	return module.ModuleResponse{
		ID:             m.ID,
		ModuleKey:      m.ModuleKey,
		Name:           m.Name,
		Category:       m.Category,
		Description:    m.Description,
		DefaultEnabled: m.DefaultEnabled,
		CreatedAt:      m.CreatedAt,
	}
}

func (s *moduleService) toCompanyModuleResponse(cm *module.CompanyModule) *module.CompanyModuleResponse {
	response := &module.CompanyModuleResponse{
		ID:        cm.ID,
		CompanyID: cm.CompanyID,
		IsEnabled: cm.IsEnabled,
		Config:    cm.Config,
		CreatedAt: cm.CreatedAt,
		UpdatedAt: cm.UpdatedAt,
	}

	if cm.Module != nil {
		response.Module = s.toModuleResponse(cm.Module)
	}

	return response
}
