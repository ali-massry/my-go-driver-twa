package service

import (
	"errors"
	"strings"

	"github.com/ali-massry/my-go-driver/internal/domain/user"
	"github.com/ali-massry/my-go-driver/pkg/hash"
	jwtpkg "github.com/ali-massry/my-go-driver/pkg/jwt"
	"gorm.io/gorm"
)

// userService implements the user.Service interface
type userService struct {
	repo       user.Repository
	jwtManager *jwtpkg.Manager
}

// NewUserService creates a new user service
func NewUserService(repo user.Repository, jwtManager *jwtpkg.Manager) user.Service {
	return &userService{
		repo:       repo,
		jwtManager: jwtManager,
	}
}

// Register creates a new user account and returns auth token
func (s *userService) Register(req user.RegisterRequest) (*user.AuthResponse, error) {
	// Hash password
	hashedPassword, err := hash.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user entity
	usr := &user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// Save to database
	if err := s.repo.Create(usr); err != nil {
		// Check if it's a duplicate email error
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}

	// Generate JWT token
	token, err := s.jwtManager.Generate(usr.ID, usr.Email)
	if err != nil {
		return nil, err
	}

	return &user.AuthResponse{
		User:  usr.ToUserResponse(),
		Token: token,
	}, nil
}

// Login authenticates a user and returns auth token
func (s *userService) Login(req user.LoginRequest) (*user.AuthResponse, error) {
	// Find user by email
	usr, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	// Verify password
	if !hash.Check(req.Password, usr.Password) {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.jwtManager.Generate(usr.ID, usr.Email)
	if err != nil {
		return nil, err
	}

	return &user.AuthResponse{
		User:  usr.ToUserResponse(),
		Token: token,
	}, nil
}

// GetAllUsers retrieves all users
func (s *userService) GetAllUsers() ([]user.User, error) {
	return s.repo.GetAll()
}

// GetUserByID retrieves a user by their ID
func (s *userService) GetUserByID(id uint) (*user.User, error) {
	return s.repo.GetByID(id)
}

// CreateUser creates a new user with hashed password
func (s *userService) CreateUser(req user.CreateUserRequest) (*user.User, error) {
	// Hash password
	hashedPassword, err := hash.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user entity
	usr := &user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// Save to database
	if err := s.repo.Create(usr); err != nil {
		// Check if it's a duplicate email error
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}

	return usr, nil
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(id uint, req user.UpdateUserRequest) (*user.User, error) {
	// Get existing user
	usr, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		usr.Name = req.Name
	}
	if req.Email != "" {
		usr.Email = req.Email
	}

	// Save changes
	if err := s.repo.Update(usr); err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}

	return usr, nil
}

// DeleteUser soft deletes a user
func (s *userService) DeleteUser(id uint) error {
	// Check if user exists
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	return s.repo.Delete(id)
}
