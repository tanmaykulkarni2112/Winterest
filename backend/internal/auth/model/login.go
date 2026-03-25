package model

type Login struct {
	HashPassword string
	SessionToken string
	CSRFToken    string
}

type RequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}