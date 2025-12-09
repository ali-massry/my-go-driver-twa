package shift

import "time"

type ShiftStatus string

const (
	ShiftStatusScheduled ShiftStatus = "scheduled"
	ShiftStatusOngoing   ShiftStatus = "ongoing"
	ShiftStatusCompleted ShiftStatus = "completed"
	ShiftStatusCancelled ShiftStatus = "cancelled"
)

// DriverShift represents a driver shift entity
type DriverShift struct {
	ID              uint64      `json:"id" gorm:"primaryKey"`
	DriverID        uint64      `json:"driver_id" gorm:"not null"`
	CompanyID       uint64      `json:"company_id" gorm:"not null"`
	ShiftDate       time.Time   `json:"shift_date" gorm:"type:date;not null"`
	StartTime       *time.Time  `json:"start_time"`
	EndTime         *time.Time  `json:"end_time"`
	Status          ShiftStatus `json:"status" gorm:"type:enum('scheduled','ongoing','completed','cancelled');default:scheduled"`
	TotalOrders     int         `json:"total_orders" gorm:"default:0"`
	CompletedOrders int         `json:"completed_orders" gorm:"default:0"`
	CancelledOrders int         `json:"cancelled_orders" gorm:"default:0"`
	TotalDistance   float64     `json:"total_distance" gorm:"type:decimal(10,2);default:0.00"`
	TotalEarnings   float64     `json:"total_earnings" gorm:"type:decimal(10,2);default:0.00"`
	Rating          float64     `json:"rating" gorm:"type:decimal(3,2);default:0.00"`
	Notes           string      `json:"notes" gorm:"type:text"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

func (DriverShift) TableName() string {
	return "driver_shifts"
}
