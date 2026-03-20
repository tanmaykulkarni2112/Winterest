package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tanmaykulkarni2112/Winterest/backend/data"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/utils"
)


func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", er)
		return
	}
	// reading from the body
	payload , err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot parse Body" , http.StatusInternalServerError)
		return
	}
	var requestpayload model.RequestPayload
	err = json.Unmarshal(payload, &requestpayload)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return 
	}
	username  := requestpayload.Username
	password  := requestpayload.Password
	if len(username) < 8 ||  len(password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(w, "Invalid user / password", er)
		return 
	}

	if _ , ok := data.Users[username]; ok {
		er := http.StatusConflict
		http.Error(w, "User already exist", er)
		return 
	}

	hashPassword , _ := utils.HashPassword(password)
	data.Users[username] = model.Login{
		HashPassword : hashPassword,
	}

	// Save users to JSON file
	err = data.SaveUsersToFile()
	if err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
    "msg": "New user created",
})
}