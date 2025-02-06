package app

import "tiny-letter/subscription/pkg/constants"

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
}

var appConfig Config = Config{
	App: AppConfig{
		JWT_secret:      constants.JWT_secret,
		JWT_exp_minutes: constants.JWT_exp_minutes,
	},
	GRPC: CommConfig{
		Domain: constants.Domain,
		Port:   constants.GRPC_port,
	},
	DB: DB_config{
		User:     constants.DB_user,
		Password: constants.DB_password,
		Name:     constants.DB_name,
		CommConfig: CommConfig{
			Domain: constants.Domain,
			Port:   constants.DB_port,
		},
	},
}

func GetConfig() Config {
	return appConfig
}
