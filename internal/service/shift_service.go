package service

import (
	"context"
	"fmt"
	"math"

	"my-go-driver/internal/domain/shift"
)

type shiftService struct {
	repo shift.Repository
}

// NewShiftService creates a new shift service
func NewShiftService(repo shift.Repository) shift.Service {
	return &shiftService{repo: repo}
}

func (s *shiftService) GetDriverShifts(ctx context.Context, driverID uint64, query shift.ListShiftsQuery) (*shift.PaginatedShiftsResponse, error) {
	shifts, total, err := s.repo.GetByDriverID(ctx, driverID, query)
	if err != nil {
		return nil, err
	}

	// Set defaults
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Limit <= 0 {
		query.Limit = 10
	}

	responses := make([]shift.ShiftResponse, len(shifts))
	for i, sh := range shifts {
		responses[i] = s.toShiftResponse(&sh)
	}

	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))

	return &shift.PaginatedShiftsResponse{
		Shifts:     responses,
		TotalCount: total,
		Page:       query.Page,
		Limit:      query.Limit,
		TotalPages: totalPages,
	}, nil
}

// Helper methods
func (s *shiftService) toShiftResponse(sh *shift.DriverShift) shift.ShiftResponse {
	response := shift.ShiftResponse{
		ID:              sh.ID,
		DriverID:        sh.DriverID,
		CompanyID:       sh.CompanyID,
		ShiftDate:       sh.ShiftDate,
		StartTime:       sh.StartTime,
		EndTime:         sh.EndTime,
		Status:          sh.Status,
		TotalOrders:     sh.TotalOrders,
		CompletedOrders: sh.CompletedOrders,
		CancelledOrders: sh.CancelledOrders,
		TotalDistance:   sh.TotalDistance,
		TotalEarnings:   sh.TotalEarnings,
		Rating:          sh.Rating,
		Notes:           sh.Notes,
		CreatedAt:       sh.CreatedAt,
		UpdatedAt:       sh.UpdatedAt,
	}

	// Calculate duration if shift is completed
	if sh.StartTime != nil && sh.EndTime != nil {
		duration := sh.EndTime.Sub(*sh.StartTime)
		hours := int(duration.Hours())
		minutes := int(duration.Minutes()) % 60
		response.Duration = fmt.Sprintf("%dh %dm", hours, minutes)
	}

	return response
}
