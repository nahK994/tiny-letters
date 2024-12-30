package db

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
