package service

import (
	"errors"
	"net/http"

	"github.com/tanmaykulkarni2112/Winterest/backend/data"
)

var AuthError = errors.New("Unauthorized")

func Authorize(username string, r *http.Request) error {
	user , ok := data.Users[username]
	if !ok {
		return AuthError
	}
	st , err := r.Cookie("session_token")
	if err != nil || st.Value != user.SessionToken{
		return AuthError
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		return AuthError
	}
	return nil
	
}
