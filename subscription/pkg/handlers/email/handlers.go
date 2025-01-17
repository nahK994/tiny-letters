package email_handlers

import pb_email "tiny-letter/subscription/cmd/grpc/pb/email"

type EmailListener struct {
	pb_email.UnimplementedEmailListenerServer
}

func GetEmailHandlers() *EmailListener {
	return &EmailListener{}
}
