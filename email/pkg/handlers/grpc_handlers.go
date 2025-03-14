package handlers

import (
	pb_email_service "tiny-letter/email/pkg/grpc/pb/email_service"
	"tiny-letter/email/pkg/models"
)

type Repository interface {
	CreateUser(data models.UserRegistrationRequest) error
}

type EmailServiceHandlers struct {
	pb_email_service.UnimplementedEmailServiceServer
	db Repository
}

func GetEmailHandlers(db Repository) *EmailServiceHandlers {
	return &EmailServiceHandlers{
		db: db,
	}
}
