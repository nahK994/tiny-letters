package db

import "fmt"

func (r *Repository) CreateUser(userInfo *CreateUserRequest) (int, error) {
	var userId int
	err := r.DB.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", userInfo.Name, userInfo.Email, userInfo.Password).Scan(&userId)
	if err != nil {
		return -1, err
	}

	_, err = r.DB.Exec("INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)", userId, userInfo.Role)
	if err != nil {
		return -1, err
	}
	return userId, err
}

func (r *Repository) GetUserInfoByEmail(email string) (*JWT_claim, error) {
	var claim JWT_claim
	err := r.DB.QueryRow(`
		SELECT id, password, role,
		FROM users
		WHERE email = $1
	`, email).Scan(&claim.Id, &claim.Password, &claim.Role)

	if err != nil {
		return nil, err
	}
	return &claim, nil
}

func (r *Repository) ConfirmPublisherSubscription(data *ConfirmPublisherSubscriptionRequest) error {
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

func (r *Repository) RevokePublisherSubscription(data *RevokePublisherSubscriptionRequest) error {
	query := `
	DELETE FROM publisher_subscriptions WHERE user_id = $1
	`
	_, err := r.DB.Exec(query, data.UserId)
	if err != nil {
		return fmt.Errorf("failed to revoke publisher subscription: %w", err)
	}
	return nil
}

func (r *Repository) ChangePublisherSubscription(data *ChangePublisherSubscriptionRequest) error {
	query := `
	SELECT id, plan_id FROM publisher_subscriptions WHERE user_id = $1
	`
	if _, err := r.DB.Exec(query, data.UserId); err != nil {
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

func (r *Repository) JoinPublication(data *JoinPublicationRequest) error {
	query := `
	INSERT INTO subscriber_subscriptions (user_id, publication_id, is_premium)
	VALUES ($1, $2, $3)
	`
	if _, err := r.DB.Exec(query, data.UserId, data.PublicationId, data.IsPremium); err != nil {
		return fmt.Errorf("failed to join publication: %w", err)
	}

	return nil
}

func (r *Repository) LeavePublication(data *LeavePublicationRequest) error {
	query := `
	DELETE FROM subscriber_subscriptions WHERE user_id = $1 AND publication_id = $2
	`

	if _, err := r.DB.Exec(query, data.UserId, data.PublicationId); err != nil {
		return fmt.Errorf("failed to join publication: %w", err)
	}

	return nil
}

func (r *Repository) ChangeSubscriberSubscription(data *ChangeSubscriberSubscriptionRequest) error {
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
