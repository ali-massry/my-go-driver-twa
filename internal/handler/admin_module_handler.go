package handler

import (
	"net/http"
	"strconv"

	"my-go-driver/internal/domain/module"
	"my-go-driver/pkg/httputil"

	"github.com/gin-gonic/gin"
)

type AdminModuleHandler struct {
	moduleService module.Service
}

func NewAdminModuleHandler(moduleService module.Service) *AdminModuleHandler {
	return &AdminModuleHandler{
		moduleService: moduleService,
	}
}

// ListAllModules lists all available modules
// @Summary List all modules
// @Tags Admin - Modules
// @Produce json
// @Success 200 {array} module.ModuleResponse
// @Router /api/v1/admin/modules [get]
func (h *AdminModuleHandler) ListAllModules(c *gin.Context) {
	result, err := h.moduleService.ListAllModules(c.Request.Context())
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to list modules", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Modules retrieved successfully", result)
}

// AssignModuleToCompany assigns a module to a company
// @Summary Assign module to company
// @Tags Admin - Modules
// @Accept json
// @Produce json
// @Param id path int true "Company ID"
// @Param request body module.AssignModuleRequest true "Module assignment request"
// @Success 201 {object} module.CompanyModuleResponse
// @Router /api/v1/admin/companies/{id}/modules [post]
func (h *AdminModuleHandler) AssignModuleToCompany(c *gin.Context) {
	companyID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid company ID", err.Error())
		return
	}

	var req module.AssignModuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	result, err := h.moduleService.AssignModuleToCompany(c.Request.Context(), companyID, req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to assign module", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusCreated, "Module assigned successfully", result)
}

// GetCompanyModules gets all modules assigned to a company
// @Summary Get company modules
// @Tags Admin - Modules
// @Produce json
// @Param id path int true "Company ID"
// @Success 200 {array} module.CompanyModuleResponse
// @Router /api/v1/admin/companies/{id}/modules [get]
func (h *AdminModuleHandler) GetCompanyModules(c *gin.Context) {
	companyID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid company ID", err.Error())
		return
	}

	result, err := h.moduleService.GetCompanyModules(c.Request.Context(), companyID)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to get company modules", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Company modules retrieved successfully", result)
}

// RemoveModuleFromCompany removes a module from a company
// @Summary Remove module from company
// @Tags Admin - Modules
// @Param id path int true "Company ID"
// @Param module_id path int true "Module ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/admin/companies/{id}/modules/{module_id} [delete]
func (h *AdminModuleHandler) RemoveModuleFromCompany(c *gin.Context) {
	companyID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid company ID", err.Error())
		return
	}

	moduleID, err := strconv.ParseUint(c.Param("module_id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid module ID", err.Error())
		return
	}

	if err := h.moduleService.RemoveModuleFromCompany(c.Request.Context(), companyID, moduleID); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to remove module", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Module removed successfully", nil)
}
