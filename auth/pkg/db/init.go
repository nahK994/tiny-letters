package db

import (
	"database/sql"
	"fmt"
	"tiny-letter/auth/pkg/app"

	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func createTables(db *sql.DB) error {
	createSubscriberTable := `
	CREATE TABLE IF NOT EXISTS subscribers (
		id SERIAL PRIMARY KEY,
		email VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
	);
	`

	createPublisherTable := `
	CREATE TABLE IF NOT EXISTS subscribers (
		id SERIAL PRIMARY KEY,
		email VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		plan_id INT NOT NULL,
	);
	`

	if _, err := db.Exec(createPublisherTable); err != nil {
		return fmt.Errorf("failed to create 'publishers' table: %w", err)
	}
	if _, err := db.Exec(createSubscriberTable); err != nil {
		return fmt.Errorf("failed to create 'subscribers' table: %w", err)
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
