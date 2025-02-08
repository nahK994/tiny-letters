package app

import "tiny-letter/content/pkg/constants"

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

type MQ_config struct {
	CommConfig
	Topic                   string
	NumberOfRetry           int
	IsProducerReturnSuccess bool
}

type Config struct {
	App  AppConfig
	DB   DB_config
	REST CommConfig
	GRPC GRPC_config
	MQ   MQ_config
}

var appConfig Config = Config{
	App: AppConfig{
		JWT_secret:      constants.JWT_secret,
		JWT_exp_minutes: constants.JWT_exp_minutes,
	},
	REST: CommConfig{
		Domain: constants.Domain,
		Port:   constants.REST_port,
	},
	GRPC: GRPC_config{
		Subscription: CommConfig{
			Domain: constants.Domain,
			Port:   constants.GRPC_subscription_port,
		},
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
	MQ: MQ_config{
		NumberOfRetry:           constants.MQ_NumberOfRetry,
		IsProducerReturnSuccess: constants.MQ_IsProducerReturnSuccess,
		Topic:                   constants.MQ_topic,
		CommConfig: CommConfig{
			Domain: constants.Domain,
			Port:   constants.MQ_port,
		},
	},
}

func GetConfig() Config {
	return appConfig
}
