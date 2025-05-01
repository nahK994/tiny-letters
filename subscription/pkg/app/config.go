package app

type DB_config struct {
	User     string
	Password string
	Name     string
	CommConfig
}

type CommConfig struct {
	Domain string
	Port   int
}

type AppConfig struct {
	JWT_secret      string
	JWT_exp_minutes int
}

type Config struct {
	App  AppConfig
	DB   DB_config
	GRPC CommConfig
	REST CommConfig
}

var appConfig Config = Config{
	App: AppConfig{
		JWT_secret:      "secret",
		JWT_exp_minutes: 60,
	},
	GRPC: CommConfig{
		Domain: "localhost",
		Port:   50002,
	},
	REST: CommConfig{
		Domain: "localhost",
		Port:   8002,
	},
	DB: DB_config{
		User:     "user",
		Password: "password",
		Name:     "subscription_db",
		CommConfig: CommConfig{
			Domain: "localhost",
			Port:   5002,
		},
	},
}

func GetConfig() Config {
	return appConfig
}
