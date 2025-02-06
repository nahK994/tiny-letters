package app

import constant "tiny-letter/notification/pkg/constants"

const (
	localhost = "localhost"
)

type CommConfig struct {
	Port   int
	Domain string
}

type QueueConfig struct {
	Topic string
}

type Consumer struct {
	Confirmation          QueueConfig
	Publication           QueueConfig
	IsConsumerReturnError bool
}

type Producer struct {
	Confirmation            QueueConfig
	Publication             QueueConfig
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
			Confirmation: QueueConfig{
				Topic: constant.ConfirmationNotification,
			},
			Publication: QueueConfig{
				Topic: constant.PublicationNotification,
			},
		},
		Producer: Producer{
			NumberOfRetry:           5,
			IsProducerReturnSuccess: true,
			Confirmation: QueueConfig{
				Topic: constant.ConfirmationEmail,
			},
			Publication: QueueConfig{
				Topic: constant.PublicationEmail,
			},
		},
	},
}

func GetConfig() Config {
	return appConfig
}
