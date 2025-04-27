package db

import "tiny-letter/email/pkg/models"

func (r *Repository) RegisterSubscriber(data *models.SubscriberRegistration) error {
	_, err := r.DB.Exec("INSERT INTO users (user_id, email) VALUES ($1, $2)", data.UserId, data.Email)
	if err != nil {
		return err
	}

	return err
}

func (r *Repository) RegisterPublisher(data *models.PublisherRegistration) error {
	_, err := r.DB.Exec("INSERT INTO users (user_id, email) VALUES ($1, $2, $3)", data.UserId, data.Email, data.PlanId)
	if err != nil {
		return err
	}

	return err
}
