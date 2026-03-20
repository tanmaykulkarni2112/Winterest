package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tanmaykulkarni2112/Winterest/backend/data"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/service"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed 
		http.Error(w, "Invalid request method", er)
		return 
	}

	res , err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	
	var loginPayload model.RequestPayload
	err = json.Unmarshal(res, &loginPayload)
	if err != nil {
		panic(err)
	}
	user := loginPayload.Username
	userExists := service.UserExist(user)
	
	if !userExists {
		er := http.StatusUnauthorized
		http.Error(w, "Invalid user or password", er)
		return
	}

	// Get the stored hash for this user
	storedUser := data.Users[user]
	if !service.CheckPassword(loginPayload.Password, storedUser.HashPassword) {
		er := http.StatusUnauthorized
		http.Error(w, "Invalid user or password", er)
		return
	}
	fmt.Fprintln(w, "User login successful")
	sessionToken := service.GenerateToken(32)
	csrfToken := service.GenerateToken(32)
	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		Expires: time.Now().Add(24 * time.Minute),
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name: "csrf_token",
		Value: csrfToken,
		Expires: time.Now().Add(24 * time.Minute),
		HttpOnly: false,
	})
	
	// Get existing user data and update with new tokens
	existingUser := data.Users[user]
	existingUser.SessionToken = sessionToken
	existingUser.CSRFToken = csrfToken
	data.Users[user] = existingUser

	// Save users to JSON file
	err = data.SaveUsersToFile()
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}
}
