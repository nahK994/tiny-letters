package app

const domain = "localhost"

type DBConfig struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     int
}

type GRPCConfig struct {
	Port   int
	Domain string
}

type CommConfig struct {
	Coordinator GRPCConfig
	Content     GRPCConfig
}

type AppConfig struct {
	CommConfig
	JWT_secret      string
	JWT_exp_minutes int
}

type Config struct {
	Port   int
	Domain string
	App    AppConfig
	DB     DBConfig
}

var appConfig Config = Config{
	App: AppConfig{
		CommConfig: CommConfig{
			Coordinator: GRPCConfig{
				Port:   50002,
				Domain: domain,
			},
			Content: GRPCConfig{
				Port:   50003,
				Domain: domain,
			},
		},

		JWT_secret:      "secret",
		JWT_exp_minutes: 60,
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
