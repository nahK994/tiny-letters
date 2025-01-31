package db

import "errors"

type JoinPublicationRequest struct {
	UserId        int  `json:"userId"`
	PublicationId int  `json:"publicationId"`
	IsPremium     bool `json:"isPremium"`
}

type RollbackJoinPublicationRequest struct {
	SubscriptionId int `json:"subscriptionId"`
}

type LeavePublicationRequest struct {
	UserId        int `json:"userId"`
	PublicationId int `json:"publicationId"`
}

type RollbackLeavePublicationRequest struct {
	UserId        int  `json:"userId"`
	PublicationId int  `json:"publicationId"`
	IsPremium     bool `json:"isPremium"`
}

type ChangeSubscriberSubscriptionRequest struct {
	UserId        int `json:"userId"`
	PublicationId int `json:"publicationId"`
}

type RollbackChangeSubscriberSubscriptionRequest struct {
	SubscriptionId int `json:"subscriptionId"`
}

type ConfirmPublisherSubscriptionRequest struct {
	UserId int `json:"userId"`
	PlanId int `json:"planId"`
}

type RollbackConfirmPublisherSubscriptionRequest struct {
	SubscriptionId int `json:"subscriptionId"`
}

type RevokePublisherSubscriptionRequest struct {
	UserId int `json:"userId"`
}

type RollbackRevokePublisherSubscriptionRequest struct {
	UserId int `json:"userId"`
	PlanId int `json:"planId"`
}

type ChangePublisherSubscriptionRequest struct {
	UserId        int `json:"userId"`
	ChangedPlanId int `json:"changedPlanId"`
}

type RollbackChangePublisherSubscriptionRequest struct {
	SubscriptionId int `json:"subscriptionId"`
	OldPlanId      int `json:"oldPlanId"`
}

type IsAuthorizedPublisherRequest struct {
	UserId        int `json:"userId"`
	PublicationId int `json:"publicationId"`
}

func (r *ConfirmPublisherSubscriptionRequest) Validate() error {
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
	if r.PublicationId <= 0 {
		return errors.New("publicationId must be greater than 0")
	}
	return nil
}

func (r *ChangePublisherSubscriptionRequest) Validate() error {
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

func (r *RevokePublisherSubscriptionRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	return nil
}

func (r *RollbackJoinPublicationRequest) Validate() error {
	if r.SubscriptionId <= 0 {
		return errors.New("subscription must be greater than 0")
	}
	return nil
}

func (r *LeavePublicationRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.PublicationId <= 0 {
		return errors.New("publicationId must be greater than 0")
	}
	return nil
}

func (r *RollbackLeavePublicationRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.PublicationId <= 0 {
		return errors.New("publicationId must be greater than 0")
	}
	return nil
}

func (r *RollbackChangeSubscriberSubscriptionRequest) Validate() error {
	if r.SubscriptionId <= 0 {
		return errors.New("subscription must be greater than 0")
	}
	return nil
}

func (r *RollbackChangePublisherSubscriptionRequest) Validate() error {
	if r.SubscriptionId <= 0 {
		return errors.New("subscription must be greater than 0")
	}
	if r.OldPlanId <= 0 {
		return errors.New("oldPlanId must be greater than 0")
	}
	return nil
}

func (r *RollbackConfirmPublisherSubscriptionRequest) Validate() error {
	if r.SubscriptionId <= 0 {
		return errors.New("subscription must be greater than 0")
	}
	return nil
}

func (r *RollbackRevokePublisherSubscriptionRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.PlanId <= 0 {
		return errors.New("planId must be greater than 0")
	}
	return nil
}

func (r *IsAuthorizedPublisherRequest) Validate() error {
	if r.UserId <= 0 {
		return errors.New("userId must be greater than 0")
	}
	if r.PublicationId <= 0 {
		return errors.New("publicationId must be greater than 0")
	}
	return nil
}
