package constants

const (
	PublisherSubscribe   = "publisher_subscribe"
	PublisherUnsubscribe = "publisher_unsubscribe"
	PublisherChangePlan  = "publisher_change_plan"

	SubscriberSubscribe   = "subscriber_subscribe"
	SubscriberUnsubscribe = "subscriber_unsubscribe"
	SubscriberChangePlan  = "subscriber_change_plan"

	ConfirmationEmail = "confirmation_email"
	PublicationEmail  = "publication_email"

	Domain = "localhost"

	MQ_port                        = 9092
	Consumer_IsConsumerReturnError = true

	GRPC_port   = 50003
	DB_port     = 5003
	DB_user     = "user"
	DB_password = "password"
	DB_name     = "email_db"
)
