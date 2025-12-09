package user

// Repository defines the interface for user data operations
// This follows the Repository pattern and allows for easy mocking in tests
type Repository interface {
	// GetAll retrieves all users from the database
	GetAll() ([]User, error)

	// GetByID retrieves a user by their ID
	GetByID(id uint) (*User, error)

	// GetByEmail retrieves a user by their email address
	GetByEmail(email string) (*User, error)

	// Create creates a new user in the database
	Create(user *User) error

	// Update updates an existing user
	Update(user *User) error

	// Delete soft deletes a user
	Delete(id uint) error
}
