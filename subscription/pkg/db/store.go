package db

import (
	"fmt"
)

type (
	oldPublisherSubscriptionPlanId    int
	subscriptionId                    int
	oldSubscriberSubscriptionPlanType bool
)

func (r *Repository) ConfirmPublisherSubscription(data *ConfirmPublisherSubscriptionRequest) (subscriptionId, error) {
	var id int
	query := `
	INSERT INTO publisher_subscriptions (user_id, plan_id)
	VALUES ($1, $2) RETURNING id
	`
	err := r.DB.QueryRow(query, data.UserId, data.PlanId).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("failed to subscribe publisher to plan: %w", err)
	}
	return subscriptionId(id), nil
}

func (r *Repository) RollbackConfirmPublisherSubscription(data *RollbackConfirmPublisherSubscriptionRequest) error {
	query := `
	DELETE FROM publisher_subscriptions WHERE id = $1
	`
	_, err := r.DB.Exec(query, data.SubscriptionId)
	if err != nil {
		return fmt.Errorf("failed to subscribe publisher to plan: %w", err)
	}
	return nil
}

func (r *Repository) RevokePublisherSubscription(data *RevokePublisherSubscriptionRequest) (oldPublisherSubscriptionPlanId, error) {
	var oldPlanId int
	query := `
	DELETE FROM publisher_subscriptions WHERE user_id = $1 RETURNING plan_id
	`
	err := r.DB.QueryRow(query, data.UserId).Scan(&oldPlanId)
	if err != nil {
		return -1, fmt.Errorf("failed to revoke publisher subscription: %w", err)
	}
	return oldPublisherSubscriptionPlanId(oldPlanId), nil
}

func (r *Repository) RollbackRevokePublisherSubscription(data *RollbackRevokePublisherSubscriptionRequest) error {
	query := `
	INSERT INTO publisher_subscriptions (user_id, plan_id) VALUES ($1, $2)
	`

	_, err := r.DB.Exec(query, data.UserId, data.PlanId)
	if err != nil {
		return fmt.Errorf("failed to rollback revoke publisher subscription: %w", err)
	}
	return nil
}

func (r *Repository) ChangePublisherSubscription(data *ChangePublisherSubscriptionRequest) (subscriptionId, oldPublisherSubscriptionPlanId, error) {
	var id, oldPlanId int
	query := `
	SELECT id, plan_id FROM publisher_subscriptions WHERE user_id = $1
	`
	if err := r.DB.QueryRow(query, data.UserId).Scan(&id, &oldPlanId); err != nil {
		return -1, -1, fmt.Errorf("failed to get publisher subscription plan: %w", err)
	}

	query = `
	UPDATE publisher_subscriptions SET plan_id = $2 WHERE user_id = $1
	`
	if _, err := r.DB.Exec(query, data.UserId, data.ChangedPlanId); err != nil {
		return -1, -1, fmt.Errorf("failed to change publisher subscription plan: %w", err)
	}

	return subscriptionId(id), oldPublisherSubscriptionPlanId(oldPlanId), nil
}

func (r *Repository) RollbackChangePublisherSubscription(data *RollbackChangePublisherSubscriptionRequest) error {
	query := `
	UPDATE publisher_subscriptions SET plan_id = $2 WHERE id = $1
	`
	if _, err := r.DB.Exec(query, data.SubscriptionId, data.OldPlanId); err != nil {
		return fmt.Errorf("failed to rollback change publisher subscription plan: %w", err)
	}
	return nil
}

func (r *Repository) JoinPublication(data *JoinPublicationRequest) (int, error) {
	query := `
	INSERT INTO subscriber_subscriptions (user_id, publication_id, is_premium)
	VALUES ($1, $2, $3) RETURNING id
	`
	var id int
	if err := r.DB.QueryRow(query, data.UserId, data.PublicationId, data.IsPremium).Scan(&id); err != nil {
		return -1, fmt.Errorf("failed to join publication: %w", err)
	}

	return id, nil
}

func (r *Repository) RollbackJoinPublication(data *RollbackJoinPublicationRequest) error {
	query := `
	DELETE FROM subscriber_subscriptions WHERE id = $1
	`
	if _, err := r.DB.Exec(query, data.SubscriptionId); err != nil {
		return fmt.Errorf("failed to leave publication: %w", err)
	}
	return nil
}

func (r *Repository) LeavePublication(data *LeavePublicationRequest) (oldSubscriberSubscriptionPlanType, error) {
	query := `
	DELETE FROM subscriber_subscriptions WHERE user_id = $1 AND publication_id = $2 RETURNING is_premium
	`

	var isPremium bool
	if err := r.DB.QueryRow(query, data.UserId, data.PublicationId).Scan(&isPremium); err != nil {
		return false, fmt.Errorf("failed to join publication: %w", err)
	}

	return oldSubscriberSubscriptionPlanType(isPremium), nil
}

func (r *Repository) RollbackLeavePublication(data *RollbackLeavePublicationRequest) error {
	query := `
	INSERT INTO subscriber_subscriptions (user_id, publication_id, is_premium) VALUES ($1, $2, $3)
	`

	if _, err := r.DB.Exec(query, data.UserId, data.PublicationId, false); err != nil {
		return fmt.Errorf("failed to rollback leave publication: %w", err)
	}
	return nil
}

func (r *Repository) ChangeSubscriberSubscription(data *ChangeSubscriberSubscriptionRequest) (subscriptionId, error) {
	var id int
	query := `
	SELECT id, is_premium FROM subscriber_subscriptions WHERE user_id = $1 AND publication_id = $2
	`
	var isPremium bool
	if err := r.DB.QueryRow(query, data.UserId, data.PublicationId).Scan(&id, &isPremium); err != nil {
		return -1, fmt.Errorf("failed to get subscription premium status: %w", err)
	}

	query = `
	UPDATE subscriber_subscriptions SET is_premium = $3 WHERE user_id = $1 AND publication_id = $2
	`
	if _, err := r.DB.Exec(query, data.UserId, data.PublicationId, !isPremium); err != nil {
		return -1, fmt.Errorf("failed to update subscription premium status: %w", err)
	}
	return subscriptionId(id), nil
}

func (r *Repository) RollbackChangeSubscriberPlan(data *RollbackChangeSubscriberSubscriptionRequest) error {
	query := `
	SELECT is_premium FROM subscriber_subscriptions WHERE id = $1
	`
	var isPremium bool
	if err := r.DB.QueryRow(query, data.SubscriptionId).Scan(&isPremium); err != nil {
		return fmt.Errorf("failed to rollback change subscription: %w", err)
	}

	query = `
	UPDATE subscriber_subscriptions SET is_premium = $2 WHERE id = $1
	`
	if _, err := r.DB.Exec(query, data.SubscriptionId, !isPremium); err != nil {
		return fmt.Errorf("failed to rollback change subscription: %w", err)
	}
	return nil
}

func (r *Repository) IsAuthorizedPublisher(data *IsAuthorizedPublisherRequest) (bool, error) {
	query := `
	SELECT EXISTS(SELECT 1 FROM publishers WHERE user_id = $1 AND publication_id = $2)
	`
	var exists bool
	if err := r.DB.QueryRow(query, data.UserId, data.PublicationId).Scan(&exists); err != nil {
		return false, fmt.Errorf("failed to check if user is authorized publisher: %w", err)
	}
	return exists, nil
}

func (r *Repository) GetContentSubscribers(data *GetContentSubscribersRequest) ([]int32, error) {
	query := "SELECT user_id FROM subscriber_subscriptions WHERE publication_id = $1"
	var subscriberIds []int32
	rows, err := r.DB.Query(query, data.PublicationId)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscribers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan subscriber ID: %w", err)
		}
		subscriberIds = append(subscriberIds, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return subscriberIds, nil
}
