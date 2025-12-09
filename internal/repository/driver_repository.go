package repository

import (
	"context"

	"my-go-driver/internal/domain/driver"

	"gorm.io/gorm"
)

type driverRepository struct {
	db *gorm.DB
}

// NewDriverRepository creates a new driver repository
func NewDriverRepository(db *gorm.DB) driver.Repository {
	return &driverRepository{db: db}
}

func (r *driverRepository) Create(ctx context.Context, d *driver.Driver) error {
	return r.db.WithContext(ctx).Create(d).Error
}

func (r *driverRepository) GetByID(ctx context.Context, id uint64) (*driver.Driver, error) {
	var d driver.Driver
	err := r.db.WithContext(ctx).First(&d, id).Error
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *driverRepository) GetByPhone(ctx context.Context, phone string, companyID uint64) (*driver.Driver, error) {
	var d driver.Driver
	err := r.db.WithContext(ctx).Where("phone = ? AND company_id = ?", phone, companyID).First(&d).Error
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *driverRepository) Update(ctx context.Context, d *driver.Driver) error {
	return r.db.WithContext(ctx).Save(d).Error
}

func (r *driverRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&driver.Driver{}, id).Error
}

func (r *driverRepository) List(ctx context.Context, query driver.ListDriversQuery) ([]driver.Driver, int64, error) {
	var drivers []driver.Driver
	var total int64

	db := r.db.WithContext(ctx).Model(&driver.Driver{})

	// Apply filters
	if query.CompanyID > 0 {
		db = db.Where("company_id = ?", query.CompanyID)
	}

	if query.StoreID > 0 {
		db = db.Where("store_id = ?", query.StoreID)
	}

	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	if query.OnlineStatus != "" {
		db = db.Where("online_status = ?", query.OnlineStatus)
	}

	if query.Search != "" {
		searchPattern := "%" + query.Search + "%"
		db = db.Where("full_name LIKE ? OR phone LIKE ? OR email LIKE ?", searchPattern, searchPattern, searchPattern)
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
	err := db.Offset(offset).Limit(query.Limit).Order("created_at DESC").Find(&drivers).Error

	return drivers, total, err
}

func (r *driverRepository) UpdateStatus(ctx context.Context, id uint64, status driver.DriverStatus) error {
	return r.db.WithContext(ctx).Model(&driver.Driver{}).Where("id = ?", id).Update("status", status).Error
}

func (r *driverRepository) GetPerformance(ctx context.Context, driverID uint64) (*driver.DriverPerformance, error) {
	var performance driver.DriverPerformance

	// Query to aggregate driver performance from shifts table
	query := `
		SELECT
			driver_id,
			COUNT(*) as total_shifts,
			SUM(CASE WHEN status = 'completed' THEN 1 ELSE 0 END) as completed_shifts,
			SUM(total_orders) as total_orders,
			SUM(completed_orders) as completed_orders,
			SUM(cancelled_orders) as cancelled_orders,
			SUM(total_distance) as total_distance,
			SUM(total_earnings) as total_earnings,
			AVG(rating) as average_rating,
			(SUM(completed_orders) * 100.0 / NULLIF(SUM(total_orders), 0)) as completion_rate,
			MAX(shift_date) as last_shift_date
		FROM driver_shifts
		WHERE driver_id = ? AND status IN ('completed', 'cancelled')
		GROUP BY driver_id
	`

	err := r.db.WithContext(ctx).Raw(query, driverID).Scan(&performance).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Return zero performance if no shifts found
			return &driver.DriverPerformance{
				DriverID: driverID,
			}, nil
		}
		return nil, err
	}

	performance.DriverID = driverID
	return &performance, nil
}
