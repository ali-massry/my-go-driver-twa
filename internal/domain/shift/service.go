package shift

import "context"

// Service defines the interface for shift business logic
type Service interface {
	GetDriverShifts(ctx context.Context, driverID uint64, query ListShiftsQuery) (*PaginatedShiftsResponse, error)
}
