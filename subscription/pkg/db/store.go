package db

import (
	"fmt"
)

func (r *Repository) SubscribePublisherPlan(data PublisherSubscriptionRequest) error {
	query := `
	INSERT INTO publisher_subscription_managements (user_id, subscription_id)
	VALUES ($1, $2)
	`
	if _, err := r.DB.Exec(query, data.UserId, data.PlanId); err != nil {
		return fmt.Errorf("failed to subscribe publisher to plan: %w", err)
	}
	return nil
}

func (r *Repository) SubscribeSubscriberPlan(data SubscriberSubscriptionRequest) error {
	query := `
	INSERT INTO subscriber_subscription_managements (user_id, subscription_id, publication_id)
	VALUES ($1, $2, $3)
	`
	_, err := r.DB.Exec(query, data.UserId, data.PlanId, data.PublicationId)
	if err != nil {
		return fmt.Errorf("failed to subscribe subscriber to plan: %w", err)
	}
	return nil
}

func (r *Repository) ChangeSubscriberSubscriptionPlan(data SubscriberChangePlanRequest) error {
	query := `
	UPDATE subscriber_subscription_managements
	SET subscription_id = $1
	WHERE user_id = $2 AND publication_id = $3
	`
	if _, err := r.DB.Exec(query, data.ChangedPlanId, data.UserId, data.PublicationId); err != nil {
		return fmt.Errorf("failed to change subscriber subscription plan: %w", err)
	}
	return nil
}

func (r *Repository) ChangePublisherSubscriptionPlan(data PublisherChangePlanRequest) error {
	query := `
	UPDATE subscriber_subscription_managements
	SET subscription_id = $1
	WHERE user_id = $2
	`
	if _, err := r.DB.Exec(query, data.ChangedPlanId, data.UserId); err != nil {
		return fmt.Errorf("failed to change publication subscription plan: %w", err)
	}
	return nil
}
