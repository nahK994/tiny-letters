package app

type CommConfig struct {
	Port   int
	Domain string
}

type Consumer struct {
	IsConsumerReturnError bool
}

type Producer struct {
	IsProducerReturnSuccess bool
	NumberOfRetry           int
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

type MQ_topic struct {
	ConfirmationEmail string
	PublicationEmail  string
}

type MQ_config struct {
	Topic MQ_topic
	CommConfig
	Consumer  Consumer
	Producer  Producer
	MsgAction MsgActionConfig
}

type Config struct {
	MQ   MQ_config
	GRPC CommConfig
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
		Consumer: Consumer{
			IsConsumerReturnError: true,
		},
		Producer: Producer{
			NumberOfRetry:           5,
			IsProducerReturnSuccess: true,
		},
		MsgAction: MsgActionConfig{
			JoinPublication:      "join_publication",
			LeavePublication:     "leave_publication",
			SubscriberChangePlan: "subscriber_change_plan",
			PublisherSubscribe:   "publisher_subscribe",
			PublisherUnsubscribe: "publisher_unsubscribe",
			PublisherChangePlan:  "publisher_change_plan",
		},
	},
	GRPC: CommConfig{
		Domain: "localhost",
		Port:   50002,
	},
}

func GetConfig() Config {
	return appConfig
}
