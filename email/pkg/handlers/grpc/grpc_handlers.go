package grpc_handlers

import (
	pb_email_service "tiny-letter/email/cmd/grpc/pb/email_service"
	"tiny-letter/email/pkg/db"
)

type EmailServiceHandlers struct {
	pb_email_service.UnimplementedEmailServiceServer
	db *db.Repository
}

func GetEmailHandlers(db *db.Repository) *EmailServiceHandlers {
	return &EmailServiceHandlers{
		db: db,
	}
}
