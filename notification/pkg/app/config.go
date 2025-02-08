package app

import "tiny-letter/notification/pkg/constants"

type CommConfig struct {
	Port   int
	Domain string
}

type Consumer struct {
	IsConsumerReturnError bool
}

type Producer struct {
	IsProducerReturnSuccess bool
	NumberOfRetry           int
}

type MQ_config struct {
	ConfirmationTopic string
	PublicationTopic  string
	CommConfig
	Consumer Consumer
	Producer Producer
}

type Config struct {
	MQ   MQ_config
	GRPC CommConfig
}

var appConfig Config = Config{
	MQ: MQ_config{
		ConfirmationTopic: constants.ConfirmationNotification,
		PublicationTopic:  constants.PublicationNotification,
		CommConfig: CommConfig{
			Port:   constants.MQ_port,
			Domain: constants.Domain,
		},
		Consumer: Consumer{
			IsConsumerReturnError: constants.Consumer_IsConsumerReturnError,
		},
		Producer: Producer{
			NumberOfRetry:           constants.Producer_NumberOfRetry,
			IsProducerReturnSuccess: constants.Producer_IsProducerReturnSuccess,
		},
	},
	GRPC: CommConfig{
		Domain: constants.Domain,
		Port:   constants.GRPC_port,
	},
}

func GetConfig() Config {
	return appConfig
}
