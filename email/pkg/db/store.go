package db

import "tiny-letter/email/pkg/models"

func (r *Repository) StoreUserContactEmail(data *models.UserContactEmail) error {
	_, err := r.DB.Exec("INSERT INTO users (user_id, email) VALUES ($1, $2)", data.UserId, data.Email)
	if err != nil {
		return err
	}

	return err
}
