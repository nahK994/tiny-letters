package db

import (
	"database/sql"
	"fmt"
	"tiny-letter/email/pkg/app"

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
		user_id SERIAL PRIMARY KEY,
		email VARCHAR(50) NOT NULL UNIQUE,
	);
	`

	if _, err := db.Exec(createUserTable); err != nil {
		return fmt.Errorf("failed to create 'users' table: %w", err)
	}

	return nil
}

func Init(dbConfig *app.DB_config) (*Repository, error) {
	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Domain, dbConfig.Port, dbConfig.Name)
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
