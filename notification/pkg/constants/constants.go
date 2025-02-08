package constants

const (
	PublisherSubscribe   = "publisher_subscribe"
	PublisherUnsubscribe = "publisher_unsubscribe"
	PublisherChangePlan  = "publisher_change_plan"

	SubscriberSubscribe   = "subscriber_subscribe"
	SubscriberUnsubscribe = "subscriber_unsubscribe"
	SubscriberChangePlan  = "subscriber_change_plan"

	ConfirmationNotification = "confirmation_notification"
	PublicationNotification  = "publication_notification"

	ConfirmationEmail = "confirmation_email"
	PublicationEmail  = "publication_email"

	Domain = "localhost"

	MQ_port                          = 9092
	Producer_NumberOfRetry           = 5
	Producer_IsProducerReturnSuccess = true
	Consumer_IsConsumerReturnError   = true

	GRPC_port = 50002
)
