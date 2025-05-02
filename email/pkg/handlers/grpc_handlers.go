package handlers

import (
	"context"
	"tiny-letter/email/pkg/db"
	pb_email_service "tiny-letter/email/pkg/grpc/pb/email_service"
	"tiny-letter/email/pkg/models"
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

func (h *EmailServiceHandlers) OnboardUser(c context.Context, req *pb_email_service.OnboardUserRequest) (*pb_email_service.Response, error) {
	user := &models.UserContactEmail{
		Email:  req.GetEmail(),
		UserId: int(req.GetUserId()),
		Role:   req.GetRole(),
	}

	if err := h.db.StoreUserContactEmail(user); err != nil {
		return nil, err
	}

	if user.Role == "subscriber" {
		// TODO: Add logic to send email to subscriber
	} else if user.Role == "publisher" {
		// TODO: Add logic to send email to publisher
	}

	return &pb_email_service.Response{
		IsSuccess: true,
	}, nil
}
