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

type MsgTopic struct {
	SubscriberRegister   string
	JoinPublication      string
	LeavePublication     string
	SubscriberChangePlan string

	PublisherRegister    string
	PublisherSubscribe   string
	PublisherUnsubscribe string
	PublisherChangePlan  string
	PublishLetter        string
}

type MQ_config struct {
	CommConfig
	Consumer Consumer
	Producer Producer
	Topic    MsgTopic
}

type Config struct {
	MQ   MQ_config
	GRPC CommConfig
}

var appConfig Config = Config{
	MQ: MQ_config{
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
		Topic: MsgTopic{
			JoinPublication:      "join_publication",
			LeavePublication:     "leave_publication",
			SubscriberChangePlan: "subscriber_change_plan",
			PublisherSubscribe:   "publisher_subscribe",
			PublisherUnsubscribe: "publisher_unsubscribe",
			PublisherChangePlan:  "publisher_change_plan",
			SubscriberRegister:   "subscriber_register",
			PublisherRegister:    "publisher_register",
			PublishLetter:        "publish_letter",
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
