package db

import "errors"

type PublisherSubscriptionRequest struct {
	UserId int `json:"userId"`
	PlanId int `json:"planId"`
}

type SubscriberSubscriptionRequest struct {
	UserId        int `json:"userId"`
	PlanId        int `json:"planId"`
	PublicationId int `json:"publicationId"`
}

type PublisherChangePlanRequest struct {
	UserId        int `json:"userId"`
	ChangedPlanId int `json:"changedPlanId"`
}

type SubscriberChangePlanRequest struct {
	UserId        int `json:"userId"`
	PublicationId int `json:"publicationId"`
	ChangedPlanId int `json:"changedPlanId"`
}

func (r PublisherSubscriptionRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.PlanId <= 0 {
		return errors.New("planId must be greater than 0")
	}
	return nil
}

func (r SubscriberSubscriptionRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.PlanId <= 0 {
		return errors.New("planId must be greater than 0")
	}
	if r.PublicationId <= 0 {
		return errors.New("publicationId must be greater than 0")
	}
	return nil
}

func (r PublisherChangePlanRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.ChangedPlanId <= 0 {
		return errors.New("changedPlanId must be greater than 0")
	}
	return nil
}

func (r SubscriberChangePlanRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.ChangedPlanId <= 0 {
		return errors.New("changedPlanId must be greater than 0")
	}
	return nil
}
