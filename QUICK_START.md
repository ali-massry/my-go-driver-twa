# Quick Start Guide - Authentication Flow

## 🔑 Complete Authentication Flow

### Step 1: Create Your First Company (NO TOKEN NEEDED) ✅

```bash
curl -X POST http://localhost:8080/api/v1/admin/companies \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Company",
    "email": "info@mycompany.com",
    "phone": "+1234567890",
    "address": "123 Main Street",
    "owner_name": "John Doe",
    "owner_email": "john@mycompany.com",
    "owner_password": "securePassword123"
  }'
```

**Response:**
```json
{
  "success": true,
  "message": "Company created successfully",
  "data": {
    "company": {
      "id": 1,
      "name": "My Company",
      "status": "active",
      ...
    },
    "owner": {
      "id": 1,
      "company_id": 1,
      "full_name": "John Doe",
      "email": "john@mycompany.com",
      "role": "owner",
      ...
    }
  }
}
```

### Step 2: Login with Owner Credentials (NO TOKEN NEEDED) ✅

```bash
curl -X POST http://localhost:8080/api/v1/admin/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@mycompany.com",
    "password": "securePassword123"
  }'
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
      "email": "john@mycompany.com",
      "role": "owner"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Step 3: Use the Token for All Other Requests 🔒

Copy the token from Step 2 and use it in the Authorization header:

```bash
# Set the token as a variable
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

# Get your profile
curl -X GET http://localhost:8080/api/v1/admin/auth/me \
  -H "Authorization: Bearer $TOKEN"

# List all companies
curl -X GET http://localhost:8080/api/v1/admin/companies \
  -H "Authorization: Bearer $TOKEN"

# Create a driver
curl -X POST http://localhost:8080/api/v1/admin/drivers \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "company_id": 1,
    "full_name": "Jane Smith",
    "phone": "+1234567891",
    "email": "jane@example.com",
    "password": "driverPassword123"
  }'
```

---

## 🔐 Public Endpoints (No Token Required)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/admin/companies` | Create company + owner |
| POST | `/api/v1/admin/auth/login` | Login to get token |

## 🔒 Protected Endpoints (Token Required)

All other endpoints require the JWT token in the Authorization header:

```
Authorization: Bearer YOUR_TOKEN_HERE
```

### Examples:

#### Get Admin Profile
```bash
curl -X GET http://localhost:8080/api/v1/admin/auth/me \
  -H "Authorization: Bearer $TOKEN"
```

#### List Drivers
```bash
curl -X GET http://localhost:8080/api/v1/admin/drivers?company_id=1 \
  -H "Authorization: Bearer $TOKEN"
```

#### Block a Driver
```bash
curl -X PUT http://localhost:8080/api/v1/admin/drivers/1/block \
  -H "Authorization: Bearer $TOKEN"
```

#### Get Driver Performance
```bash
curl -X GET http://localhost:8080/api/v1/admin/drivers/1/performance \
  -H "Authorization: Bearer $TOKEN"
```

---

## 📋 Summary

1. **Create Company** → Automatically creates owner account
2. **Login** → Get JWT token
3. **Use Token** → Access all other endpoints

**Token expires after 24 hours** - just login again to get a new token!

---

## ✅ Test the Flow

Here's a complete test script you can run:

```bash
#!/bin/bash

# Step 1: Create company
echo "Creating company..."
RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/admin/companies \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Company",
    "owner_name": "Test Admin",
    "owner_email": "test@example.com",
    "owner_password": "testpass123"
  }')
echo $RESPONSE | python3 -m json.tool

# Step 2: Login
echo -e "\n\nLogging in..."
LOGIN=$(curl -s -X POST http://localhost:8080/api/v1/admin/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "testpass123"
  }')
echo $LOGIN | python3 -m json.tool

# Extract token
TOKEN=$(echo $LOGIN | python3 -c "import sys, json; print(json.load(sys.stdin)['data']['token'])")
echo -e "\n\nToken: $TOKEN"

# Step 3: Test with token
echo -e "\n\nGetting profile..."
curl -s -X GET http://localhost:8080/api/v1/admin/auth/me \
  -H "Authorization: Bearer $TOKEN" | python3 -m json.tool
```

Save this as `test_auth.sh`, make it executable (`chmod +x test_auth.sh`), and run it!

---

## 🚨 Troubleshooting

### "Authorization header required"
- Make sure you're using `Bearer` prefix: `Authorization: Bearer YOUR_TOKEN`
- Check that the token hasn't expired (24 hours)

### "Invalid or expired token"
- Login again to get a new token

### "Company already exists"
- The email or company might already be registered
- Use a different email address

---

## 🎯 What's Next?

Now you can:
- ✅ Create and manage companies
- ✅ Create and manage drivers
- ✅ Assign modules to companies
- ✅ Block/unblock drivers
- ✅ View driver performance and shift history
- ✅ Update company branding

Check the **API_DOCUMENTATION.md** for all 24 available endpoints!
