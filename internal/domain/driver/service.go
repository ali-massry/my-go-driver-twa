package driver

import "context"

// Service defines the interface for driver business logic
type Service interface {
	CreateDriver(ctx context.Context, req CreateDriverRequest) (*DriverResponse, error)
	GetDriver(ctx context.Context, id uint64) (*DriverResponse, error)
	UpdateDriver(ctx context.Context, id uint64, req UpdateDriverRequest) (*DriverResponse, error)
	DeleteDriver(ctx context.Context, id uint64) error
	ListDrivers(ctx context.Context, query ListDriversQuery) (*PaginatedDriversResponse, error)
	AssignToCompany(ctx context.Context, driverID uint64, req AssignDriverToCompanyRequest) (*DriverResponse, error)
	BlockDriver(ctx context.Context, driverID uint64) error
	UnblockDriver(ctx context.Context, driverID uint64) error
	GetDriverPerformance(ctx context.Context, driverID uint64) (*DriverPerformance, error)
}
