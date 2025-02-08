package constants

const (
	Domain = "localhost"

	JWT_secret      = "secret"
	JWT_exp_minutes = 60

	REST_port              = 8001
	GRPC_subscription_port = 50002
	DB_port                = 5001
	MQ_port                = 9092

	DB_user     = "user"
	DB_password = "password"
	DB_name     = "content_db"

	MQ_NumberOfRetry           = 5
	MQ_IsProducerReturnSuccess = true
	MQ_topic                   = "publication_notification"
)
