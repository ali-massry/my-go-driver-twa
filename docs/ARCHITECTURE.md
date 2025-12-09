# Architecture Documentation

## Overview

This project follows **Clean Architecture** (also known as Hexagonal Architecture or Ports and Adapters) principles combined with **Domain-Driven Design (DDD)** patterns. The architecture ensures maintainability, testability, and scalability.

## Core Principles

### 1. Separation of Concerns
Each layer has a specific responsibility and doesn't know about layers above it.

### 2. Dependency Rule
Dependencies point inward. Inner layers don't depend on outer layers:
```
Handler → Service → Repository → Database
(outer)                          (inner)
```

### 3. Interface-Based Design
All layer interactions happen through interfaces, enabling:
- Easy testing with mocks
- Flexibility to swap implementations
- Loose coupling between components

## Project Structure Explained

### `/cmd/api/`
**Purpose:** Application entry points

- `main.go` - Bootstraps the application
- Initializes configuration
- Sets up dependency injection
- Starts the HTTP server
- Handles graceful shutdown

**Why separate?**
- Allows multiple entry points (API, CLI, workers)
- Keeps main package minimal

### `/internal/`
**Purpose:** Private application code

Code here cannot be imported by other projects. This enforces API boundaries.

#### `/internal/domain/`
**The Core - Business Domain Layer**

Contains:
- **Entities** (`entity.go`) - Core business objects
- **DTOs** (`dto.go`) - Data transfer objects for API requests/responses
- **Interfaces** (`repository.go`, `service.go`) - Contracts that outer layers must implement

**Why interfaces here?**
- Domain layer defines what it needs
- Infrastructure layers implement these needs
- Dependency Inversion Principle (SOLID)

Example:
```go
// Domain defines the contract
type Repository interface {
    GetAll() ([]User, error)
    Create(user *User) error
}

// Infrastructure implements it
type userRepository struct {
    db *gorm.DB
}
func (r *userRepository) GetAll() ([]User, error) { ... }
```

#### `/internal/repository/`
**Data Access Layer**

- Implements domain repository interfaces
- Handles database operations (GORM)
- Translates between domain entities and database tables

**Responsibilities:**
- SQL queries
- Database transactions
- Data persistence

**Does NOT contain:**
- Business logic
- HTTP handling
- Validation rules

#### `/internal/service/`
**Business Logic Layer**

- Implements domain service interfaces
- Contains business rules and workflows
- Orchestrates repository calls
- Manages transactions

**Example Flow:**
```go
func (s *userService) Register(req RegisterRequest) (*AuthResponse, error) {
    // 1. Business rule: Hash password
    hashedPassword := hash.Hash(req.Password)

    // 2. Create domain entity
    user := &User{...}

    // 3. Persist via repository
    s.repo.Create(user)

    // 4. Business rule: Generate token
    token := s.jwtManager.Generate(user.ID)

    return &AuthResponse{User: user, Token: token}
}
```

#### `/internal/handler/`
**HTTP Presentation Layer**

- Handles HTTP requests/responses
- Validates input (using Gin bindings)
- Calls service methods
- Formats responses

**Responsibilities:**
- Parse request bodies
- Validate input
- Call services
- Return JSON responses

**Does NOT contain:**
- Business logic
- Database queries
- Direct DB access

#### `/internal/middleware/`
**HTTP Middleware**

Cross-cutting concerns:
- Authentication (`auth.go`)
- Logging (`logger.go`)
- Error recovery
- CORS (if added)

#### `/internal/router/`
**Route Configuration**

- Maps URLs to handlers
- Applies middleware
- Groups related routes

#### `/internal/config/`
**Configuration Management**

- Loads environment variables
- Validates configuration
- Provides typed config objects
- Database connection setup

#### `/internal/app/`
**Dependency Injection Container**

The "glue" that wires everything together:

```go
func NewContainer(cfg *Config) (*Container, error) {
    // 1. Initialize infrastructure
    db := config.NewDatabase(cfg.Database)
    jwtManager := jwt.NewManager(cfg.JWT.Secret)

    // 2. Create repositories (data layer)
    userRepo := repository.NewUserRepository(db)

    // 3. Create services (business layer)
    userService := service.NewUserService(userRepo, jwtManager)

    // 4. Create handlers (presentation layer)
    userHandler := handler.NewUserHandler(userService)

    return &Container{...}
}
```

**Why a container?**
- Single source of truth for dependencies
- Easy to modify wiring
- Facilitates testing
- Clear dependency graph

### `/pkg/`
**Purpose:** Public, reusable packages

Can be imported by other projects. Contains utilities:

- `hash/` - Password hashing (bcrypt)
- `jwt/` - JWT token management
- `logger/` - Structured logging (zerolog)
- `httputil/` - HTTP response helpers

**Guidelines:**
- No application-specific logic
- Fully self-contained
- Well-documented
- Tested independently

### `/migrations/`
**Database Schema Versioning**

SQL migrations for schema changes:
- `*.up.sql` - Apply changes
- `*.down.sql` - Rollback changes

**Benefits:**
- Version control for database
- Repeatable deployments
- Rollback capability
- Team collaboration

## Data Flow

### Request Flow Example: Register User

```
1. HTTP Request
   POST /api/v1/auth/register
   Body: {"name": "John", "email": "...", "password": "..."}

   ↓

2. Router
   router.SetupRoutes()
   Matches route → authHandler.Register

   ↓

3. Handler (Presentation Layer)
   handler/auth_handler.go:Register()
   - Validates request body
   - Calls service.Register()

   ↓

4. Service (Business Logic Layer)
   service/user_service.go:Register()
   - Hashes password
   - Creates User entity
   - Calls repo.Create()
   - Generates JWT token
   - Returns AuthResponse

   ↓

5. Repository (Data Access Layer)
   repository/user_repository.go:Create()
   - Executes SQL INSERT
   - Returns saved user

   ↓

6. Database
   MySQL stores the user

   ↓

7. Response Flow (back up the chain)
   Service → Handler → Router → HTTP Response

   ↓

8. HTTP Response
   201 Created
   Body: {"success": true, "data": {..., "token": "..."}}
```

## Design Patterns Used

### 1. Repository Pattern
**Problem:** Direct database access spreads throughout the codebase
**Solution:** Centralize data access in repository interfaces

```go
// Interface in domain layer
type Repository interface {
    GetAll() ([]User, error)
    Create(user *User) error
}

// Implementation in repository layer
type userRepository struct {
    db *gorm.DB
}
```

**Benefits:**
- Swap databases without changing business logic
- Easy to mock for testing
- Centralized query logic

### 2. Service Pattern
**Problem:** Business logic mixed with HTTP handlers
**Solution:** Extract business logic into service layer

```go
type Service interface {
    Register(req RegisterRequest) (*AuthResponse, error)
    Login(req LoginRequest) (*AuthResponse, error)
}
```

**Benefits:**
- Reusable business logic
- Testable without HTTP
- Multiple interfaces can use same service

### 3. Dependency Injection (Constructor-Based)
**Problem:** Hard-coded dependencies make testing difficult
**Solution:** Inject dependencies through constructors

```go
type userService struct {
    repo       user.Repository  // Interface!
    jwtManager *jwt.Manager
}

func NewUserService(repo user.Repository, jwtManager *jwt.Manager) user.Service {
    return &userService{repo: repo, jwtManager: jwtManager}
}
```

**Benefits:**
- Easy to mock dependencies
- Clear dependency graph
- Loose coupling

### 4. DTO Pattern
**Problem:** Exposing domain entities directly in API
**Solution:** Use Data Transfer Objects for API layer

```go
// Request DTO
type RegisterRequest struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

// Response DTO
type UserResponse struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    // Password NOT exposed
}
```

**Benefits:**
- API contract independent of domain
- Hide sensitive fields (passwords)
- Versioning flexibility

### 5. Middleware Pattern
**Problem:** Cross-cutting concerns (auth, logging) duplicated
**Solution:** Chain of responsibility with middleware

```go
func Auth(jwtManager *jwt.Manager) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Validate token
        // Set user in context
        c.Next() // Continue to next handler
    }
}
```

### 6. Factory Pattern
**Problem:** Complex object construction
**Solution:** Constructor functions

```go
func NewUserRepository(db *gorm.DB) user.Repository {
    return &userRepository{db: db}
}
```

## SOLID Principles Applied

### S - Single Responsibility Principle
Each struct/package has one reason to change:
- `userRepository` - Only changes if data access logic changes
- `userService` - Only changes if business rules change
- `userHandler` - Only changes if API contract changes

### O - Open/Closed Principle
Open for extension, closed for modification:
- Want to add Redis caching? Create new repository wrapper
- Want to add email service? Inject it into service

### L - Liskov Substitution Principle
Interfaces can be swapped with implementations:
```go
var repo user.Repository
repo = &MySQLUserRepository{db}    // Works
repo = &PostgresUserRepository{db} // Also works
repo = &MockUserRepository{}       // Also works in tests
```

### I - Interface Segregation Principle
Small, focused interfaces:
```go
// Good: Specific interface
type Repository interface {
    GetAll() ([]User, error)
    Create(user *User) error
}

// Bad: God interface
type Repository interface {
    GetAll() ([]User, error)
    Create(user *User) error)
    SendEmail()
    GenerateReport()
    // ... 20 more methods
}
```

### D - Dependency Inversion Principle
Depend on abstractions (interfaces), not concretions:
```go
// Service depends on interface, not concrete implementation
type userService struct {
    repo user.Repository  // Interface!
}
```

## Testing Strategy

### Unit Tests
Test each layer independently:

```go
// Test service with mock repository
func TestUserService_Register(t *testing.T) {
    mockRepo := &MockUserRepository{} // Mock implements interface
    service := NewUserService(mockRepo, jwtManager)

    result, err := service.Register(RegisterRequest{...})

    assert.NoError(t, err)
    assert.NotNil(t, result.Token)
}
```

### Integration Tests
Test with real database:
```go
func TestUserRepository_Create(t *testing.T) {
    db := setupTestDB()
    repo := NewUserRepository(db)

    user := &User{Name: "Test", Email: "test@test.com"}
    err := repo.Create(user)

    assert.NoError(t, err)
    assert.NotZero(t, user.ID)
}
```

### API Tests
Test complete request/response:
```go
func TestRegisterEndpoint(t *testing.T) {
    router := setupTestRouter()

    w := httptest.NewRecorder()
    req := httptest.NewRequest("POST", "/api/v1/auth/register", body)
    router.ServeHTTP(w, req)

    assert.Equal(t, 201, w.Code)
}
```

## Scalability Considerations

### Horizontal Scaling
- Stateless design (JWT tokens)
- No in-memory sessions
- Database connection pooling
- Can run multiple instances

### Vertical Scaling
- Efficient database queries
- Connection pooling
- Goroutine-based concurrency (Gin)

### Future Enhancements

1. **Caching Layer**
   ```go
   type CachedUserRepository struct {
       repo  user.Repository
       cache *redis.Client
   }
   ```

2. **Event-Driven Architecture**
   ```go
   type EventPublisher interface {
       Publish(event Event) error
   }

   func (s *userService) Register(req RegisterRequest) {
       user := s.repo.Create(user)
       s.publisher.Publish(UserRegisteredEvent{UserID: user.ID})
   }
   ```

3. **CQRS (Command Query Responsibility Segregation)**
   - Separate read and write models
   - Optimize queries independently

## Why This Architecture?

### Maintainability
- Clear boundaries between layers
- Easy to find and fix bugs
- Self-documenting code structure

### Testability
- Mock any layer easily
- Test business logic without HTTP
- Test data access without business logic

### Flexibility
- Swap implementations (MySQL → PostgreSQL)
- Add features without modifying existing code
- Support multiple interfaces (REST, GraphQL, gRPC)

### Team Collaboration
- Developers can work on different layers simultaneously
- Clear contracts (interfaces)
- Reduced merge conflicts

### Long-term Benefits
- Easier onboarding for new developers
- Reduced technical debt
- Prepared for future requirements

## Common Pitfalls to Avoid

1. **❌ Business Logic in Handlers**
   ```go
   // BAD
   func (h *Handler) CreateUser(c *gin.Context) {
       hashedPassword := hash.Hash(req.Password) // Business logic!
       db.Create(&User{...})                     // Direct DB access!
   }
   ```

2. **❌ Direct Database Access in Services**
   ```go
   // BAD
   func (s *Service) GetUser(id uint) {
       db.First(&user, id) // Should use repository!
   }
   ```

3. **❌ Bypassing Layers**
   ```go
   // BAD
   Handler → Repository (skipping Service)
   Service → Database (skipping Repository)
   ```

4. **❌ Circular Dependencies**
   ```go
   // BAD
   package service
   import "myapp/handler" // Service shouldn't know about handlers!
   ```

## Summary

This architecture provides:
- ✅ Clear separation of concerns
- ✅ Testability at every layer
- ✅ Flexibility for future changes
- ✅ Maintainable codebase
- ✅ SOLID principles
- ✅ Industry best practices

It may seem like "over-engineering" for small projects, but it pays dividends as the project grows and requirements evolve.
