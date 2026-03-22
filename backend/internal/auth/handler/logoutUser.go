package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tanmaykulkarni2112/Winterest/backend/data"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/service"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	// Get the session token from cookie
	sessionCookie, err := r.Cookie("session_token")
	if err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return
	}

	sessionToken := sessionCookie.Value

	// Find the user with this session token
	var logoutUsername string
	for username, user := range data.Users {
		if user.SessionToken == sessionToken {
			logoutUsername = username
			break
		}
	}

	// Verify the user exists and has a valid session
	if logoutUsername == "" {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return
	}

	if err := service.Authorize(logoutUsername, r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return
	}

	// Clear tokens in cookies
	http.SetCookie(w, &http.Cookie{
		Name : "session_token",
		Value : "",
		Expires: time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	http.SetCookie(w, &http.Cookie{
		Name : "csrf_token",
		Value : "",
		Expires: time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	// Clear tokens from the database
	user := data.Users[logoutUsername]
	user.SessionToken = ""
	user.CSRFToken = ""
	data.Users[logoutUsername] = user

	// Save changes to JSON file
	if err := data.SaveUsersToFile(); err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "Logged out successfully",
	})
}

