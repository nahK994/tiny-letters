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

type Config struct {
	Port     int
	Domain   string
	Consumer Consumer
	Producer Producer
}

var appConfig Config = Config{
	Domain: localhost,
	Port:   9092,
	Consumer: Consumer{
		IsConsumerReturnError: true,
		Confirmation: QueueConfig{
			Topic: constant.NotificationConfirmation,
		},
		Publication: QueueConfig{
			Topic: constant.NotificationPublication,
		},
	},
	Producer: Producer{
		NumberOfRetry:           5,
		IsProducerReturnSuccess: true,
		Confirmation: QueueConfig{
			Topic: constant.EmailConfirmation,
		},
		Publication: QueueConfig{
			Topic: constant.EmailPublication,
		},
	},
}

func GetConfig() Config {
	return appConfig
}
