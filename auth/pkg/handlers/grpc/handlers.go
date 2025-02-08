package grpc_handlers

import (
	"context"
	pb_auth_manager "tiny-letter/auth/cmd/grpc/pb/auth_manager"
	"tiny-letter/auth/pkg/db"
	"tiny-letter/auth/pkg/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CoordinatorListener struct {
	pb_auth_manager.UnimplementedAuthManagerServer
	db *db.Repository
}

func GetCoordinatorHandlers(db *db.Repository) *CoordinatorListener {
	return &CoordinatorListener{
		db: db,
	}
}

func (l *CoordinatorListener) JoinPublication(c context.Context, req *pb_auth_manager.JoinPublicationRequest) (*pb_auth_manager.Response, error) {
	data := &models.JoinPublicationRequest{
		UserId:        int(req.GetUserId()),
		PublicationId: int(req.GetPublicationId()),
		IsPremium:     req.GetIsPremium(),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.JoinPublication(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to join publication: %v", err)
	}

	return &pb_auth_manager.Response{IsSuccess: true}, nil

}
func (l *CoordinatorListener) LeavePublication(c context.Context, req *pb_auth_manager.LeavePublicationRequest) (*pb_auth_manager.Response, error) {
	data := &models.LeavePublicationRequest{
		UserId:        int(req.GetUserId()),
		PublicationId: int(req.GetPublicationId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.LeavePublication(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to leave publication: %v", err)
	}

	return &pb_auth_manager.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) ChangeSubscriberSubscription(c context.Context, req *pb_auth_manager.ChangeSubscriberSubscriptionRequest) (*pb_auth_manager.Response, error) {
	data := &models.ChangeSubscriberSubscriptionRequest{
		UserId:        int(req.GetUserId()),
		PublicationId: int(req.GetPublicationId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.ChangeSubscriberSubscription(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to change subscriber subscription: %v", err)
	}

	return &pb_auth_manager.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) ConfirmPublisherSubscription(c context.Context, req *pb_auth_manager.ConfirmPublisherSubscriptionRequest) (*pb_auth_manager.Response, error) {
	data := &models.ConfirmPublisherSubscriptionRequest{
		UserId: int(req.GetUserId()),
		PlanId: int(req.GetPlanId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.ConfirmPublisherSubscription(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to confirm publisher subscription: %v", err)
	}

	return &pb_auth_manager.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) RevokePublisherSubscription(c context.Context, req *pb_auth_manager.RevokePublisherSubscriptionRequest) (*pb_auth_manager.Response, error) {
	data := &models.RevokePublisherSubscriptionRequest{
		UserId: int(req.GetUserId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.RevokePublisherSubscription(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to revoke publisher subscription: %v", err)
	}

	return &pb_auth_manager.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) ChangePublisherSubscription(c context.Context, req *pb_auth_manager.ChangePublisherSubscriptionRequest) (*pb_auth_manager.Response, error) {
	data := &models.ChangePublisherSubscriptionRequest{
		UserId:        int(req.GetUserId()),
		ChangedPlanId: int(req.GetPlanId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.ChangePublisherSubscription(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to change publisher subscription: %v", err)
	}

	return &pb_auth_manager.Response{IsSuccess: true}, nil
}
