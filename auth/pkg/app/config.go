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
	Port   int
	Domain string
}

type AppConfig struct {
	JWT_secret      string
	JWT_exp_minutes int
}

type Config struct {
	App  AppConfig
	DB   DBConfig
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
		Domain: domain,
	},
	GRPC: CommConfig{
		Domain: domain,
		Port:   50000,
	},
	DB: DBConfig{
		User:     "user",
		Password: "password",
		Name:     "auth_db",
		Port:     5000,
		Host:     domain,
	},
}

func GetConfig() Config {
	return appConfig
}
