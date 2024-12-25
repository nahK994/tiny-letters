package db

import (
	"database/sql"
	"fmt"
	"tiny-letter-user/pkg/app"

	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func createTables(db *sql.DB) error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		email VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
	);
	`

	roleTable := `
	CREATE TABLE IF NOT EXISTS roles (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
	);
	`

	userRoleTable := `
	CREATE TABLE IF NOT EXISTS user_roles (
		user_id INT NOT NULL,
		role_id INT NOT NULL,
		CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
		INDEX idx_user_id (user_id),
		INDEX idx_role_id (role_id)
	);
	`

	if _, err := db.Exec(createUserTable); err != nil {
		return fmt.Errorf("could not create vendors table: %w", err)
	}

	if _, err := db.Exec(roleTable); err != nil {
		return fmt.Errorf("could not create vendors table: %w", err)
	}

	if _, err := db.Exec(userRoleTable); err != nil {
		return fmt.Errorf("could not create vendors table: %w", err)
	}

	return nil
}

func Init(dbConfig app.DBConfig) (*Repository, error) {
	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	if err := createTables(db); err != nil {
		return nil, err
	}

	return &Repository{
		DB: db,
	}, nil
}
