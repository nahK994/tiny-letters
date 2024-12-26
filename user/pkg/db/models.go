package db

import "github.com/lib/pq"

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
	Email          string   `json:"email"`
	Password       string   `json:"password"`
	Name           string   `json:"name"`
	Roles          []string `json:"roles"`
	SubscriptionId int      `json:"subscriptionId"`
}

type GenerateTokenRequest struct {
	Id             int
	Roles          []string
	SubscriptionId int
}

type JWT_claim struct {
	Id             int
	SubscriptionId int
	Password       string
	Roles          pq.StringArray
}
