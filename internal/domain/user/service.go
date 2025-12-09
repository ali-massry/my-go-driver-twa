package user

// Service defines the interface for user business logic
// This follows the Service pattern and separates business logic from HTTP handlers
type Service interface {
	// Authentication
	// Register creates a new user account and returns auth token
	Register(req RegisterRequest) (*AuthResponse, error)

	// Login authenticates a user and returns auth token
	Login(req LoginRequest) (*AuthResponse, error)

	// User Management
	// GetAllUsers retrieves all users
	GetAllUsers() ([]User, error)

	// GetUserByID retrieves a user by their ID
	GetUserByID(id uint) (*User, error)

	// CreateUser creates a new user with hashed password
	CreateUser(req CreateUserRequest) (*User, error)

	// UpdateUser updates an existing user
	UpdateUser(id uint, req UpdateUserRequest) (*User, error)

	// DeleteUser soft deletes a user
	DeleteUser(id uint) error
}
