package app

type CommConfig struct {
	Port   int
	Domain string
}

type Producer struct {
	IsProducerReturnSuccess bool
	NumberOfRetry           int
}

type MQ_consumer struct {
	IsConsumerReturnError bool
}

type Topic struct {
	JoinPublication      string
	LeavePublication     string
	SubscriberChangePlan string

	PublisherSubscribe   string
	PublisherUnsubscribe string
	PublisherChangePlan  string
	PublishLetter        string
}

type MQ_config struct {
	Topic Topic
	CommConfig
	Consumer MQ_consumer
}

type DB_config struct {
	User     string
	Password string
	Name     string
	CommConfig
}

type MailConfig struct {
	HostEmail        string
	TemplatePath     string
	EmailAppPassword string
	RateLimit        int
	BatchSize        int
}

type Config struct {
	DB   DB_config
	MQ   MQ_config
	GRPC CommConfig
	Mail MailConfig
}

var appConfig Config = Config{
	MQ: MQ_config{
		CommConfig: CommConfig{
			Port:   9092,
			Domain: "localhost",
		},
		Consumer: MQ_consumer{
			IsConsumerReturnError: true,
		},
		Topic: Topic{
			JoinPublication:      "join_publication",
			LeavePublication:     "leave_publication",
			SubscriberChangePlan: "subscriber_change_plan",
			PublisherSubscribe:   "publisher_subscribe",
			PublisherUnsubscribe: "publisher_unsubscribe",
			PublisherChangePlan:  "publisher_change_plan",
			PublishLetter:        "publish_letter",
		},
	},
	GRPC: CommConfig{
		Domain: "localhost",
		Port:   50003,
	},
	DB: DB_config{
		User:     "user",
		Password: "password",
		Name:     "email_db",
		CommConfig: CommConfig{
			Port:   5003,
			Domain: "localhost",
		},
	},
	Mail: MailConfig{
		HostEmail:        "tinyletter043@gmail.com",
		TemplatePath:     "templates/",
		EmailAppPassword: "ddyv ctjj hgrq hejt",
		RateLimit:        2,
		BatchSize:        2,
	},
}

func GetConfig() Config {
	return appConfig
}
