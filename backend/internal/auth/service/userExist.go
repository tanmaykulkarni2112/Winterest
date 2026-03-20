package service

import (
	"github.com/tanmaykulkarni2112/Winterest/backend/data"
)

func UserExist(user string) bool {
	_ , ok := data.Users[user]
	if !ok {
		return false 
	}
	return true
}