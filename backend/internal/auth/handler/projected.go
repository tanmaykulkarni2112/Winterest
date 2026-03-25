package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/factory"
)

// ProtectedHandler wraps the Protected handler with dependencies
type ProtectedHandler struct {
	authService factory.AuthService
}

// NewProtectedHandler creates a new Protected handler
func NewProtectedHandler(authService factory.AuthService) *ProtectedHandler {
	return &ProtectedHandler{
		authService: authService,
	}
}

// ServeHTTP handles requests to protected resources
func (h *ProtectedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Cannot access", http.StatusMethodNotAllowed)
		return
	}

	reqPayload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var requestPayload model.RequestPayload
	err = json.Unmarshal(reqPayload, &requestPayload)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	username := requestPayload.Username
	if err := h.authService.Authorize(username, r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg":      "Access granted to protected resource",
		"username": username,
	})
}

// Protected is a convenience function for backwards compatibility
func Protected(w http.ResponseWriter, r *http.Request) {
	// This will be replaced in main.go with proper factory injection
	http.Error(w, "Handler not properly initialized", http.StatusInternalServerError)
}