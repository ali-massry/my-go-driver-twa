package driver

import "time"

// CreateDriverRequest represents request to create a new driver
type CreateDriverRequest struct {
	CompanyID    uint64  `json:"company_id" binding:"required"`
	StoreID      *uint64 `json:"store_id"`
	FullName     string  `json:"full_name" binding:"required,min=2,max=255"`
	Phone        string  `json:"phone" binding:"required,max=50"`
	Email        string  `json:"email" binding:"omitempty,email"`
	Password     string  `json:"password" binding:"required,min=8"`
	ProfilePhoto string  `json:"profile_photo" binding:"omitempty,url"`
}

// UpdateDriverRequest represents request to update driver
type UpdateDriverRequest struct {
	FullName     string  `json:"full_name" binding:"omitempty,min=2,max=255"`
	Phone        string  `json:"phone" binding:"omitempty,max=50"`
	Email        string  `json:"email" binding:"omitempty,email"`
	StoreID      *uint64 `json:"store_id"`
	ProfilePhoto string  `json:"profile_photo" binding:"omitempty,url"`
}

// AssignDriverToCompanyRequest represents request to assign driver to company
type AssignDriverToCompanyRequest struct {
	CompanyID uint64  `json:"company_id" binding:"required"`
	StoreID   *uint64 `json:"store_id"`
}

// DriverResponse represents driver response
type DriverResponse struct {
	ID           uint64       `json:"id"`
	CompanyID    uint64       `json:"company_id"`
	StoreID      *uint64      `json:"store_id"`
	FullName     string       `json:"full_name"`
	Phone        string       `json:"phone"`
	Email        string       `json:"email"`
	Status       DriverStatus `json:"status"`
	OnlineStatus OnlineStatus `json:"online_status"`
	Rating       float64      `json:"rating"`
	ProfilePhoto string       `json:"profile_photo"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

// ListDriversQuery represents query parameters for listing drivers
type ListDriversQuery struct {
	Page         int          `form:"page" binding:"omitempty,min=1"`
	Limit        int          `form:"limit" binding:"omitempty,min=1,max=100"`
	CompanyID    uint64       `form:"company_id" binding:"omitempty"`
	StoreID      uint64       `form:"store_id" binding:"omitempty"`
	Status       DriverStatus `form:"status" binding:"omitempty,oneof=active off_duty suspended"`
	OnlineStatus OnlineStatus `form:"online_status" binding:"omitempty,oneof=online offline"`
	Search       string       `form:"search" binding:"omitempty"`
}

// PaginatedDriversResponse represents paginated drivers response
type PaginatedDriversResponse struct {
	Drivers    []DriverResponse `json:"drivers"`
	TotalCount int64            `json:"total_count"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
	TotalPages int              `json:"total_pages"`
}

// DriverPerformance represents driver performance metrics
type DriverPerformance struct {
	DriverID          uint64  `json:"driver_id"`
	TotalShifts       int     `json:"total_shifts"`
	CompletedShifts   int     `json:"completed_shifts"`
	TotalOrders       int     `json:"total_orders"`
	CompletedOrders   int     `json:"completed_orders"`
	CancelledOrders   int     `json:"cancelled_orders"`
	TotalDistance     float64 `json:"total_distance"`
	TotalEarnings     float64 `json:"total_earnings"`
	AverageRating     float64 `json:"average_rating"`
	CompletionRate    float64 `json:"completion_rate"`
	LastShiftDate     *time.Time `json:"last_shift_date"`
}
