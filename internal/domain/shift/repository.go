package shift

import "context"

// Repository defines the interface for shift data access
type Repository interface {
	GetByDriverID(ctx context.Context, driverID uint64, query ListShiftsQuery) ([]DriverShift, int64, error)
}
