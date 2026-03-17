package main

import (
	"net/http"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/home"
)

func main() {
	http.HandleFunc("/home", home.HomeFunc)
	http.ListenAndServe(":8080", nil)
}