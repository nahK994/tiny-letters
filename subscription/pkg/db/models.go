package db

import "errors"

type SubscriberSubscriptionRequest struct {
	UserId        int  `json:"userId"`
	PublicationId int  `json:"publicationId"`
	IsPremium     bool `json:"isPremium"`
}

type RevokeSubscriberSubscriptionRequest struct {
	UserId        int `json:"userId"`
	PublicationId int `json:"publicationId"`
}

type ChangeSubscriberSubscriptionRequest struct {
	UserId        int `json:"userId"`
	PublicationId int `json:"publicationId"`
}

type ChangePublisherPlanRequest struct {
	UserId        int `json:"userId"`
	ChangedPlanId int `json:"changedPlanId"`
}

type PublisherSubscriptionRequest struct {
	UserId int `json:"userId"`
	PlanId int `json:"planId"`
}

type RevokePublisherSubscribeRequest struct {
	UserId int `json:"userId"`
}

func (r *PublisherSubscriptionRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.PlanId <= 0 {
		return errors.New("planId must be greater than 0")
	}
	return nil
}

func (r *SubscriberSubscriptionRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.PublicationId <= 0 {
		return errors.New("publicationId must be greater than 0")
	}
	return nil
}

func (r *ChangePublisherPlanRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.ChangedPlanId <= 0 {
		return errors.New("changedPlanId must be greater than 0")
	}
	return nil
}

func (r *ChangeSubscriberSubscriptionRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.PublicationId <= 0 {
		return errors.New("publicationId must be greater than 0")
	}
	return nil
}

func (r *RevokePublisherSubscribeRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	return nil
}

func (r *RevokeSubscriberSubscriptionRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.PublicationId <= 0 {
		return errors.New("publicationId must be greater than 0")
	}
	return nil
}
