package app

import "tiny-letter/orchestrator/pkg/constants"

type CommConfig struct {
	Port   int
	Domain string
}

type GRPC struct {
	Subscription CommConfig
	Auth         CommConfig
}

type MQ_config struct {
	CommConfig
	Topic                   string
	NumberOfRetry           int
	IsProducerReturnSuccess bool
}

type Config struct {
	GRPC
	REST CommConfig
	MQ   MQ_config
}

var appConfig Config = Config{
	GRPC: GRPC{
		Subscription: CommConfig{
			Port:   constants.GRPC_subscription_port,
			Domain: constants.Domain,
		},
		Auth: CommConfig{
			Port:   constants.GRPC_auth_port,
			Domain: constants.Domain,
		},
	},
	REST: CommConfig{
		Port:   constants.REST_port,
		Domain: constants.Domain,
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
