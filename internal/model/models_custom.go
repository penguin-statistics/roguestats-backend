package model

type LoginInput struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	TurnstileResponse string `json:"turnstileResponse"`
}
