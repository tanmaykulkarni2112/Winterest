package factory

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// AuthServiceImpl implements the AuthService interface
type AuthServiceImpl struct {
	dataService DataService
}

// NewAuthService creates a new AuthService instance
func NewAuthService(dataService DataService) AuthService {
	return &AuthServiceImpl{
		dataService: dataService,
	}
}

var AuthError = errors.New("Unauthorized")

// HashPassword hashes a password using bcrypt
func (as *AuthServiceImpl) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPassword compares a password with its hash
func (as *AuthServiceImpl) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken generates a random token of specified length
func (as *AuthServiceImpl) GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}

// Authorize verifies user session and CSRF tokens
func (as *AuthServiceImpl) Authorize(username string, r *http.Request) error {
	user, ok := as.dataService.GetUser(username)
	if !ok {
		return AuthError
	}

	st, err := r.Cookie("session_token")
	if err != nil || st.Value != user.SessionToken {
		return AuthError
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		return AuthError
	}

	return nil
}
