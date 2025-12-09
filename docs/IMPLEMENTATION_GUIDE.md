# Implementation Guide - Driver-Based Delivery Platform

## Overview

This is a **multi-tenant driver-based delivery/logistics platform** with comprehensive features for managing companies, stores, drivers, vehicles, orders, and deliveries.

## Database Schema

The system has **13 migration files** creating **24 tables**:

### Core Tables
1. **companies** - Multi-tenant company management
2. **company_admins** - Company admin users
3. **modules_master** - Available system modules
4. **company_modules** - Company-specific module enablement
5. **stores** - Store locations per company
6. **store_modules** - Store-specific module configuration
7. **vehicles** - Delivery vehicles
8. **drivers** - Delivery drivers
9. **driver_vehicle_assignments** - Driver-vehicle relationships
10. **driver_locations** - Real-time GPS tracking
11. **clients** - Delivery clients/customers
12. **products** - Product catalog
13. **vehicle_stock** - Inventory in vehicles
14. **stock_logs** - Stock movement history
15. **orders** - Delivery orders
16. **order_items** - Order line items
17. **proof_of_delivery** - POD (photos/signatures/OTP)
18. **order_tracking_logs** - Order status history
19. **payments** - Payment transactions
20. **driver_cash_reconciliation** - Daily cash reconciliation
21. **notifications** - Push notifications
22. **chat_messages** - In-app chat
23. **activity_logs** - Audit trail
24. **users** - System users (from migration 001)

## Current Implementation Status

### ✅ Completed
- **Database Migrations** - All 13 migrations created
- **User Module** - Complete authentication system
  - Register/Login with JWT
  - User management CRUD
  - Password hashing with bcrypt
  - Middleware authentication

### 📝 To Be Implemented

The system currently has a complete **User authentication module** as a reference implementation.

To implement the remaining modules, follow the same pattern:

## Implementation Pattern

Each module follows Clean Architecture:

```
internal/domain/{module}/
├── entity.go      # Domain entity (database model)
├── dto.go         # Data Transfer Objects (API contracts)
├── repository.go  # Repository interface
└── service.go     # Service interface

internal/repository/
└── {module}_repository.go  # Repository implementation

internal/service/
└── {module}_service.go     # Service implementation

internal/handler/
└── {module}_handler.go     # HTTP handlers
```

## Example: Implementing Company Module

### 1. Create Entity (`internal/domain/company/entity.go`)

```go
package company

import (
    "database/sql/driver"
    "encoding/json"
    "time"
    "gorm.io/gorm"
)

type Company struct {
    ID           uint           `json:"id" gorm:"primarykey"`
    Name         string         `json:"name" gorm:"not null"`
    Email        string         `json:"email"`
    Phone        string         `json:"phone"`
    Address      string         `json:"address"`
    Timezone     string         `json:"timezone" gorm:"default:'UTC'"`
    LogoURL      string         `json:"logo_url"`
    ColorPalette ColorPalette   `json:"color_palette" gorm:"type:json"`
    FontFamily   string         `json:"font_family"`
    Status       string         `json:"status" gorm:"type:ENUM('active','suspended');default:'active'"`
    CreatedAt    time.Time      `json:"created_at"`
    UpdatedAt    time.Time      `json:"updated_at"`
}

type ColorPalette struct {
    Primary   string `json:"primary"`
    Secondary string `json:"secondary"`
    Accent    string `json:"accent"`
}

// Implement Scanner and Valuer for JSON field
func (c *ColorPalette) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return nil
    }
    return json.Unmarshal(bytes, c)
}

func (c ColorPalette) Value() (driver.Value, error) {
    return json.Marshal(c)
}
```

### 2. Create DTOs (`internal/domain/company/dto.go`)

```go
package company

type CreateCompanyRequest struct {
    Name    string `json:"name" binding:"required,min=2,max=255"`
    Email   string `json:"email" binding:"omitempty,email"`
    Phone   string `json:"phone"`
    Address string `json:"address"`
}

type UpdateCompanyRequest struct {
    Name    string `json:"name" binding:"omitempty,min=2,max=255"`
    Email   string `json:"email" binding:"omitempty,email"`
    Phone   string `json:"phone"`
    Address string `json:"address"`
}

type CompanyResponse struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}

func (c *Company) ToResponse() CompanyResponse {
    return CompanyResponse{
        ID:        c.ID,
        Name:      c.Name,
        Email:     c.Email,
        Phone:     c.Phone,
        Status:    c.Status,
        CreatedAt: c.CreatedAt,
    }
}
```

### 3. Create Repository Interface (`internal/domain/company/repository.go`)

```go
package company

type Repository interface {
    GetAll() ([]Company, error)
    GetByID(id uint) (*Company, error)
    Create(company *Company) error
    Update(company *Company) error
    Delete(id uint) error
}
```

### 4. Create Service Interface (`internal/domain/company/service.go`)

```go
package company

type Service interface {
    GetAllCompanies() ([]Company, error)
    GetCompanyByID(id uint) (*Company, error)
    CreateCompany(req CreateCompanyRequest) (*Company, error)
    UpdateCompany(id uint, req UpdateCompanyRequest) (*Company, error)
    DeleteCompany(id uint) error
}
```

### 5. Implement Repository (`internal/repository/company_repository.go`)

```go
package repository

import (
    "github.com/ali-massry/my-go-driver/internal/domain/company"
    "gorm.io/gorm"
)

type companyRepository struct {
    db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) company.Repository {
    return &companyRepository{db: db}
}

func (r *companyRepository) GetAll() ([]company.Company, error) {
    var companies []company.Company
    err := r.db.Find(&companies).Error
    return companies, err
}

func (r *companyRepository) GetByID(id uint) (*company.Company, error) {
    var c company.Company
    err := r.db.First(&c, id).Error
    return &c, err
}

func (r *companyRepository) Create(c *company.Company) error {
    return r.db.Create(c).Error
}

func (r *companyRepository) Update(c *company.Company) error {
    return r.db.Save(c).Error
}

func (r *companyRepository) Delete(id uint) error {
    return r.db.Delete(&company.Company{}, id).Error
}
```

### 6. Implement Service (`internal/service/company_service.go`)

```go
package service

import (
    "errors"
    "github.com/ali-massry/my-go-driver/internal/domain/company"
    "gorm.io/gorm"
)

type companyService struct {
    repo company.Repository
}

func NewCompanyService(repo company.Repository) company.Service {
    return &companyService{repo: repo}
}

func (s *companyService) GetAllCompanies() ([]company.Company, error) {
    return s.repo.GetAll()
}

func (s *companyService) GetCompanyByID(id uint) (*company.Company, error) {
    return s.repo.GetByID(id)
}

func (s *companyService) CreateCompany(req company.CreateCompanyRequest) (*company.Company, error) {
    c := &company.Company{
        Name:    req.Name,
        Email:   req.Email,
        Phone:   req.Phone,
        Address: req.Address,
    }

    if err := s.repo.Create(c); err != nil {
        return nil, err
    }
    return c, nil
}

func (s *companyService) UpdateCompany(id uint, req company.UpdateCompanyRequest) (*company.Company, error) {
    c, err := s.repo.GetByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("company not found")
        }
        return nil, err
    }

    if req.Name != "" {
        c.Name = req.Name
    }
    if req.Email != "" {
        c.Email = req.Email
    }
    if req.Phone != "" {
        c.Phone = req.Phone
    }

    if err := s.repo.Update(c); err != nil {
        return nil, err
    }
    return c, nil
}

func (s *companyService) DeleteCompany(id uint) error {
    _, err := s.repo.GetByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return errors.New("company not found")
        }
        return err
    }
    return s.repo.Delete(id)
}
```

### 7. Create Handler (`internal/handler/company_handler.go`)

```go
package handler

import (
    "net/http"
    "github.com/ali-massry/my-go-driver/internal/domain/company"
    "github.com/ali-massry/my-go-driver/pkg/httputil"
    "github.com/gin-gonic/gin"
)

type CompanyHandler struct {
    service company.Service
}

func NewCompanyHandler(service company.Service) *CompanyHandler {
    return &CompanyHandler{service: service}
}

func (h *CompanyHandler) GetCompanies(c *gin.Context) {
    companies, err := h.service.GetAllCompanies()
    if err != nil {
        httputil.RespondError(c, http.StatusInternalServerError, "Failed to fetch companies", nil)
        return
    }

    httputil.RespondSuccess(c, http.StatusOK, "Companies retrieved successfully", companies)
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
    var req company.CreateCompanyRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        validationErrors := httputil.FormatValidationErrors(err)
        httputil.RespondError(c, http.StatusBadRequest, "Validation failed", validationErrors)
        return
    }

    comp, err := h.service.CreateCompany(req)
    if err != nil {
        httputil.RespondError(c, http.StatusInternalServerError, "Failed to create company", nil)
        return
    }

    httputil.RespondSuccess(c, http.StatusCreated, "Company created successfully", comp.ToResponse())
}

// Add GetByID, Update, Delete similarly...
```

### 8. Wire in Container (`internal/app/container.go`)

```go
// Add to Container struct
CompanyHandler *handler.CompanyHandler

// Add to NewContainer function
companyRepo := repository.NewCompanyRepository(db)
companyService := service.NewCompanyService(companyRepo)
companyHandler := handler.NewCompanyHandler(companyService)
```

### 9. Add Routes (`internal/router/router.go`)

```go
// Add to protected routes
companies := protected.Group("/companies")
{
    companies.GET("", companyHandler.GetCompanies)
    companies.POST("", companyHandler.CreateCompany)
    companies.GET("/:id", companyHandler.GetCompanyByID)
    companies.PUT("/:id", companyHandler.UpdateCompany)
    companies.DELETE("/:id", companyHandler.DeleteCompany)
}
```

## Modules to Implement

Follow the same pattern for:

1. **Company Module** ✅ (See above)
2. **Store Module** - Store management per company
3. **Driver Module** - Driver management with authentication
4. **Vehicle Module** - Vehicle fleet management
5. **Client Module** - Customer/client management
6. **Product Module** - Product catalog
7. **Order Module** - Order processing workflow
8. **Payment Module** - Payment and reconciliation
9. **Tracking Module** - Real-time GPS tracking
10. **Notification Module** - Push notifications

## Quick Start Development

1. **Run migrations:**
   ```bash
   # Migrations are ready in /migrations folder
   # Use golang-migrate or let GORM AutoMigrate handle it
   ```

2. **Implement modules incrementally:**
   - Start with Company
   - Then Store
   - Then Driver
   - Then Order
   - Add others as needed

3. **Test each module:**
   ```bash
   make test
   ```

## API Structure

```
/api/v1/
├── /auth/
│   ├── POST /register
│   ├── POST /login
│   └── GET  /me
├── /companies/        [Protected]
├── /stores/           [Protected]
├── /drivers/          [Protected]
├── /vehicles/         [Protected]
├── /clients/          [Protected]
├── /products/         [Protected]
├── /orders/           [Protected]
├── /payments/         [Protected]
└── /tracking/         [Protected]
```

## Next Steps

1. Implement Company module (use example above)
2. Implement Driver module (similar pattern + authentication)
3. Implement Order module (includes order items, POD)
4. Add real-time features (WebSocket for tracking)
5. Add notification system
6. Add analytics/reporting

## Notes

- All tables are created via migrations
- Follow the User module as reference
- Use dependency injection throughout
- Keep business logic in service layer
- Keep HTTP logic in handlers
- Use DTOs for API contracts
- Implement proper error handling

## Database Connection

The application will create all tables automatically via AutoMigrate in development. For production, use the migration files in `/migrations`.
