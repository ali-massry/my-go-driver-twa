package handler

import (
	"net/http"
	"strconv"

	"my-go-driver/internal/domain/driver"
	"my-go-driver/internal/domain/shift"
	"my-go-driver/pkg/httputil"

	"github.com/gin-gonic/gin"
)

type AdminDriverHandler struct {
	driverService driver.Service
	shiftService  shift.Service
}

func NewAdminDriverHandler(driverService driver.Service, shiftService shift.Service) *AdminDriverHandler {
	return &AdminDriverHandler{
		driverService: driverService,
		shiftService:  shiftService,
	}
}

// CreateDriver creates a new driver
// @Summary Create driver
// @Tags Admin - Drivers
// @Accept json
// @Produce json
// @Param request body driver.CreateDriverRequest true "Driver creation request"
// @Success 201 {object} driver.DriverResponse
// @Router /api/v1/admin/drivers [post]
func (h *AdminDriverHandler) CreateDriver(c *gin.Context) {
	var req driver.CreateDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	result, err := h.driverService.CreateDriver(c.Request.Context(), req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to create driver", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusCreated, "Driver created successfully", result)
}

// GetDriver retrieves a driver by ID
// @Summary Get driver
// @Tags Admin - Drivers
// @Produce json
// @Param id path int true "Driver ID"
// @Success 200 {object} driver.DriverResponse
// @Router /api/v1/admin/drivers/{id} [get]
func (h *AdminDriverHandler) GetDriver(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid driver ID", err.Error())
		return
	}

	result, err := h.driverService.GetDriver(c.Request.Context(), id)
	if err != nil {
		httputil.RespondError(c, http.StatusNotFound, "Driver not found", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Driver retrieved successfully", result)
}

// UpdateDriver updates a driver
// @Summary Update driver
// @Tags Admin - Drivers
// @Accept json
// @Produce json
// @Param id path int true "Driver ID"
// @Param request body driver.UpdateDriverRequest true "Driver update request"
// @Success 200 {object} driver.DriverResponse
// @Router /api/v1/admin/drivers/{id} [put]
func (h *AdminDriverHandler) UpdateDriver(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid driver ID", err.Error())
		return
	}

	var req driver.UpdateDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	result, err := h.driverService.UpdateDriver(c.Request.Context(), id, req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to update driver", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Driver updated successfully", result)
}

// DeleteDriver deletes a driver
// @Summary Delete driver
// @Tags Admin - Drivers
// @Param id path int true "Driver ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/admin/drivers/{id} [delete]
func (h *AdminDriverHandler) DeleteDriver(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid driver ID", err.Error())
		return
	}

	if err := h.driverService.DeleteDriver(c.Request.Context(), id); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to delete driver", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Driver deleted successfully", nil)
}

// ListDrivers lists all drivers with pagination
// @Summary List drivers
// @Tags Admin - Drivers
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param company_id query int false "Company ID"
// @Param store_id query int false "Store ID"
// @Param status query string false "Driver status"
// @Param online_status query string false "Online status"
// @Param search query string false "Search term"
// @Success 200 {object} driver.PaginatedDriversResponse
// @Router /api/v1/admin/drivers [get]
func (h *AdminDriverHandler) ListDrivers(c *gin.Context) {
	var query driver.ListDriversQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	result, err := h.driverService.ListDrivers(c.Request.Context(), query)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to list drivers", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Drivers retrieved successfully", result)
}

// AssignDriverToCompany assigns driver to a company
// @Summary Assign driver to company
// @Tags Admin - Drivers
// @Accept json
// @Produce json
// @Param id path int true "Driver ID"
// @Param request body driver.AssignDriverToCompanyRequest true "Assignment request"
// @Success 200 {object} driver.DriverResponse
// @Router /api/v1/admin/drivers/{id}/assign-company [put]
func (h *AdminDriverHandler) AssignDriverToCompany(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid driver ID", err.Error())
		return
	}

	var req driver.AssignDriverToCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	result, err := h.driverService.AssignToCompany(c.Request.Context(), id, req)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to assign driver", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Driver assigned successfully", result)
}

// BlockDriver blocks a driver
// @Summary Block driver
// @Tags Admin - Drivers
// @Param id path int true "Driver ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/admin/drivers/{id}/block [put]
func (h *AdminDriverHandler) BlockDriver(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid driver ID", err.Error())
		return
	}

	if err := h.driverService.BlockDriver(c.Request.Context(), id); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to block driver", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Driver blocked successfully", nil)
}

// UnblockDriver unblocks a driver
// @Summary Unblock driver
// @Tags Admin - Drivers
// @Param id path int true "Driver ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/admin/drivers/{id}/unblock [put]
func (h *AdminDriverHandler) UnblockDriver(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid driver ID", err.Error())
		return
	}

	if err := h.driverService.UnblockDriver(c.Request.Context(), id); err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to unblock driver", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Driver unblocked successfully", nil)
}

// GetDriverPerformance gets driver performance metrics
// @Summary Get driver performance
// @Tags Admin - Drivers
// @Produce json
// @Param id path int true "Driver ID"
// @Success 200 {object} driver.DriverPerformance
// @Router /api/v1/admin/drivers/{id}/performance [get]
func (h *AdminDriverHandler) GetDriverPerformance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid driver ID", err.Error())
		return
	}

	result, err := h.driverService.GetDriverPerformance(c.Request.Context(), id)
	if err != nil {
		httputil.RespondError(c, http.StatusNotFound, "Failed to get performance", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Performance retrieved successfully", result)
}

// GetDriverShifts gets driver shift history
// @Summary Get driver shifts
// @Tags Admin - Drivers
// @Produce json
// @Param id path int true "Driver ID"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param status query string false "Shift status"
// @Param start_date query string false "Start date (YYYY-MM-DD)"
// @Param end_date query string false "End date (YYYY-MM-DD)"
// @Success 200 {object} shift.PaginatedShiftsResponse
// @Router /api/v1/admin/drivers/{id}/shifts [get]
func (h *AdminDriverHandler) GetDriverShifts(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid driver ID", err.Error())
		return
	}

	var query shift.ListShiftsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	result, err := h.shiftService.GetDriverShifts(c.Request.Context(), id, query)
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to get shifts", err.Error())
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Shifts retrieved successfully", result)
}
