package db

func (r *Repository) CreateUser(userInfo *CreateUserRequest) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		// Rollback if the transaction is not committed
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after rollback
		} else if err != nil {
			tx.Rollback() // rollback if error occurred
		} else {
			err = tx.Commit() // commit if no error occurred
		}
	}()

	var userId int
	err = tx.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", userInfo.Name, userInfo.Email, userInfo.Password).Scan(&userId)
	if err != nil {
		return err
	}

	for _, role := range userInfo.Roles {
		_, err = tx.Exec("INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)", userId, role)
		if err != nil {
			return err
		}
	}

	return nil
}
