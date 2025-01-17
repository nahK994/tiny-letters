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
	Orchestrator GRPCConfig
	Email        GRPCConfig
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
			Orchestrator: GRPCConfig{
				Port:   50002,
				Domain: domain,
			},
			Email: GRPCConfig{
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
