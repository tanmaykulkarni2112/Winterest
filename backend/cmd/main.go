package main

import (
	"log"
	"net/http"

	"github.com/tanmaykulkarni2112/Winterest/backend/data"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/handler"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/home"
)

func main() {
	// Load users from JSON file on startup
	err := data.LoadUsersFromFile()
	if err != nil {
		log.Println("Warning: Could not load users from file:", err)
	}

	http.HandleFunc("/home", home.HomeFunc)
	http.HandleFunc("/register", handler.RegisterUser)
	http.HandleFunc("/login", handler.LoginUser)
	http.HandleFunc("/logout", handler.LogoutUser)
	http.HandleFunc("/protected", handler.Protected)
	http.ListenAndServe(":8080", nil)
}