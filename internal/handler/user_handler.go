package handler

import (
	"errors"
	"net/http"

	"github.com/ali-massry/my-go-driver/internal/domain/user"
	"github.com/ali-massry/my-go-driver/pkg/httputil"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	service user.Service
}

// NewUserHandler creates a new user handler
func NewUserHandler(service user.Service) *UserHandler {
	return &UserHandler{service: service}
}

// GetUsers handles GET /users - retrieves all users
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to fetch users", nil)
		return
	}

	// Convert to response DTOs
	userResponses := user.ToUserResponseList(users)
	httputil.RespondSuccess(c, http.StatusOK, "Users retrieved successfully", gin.H{"users": userResponses})
}

// CreateUser handles POST /users - creates a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req user.CreateUserRequest

	// Validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := httputil.FormatValidationErrors(err)
		httputil.RespondError(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	// Create user
	usr, err := h.service.CreateUser(req)
	if err != nil {
		if err.Error() == "email already exists" {
			httputil.RespondError(c, http.StatusConflict, "Email already exists", nil)
			return
		}
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to create user", nil)
		return
	}

	// Return response
	httputil.RespondSuccess(c, http.StatusCreated, "User created successfully", usr.ToUserResponse())
}

// GetUserByID handles GET /users/:id - retrieves a user by ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	var idParam struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&idParam); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	usr, err := h.service.GetUserByID(idParam.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httputil.RespondError(c, http.StatusNotFound, "User not found", nil)
			return
		}
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to fetch user", nil)
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "User retrieved successfully", usr.ToUserResponse())
}

// UpdateUser handles PUT /users/:id - updates a user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var idParam struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&idParam); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	var req user.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := httputil.FormatValidationErrors(err)
		httputil.RespondError(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	usr, err := h.service.UpdateUser(idParam.ID, req)
	if err != nil {
		if err.Error() == "user not found" {
			httputil.RespondError(c, http.StatusNotFound, "User not found", nil)
			return
		}
		if err.Error() == "email already exists" {
			httputil.RespondError(c, http.StatusConflict, "Email already exists", nil)
			return
		}
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to update user", nil)
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "User updated successfully", usr.ToUserResponse())
}

// DeleteUser handles DELETE /users/:id - deletes a user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	var idParam struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&idParam); err != nil {
		httputil.RespondError(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	if err := h.service.DeleteUser(idParam.ID); err != nil {
		if err.Error() == "user not found" {
			httputil.RespondError(c, http.StatusNotFound, "User not found", nil)
			return
		}
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to delete user", nil)
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "User deleted successfully", nil)
}
