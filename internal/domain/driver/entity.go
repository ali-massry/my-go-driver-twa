package driver

import "time"

type DriverStatus string

const (
	DriverStatusActive    DriverStatus = "active"
	DriverStatusOffDuty   DriverStatus = "off_duty"
	DriverStatusSuspended DriverStatus = "suspended"
)

type OnlineStatus string

const (
	OnlineStatusOnline  OnlineStatus = "online"
	OnlineStatusOffline OnlineStatus = "offline"
)

// Driver represents a driver entity
type Driver struct {
	ID           uint64       `json:"id" gorm:"primaryKey"`
	CompanyID    uint64       `json:"company_id" gorm:"not null"`
	StoreID      *uint64      `json:"store_id"`
	FullName     string       `json:"full_name" gorm:"not null"`
	Phone        string       `json:"phone" gorm:"not null"`
	Email        string       `json:"email"`
	PasswordHash string       `json:"-" gorm:"not null"`
	Status       DriverStatus `json:"status" gorm:"type:enum('active','off_duty','suspended');default:active"`
	OnlineStatus OnlineStatus `json:"online_status" gorm:"type:enum('online','offline');default:offline"`
	Rating       float64      `json:"rating" gorm:"type:decimal(3,2);default:0.00"`
	ProfilePhoto string       `json:"profile_photo"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    *time.Time   `json:"deleted_at,omitempty" gorm:"index"`
}

func (Driver) TableName() string {
	return "drivers"
}
