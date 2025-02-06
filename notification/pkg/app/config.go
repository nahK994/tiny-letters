package app

import "tiny-letter/notification/pkg/constants"

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

type Config struct {
	CommConfig
	Consumer Consumer
	Producer Producer
}

var appConfig Config = Config{
	CommConfig: CommConfig{
		Port:   constants.MQ_port,
		Domain: constants.Domain,
	},
	Consumer: Consumer{
		IsConsumerReturnError: constants.Consumer_IsConsumerReturnError,
		ConfirmationTopic:     constants.ConfirmationNotification,
		PublicationTopic:      constants.PublicationNotification,
	},
	Producer: Producer{
		NumberOfRetry:           constants.Producer_NumberOfRetry,
		IsProducerReturnSuccess: constants.Producer_IsProducerReturnSuccess,
		ConfirmationTopic:       constants.ConfirmationEmail,
		PublicationTopic:        constants.PublicationEmail,
	},
}

func GetConfig() Config {
	return appConfig
}
