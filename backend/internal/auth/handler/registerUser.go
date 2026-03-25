package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/factory"
)

// RegisterUserHandler wraps the RegisterUser handler with dependencies
type RegisterUserHandler struct {
	userService factory.UserService
}

// NewRegisterUserHandler creates a new RegisterUser handler
func NewRegisterUserHandler(userService factory.UserService) *RegisterUserHandler {
	return &RegisterUserHandler{
		userService: userService,
	}
}

// ServeHTTP handles user registration requests
func (h *RegisterUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot parse Body", http.StatusInternalServerError)
		return
	}

	var requestPayload model.RequestPayload
	err = json.Unmarshal(payload, &requestPayload)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	username := requestPayload.Username
	password := requestPayload.Password

	err = h.userService.RegisterUser(username, password)
	if err != nil {
		switch err {
		case factory.ErrInvalidLength:
			http.Error(w, "Invalid user / password", http.StatusNotAcceptable)
		case factory.ErrUserExists:
			http.Error(w, "User already exist", http.StatusConflict)
		default:
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "New user created",
	})
}

// RegisterUser is a convenience function for backwards compatibility
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// This will be replaced in main.go with proper factory injection
	http.Error(w, "Handler not properly initialized", http.StatusInternalServerError)
}