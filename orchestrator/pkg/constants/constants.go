package constants

const (
	PublisherSubscribe   = "publisher_subscribe"
	PublisherUnsubscribe = "publisher_unsubscribe"
	PublisherChangePlan  = "publisher_change_plan"

	JoinPublication      = "join_publication"
	LeavePublication     = "leave_publication"
	SubscriberChangePlan = "subscriber_change_plan"

	RegisterSubscriber = "register_subscriber"
	RegisterPublisher  = "register_publisher"

	Domain                 = "localhost"
	GRPC_subscription_port = 50002
	GRPC_auth_port         = 50000
	REST_port              = 8080
	MQ_port                = 9092

	MQ_NumberOfRetry           = 5
	MQ_IsProducerReturnSuccess = true
	MQ_topic                   = "confirmation_notification"
)
