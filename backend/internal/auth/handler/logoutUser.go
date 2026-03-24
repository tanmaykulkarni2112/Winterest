package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/factory"
)

// LogoutUserHandler wraps the LogoutUser handler with dependencies
type LogoutUserHandler struct {
	userService factory.UserService
	authService factory.AuthService
}

// NewLogoutUserHandler creates a new LogoutUser handler
func NewLogoutUserHandler(userService factory.UserService, authService factory.AuthService) *LogoutUserHandler {
	return &LogoutUserHandler{
		userService: userService,
		authService: authService,
	}
}

// ServeHTTP handles user logout requests
func (h *LogoutUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the session token from cookie
	sessionCookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	sessionToken := sessionCookie.Value

	// Find the user with this session token
	logoutUsername, ok := h.userService.GetUserBySessionToken(sessionToken)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Verify the user has valid tokens
	if err := h.authService.Authorize(logoutUsername, r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Clear tokens from the user
	err = h.userService.LogoutUser(logoutUsername)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	// Clear tokens in cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "Logged out successfully",
	})
}

// LogoutUser is a convenience function for backwards compatibility
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	// This will be replaced in main.go with proper factory injection
	http.Error(w, "Handler not properly initialized", http.StatusInternalServerError)
}

