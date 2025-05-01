package handlers

import (
	"context"
	"tiny-letter/auth/pkg/db"
	pb_auth_manager "tiny-letter/auth/pkg/grpc/pb/auth_manager"
	"tiny-letter/auth/pkg/models"
)

type GRPC_Handler struct {
	pb_auth_manager.UnimplementedAuthManagerServer
	db *db.Repository
}

func GetCoordinatorHandlers(db *db.Repository) *GRPC_Handler {
	return &GRPC_Handler{
		db: db,
	}
}

func (h *GRPC_Handler) CreatePublisher(c context.Context, req *pb_auth_manager.CreatePublisherRequest) (*pb_auth_manager.Response, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user_id, err := h.db.CreateUser(&models.UserRegistration{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "publisher",
	})
	if err != nil {
		return nil, err
	}
	return &pb_auth_manager.Response{
		UserId: int32(user_id),
	}, nil
}

func (h *GRPC_Handler) RollbackCreatePublisher(c context.Context, req *pb_auth_manager.RollbackCreatePublisherRequest) (*pb_auth_manager.Response, error) {
	err := h.db.RollbackCreateUser(int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &pb_auth_manager.Response{
		UserId: req.UserId,
	}, nil
}
