# TWA Driver API - Admin Dashboard Documentation

## Overview

This API provides comprehensive admin dashboard functionality for managing companies, drivers, modules, and monitoring driver performance.

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication

Most endpoints require JWT authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

---

## API Endpoints

### Authentication

#### 1. Admin Login
**POST** `/admin/auth/login`

Login as a company admin/owner.

**Request Body:**
```json
{
  "email": "owner@company.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "admin": {
      "id": 1,
      "company_id": 1,
      "full_name": "John Doe",
      "email": "owner@company.com",
      "phone": "+1234567890",
      "role": "owner",
      "is_active": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### 2. Get Admin Profile
**GET** `/admin/auth/me` 🔒

Get the authenticated admin's profile.

**Response:**
```json
{
  "success": true,
  "message": "Profile retrieved successfully",
  "data": {
    "id": 1,
    "company_id": 1,
    "full_name": "John Doe",
    "email": "owner@company.com",
    "phone": "+1234567890",
    "role": "owner",
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### Company Management

#### 3. Create Company
**POST** `/admin/companies` 🔒

Create a new company with an owner account.

**Request Body:**
```json
{
  "name": "Acme Corporation",
  "email": "info@acme.com",
  "phone": "+1234567890",
  "address": "123 Main St, City, Country",
  "timezone": "UTC",
  "owner_name": "John Doe",
  "owner_email": "john@acme.com",
  "owner_phone": "+1234567891",
  "owner_password": "securePassword123"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Company created successfully",
  "data": {
    "company": {
      "id": 1,
      "name": "Acme Corporation",
      "email": "info@acme.com",
      "phone": "+1234567890",
      "address": "123 Main St, City, Country",
      "timezone": "UTC",
      "logo_url": "",
      "color_palette": null,
      "font_family": "",
      "status": "active",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    "owner": {
      "id": 1,
      "company_id": 1,
      "full_name": "John Doe",
      "email": "john@acme.com",
      "phone": "+1234567891",
      "role": "owner",
      "is_active": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  }
}
```

#### 4. List Companies
**GET** `/admin/companies` 🔒

List all companies with pagination and filtering.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10, max: 100)
- `status` (optional): Filter by status (`active` or `suspended`)
- `search` (optional): Search by name, email, or phone

**Response:**
```json
{
  "success": true,
  "message": "Companies retrieved successfully",
  "data": {
    "companies": [
      {
        "id": 1,
        "name": "Acme Corporation",
        "email": "info@acme.com",
        "phone": "+1234567890",
        "address": "123 Main St",
        "timezone": "UTC",
        "logo_url": "",
        "color_palette": null,
        "font_family": "",
        "status": "active",
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ],
    "total_count": 1,
    "page": 1,
    "limit": 10,
    "total_pages": 1
  }
}
```

#### 5. Get Company
**GET** `/admin/companies/:id` 🔒

Get a specific company by ID.

#### 6. Update Company
**PUT** `/admin/companies/:id` 🔒

Update company details.

**Request Body:**
```json
{
  "name": "Updated Company Name",
  "email": "newemail@company.com",
  "phone": "+9876543210",
  "address": "456 New Address",
  "timezone": "America/New_York"
}
```

#### 7. Delete Company
**DELETE** `/admin/companies/:id` 🔒

Delete a company (also deletes all associated data).

#### 8. Update Company Branding
**PUT** `/admin/companies/:id/branding` 🔒

Update company branding settings.

**Request Body:**
```json
{
  "logo_url": "https://example.com/logo.png",
  "color_palette": {
    "primary": "#007bff",
    "secondary": "#6c757d",
    "accent": "#ffc107"
  },
  "font_family": "Roboto"
}
```

#### 9. Suspend Company
**PUT** `/admin/companies/:id/suspend` 🔒

Suspend a company (prevents login and operations).

#### 10. Activate Company
**PUT** `/admin/companies/:id/activate` 🔒

Reactivate a suspended company.

---

### Module Management

#### 11. List All Modules
**GET** `/admin/modules` 🔒

Get all available modules.

**Response:**
```json
{
  "success": true,
  "message": "Modules retrieved successfully",
  "data": [
    {
      "id": 1,
      "module_key": "delivery_tracking",
      "name": "Delivery Tracking",
      "category": "Operations",
      "description": "Real-time delivery tracking and updates",
      "default_enabled": true,
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### 12. Assign Module to Company
**POST** `/admin/companies/:company_id/modules` 🔒

Assign a module to a company.

**Request Body:**
```json
{
  "module_id": 1,
  "is_enabled": true,
  "config": {
    "setting1": "value1",
    "setting2": "value2"
  }
}
```

#### 13. Get Company Modules
**GET** `/admin/companies/:company_id/modules` 🔒

Get all modules assigned to a company.

#### 14. Remove Module from Company
**DELETE** `/admin/companies/:company_id/modules/:module_id` 🔒

Remove a module from a company.

---

### Driver Management

#### 15. Create Driver
**POST** `/admin/drivers` 🔒

Create a new driver.

**Request Body:**
```json
{
  "company_id": 1,
  "store_id": 1,
  "full_name": "Jane Smith",
  "phone": "+1234567890",
  "email": "jane@example.com",
  "password": "driverPassword123",
  "profile_photo": "https://example.com/photo.jpg"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Driver created successfully",
  "data": {
    "id": 1,
    "company_id": 1,
    "store_id": 1,
    "full_name": "Jane Smith",
    "phone": "+1234567890",
    "email": "jane@example.com",
    "status": "active",
    "online_status": "offline",
    "rating": 0.0,
    "profile_photo": "https://example.com/photo.jpg",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 16. List Drivers
**GET** `/admin/drivers` 🔒

List all drivers with filtering and pagination.

**Query Parameters:**
- `page` (optional): Page number
- `limit` (optional): Items per page
- `company_id` (optional): Filter by company
- `store_id` (optional): Filter by store
- `status` (optional): Filter by status (`active`, `off_duty`, `suspended`)
- `online_status` (optional): Filter by online status (`online`, `offline`)
- `search` (optional): Search by name, phone, or email

**Response:**
```json
{
  "success": true,
  "message": "Drivers retrieved successfully",
  "data": {
    "drivers": [
      {
        "id": 1,
        "company_id": 1,
        "store_id": 1,
        "full_name": "Jane Smith",
        "phone": "+1234567890",
        "email": "jane@example.com",
        "status": "active",
        "online_status": "offline",
        "rating": 4.5,
        "profile_photo": "https://example.com/photo.jpg",
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ],
    "total_count": 1,
    "page": 1,
    "limit": 10,
    "total_pages": 1
  }
}
```

#### 17. Get Driver
**GET** `/admin/drivers/:id` 🔒

Get a specific driver by ID.

#### 18. Update Driver
**PUT** `/admin/drivers/:id` 🔒

Update driver details.

**Request Body:**
```json
{
  "full_name": "Jane Updated",
  "phone": "+9876543210",
  "email": "jane.updated@example.com",
  "store_id": 2,
  "profile_photo": "https://example.com/newphoto.jpg"
}
```

#### 19. Delete Driver
**DELETE** `/admin/drivers/:id` 🔒

Delete a driver.

#### 20. Assign Driver to Company
**PUT** `/admin/drivers/:id/assign-company` 🔒

Assign/reassign a driver to a different company.

**Request Body:**
```json
{
  "company_id": 2,
  "store_id": 3
}
```

#### 21. Block Driver
**PUT** `/admin/drivers/:id/block` 🔒

Block/suspend a driver (sets status to `suspended`).

**Response:**
```json
{
  "success": true,
  "message": "Driver blocked successfully",
  "data": null
}
```

#### 22. Unblock Driver
**PUT** `/admin/drivers/:id/unblock` 🔒

Unblock a driver (sets status to `active`).

**Response:**
```json
{
  "success": true,
  "message": "Driver unblocked successfully",
  "data": null
}
```

#### 23. Get Driver Performance
**GET** `/admin/drivers/:id/performance` 🔒

Get driver performance metrics.

**Response:**
```json
{
  "success": true,
  "message": "Performance retrieved successfully",
  "data": {
    "driver_id": 1,
    "total_shifts": 120,
    "completed_shifts": 115,
    "total_orders": 850,
    "completed_orders": 820,
    "cancelled_orders": 30,
    "total_distance": 2500.50,
    "total_earnings": 12500.00,
    "average_rating": 4.7,
    "completion_rate": 96.47,
    "last_shift_date": "2024-01-15T00:00:00Z"
  }
}
```

#### 24. Get Driver Shifts
**GET** `/admin/drivers/:id/shifts` 🔒

Get driver shift history.

**Query Parameters:**
- `page` (optional): Page number
- `limit` (optional): Items per page
- `status` (optional): Filter by status (`scheduled`, `ongoing`, `completed`, `cancelled`)
- `start_date` (optional): Filter from date (YYYY-MM-DD)
- `end_date` (optional): Filter to date (YYYY-MM-DD)

**Response:**
```json
{
  "success": true,
  "message": "Shifts retrieved successfully",
  "data": {
    "shifts": [
      {
        "id": 1,
        "driver_id": 1,
        "company_id": 1,
        "shift_date": "2024-01-15",
        "start_time": "2024-01-15T08:00:00Z",
        "end_time": "2024-01-15T17:00:00Z",
        "status": "completed",
        "total_orders": 25,
        "completed_orders": 23,
        "cancelled_orders": 2,
        "total_distance": 150.5,
        "total_earnings": 350.00,
        "rating": 4.8,
        "notes": "",
        "duration": "9h 0m",
        "created_at": "2024-01-15T00:00:00Z",
        "updated_at": "2024-01-15T17:00:00Z"
      }
    ],
    "total_count": 120,
    "page": 1,
    "limit": 10,
    "total_pages": 12
  }
}
```

---

## Database Schema

The API uses the following main tables:

- **companies**: Company information and branding
- **company_admins**: Company owners and administrators
- **drivers**: Driver accounts and profiles
- **driver_shifts**: Driver shift history and performance
- **modules_master**: Available system modules
- **company_modules**: Modules assigned to companies
- **stores**: Store/branch locations
- **vehicles**: Company vehicles
- **driver_vehicle_assignments**: Vehicle assignments to drivers
- **driver_locations**: GPS location tracking

---

## Error Responses

All errors follow this format:

```json
{
  "success": false,
  "message": "Error message",
  "errors": "Detailed error information"
}
```

**Common HTTP Status Codes:**
- `200 OK`: Success
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request data
- `401 Unauthorized`: Missing or invalid authentication
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

---

## Getting Started

### 1. Start the Server
```bash
make run
```

### 2. Create Your First Company
```bash
curl -X POST http://localhost:8080/api/v1/admin/companies \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Company",
    "email": "info@mycompany.com",
    "owner_name": "Admin User",
    "owner_email": "admin@mycompany.com",
    "owner_password": "securePassword123"
  }'
```

### 3. Login as Owner
```bash
curl -X POST http://localhost:8080/api/v1/admin/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@mycompany.com",
    "password": "securePassword123"
  }'
```

### 4. Use the Token
Copy the token from the login response and use it in subsequent requests:

```bash
curl -X GET http://localhost:8080/api/v1/admin/companies \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

---

## Notes

- 🔒 indicates endpoints that require authentication
- All timestamps are in UTC
- Phone numbers should include country code
- All IDs are unsigned 64-bit integers
- Pagination defaults: page=1, limit=10
- Maximum limit per page: 100

---

## Support

For issues or questions, please check the README.md or create an issue on GitHub.
