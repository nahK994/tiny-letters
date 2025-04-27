package models

import "encoding/json"

type MessageItem struct {
	Topic string `json:"action"`
	Msg   json.RawMessage
}

type SubscriberRegistration struct {
	UserId int `json:"user_id"`
	Email  int `json:"email"`
}

type JoinPublication struct {
	UserId        int  `json:"user_id"`
	PublicationId int  `json:"publication_id"`
	PlanType      bool `json:"plan_type"`
}

type LeavePublication struct {
	UserId        int `json:"user_id"`
	PublicationId int `json:"publication_id"`
}

type ChangeSubscriberSubscription struct {
	UserId        int `json:"user_id"`
	PublicationId int `json:"publication_id"`
}

type PublisherRegistration struct {
	UserId int `json:"user_id"`
	Email  int `json:"email"`
	PlanId int `json:"plan_id"`
}

type PublisherSubscription struct {
	UserId int `json:"user_id"`
	PlanId int `json:"plan_id"`
}

type PublisherUnsubscription struct {
	UserId int `json:"user_id"`
}

type ChangePublisherSubscription struct {
	UserId    int `json:"user_id"`
	NewPlanId int `json:"new_plan_id"`
	OldPlanId int `json:"old_plan_id"`
}

type PublishLetter struct {
	Content       string  `json:"content"`
	SubscriberIds []int32 `json:"subscriber_ids"`
}
