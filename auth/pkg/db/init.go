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
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		email VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		subscription_id INT NOT NULL
	);
	`

	roleTable := `
	CREATE TABLE IF NOT EXISTS roles (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL
	);
	`

	userRoleTable := `
	CREATE TABLE IF NOT EXISTS user_roles (
		user_id INT NOT NULL,
		role_id INT NOT NULL,
		CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
	);

	-- Indexes for faster lookups
	CREATE INDEX IF NOT EXISTS idx_user_id ON user_roles(user_id);
	CREATE INDEX IF NOT EXISTS idx_role_id ON user_roles(role_id);
	`

	createSubscriptionPlanTable := `
	CREATE TABLE IF NOT EXISTS subscription_plans (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		order INT NOT NULL
	);
	`

	createPermissionTable := `
	CREATE TABLE IF NOT EXISTS permissions (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		subscription_id INT NOT NULL,
		CONSTRAINT fk_permissions_subscription FOREIGN KEY (subscription_id) REFERENCES subscription_plans(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_permissions_subscription_id ON permissions(subscription_id);
	`

	createAudienceTable := `
	CREATE TABLE IF NOT EXISTS plan_audiences (
		id SERIAL PRIMARY KEY,
		subscription_id INT NOT NULL,
		size INT NOT NULL,
		CONSTRAINT fk_plan_audiences_subscription FOREIGN KEY (subscription_id) REFERENCES subscription_plans(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_plan_audiences_subscription_id ON plan_audiences(subscription_id);
	`

	createPublisherSubscriptionManagementTable := `
	CREATE TABLE IF NOT EXISTS publisher_subscription_managements (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		subscription_id INT NOT NULL,
		CONSTRAINT fk_publisher_subscription_managements_subscription FOREIGN KEY (subscription_id) REFERENCES subscription_plans(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_publisher_subscription_management_subscription_id ON publisher_subscription_managements(subscription_id);
	`

	createSubscriberSubscriptionManagementTable := `
	CREATE TABLE IF NOT EXISTS subscriber_subscription_managements (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		subscription_id INT NOT NULL,
		publication_id INT NOT NULL,
		CONSTRAINT fk_subscriber_subscription_managements_subscription FOREIGN KEY (subscription_id) REFERENCES subscription_plans(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS idx_subscriber_subscription_management_subscription_id ON subscriber_subscription_managements(subscription_id);
	CREATE INDEX IF NOT EXISTS idx_subscriber_subscription_management_publication_id ON subscriber_subscription_managements(publication_id);
	`

	if _, err := db.Exec(createSubscriptionPlanTable); err != nil {
		return fmt.Errorf("failed to create 'subscription_plans' table: %w", err)
	}
	if _, err := db.Exec(createPermissionTable); err != nil {
		return fmt.Errorf("failed to create 'permissions' table: %w", err)
	}
	if _, err := db.Exec(createAudienceTable); err != nil {
		return fmt.Errorf("failed to create 'plan_audiences' table: %w", err)
	}
	if _, err := db.Exec(createUserTable); err != nil {
		return fmt.Errorf("failed to create 'users' table: %w", err)
	}
	if _, err := db.Exec(roleTable); err != nil {
		return fmt.Errorf("failed to create 'roles' table: %w", err)
	}
	if _, err := db.Exec(userRoleTable); err != nil {
		return fmt.Errorf("failed to create 'user_roles' table: %w", err)
	}
	if _, err := db.Exec(createPublisherSubscriptionManagementTable); err != nil {
		return fmt.Errorf("failed to create 'publisher_subscription_managements' table: %w", err)
	}
	if _, err := db.Exec(createSubscriberSubscriptionManagementTable); err != nil {
		return fmt.Errorf("failed to create 'subscriber_subscription_managements' table: %w", err)
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
