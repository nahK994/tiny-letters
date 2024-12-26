package app

type DBConfig struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     int
}

type AppConfig struct {
	Port            string
	Domain          string
	JWT_secret      string
	JWT_exp_minutes int
}

type Config struct {
	App AppConfig
	DB  DBConfig
}

var appConfig Config = Config{
	App: AppConfig{
		Port:            "8002",
		Domain:          "localhost",
		JWT_secret:      "secret",
		JWT_exp_minutes: 60,
	},
	DB: DBConfig{
		User:     "user",
		Password: "password",
		Name:     "publisher_db",
		Port:     5002,
		Host:     "localhost",
	},
}

func GetConfig() Config {
	return appConfig
}
