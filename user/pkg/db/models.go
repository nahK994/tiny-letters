package db

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	Name             string `json:"name"`
	Role             string `json:"role"`
	SubscriptionType int    `json:"subscriptionType"`
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
