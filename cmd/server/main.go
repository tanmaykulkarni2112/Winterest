package main

import (
	"net/http"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth"
	"github.com/tanmaykulkarni2112/Winterest/backend/internal/home"
)

func main() {
	http.HandleFunc("/home", home.HomeFunc)
	http.HandleFunc("/api/register", auth.RegisterUser)
	http.ListenAndServe(":8080", nil)
}