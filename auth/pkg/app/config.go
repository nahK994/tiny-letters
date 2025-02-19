package app

type DB_config struct {
	User     string
	Password string
	Name     string
	CommConfig
}

type CommConfig struct {
	Port   int
	Domain string
}

type AppConfig struct {
	JWT_secret      string
	JWT_exp_minutes int
}

type Config struct {
	App  AppConfig
	DB   DB_config
	REST CommConfig
	GRPC CommConfig
}

var appConfig Config = Config{
	App: AppConfig{
		JWT_secret:      "secret",
		JWT_exp_minutes: 60,
	},
	REST: CommConfig{
		Port:   8000,
		Domain: "localhost",
	},
	GRPC: CommConfig{
		Domain: "localhost",
		Port:   50000,
	},
	DB: DB_config{
		User:     "user",
		Password: "password",
		Name:     "auth_db",
		CommConfig: CommConfig{
			Port:   5000,
			Domain: "localhost",
		},
	},
}

func GetConfig() Config {
	return appConfig
}
