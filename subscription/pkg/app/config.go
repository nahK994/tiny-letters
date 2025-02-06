package app

const (
	domain                   = "localhost"
	ConfirmationNotification = "confirmation_notification"
)

type DB_config struct {
	User     string
	Password string
	Name     string
	CommConfig
}

type CommConfig struct {
	Domain string
	Port   int
}

type AppConfig struct {
	JWT_secret      string
	JWT_exp_minutes int
}

type MQ_config struct {
	CommConfig
	Topic                   string
	NumberOfRetry           int
	IsProducerReturnSuccess bool
}

type Config struct {
	App  AppConfig
	DB   DB_config
	GRPC CommConfig
	MQ   MQ_config
}

var appConfig Config = Config{
	App: AppConfig{
		JWT_secret:      "secret",
		JWT_exp_minutes: 60,
	},
	GRPC: CommConfig{
		Domain: domain,
		Port:   50002,
	},
	DB: DB_config{
		User:     "user",
		Password: "password",
		Name:     "subscription_db",
		CommConfig: CommConfig{
			Domain: domain,
			Port:   5002,
		},
	},
	MQ: MQ_config{
		NumberOfRetry:           5,
		IsProducerReturnSuccess: true,
		Topic:                   ConfirmationNotification,
		CommConfig: CommConfig{
			Domain: domain,
			Port:   9092,
		},
	},
}

func GetConfig() Config {
	return appConfig
}
