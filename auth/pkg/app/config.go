package app

const domain = "localhost"

type DBConfig struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     int
}

type BaseCommConfig struct {
	Port   int
	Domain string
}

type CommConfig struct {
	REST BaseCommConfig
	GRPC BaseCommConfig
}

type AppConfig struct {
	CommConfig
	JWT_secret      string
	JWT_exp_minutes int
}

type Config struct {
	App AppConfig
	DB  DBConfig
}

var appConfig Config = Config{
	App: AppConfig{
		CommConfig: CommConfig{
			REST: BaseCommConfig{
				Port:   8000,
				Domain: domain,
			},
			GRPC: BaseCommConfig{
				Domain: domain,
				Port:   50000,
			},
		},
		JWT_secret:      "secret",
		JWT_exp_minutes: 60,
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
