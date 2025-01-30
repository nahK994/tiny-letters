package db

import (
	"database/sql"
	"fmt"
	"tiny-letter/subscription/pkg/app"

	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func createTables(db *sql.DB) error {
	createPublisherSubscriptionPlansTable := `
	CREATE TABLE IF NOT EXISTS publisher_subscription_plans (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		order INT NOT NULL
	);
	`

	createPublisherSubscriptionsTable := `
	CREATE TABLE IF NOT EXISTS publisher_subscriptions (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL UNIQUE,
		plan_id INT NOT NULL,
		CONSTRAINT fk_publisher_subscription FOREIGN KEY (plan_id) REFERENCES publisher_subscription_plans(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_publisher_subscription_plan_id ON publisher_subscriptions(plan_id);
	CREATE INDEX IF NOT EXISTS idx_publisher_subscription_user_id ON publisher_subscriptions(user_id);
	`

	createSubscriberSubscribersTable := `
	CREATE TABLE IF NOT EXISTS subscriber_subscriptions (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		publication_id INT NOT NULL,
		is_premium BOOLEAN NOT NULL,
	);
	CREATE INDEX IF NOT EXISTS idx_publisher_subscriber_user_id ON subscriber_subscribers(user_id);
	CREATE INDEX IF NOT EXISTS idx_publisher_subscriber_publication_id ON subscriber_subscribers(publication_id);
	`

	createAudienceLimitsTable := `
	CREATE TABLE IF NOT EXISTS audience_limits (
		id SERIAL PRIMARY KEY,
		plan_id INT NOT NULL,
		size INT NOT NULL,
	CONSTRAINT fk_audience_limit FOREIGN KEY (plan_id) REFERENCES publisher_subscription_plans(id) ON DELETE CASCADE
	);
	`

	if _, err := db.Exec(createPublisherSubscriptionPlansTable); err != nil {
		return fmt.Errorf("could not create publisher_subscription_plans table: %w", err)
	}
	if _, err := db.Exec(createPublisherSubscriptionsTable); err != nil {
		return fmt.Errorf("could not create publisher_subscriptions table: %w", err)
	}
	if _, err := db.Exec(createSubscriberSubscribersTable); err != nil {
		return fmt.Errorf("could not create subscriber_subscribers table: %w", err)
	}
	if _, err := db.Exec(createAudienceLimitsTable); err != nil {
		return fmt.Errorf("could not create audience_limits table: %w", err)
	}

	return nil
}

func Init(dbConfig app.DB_config) (*Repository, error) {
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
