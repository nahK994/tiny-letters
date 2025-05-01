package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegistration struct {
	Email    string
	Password string
	Role     string
}

type GenerateTokenRequest struct {
	Id   int
	Role string
}

type JWT_claim struct {
	Id       int
	Password string
	Role     string
}
