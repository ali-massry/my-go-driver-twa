package repository

import (
	"github.com/ali-massry/my-go-driver/internal/domain/user"
	"gorm.io/gorm"
)

// userRepository implements the user.Repository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

// GetAll retrieves all users from the database
func (r *userRepository) GetAll() ([]user.User, error) {
	var users []user.User
	err := r.db.Find(&users).Error
	return users, err
}

// GetByID retrieves a user by their ID
func (r *userRepository) GetByID(id uint) (*user.User, error) {
	var usr user.User
	err := r.db.First(&usr, id).Error
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

// GetByEmail retrieves a user by their email address
func (r *userRepository) GetByEmail(email string) (*user.User, error) {
	var usr user.User
	err := r.db.Where("email = ?", email).First(&usr).Error
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

// Create creates a new user in the database
func (r *userRepository) Create(usr *user.User) error {
	return r.db.Create(usr).Error
}

// Update updates an existing user
func (r *userRepository) Update(usr *user.User) error {
	return r.db.Save(usr).Error
}

// Delete soft deletes a user
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&user.User{}, id).Error
}
