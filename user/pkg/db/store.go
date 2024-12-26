package db

import (
	grpc_client "tiny-letter-user/cmd/grpc/client"

	"github.com/lib/pq"
)

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

func (d *Repository) GetUserInfoByEmail(email string) (*JWT_claim, error) {
	var claim JWT_claim
	err := d.DB.QueryRow(`
		SELECT 
			u.id, 
			u.subscription_id, 
			u.password, 
			array_agg(r.name) AS roles
		FROM 
			users AS u
		INNER JOIN 
			user_roles AS ur ON u.id = ur.user_id
		INNER JOIN 
			roles AS r ON ur.role_id = r.id
		WHERE 
			u.email = $1
		GROUP BY 
			u.id, u.subscription_id, u.password
	`, email).Scan(&claim.Id, &claim.SubscriptionId, &claim.Password, pq.Array(&claim.Roles))

	if err != nil {
		return nil, err
	}
	return &claim, nil
}
