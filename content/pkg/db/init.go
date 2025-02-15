package db

import (
	"database/sql"
	"fmt"
	"tiny-letter/content/pkg/app"

	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func createTables(db *sql.DB) error {
	createPublicationsTable := `
	CREATE TABLE IF NOT EXISTS publications (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		publisher_id INT NOT NULL,
	);
	`

	createPostsTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		content TEXT NOT NULL,
		publication_id INT NOT NULL,
		is_premium BOOLEAN NOT NULL,
		CONSTRAINT fk_post_publication FOREIGN KEY (publication_id) REFERENCES publications(id) ON DELETE CASCADE	
	);
	`

	if _, err := db.Exec(createPublicationsTable); err != nil {
		return fmt.Errorf("could not create publications table: %w", err)
	}
	if _, err := db.Exec(createPostsTable); err != nil {
		return fmt.Errorf("could not create posts table: %w", err)
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
