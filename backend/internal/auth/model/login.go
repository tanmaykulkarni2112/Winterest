package model

type Login struct {
	HashPassword string
	SessionToken string
	CSRFToken    string
}