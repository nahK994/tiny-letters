package db

import (
	"database/sql"
	"fmt"
	"tiny-letter-subscription/pkg/app"

	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func createTables(db *sql.DB) error {
	createSubscriptionPlanTable := `
	CREATE TABLE IF NOT EXISTS subscription_plans (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL
	);
	`

	createPublisherSubscriptionManagementTable := `
	CREATE TABLE IF NOT EXISTS publisher_subscription_managements (
		id SERIAL PRIMARY KEY,
		subscription_id INT NOT NULL,
		CONSTRAINT fk_publisher_subscription_managements_subscription FOREIGN KEY (subscription_id) REFERENCES subscription_plans(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_publisher_subscription_management_subscription_id ON publisher_subscription_managements(subscription_id);
	`

	createSubscriberSubscriptionManagementTable := `
	CREATE TABLE IF NOT EXISTS subscriber_subscription_managements (
		id SERIAL PRIMARY KEY,
		subscription_id INT NOT NULL,
		publication_id INT NOT NULL,
		CONSTRAINT fk_subscriber_subscription_managements_subscription FOREIGN KEY (subscription_id) REFERENCES subscription_plans(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_subscriber_subscription_management_subscription_id ON subscriber_subscription_managements(subscription_id);
	CREATE INDEX IF NOT EXISTS idx_subscriber_subscription_management_publication_id ON subscriber_subscription_managements(publication_id);
	`

	if _, err := db.Exec(createSubscriptionPlanTable); err != nil {
		return fmt.Errorf("could not create subscription_plans table: %w", err)
	}
	if _, err := db.Exec(createPublisherSubscriptionManagementTable); err != nil {
		return fmt.Errorf("could not create publisher_subscription_managements table: %w", err)
	}
	if _, err := db.Exec(createSubscriberSubscriptionManagementTable); err != nil {
		return fmt.Errorf("could not create subscriber_subscription_managements table: %w", err)
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
