package hash

import "golang.org/x/crypto/bcrypt"

// Hash hashes a password using bcrypt
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Check compares a hashed password with a plain text password
func Check(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword is an alias for Hash
func HashPassword(password string) (string, error) {
	return Hash(password)
}

// CheckPasswordHash is an alias for Check
func CheckPasswordHash(password, hash string) bool {
	return Check(password, hash)
}
