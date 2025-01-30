package app

const domain = "localhost"

type DB_config struct {
	User     string
	Password string
	Name     string
	CommConfig
}

type AppConfig struct {
	JWT_secret      string
	JWT_exp_minutes int
}

type CommConfig struct {
	Domain string
	Port   int
}

type GRPC_config struct {
	Subscription CommConfig
}

type Config struct {
	App  AppConfig
	DB   DB_config
	REST CommConfig
	GRPC GRPC_config
}

var appConfig Config = Config{
	App: AppConfig{
		JWT_secret:      "secret",
		JWT_exp_minutes: 60,
	},
	REST: CommConfig{
		Domain: domain,
		Port:   8001,
	},
	GRPC: GRPC_config{
		Subscription: CommConfig{
			Domain: domain,
			Port:   50002,
		},
	},
	DB: DB_config{
		User:     "user",
		Password: "password",
		Name:     "content_db",
		CommConfig: CommConfig{
			Domain: domain,
			Port:   5001,
		},
	},
}

func GetConfig() Config {
	return appConfig
}
