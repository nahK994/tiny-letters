package db

import (
	"tiny-letter/auth/pkg/models"
)

func (r *Repository) GetUserInfoByEmail(email string) (*models.JWT_claim, error) {
	var claim models.JWT_claim
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

func (r *Repository) CreateUser(userInfo *models.UserRegistration) (int, error) {
	var userId int
	err := r.DB.QueryRow("INSERT INTO users (email, password, role) VALUES ($1, $2, $3) RETURNING id", userInfo.Email, userInfo.Password, userInfo.Role).Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, err
}

func (r *Repository) RollbackCreateUser(userId int) error {
	_, err := r.DB.Exec("DELETE FROM users where id=$1", userId)

	return err
}
