package app

type DB_config struct {
	User     string
	Password string
	Name     string
	CommConfig
}

type AppConfig struct {
	JWT_secret      string
	JWT_exp_minutes int
}

type CommConfig struct {
	Domain string
	Port   int
}

type GRPC_config struct {
	Subscription CommConfig
}

type MQ_producer struct {
	NumberOfRetry           int
	IsProducerReturnSuccess bool
}

type MsgActionConfig struct {
	PublishContent string
}

type MQ_topic struct {
	PublicationNotification string
}

type MQ_config struct {
	CommConfig
	Topic     MQ_topic
	Producer  MQ_producer
	MsgAction MsgActionConfig
}

type Config struct {
	App  AppConfig
	DB   DB_config
	REST CommConfig
	GRPC GRPC_config
	MQ   MQ_config
}

var appConfig Config = Config{
	App: AppConfig{
		JWT_secret:      "secret",
		JWT_exp_minutes: 60,
	},
	REST: CommConfig{
		Domain: "localhost",
		Port:   8001,
	},
	GRPC: GRPC_config{
		Subscription: CommConfig{
			Domain: "localhost",
			Port:   50001,
		},
	},
	DB: DB_config{
		User:     "user",
		Password: "password",
		Name:     "content_db",
		CommConfig: CommConfig{
			Domain: "localhost",
			Port:   5001,
		},
	},
	MQ: MQ_config{
		Producer: MQ_producer{
			NumberOfRetry:           5,
			IsProducerReturnSuccess: true,
		},
		Topic: MQ_topic{
			PublicationNotification: "publication_notification",
		},
		CommConfig: CommConfig{
			Domain: "localhost",
			Port:   9092,
		},
		MsgAction: MsgActionConfig{
			PublishContent: "publish_content",
		},
	},
}

func GetConfig() Config {
	return appConfig
}
