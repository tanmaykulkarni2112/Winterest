package factory

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
)

// MockDataService is a mock implementation for testing
type MockDataService struct {
	Users map[string]model.Login
}

func (m *MockDataService) GetUser(username string) (model.Login, bool) {
	user, ok := m.Users[username]
	return user, ok
}

func (m *MockDataService) UserExists(username string) bool {
	_, ok := m.Users[username]
	return ok
}

func (m *MockDataService) SaveUser(username string, user model.Login) error {
	m.Users[username] = user
	return nil
}

func (m *MockDataService) LoadUsersFromFile() error {
	return nil
}

func (m *MockDataService) SaveUsersToFile() error {
	return nil
}

func (m *MockDataService) UpdateUser(username string, user model.Login) error {
	m.Users[username] = user
	return nil
}

// Example test showing how factory pattern improves testability
func TestUserServiceWithMockData(t *testing.T) {
	// Create mock data service
	mockData := &MockDataService{
		Users: make(map[string]model.Login),
	}

	// Create real auth service with mock data
	authService := NewAuthService(mockData)

	// Create user service
	userService := NewUserService(mockData, authService)

	// Test registration
	err := userService.RegisterUser("testuser123", "password123")
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	// Verify user was saved
	if !mockData.UserExists("testuser123") {
		t.Fatal("User was not saved")
	}

	// Test login
	sessionToken, csrfToken, err := userService.LoginUser("testuser123", "password123")
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	if sessionToken == "" || csrfToken == "" {
		t.Fatal("Tokens were not generated")
	}

	// Verify tokens are stored
	user, _ := mockData.GetUser("testuser123")
	if user.SessionToken != sessionToken {
		t.Fatal("Session token not stored correctly")
	}
}

// Example test for authorization
func TestAuthorizeWithValidToken(t *testing.T) {
	mockData := &MockDataService{
		Users: map[string]model.Login{
			"testuser": {
				HashPassword: "$2a$10$mocked_hash",
				SessionToken: "valid_session_token",
				CSRFToken:    "valid_csrf_token",
			},
		},
	}

	authService := NewAuthService(mockData)

	// Create mock request with cookies
	req := httptest.NewRequest("POST", "/test", nil)
	req.AddCookie(&http.Cookie{
		Name:  "session_token",
		Value: "valid_session_token",
	})
	req.Header.Set("X-CSRF-Token", "valid_csrf_token")

	// Test authorization
	err := authService.Authorize("testuser", req)
	if err != nil {
		t.Fatalf("Authorization failed: %v", err)
	}
}

// Example test for invalid token
func TestAuthorizeWithInvalidToken(t *testing.T) {
	mockData := &MockDataService{
		Users: map[string]model.Login{
			"testuser": {
				SessionToken: "valid_session_token",
				CSRFToken:    "valid_csrf_token",
			},
		},
	}

	authService := NewAuthService(mockData)

	req := httptest.NewRequest("POST", "/test", nil)
	req.AddCookie(&http.Cookie{
		Name:  "session_token",
		Value: "invalid_session_token",
	})

	err := authService.Authorize("testuser", req)
	if err != AuthError {
		t.Fatalf("Expected AuthError, got %v", err)
	}
}
