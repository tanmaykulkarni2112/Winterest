package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/service"
)

func Protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Cannot access", er)
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
	if err := service.Authorize(username, r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "Access granted to protected resource",
		"username": username,
	})
}