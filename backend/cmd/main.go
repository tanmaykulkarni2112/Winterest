package main

import (
	"log"
	"net/http"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/handler"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/factory"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/home"
)

func main() {
	// Initialize the factory with the data file path
	factory := factory.NewFactory("data/users.json")

	// Load users from JSON file on startup
	err := factory.Initialize()
	if err != nil {
		log.Println("Warning: Could not load users from file:", err)
	}

	// Get service instances from the factory
	authService := factory.GetAuthService()
	userService := factory.GetUserService()

	// Create handlers with injected dependencies
	homeHandler := home.NewHomeHandler()
	registerHandler := handler.NewRegisterUserHandler(userService)
	loginHandler := handler.NewLoginUserHandler(userService, authService)
	logoutHandler := handler.NewLogoutUserHandler(userService, authService)
	protectedHandler := handler.NewProtectedHandler(authService)

	// Register HTTP handlers
	http.Handle("/home", homeHandler)
	http.Handle("/register", registerHandler)
	http.Handle("/login", loginHandler)
	http.Handle("/logout", logoutHandler)
	http.Handle("/protected", protectedHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}