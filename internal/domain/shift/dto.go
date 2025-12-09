package shift

import "time"

// ShiftResponse represents a shift response
type ShiftResponse struct {
	ID              uint64      `json:"id"`
	DriverID        uint64      `json:"driver_id"`
	CompanyID       uint64      `json:"company_id"`
	ShiftDate       time.Time   `json:"shift_date"`
	StartTime       *time.Time  `json:"start_time"`
	EndTime         *time.Time  `json:"end_time"`
	Status          ShiftStatus `json:"status"`
	TotalOrders     int         `json:"total_orders"`
	CompletedOrders int         `json:"completed_orders"`
	CancelledOrders int         `json:"cancelled_orders"`
	TotalDistance   float64     `json:"total_distance"`
	TotalEarnings   float64     `json:"total_earnings"`
	Rating          float64     `json:"rating"`
	Notes           string      `json:"notes"`
	Duration        string      `json:"duration,omitempty"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

// ListShiftsQuery represents query parameters for listing shifts
type ListShiftsQuery struct {
	Page      int         `form:"page" binding:"omitempty,min=1"`
	Limit     int         `form:"limit" binding:"omitempty,min=1,max=100"`
	DriverID  uint64      `form:"driver_id" binding:"omitempty"`
	CompanyID uint64      `form:"company_id" binding:"omitempty"`
	Status    ShiftStatus `form:"status" binding:"omitempty,oneof=scheduled ongoing completed cancelled"`
	StartDate string      `form:"start_date" binding:"omitempty"`
	EndDate   string      `form:"end_date" binding:"omitempty"`
}

// PaginatedShiftsResponse represents paginated shifts response
type PaginatedShiftsResponse struct {
	Shifts     []ShiftResponse `json:"shifts"`
	TotalCount int64           `json:"total_count"`
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
	TotalPages int             `json:"total_pages"`
}
