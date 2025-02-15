package app

import "tiny-letter/email/pkg/constants"

type CommConfig struct {
	Port   int
	Domain string
}

type Producer struct {
	IsProducerReturnSuccess bool
	NumberOfRetry           int
}

type MQ_config struct {
	ConfirmationTopic string
	PublicationTopic  string
	CommConfig
	IsConsumerReturnError bool
}

type DB_config struct {
	User     string
	Password string
	Name     string
	CommConfig
}

type Config struct {
	DB   DB_config
	MQ   MQ_config
	GRPC CommConfig
}

var appConfig Config = Config{
	MQ: MQ_config{
		ConfirmationTopic: constants.ConfirmationEmail,
		PublicationTopic:  constants.PublicationEmail,
		CommConfig: CommConfig{
			Port:   constants.MQ_port,
			Domain: constants.Domain,
		},
		IsConsumerReturnError: constants.Consumer_IsConsumerReturnError,
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
