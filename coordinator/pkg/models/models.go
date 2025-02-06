package models

import "encoding/json"

type ConfirmationMessage struct {
	Action string `json:"action"`
	Data   json.RawMessage
}

type JoinPublicationData struct {
	UserId        int  `json:"user_id"`
	PublicationId int  `json:"publication_id"`
	PlanType      bool `json:"plan_type"`
}

type LeavePublicationData struct {
	UserId        int `json:"user_id"`
	PublicationId int `json:"publication_id"`
}

type ChangeSubscriberSubscriptionData struct {
	UserId        int `json:"user_id"`
	PublicationId int `json:"publication_id"`
}
