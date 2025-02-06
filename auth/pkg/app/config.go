package app

import "tiny-letter/auth/pkg/constants"

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
		JWT_secret:      constants.JWT_secret,
		JWT_exp_minutes: constants.JWT_exp_minutes,
	},
	REST: CommConfig{
		Port:   constants.REST_port,
		Domain: constants.Domain,
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
			Port:   constants.DB_port,
			Domain: constants.Domain,
		},
	},
}

func GetConfig() Config {
	return appConfig
}
