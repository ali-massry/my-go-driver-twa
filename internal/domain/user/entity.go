package user

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user entity
type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"-"`
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}
