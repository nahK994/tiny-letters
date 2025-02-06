package app

import constant "tiny-letter/notification/pkg/constants"

const (
	localhost = "localhost"
)

type CommConfig struct {
	Port   int
	Domain string
}

type Consumer struct {
	ConfirmationTopic     string
	PublicationTopic      string
	IsConsumerReturnError bool
}

type Producer struct {
	ConfirmationTopic       string
	PublicationTopic        string
	IsProducerReturnSuccess bool
	NumberOfRetry           int
}

type MQ_config struct {
	CommConfig
	Consumer Consumer
	Producer Producer
}

type Config struct {
	MQ   MQ_config
	GRPC CommConfig
}

var appConfig Config = Config{
	GRPC: CommConfig{
		Port:   50002,
		Domain: localhost,
	},
	MQ: MQ_config{
		CommConfig: CommConfig{
			Port:   9092,
			Domain: localhost,
		},
		Consumer: Consumer{
			IsConsumerReturnError: true,
			ConfirmationTopic:     constant.ConfirmationNotification,
			PublicationTopic:      constant.PublicationNotification,
		},
		Producer: Producer{
			NumberOfRetry:           5,
			IsProducerReturnSuccess: true,
			ConfirmationTopic:       constant.ConfirmationEmail,
			PublicationTopic:        constant.PublicationEmail,
		},
	},
}

func GetConfig() Config {
	return appConfig
}
