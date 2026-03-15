package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := HashPass(req.Password)
	if err != nil {
		http.Error(w, "Password hashing failed", http.StatusInternalServerError)
		return
	}

	req.Password = hashedPassword

	err = WriteInData(req)
	if err != nil {
		http.Error(w, "Failed to store user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(Response{
		Message: "User registered successfully",
	})
}

func HashPass(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func WriteInData(payload RegisterRequest) error {

	jsonData, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err
	}

	filePath := "./data/users.json"

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}