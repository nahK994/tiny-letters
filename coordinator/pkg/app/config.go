package app

const (
	localhost                = "localhost"
	ConfirmationNotification = "confirmation_notification"
)

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
			Port:   50002,
			Domain: localhost,
		},
		Auth: CommConfig{
			Port:   50000,
			Domain: localhost,
		},
	},
	REST: CommConfig{
		Port:   8080,
		Domain: localhost,
	},
	MQ: MQ_config{
		NumberOfRetry:           5,
		IsProducerReturnSuccess: true,
		Topic:                   ConfirmationNotification,
		CommConfig: CommConfig{
			Domain: localhost,
			Port:   9092,
		},
	},
}

func GetConfig() Config {
	return appConfig
}
