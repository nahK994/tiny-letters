package db

import (
	"database/sql"
	"fmt"
	"tiny-letter/subscription/pkg/models"
)

type (
	subscriptionId int
)

func (r *Repository) CreateSubscriptionForPublisher(data *models.ConfirmPublisherSubscriptionRequest) (subscriptionId, error) {
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

func (r *Repository) ChangePublisherSubscription(data *models.ChangePublisherSubscriptionRequest) error {
	var id, oldPlanId int
	query := `
	SELECT id, plan_id FROM publisher_subscriptions WHERE user_id = $1
	`
	if err := r.DB.QueryRow(query, data.UserId).Scan(&id, &oldPlanId); err != nil {
		return fmt.Errorf("failed to get publisher subscription plan: %w", err)
	}

	query = `
	UPDATE publisher_subscriptions SET plan_id = $2 WHERE user_id = $1
	`
	if _, err := r.DB.Exec(query, data.UserId, data.ChangedPlanId); err != nil {
		return fmt.Errorf("failed to change publisher subscription plan: %w", err)
	}

	return nil
}

func (r *Repository) JoinPublication(data *models.JoinPublicationRequest) (int, error) {
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

func (r *Repository) LeavePublication(data *models.LeavePublicationRequest) error {
	query := `
	DELETE FROM subscriber_subscriptions WHERE user_id = $1 AND publication_id = $2 RETURNING is_premium
	`

	var isPremium bool
	if err := r.DB.QueryRow(query, data.UserId, data.PublicationId).Scan(&isPremium); err != nil {
		return fmt.Errorf("failed to leave publication: %w", err)
	}

	return nil
}

func (r *Repository) ChangeSubscriberSubscription(data *models.ChangeSubscriberSubscriptionRequest) error {
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

func (r *Repository) GetContentSubscribers(data *models.GetContentSubscribersRequest) ([]int32, error) {
	var subscriberIds []int32
	var rows *sql.Rows
	var err error
	if !data.ContentIsPremium {
		query := "SELECT user_id FROM subscriber_subscriptions WHERE publication_id = $1"
		rows, err = r.DB.Query(query, data.PublicationId)
		if err != nil {
			return nil, fmt.Errorf("failed to get subscribers: %w", err)
		}
	} else {
		query := "SELECT user_id FROM subscriber_subscriptions WHERE publication_id = $1 AND is_premium = true"
		rows, err = r.DB.Query(query, data.PublicationId)
		if err != nil {
			return nil, fmt.Errorf("failed to get premium subscribers: %w", err)
		}
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
