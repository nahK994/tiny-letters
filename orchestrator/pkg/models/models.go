package models

import "encoding/json"

type MessageItem struct {
	Topic string `json:"topic"`
	Data  json.RawMessage
}

type OnboardUserData struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}
