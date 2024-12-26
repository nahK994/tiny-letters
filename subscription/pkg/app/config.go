package app

type DBConfig struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     int
}

type BaseCommConfig struct {
	Port   string
	Domain string
}

type GRPCConfig struct {
	User_Subscriber BaseCommConfig
}

type CommConfig struct {
	REST BaseCommConfig
	GRPC GRPCConfig
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
				Port:   "8002",
				Domain: "localhost",
			},
			GRPC: GRPCConfig{
				User_Subscriber: BaseCommConfig{
					Port:   "50000",
					Domain: "localhost",
				},
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
		Host:     "localhost",
	},
}

func GetConfig() Config {
	return appConfig
}
