package db

import "errors"

type PublisherSubscriptionRequest struct {
	UserId int `json:"userId"`
	PlanId int `json:"planId"`
}

type JoinPublicationRequest struct {
	UserId        int `json:"userId"`
	PlanId        int `json:"planId"`
	PublicationId int `json:"publicationId"`
}

type ChangePublisherPlanRequest struct {
	UserId        int `json:"userId"`
	ChangedPlanId int `json:"changedPlanId"`
}

type ChangePublicationSubscriptionRequest struct {
	UserId        int `json:"userId"`
	PublicationId int `json:"publicationId"`
	ChangedPlanId int `json:"changedPlanId"`
}

type UnsubscribePublisherRequest struct {
	UserId int `json:"userId"`
}

type LeavePublicationRequest struct {
	UserId        int `json:"userId"`
	PlanId        int `json:"planId"`
	PublicationId int `json:"publicationId"`
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

func (r *JoinPublicationRequest) Validate() error {
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

func (r *ChangePublisherPlanRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.ChangedPlanId <= 0 {
		return errors.New("changedPlanId must be greater than 0")
	}
	return nil
}

func (r *ChangePublicationSubscriptionRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.ChangedPlanId <= 0 {
		return errors.New("changedPlanId must be greater than 0")
	}
	if r.PublicationId <= 0 {
		return errors.New("publicationId must be greater than 0")
	}
	return nil
}

func (r *UnsubscribePublisherRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	return nil
}

func (r *LeavePublicationRequest) Validate() error {
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
