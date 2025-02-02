package models

type (
	ActionType string
	IsPremium  bool
)

type PublicationMessage struct {
	PublisherId int
	Content     string
}

type PublicationBroadcastMessage struct {
	SubscriberIds []int
	Content       string
}

type PublisherConfirmationMessage struct {
	UserId      int        `json:"user_id"`
	Action      ActionType `json:"action"`
	PlanName    string     `json:"plan_name,omitempty"`
	OldPlanName string     `json:"old_plan_name,omitempty"`
	NewPlanName string     `json:"new_plan_name,omitempty"`
}

type SubscriberConfirmationMessage struct {
	UserId      int        `json:"user_id"`
	Action      ActionType `json:"action"`
	PlanType    IsPremium  `json:"plan_type,omitempty"`
	OldPlanType IsPremium  `json:"old_plan_type,omitempty"`
	NewPlanType IsPremium  `json:"new_plan_type,omitempty"`
}
