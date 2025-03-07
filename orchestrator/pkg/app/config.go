package app

type CommConfig struct {
	Port   int
	Domain string
}

type GRPC struct {
	Subscription CommConfig
	Auth         CommConfig
}

type Producer struct {
	IsProducerReturnSuccess bool
	NumberOfRetry           int
}

type MQ_topic struct {
	ConfirmationNotification string
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
	CommConfig
	Topic     MQ_topic
	Producer  Producer
	MsgAction MsgActionConfig
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
			Domain: "localhost",
		},
		Auth: CommConfig{
			Port:   50000,
			Domain: "localhost",
		},
	},
	REST: CommConfig{
		Port:   8080,
		Domain: "localhost",
	},
	MQ: MQ_config{
		Producer: Producer{
			NumberOfRetry:           5,
			IsProducerReturnSuccess: true,
		},
		Topic: MQ_topic{
			ConfirmationNotification: "confirmation_notification",
		},
		CommConfig: CommConfig{
			Domain: "localhost",
			Port:   9092,
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
}

func GetConfig() Config {
	return appConfig
}
