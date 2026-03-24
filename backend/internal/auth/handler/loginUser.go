package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/factory"
)

// LoginUserHandler wraps the LoginUser handler with dependencies
type LoginUserHandler struct {
	userService factory.UserService
	authService factory.AuthService
}

// NewLoginUserHandler creates a new LoginUser handler
func NewLoginUserHandler(userService factory.UserService, authService factory.AuthService) *LoginUserHandler {
	return &LoginUserHandler{
		userService: userService,
		authService: authService,
	}
}

// ServeHTTP handles user login requests
func (h *LoginUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	res, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var loginPayload model.RequestPayload
	err = json.Unmarshal(res, &loginPayload)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	user := loginPayload.Username

	sessionToken, csrfToken, err := h.userService.LoginUser(user, loginPayload.Password)
	if err != nil {
		http.Error(w, "Invalid user or password", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Minute),
		HttpOnly: false,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Minute),
		HttpOnly: false,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "User login successful",
	})
}

// LoginUser is a convenience function for backwards compatibility
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// This will be replaced in main.go with proper factory injection
	http.Error(w, "Handler not properly initialized", http.StatusInternalServerError)
}
