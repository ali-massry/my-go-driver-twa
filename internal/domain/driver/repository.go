package driver

import "context"

// Repository defines the interface for driver data access
type Repository interface {
	Create(ctx context.Context, driver *Driver) error
	GetByID(ctx context.Context, id uint64) (*Driver, error)
	GetByPhone(ctx context.Context, phone string, companyID uint64) (*Driver, error)
	Update(ctx context.Context, driver *Driver) error
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, query ListDriversQuery) ([]Driver, int64, error)
	UpdateStatus(ctx context.Context, id uint64, status DriverStatus) error
	GetPerformance(ctx context.Context, driverID uint64) (*DriverPerformance, error)
}
