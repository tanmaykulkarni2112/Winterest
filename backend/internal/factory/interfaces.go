package factory

import (
	"net/http"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
)

// DataService defines the interface for user data operations
type DataService interface {
	GetUser(username string) (model.Login, bool)
	UserExists(username string) bool
	SaveUser(username string, user model.Login) error
	LoadUsersFromFile() error
	SaveUsersToFile() error
	UpdateUser(username string, user model.Login) error
}

// AuthService defines the interface for authentication operations
type AuthService interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hash string) bool
	GenerateToken(length int) string
	Authorize(username string, r *http.Request) error
}

// UserService defines the interface for user-related operations
type UserService interface {
	RegisterUser(username, password string) error
	LoginUser(username, password string) (sessionToken, csrfToken string, err error)
	LogoutUser(username string) error
	GetUserBySessionToken(sessionToken string) (string, bool)
}

// HandlerDependencies groups all dependencies for HTTP handlers
type HandlerDependencies struct {
	DataService DataService
	AuthService AuthService
	UserService UserService
}
