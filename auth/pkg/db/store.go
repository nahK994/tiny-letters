package db

import (
	"database/sql"
	grpc_client "tiny-letter/auth/cmd/grpc/client"
	"tiny-letter/auth/pkg/utils"
)

func (r *Repository) createUser(userInfo *CreateBaseUserRequest) (int, *sql.Tx, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return -1, nil, err
	}
	var userId int
	err = tx.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", userInfo.Name, userInfo.Email, userInfo.Password).Scan(&userId)
	if err != nil {
		return -1, nil, err
	}

	_, err = tx.Exec("INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)", userId, userInfo.Role)
	if err != nil {
		return -1, nil, err
	}
	return userId, tx, err
}

func (r *Repository) CreateSubscriber(userInfo *CreateBaseUserRequest) error {
	_, tx, err := r.createUser(userInfo)
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	return nil
}

func (r *Repository) CreatePublisher(userInfo *CreatePublisherRequest) error {
	userId, tx, err := r.createUser(&CreateBaseUserRequest{
		Email:    userInfo.Email,
		Password: userInfo.Password,
		Name:     userInfo.Name,
		Role:     userInfo.Role,
	})
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	if userInfo.Role == utils.Publisher {
		err = grpc_client.NotifyPublisherSubscription(userId, userInfo.PlanId, userInfo.Role)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Repository) GetUserInfoByEmail(email string) (*JWT_claim, error) {
	var claim JWT_claim
	err := d.DB.QueryRow(`
		SELECT 
			id, 
			password, 
			role,
		FROM 
			users
		WHERE 
			email = $1
	`, email).Scan(&claim.Id, &claim.Password, &claim.Role)

	if err != nil {
		return nil, err
	}
	return &claim, nil
}
