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

type Config struct {
	DB   CommConfig
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
	DB: CommConfig{
		Domain: constants.Domain,
		Port:   constants.DB_port,
	},
}

func GetConfig() Config {
	return appConfig
}
