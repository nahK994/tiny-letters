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

type MQ_topic struct {
	ConfirmationEmail string
	PublicationEmail  string
}

type MsgActionConfig struct {
	JoinPublication      string
	LeavePublication     string
	SubscriberChangePlan string
	PublisherSubscribe   string
	PublisherUnsubscribe string
	PublisherChangePlan  string
	RegisterUser         string
}

type MQ_config struct {
	Topic MQ_topic
	CommConfig
	Consumer  MQ_consumer
	MsgAction MsgActionConfig
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
		Topic: MQ_topic{
			ConfirmationEmail: "confirmation_email",
			PublicationEmail:  "publication_email",
		},
		CommConfig: CommConfig{
			Port:   9092,
			Domain: "localhost",
		},
		Consumer: MQ_consumer{
			IsConsumerReturnError: true,
		},
		MsgAction: MsgActionConfig{
			JoinPublication:      "join_publication",
			LeavePublication:     "leave_publication",
			SubscriberChangePlan: "subscriber_change_plan",
			PublisherSubscribe:   "publisher_subscribe",
			PublisherUnsubscribe: "publisher_unsubscribe",
			PublisherChangePlan:  "publisher_change_plan",
			RegisterUser:         "register_user",
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
