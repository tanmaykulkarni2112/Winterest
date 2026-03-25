package home

import (
	"encoding/json"
	"net/http"
)

// HomeHandler handles the home endpoint
type HomeHandler struct{}

// NewHomeHandler creates a new HomeHandler
func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

// ServeHTTP handles home requests
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Message": "Hello world"})
}

// HomeFunc is a convenience function for backwards compatibility
var HomeFunc = func(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Message": "Hello world"})
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}