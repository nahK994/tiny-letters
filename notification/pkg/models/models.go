package models

import "encoding/json"

type MessageItem struct {
	Topic string `json:"action"`
	Msg   json.RawMessage
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

type ConfirmPublisherSubscriptionData struct {
	UserId int `json:"user_id"`
	PlanId int `json:"plan_id"`
}

type RevokePublisherSubscriptionData struct {
	UserId int `json:"user_id"`
}

type ChangePublisherSubscriptionData struct {
	UserId    int `json:"user_id"`
	NewPlanId int `json:"new_plan_id"`
	OldPlanId int `json:"old_plan_id"`
}

type ContentData struct {
	ContentId int    `json:"content_id"`
	Content   string `json:"content"`
}

type ConsumedContentMessage struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type PublishedContentMessage struct {
	Content       string  `json:"content"`
	SubscriberIds []int32 `json:"subscriber_ids"`
}
