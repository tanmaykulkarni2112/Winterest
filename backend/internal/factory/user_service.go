package factory

import (
	"errors"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
)

// UserServiceImpl implements the UserService interface
type UserServiceImpl struct {
	dataService DataService
	authService AuthService
}

// NewUserService creates a new UserService instance
func NewUserService(dataService DataService, authService AuthService) UserService {
	return &UserServiceImpl{
		dataService: dataService,
		authService: authService,
	}
}

var (
	ErrUserExists      = errors.New("user already exists")
	ErrInvalidPassword = errors.New("invalid username or password")
	ErrInvalidLength   = errors.New("username and password must be at least 8 characters")
	ErrUnauthorized    = errors.New("unauthorized")
)

// RegisterUser registers a new user
func (us *UserServiceImpl) RegisterUser(username, password string) error {
	if len(username) < 8 || len(password) < 8 {
		return ErrInvalidLength
	}

	if us.dataService.UserExists(username) {
		return ErrUserExists
	}

	hashPassword, err := us.authService.HashPassword(password)
	if err != nil {
		return err
	}

	user := model.Login{
		HashPassword: hashPassword,
	}

	return us.dataService.SaveUser(username, user)
}

// LoginUser logs in a user and returns session and CSRF tokens
func (us *UserServiceImpl) LoginUser(username, password string) (sessionToken, csrfToken string, err error) {
	if !us.dataService.UserExists(username) {
		return "", "", ErrInvalidPassword
	}

	storedUser, _ := us.dataService.GetUser(username)
	if !us.authService.CheckPassword(password, storedUser.HashPassword) {
		return "", "", ErrInvalidPassword
	}

	sessionToken = us.authService.GenerateToken(32)
	csrfToken = us.authService.GenerateToken(32)

	storedUser.SessionToken = sessionToken
	storedUser.CSRFToken = csrfToken

	err = us.dataService.UpdateUser(username, storedUser)
	if err != nil {
		return "", "", err
	}

	return sessionToken, csrfToken, nil
}

// LogoutUser clears the session for a user
func (us *UserServiceImpl) LogoutUser(username string) error {
	user, ok := us.dataService.GetUser(username)
	if !ok {
		return ErrUnauthorized
	}

	user.SessionToken = ""
	user.CSRFToken = ""

	return us.dataService.UpdateUser(username, user)
}

// GetUserBySessionToken finds a user by their session token
func (us *UserServiceImpl) GetUserBySessionToken(sessionToken string) (string, bool) {
	if sessionToken == "" {
		return "", false
	}

	// This is a simple implementation; in production, you'd want an index
	allUsers := us.dataService.(*DataServiceImpl).GetAllUsers()
	for username, user := range allUsers {
		if user.SessionToken == sessionToken {
			return username, true
		}
	}

	return "", false
}
