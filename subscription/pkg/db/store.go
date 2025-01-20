package db

import (
	"fmt"
)

func (r *Repository) SubscribePublisherPlan(data *PublisherSubscriptionRequest) error {
	query := `
	INSERT INTO publisher_subscriptions (user_id, plan_id)
	VALUES ($1, $2)
	`
	_, err := r.DB.Exec(query, data.UserId, data.PlanId)
	if err != nil {
		return fmt.Errorf("failed to subscribe publisher to plan: %w", err)
	}
	return nil
}

func (r *Repository) UnsubscribePublisherPlan(data *RevokePublisherSubscribeRequest) error {
	query := `
	DELETE FROM publisher_subscriptions WHERE user_id = $1
	`
	if _, err := r.DB.Exec(query, data.UserId); err != nil {
		return fmt.Errorf("failed to revoke publisher subscription: %w", err)
	}
	return nil
}

func (r *Repository) ChangePublisherPlan(data *ChangePublisherPlanRequest) error {
	query := `
	UPDATE publisher_subscriptions SET plan_id = $2 WHERE user_id = $1
	`
	if _, err := r.DB.Exec(query, data.UserId, data.ChangedPlanId); err != nil {
		return fmt.Errorf("failed to change publisher subscription plan: %w", err)
	}
	return nil
}

func (r *Repository) JoinPublication(data *SubscriberSubscriptionRequest) error {
	query := `
	INSERT INTO subscriber_subscriptions (user_id, publication_id, is_premium)
	VALUES ($1, $2, $3)
	`
	_, err := r.DB.Exec(query, data.UserId, data.PublicationId, data.IsPremium)
	if err != nil {
		return fmt.Errorf("failed to join publication: %w", err)
	}
	return nil
}

func (r *Repository) LeavePublication(data *RevokeSubscriberSubscriptionRequest) error {
	query := `
	DELETE FROM subscriber_subscriptions WHERE user_id = $1 AND publication_id = $2
	`
	if _, err := r.DB.Exec(query, data.UserId, data.PublicationId); err != nil {
		return fmt.Errorf("failed to leave publication: %w", err)
	}
	return nil
}

func (r *Repository) ChangeSubscriberPlan(data *ChangeSubscriberSubscriptionRequest) error {
	query := `
	SELECT is_premium FROM subscriber_subscriptions WHERE user_id = $1 AND publication_id = $2
	`
	var isPremium bool
	if err := r.DB.QueryRow(query, data.UserId, data.PublicationId).Scan(&isPremium); err != nil {
		return fmt.Errorf("failed to get subscription premium status: %w", err)
	}

	query = `
	UPDATE subscriber_subscriptions SET is_premium = $3 WHERE user_id = $1 AND publication_id = $2
	`
	if _, err := r.DB.Exec(query, data.UserId, data.PublicationId, !isPremium); err != nil {
		return fmt.Errorf("failed to update subscription premium status: %w", err)
	}
	return nil
}
