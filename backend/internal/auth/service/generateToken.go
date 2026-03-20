package service

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _ , err := rand.Read(bytes); err != nil {
		log.Fatalln("Failed to generate token:", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}