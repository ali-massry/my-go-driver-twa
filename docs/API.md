# API Documentation

## Base URL

```
http://localhost:8080/api/v1
```

## Authentication

All protected endpoints require a JWT token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

## Response Format

### Success Response

```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error message",
  "errors": [
    {
      "field": "email",
      "message": "Invalid email format"
    }
  ]
}
```

## Endpoints

### Health Check

Check if the API is running.

**Endpoint:** `GET /health`

**Authentication:** None

**Response:**
```json
{
  "status": "healthy"
}
```

---

## Authentication Endpoints

### Register

Create a new user account.

**Endpoint:** `POST /api/v1/auth/register`

**Authentication:** None

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Validation Rules:**
- `name`: required, min 2 chars, max 100 chars
- `email`: required, valid email format
- `password`: required, min 6 chars, max 100 chars

**Success Response:** `201 Created`
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Error Responses:**

`400 Bad Request` - Validation errors
```json
{
  "success": false,
  "message": "Validation failed",
  "errors": [
    {
      "field": "email",
      "message": "Invalid email format"
    }
  ]
}
```

`409 Conflict` - Email already exists
```json
{
  "success": false,
  "message": "Email already exists",
  "errors": null
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

---

### Login

Authenticate a user and receive a JWT token.

**Endpoint:** `POST /api/v1/auth/login`

**Authentication:** None

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Validation Rules:**
- `email`: required, valid email format
- `password`: required

**Success Response:** `200 OK`
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Error Responses:**

`401 Unauthorized` - Invalid credentials
```json
{
  "success": false,
  "message": "Invalid credentials",
  "errors": null
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

---

### Get Current User

Get the authenticated user's profile.

**Endpoint:** `GET /api/v1/auth/me`

**Authentication:** Required

**Success Response:** `200 OK`
```json
{
  "success": true,
  "message": "User profile retrieved successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

**Error Responses:**

`401 Unauthorized` - Missing or invalid token
```json
{
  "success": false,
  "message": "Authorization header required",
  "errors": null
}
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## User Management Endpoints

All user management endpoints require authentication.

### List All Users

Get a list of all users.

**Endpoint:** `GET /api/v1/users`

**Authentication:** Required

**Success Response:** `200 OK`
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": {
    "users": [
      {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      },
      {
        "id": 2,
        "name": "Jane Smith",
        "email": "jane@example.com",
        "created_at": "2024-01-02T10:00:00Z",
        "updated_at": "2024-01-02T10:00:00Z"
      }
    ]
  }
}
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### Create User

Create a new user (admin function).

**Endpoint:** `POST /api/v1/users`

**Authentication:** Required

**Request Body:**
```json
{
  "name": "Jane Smith",
  "email": "jane@example.com",
  "password": "password123"
}
```

**Success Response:** `201 Created`
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane@example.com",
    "created_at": "2024-01-02T10:00:00Z",
    "updated_at": "2024-01-02T10:00:00Z"
  }
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "email": "jane@example.com",
    "password": "password123"
  }'
```

---

### Get User by ID

Get a specific user's details.

**Endpoint:** `GET /api/v1/users/:id`

**Authentication:** Required

**URL Parameters:**
- `id` (integer, required): User ID

**Success Response:** `200 OK`
```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

**Error Responses:**

`404 Not Found` - User doesn't exist
```json
{
  "success": false,
  "message": "User not found",
  "errors": null
}
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### Update User

Update a user's information.

**Endpoint:** `PUT /api/v1/users/:id`

**Authentication:** Required

**URL Parameters:**
- `id` (integer, required): User ID

**Request Body:**
```json
{
  "name": "John Updated",
  "email": "john.updated@example.com"
}
```

**Validation Rules:**
- `name`: optional, min 2 chars, max 100 chars (if provided)
- `email`: optional, valid email format (if provided)

**Success Response:** `200 OK`
```json
{
  "success": true,
  "message": "User updated successfully",
  "data": {
    "id": 1,
    "name": "John Updated",
    "email": "john.updated@example.com",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T11:00:00Z"
  }
}
```

**Error Responses:**

`404 Not Found` - User doesn't exist
`409 Conflict` - Email already taken by another user

**cURL Example:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated"
  }'
```

---

### Delete User

Soft delete a user.

**Endpoint:** `DELETE /api/v1/users/:id`

**Authentication:** Required

**URL Parameters:**
- `id` (integer, required): User ID

**Success Response:** `200 OK`
```json
{
  "success": true,
  "message": "User deleted successfully",
  "data": null
}
```

**Error Responses:**

`404 Not Found` - User doesn't exist

**cURL Example:**
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## Error Codes

| Code | Description |
|------|-------------|
| 200  | Success |
| 201  | Created |
| 400  | Bad Request (validation errors) |
| 401  | Unauthorized (missing/invalid token) |
| 404  | Not Found |
| 409  | Conflict (duplicate resource) |
| 500  | Internal Server Error |

## Rate Limiting

Currently not implemented. Consider adding rate limiting middleware for production.

## Pagination

Currently not implemented. All list endpoints return all records. Consider adding pagination for large datasets.

## Testing with Postman

Import this collection to test the API:

1. Create a new collection in Postman
2. Add base URL as a variable: `{{base_url}} = http://localhost:8080/api/v1`
3. Create a request for each endpoint above
4. After login/register, save the token and use `{{token}}` variable in Authorization header

## Testing with HTTPie

```bash
# Install HTTPie
pip install httpie

# Register
http POST localhost:8080/api/v1/auth/register name="John" email="john@test.com" password="test123"

# Login
http POST localhost:8080/api/v1/auth/login email="john@test.com" password="test123"

# Get current user (replace TOKEN)
http GET localhost:8080/api/v1/auth/me Authorization:"Bearer TOKEN"

# List users
http GET localhost:8080/api/v1/users Authorization:"Bearer TOKEN"
```
