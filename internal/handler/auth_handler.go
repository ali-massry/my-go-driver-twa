package handler

import (
	"net/http"

	"github.com/ali-massry/my-go-driver/internal/domain/user"
	"github.com/ali-massry/my-go-driver/internal/middleware"
	"github.com/ali-massry/my-go-driver/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// AuthHandler handles HTTP requests for authentication operations
type AuthHandler struct {
	service user.Service
}

// NewAuthHandler creates a new authentication handler
func NewAuthHandler(service user.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

// Register handles POST /auth/register - creates a new user account
func (h *AuthHandler) Register(c *gin.Context) {
	var req user.RegisterRequest

	// Validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := httputil.FormatValidationErrors(err)
		httputil.RespondError(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	// Register user
	authResp, err := h.service.Register(req)
	if err != nil {
		if err.Error() == "email already exists" {
			httputil.RespondError(c, http.StatusConflict, "Email already exists", nil)
			return
		}
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to register user", nil)
		return
	}

	httputil.RespondSuccess(c, http.StatusCreated, "User registered successfully", authResp)
}

// Login handles POST /auth/login - authenticates a user
func (h *AuthHandler) Login(c *gin.Context) {
	var req user.LoginRequest

	// Validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := httputil.FormatValidationErrors(err)
		httputil.RespondError(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	// Authenticate user
	authResp, err := h.service.Login(req)
	if err != nil {
		if err.Error() == "invalid credentials" {
			httputil.RespondError(c, http.StatusUnauthorized, "Invalid credentials", nil)
			return
		}
		httputil.RespondError(c, http.StatusInternalServerError, "Failed to login", nil)
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "Login successful", authResp)
}

// Me handles GET /auth/me - retrieves the authenticated user's profile
func (h *AuthHandler) Me(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		httputil.RespondError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	usr, err := h.service.GetUserByID(userID)
	if err != nil {
		httputil.RespondError(c, http.StatusNotFound, "User not found", nil)
		return
	}

	httputil.RespondSuccess(c, http.StatusOK, "User profile retrieved successfully", usr.ToUserResponse())
}
