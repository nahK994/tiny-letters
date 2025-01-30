package app

const domain = "localhost"

type DBConfig struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     int
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
	DB   DBConfig
	GRPC CommConfig
}

var appConfig Config = Config{
	App: AppConfig{
		JWT_secret:      "secret",
		JWT_exp_minutes: 60,
	},
	GRPC: CommConfig{
		Domain: domain,
		Port:   50002,
	},
	DB: DBConfig{
		User:     "user",
		Password: "password",
		Name:     "subscription_db",
		Port:     5002,
		Host:     domain,
	},
}

func GetConfig() Config {
	return appConfig
}
