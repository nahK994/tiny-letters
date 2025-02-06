package constants

const (
	Subscribe   = "subscribe"
	Unsubscribe = "unsubscribe"
	ChangePlan  = "change_plan"

	ConfirmationNotification = "confirmation_notification"
	PublicationNotification  = "publication_notification"

	ConfirmationEmail = "confirmation_email"
	PublicationEmail  = "publication_email"

	Domain = "localhost"

	MQ_port                          = 9092
	Producer_NumberOfRetry           = 5
	Producer_IsProducerReturnSuccess = true
	Consumer_IsConsumerReturnError   = true
)
