# TWA Driver API

A production-ready RESTful API built with Go, following clean architecture principles and industry best practices.

## Features

- **Clean Architecture** - Separation of concerns with domain-driven design
- **Dependency Injection** - Loose coupling and easy testing
- **JWT Authentication** - Secure token-based authentication
- **Structured Logging** - JSON logging with zerolog
- **Database Migrations** - Version-controlled schema changes
- **Docker Support** - Containerized deployment
- **Graceful Shutdown** - Proper cleanup on termination
- **Input Validation** - Request validation with detailed error messages
- **Health Checks** - Endpoint for monitoring

## Project Structure

```
my-go-driver/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/                    # Private application code
│   ├── app/
│   │   └── container.go         # Dependency injection container
│   ├── config/
│   │   ├── config.go            # Configuration management
│   │   └── database.go          # Database setup
│   ├── domain/                  # Business domain layer
│   │   └── user/
│   │       ├── entity.go        # Domain entities
│   │       ├── dto.go           # Data transfer objects
│   │       ├── repository.go    # Repository interface
│   │       └── service.go       # Service interface
│   ├── repository/              # Data access implementations
│   │   └── user_repository.go
│   ├── service/                 # Business logic implementations
│   │   └── user_service.go
│   ├── handler/                 # HTTP handlers (controllers)
│   │   ├── user_handler.go
│   │   └── auth_handler.go
│   ├── middleware/              # HTTP middleware
│   │   ├── auth.go
│   │   └── logger.go
│   └── router/                  # Route definitions
│       └── router.go
├── pkg/                         # Public reusable packages
│   ├── hash/                    # Password hashing utilities
│   ├── jwt/                     # JWT token management
│   ├── logger/                  # Structured logging
│   └── httputil/                # HTTP utilities
├── migrations/                  # Database migrations
│   ├── 000001_create_users_table.up.sql
│   └── 000001_create_users_table.down.sql
├── .env.example                 # Environment variables template
├── Dockerfile                   # Docker image definition
├── docker-compose.yml           # Docker Compose configuration
├── Makefile                     # Build automation
└── README.md                    # This file
```

## Architecture

The project follows **Clean Architecture** principles:

```
┌─────────────────────────────────────────────────────┐
│                   HTTP Layer                         │
│              (Handlers/Controllers)                  │
└────────────────────┬────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────┐
│                Service Layer                         │
│              (Business Logic)                        │
└────────────────────┬────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────┐
│              Repository Layer                        │
│              (Data Access)                           │
└────────────────────┬────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────┐
│                  Database                            │
└─────────────────────────────────────────────────────┘
```

### Key Patterns

1. **Repository Pattern** - Abstracts data access logic
2. **Service Pattern** - Encapsulates business logic
3. **Dependency Injection** - Constructor-based injection for loose coupling
4. **Interface Segregation** - Small, focused interfaces
5. **DTO Pattern** - Separate data transfer objects from domain entities

## Getting Started

### Prerequisites

- Go 1.23 or higher
- MySQL 8.0 or higher
- Docker and Docker Compose (optional)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd my-go-driver
   ```

2. **Copy environment variables**
   ```bash
   cp .env.example .env
   ```

3. **Edit `.env` and configure your database**
   ```env
   DB_DSN=root:password@tcp(127.0.0.1:3306)/twa-driver-app?charset=utf8mb4&parseTime=True&loc=Local
   JWT_SECRET=your-super-secret-jwt-key
   ```

4. **Install dependencies**
   ```bash
   make install
   ```

5. **Run the application**
   ```bash
   make run
   ```

### Using Docker

1. **Start the application with Docker Compose**
   ```bash
   docker-compose up -d
   ```

2. **View logs**
   ```bash
   docker-compose logs -f app
   ```

3. **Stop the application**
   ```bash
   docker-compose down
   ```

## API Endpoints

### Public Endpoints

| Method | Endpoint           | Description              |
|--------|--------------------|--------------------------|
| GET    | `/health`          | Health check             |
| POST   | `/api/v1/auth/register` | Register new user   |
| POST   | `/api/v1/auth/login`    | Login user          |

### Protected Endpoints (Require JWT Token)

| Method | Endpoint                | Description              |
|--------|-------------------------|--------------------------|
| GET    | `/api/v1/auth/me`       | Get current user profile |
| GET    | `/api/v1/users`         | List all users           |
| POST   | `/api/v1/users`         | Create new user          |
| GET    | `/api/v1/users/:id`     | Get user by ID           |
| PUT    | `/api/v1/users/:id`     | Update user              |
| DELETE | `/api/v1/users/:id`     | Delete user              |

## API Examples

### 1. Register a new user

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Response:**
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

### 2. Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 3. Get current user profile (Protected)

```bash
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 4. List all users (Protected)

```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Development

### Available Make Commands

```bash
make help              # Show all available commands
make install           # Install dependencies
make build             # Build the application
make run               # Run the application
make test              # Run tests
make test-coverage     # Run tests with coverage report
make clean             # Clean build artifacts
make fmt               # Format code
make vet               # Run go vet
make lint              # Run golangci-lint
make install-tools     # Install development tools
make docker-build      # Build Docker image
make docker-run        # Run with Docker Compose
make docker-stop       # Stop Docker containers
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific test
go test -v ./internal/service/...
```

### Code Quality

```bash
# Format code
make fmt

# Run linter
make lint

# Run static analysis
make vet
```

## Database Migrations

Migrations are located in the `migrations/` directory. Each migration has an `up` and `down` SQL file.

### Creating a new migration

```bash
make migrate-create name=add_users_table
```

This creates:
- `migrations/XXXXXX_add_users_table.up.sql`
- `migrations/XXXXXX_add_users_table.down.sql`

**Note:** Currently, the application uses GORM's AutoMigrate for development. For production, use proper migration tools like [golang-migrate](https://github.com/golang-migrate/migrate).

## Configuration

Configuration is managed through environment variables using Viper. See `.env.example` for all available options.

| Variable              | Description                    | Default     |
|-----------------------|--------------------------------|-------------|
| SERVER_PORT           | Server port                    | 8080        |
| SERVER_ENVIRONMENT    | Environment (dev/prod)         | development |
| DB_DSN                | Database connection string     | required    |
| DB_MAX_OPEN_CONNS     | Max open DB connections        | 25          |
| DB_MAX_IDLE_CONNS     | Max idle DB connections        | 5           |
| JWT_SECRET            | JWT signing secret             | required    |
| JWT_EXPIRATION        | Token expiration duration      | 24h         |

## Security Best Practices

- ✅ Passwords are hashed using bcrypt
- ✅ JWT tokens for stateless authentication
- ✅ Input validation on all endpoints
- ✅ SQL injection prevention via GORM
- ✅ CORS configuration ready
- ✅ Rate limiting ready (add middleware)
- ✅ Graceful shutdown for data integrity

## Deployment

### Production Checklist

1. ✅ Set strong `JWT_SECRET`
2. ✅ Set `SERVER_ENVIRONMENT=production`
3. ✅ Use proper database credentials
4. ✅ Enable HTTPS/TLS
5. ✅ Configure CORS properly
6. ✅ Set up monitoring and logging
7. ✅ Use database migrations instead of AutoMigrate
8. ✅ Set up backup strategy
9. ✅ Configure rate limiting
10. ✅ Review and update security headers

### Docker Production Deployment

```bash
# Build for production
docker build -t twa-driver-api:latest .

# Run with production settings
docker run -d \
  -p 8080:8080 \
  -e SERVER_ENVIRONMENT=production \
  -e DB_DSN="user:pass@tcp(db:3306)/dbname" \
  -e JWT_SECRET="your-production-secret" \
  twa-driver-api:latest
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Support

For issues, questions, or contributions, please open an issue on GitHub.
