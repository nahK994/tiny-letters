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

type PublisherRegistrationRequest struct {
	Email              string   `json:"email"`
	Password           string   `json:"password"`
	Name               string   `json:"name"`
	Roles              []string `json:"roles"`
	SubscriptionPlanId int      `json:"subscriptionPlanId"`
}
