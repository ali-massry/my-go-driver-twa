package handler

import (
	"net/http"
	"strconv"

	"my-go-driver/internal/domain/company"
	"my-go-driver/pkg/httputil"

	"github.com/gin-gonic/gin"
)

type AdminCompanyHandler struct {
	companyService company.Service
}

func NewAdminCompanyHandler(companyService company.Service) *AdminCompanyHandler {
	return &AdminCompanyHandler{
		companyService: companyService,
	}
}

// CreateCompany creates a new company with owner
// @Summary Create company
// @Tags Admin - Companies
// @Accept json
// @Produce json
// @Param request body company.CreateCompanyRequest true "Company creation request"
// @Success 201 {object} company.CompanyWithOwnerResponse
// @Router /api/v1/admin/companies [post]
func (h *AdminCompanyHandler) CreateCompany(c *gin.Context) {
	var req company.CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	result, err := h.companyService.CreateCompany(c.Request.Context(), req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to create company", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusCreated, "Company created successfully", result)
}

// GetCompany retrieves a company by ID
// @Summary Get company
// @Tags Admin - Companies
// @Produce json
// @Param id path int true "Company ID"
// @Success 200 {object} company.CompanyResponse
// @Router /api/v1/admin/companies/{id} [get]
func (h *AdminCompanyHandler) GetCompany(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid company ID", err.Error())
		return
	}

	result, err := h.companyService.GetCompany(c.Request.Context(), id)
	if err != nil {
		httputil.RespondError(c, http.StatusNotFound, "Company not found", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Company retrieved successfully", result)
}

// UpdateCompany updates a company
// @Summary Update company
// @Tags Admin - Companies
// @Accept json
// @Produce json
// @Param id path int true "Company ID"
// @Param request body company.UpdateCompanyRequest true "Company update request"
// @Success 200 {object} company.CompanyResponse
// @Router /api/v1/admin/companies/{id} [put]
func (h *AdminCompanyHandler) UpdateCompany(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid company ID", err.Error())
		return
	}

	var req company.UpdateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	result, err := h.companyService.UpdateCompany(c.Request.Context(), id, req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to update company", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Company updated successfully", result)
}

// DeleteCompany deletes a company
// @Summary Delete company
// @Tags Admin - Companies
// @Param id path int true "Company ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/admin/companies/{id} [delete]
func (h *AdminCompanyHandler) DeleteCompany(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid company ID", err.Error())
		return
	}

	if err := h.companyService.DeleteCompany(c.Request.Context(), id); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to delete company", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Company deleted successfully", nil)
}

// ListCompanies lists all companies with pagination
// @Summary List companies
// @Tags Admin - Companies
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param status query string false "Company status"
// @Param search query string false "Search term"
// @Success 200 {object} company.PaginatedCompaniesResponse
// @Router /api/v1/admin/companies [get]
func (h *AdminCompanyHandler) ListCompanies(c *gin.Context) {
	var query company.ListCompaniesQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	result, err := h.companyService.ListCompanies(c.Request.Context(), query)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to list companies", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Companies retrieved successfully", result)
}

// UpdateBranding updates company branding
// @Summary Update company branding
// @Tags Admin - Companies
// @Accept json
// @Produce json
// @Param id path int true "Company ID"
// @Param request body company.UpdateBrandingRequest true "Branding update request"
// @Success 200 {object} company.CompanyResponse
// @Router /api/v1/admin/companies/{id}/branding [put]
func (h *AdminCompanyHandler) UpdateBranding(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid company ID", err.Error())
		return
	}

	var req company.UpdateBrandingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	result, err := h.companyService.UpdateBranding(c.Request.Context(), id, req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to update branding", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Branding updated successfully", result)
}

// SuspendCompany suspends a company
// @Summary Suspend company
// @Tags Admin - Companies
// @Param id path int true "Company ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/admin/companies/{id}/suspend [put]
func (h *AdminCompanyHandler) SuspendCompany(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid company ID", err.Error())
		return
	}

	if err := h.companyService.SuspendCompany(c.Request.Context(), id); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to suspend company", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Company suspended successfully", nil)
}

// ActivateCompany activates a company
// @Summary Activate company
// @Tags Admin - Companies
// @Param id path int true "Company ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/admin/companies/{id}/activate [put]
func (h *AdminCompanyHandler) ActivateCompany(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid company ID", err.Error())
		return
	}

	if err := h.companyService.ActivateCompany(c.Request.Context(), id); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to activate company", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Company activated successfully", nil)
}

// LoginAdmin handles company admin login
// @Summary Company admin login
// @Tags Admin - Auth
// @Accept json
// @Produce json
// @Param request body company.LoginRequest true "Login credentials"
// @Success 200 {object} company.LoginResponse
// @Router /api/v1/admin/auth/login [post]
func (h *AdminCompanyHandler) LoginAdmin(c *gin.Context) {
	var req company.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	result, err := h.companyService.LoginAdmin(c.Request.Context(), req)
	if err != nil {
		httputil.RespondError(c, http.StatusUnauthorized, "Login failed", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Login successful", result)
}

// GetAdminProfile retrieves admin profile
// @Summary Get admin profile
// @Tags Admin - Auth
// @Produce json
// @Success 200 {object} company.CompanyAdminResponse
// @Router /api/v1/admin/auth/me [get]
func (h *AdminCompanyHandler) GetAdminProfile(c *gin.Context) {
	adminID, exists := c.Get("user_id")
	if !exists {
		httputil.RespondError(c, http.StatusUnauthorized, "Unauthorized", "User ID not found in context")
		return
	}

	result, err := h.companyService.GetAdminProfile(c.Request.Context(), adminID.(uint64))
	if err != nil {
		httputil.RespondError(c, http.StatusNotFound, "Admin not found", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Profile retrieved successfully", result)
}
