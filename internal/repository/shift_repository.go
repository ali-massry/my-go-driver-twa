package repository

import (
	"context"
	"time"

	"my-go-driver/internal/domain/shift"

	"gorm.io/gorm"
)

type shiftRepository struct {
	db *gorm.DB
}

// NewShiftRepository creates a new shift repository
func NewShiftRepository(db *gorm.DB) shift.Repository {
	return &shiftRepository{db: db}
}

func (r *shiftRepository) GetByDriverID(ctx context.Context, driverID uint64, query shift.ListShiftsQuery) ([]shift.DriverShift, int64, error) {
	var shifts []shift.DriverShift
	var total int64

	db := r.db.WithContext(ctx).Model(&shift.DriverShift{}).Where("driver_id = ?", driverID)

	// Apply filters
	if query.CompanyID > 0 {
		db = db.Where("company_id = ?", query.CompanyID)
	}

	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	if query.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", query.StartDate)
		if err == nil {
			db = db.Where("shift_date >= ?", startDate)
		}
	}

	if query.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", query.EndDate)
		if err == nil {
			db = db.Where("shift_date <= ?", endDate)
		}
	}

	// Get total count
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Set defaults
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Limit <= 0 {
		query.Limit = 10
	}

	// Apply pagination
	offset := (query.Page - 1) * query.Limit
	err := db.Offset(offset).Limit(query.Limit).Order("shift_date DESC, created_at DESC").Find(&shifts).Error

	return shifts, total, err
}
