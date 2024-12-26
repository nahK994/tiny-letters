package db

import grpc_client "tiny-letter-user/cmd/grpc/client"

func (r *Repository) CreateUser(userInfo *CreateUserRequest) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var userId int
	err = tx.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", userInfo.Name, userInfo.Email, userInfo.Password).Scan(&userId)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)", userId, userInfo.Role)
	if err != nil {
		return err
	}

	err = grpc_client.NotifySubscription(userId, userInfo.SubscriptionType, userInfo.Role)
	if err != nil {
		return err
	}

	return nil
}
