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
	Confirmation QueueConfig
	Publication  QueueConfig
}

type Producer struct {
	Confirmation QueueConfig
	Publication  QueueConfig
}

type Config struct {
	CommConfig
	Consumer
	Producer
}

var appConfig Config = Config{
	CommConfig: CommConfig{
		Domain: localhost,
		Port:   9092,
	},
	Consumer: Consumer{
		Confirmation: QueueConfig{
			Topic: constant.NotificationConfirmation,
		},
		Publication: QueueConfig{
			Topic: constant.NotificationPublication,
		},
	},
	Producer: Producer{
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
