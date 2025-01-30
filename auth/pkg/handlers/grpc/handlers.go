package grpc_handlers

import (
	"context"
	pb_subscription "tiny-letter/auth/cmd/grpc/pb/subscription"
	"tiny-letter/auth/pkg/db"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CoordinatorListener struct {
	pb_subscription.UnimplementedNotifyAuthServer
	db *db.Repository
}

func GetCoordinatorHandlers(db *db.Repository) *CoordinatorListener {
	return &CoordinatorListener{
		db: db,
	}
}

func (l *CoordinatorListener) JoinPublication(c context.Context, req *pb_subscription.JoinPublicationRequest) (*pb_subscription.Response, error) {
	data := &db.JoinPublicationRequest{
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

	return &pb_subscription.Response{IsSuccess: true}, nil

}
func (l *CoordinatorListener) LeavePublication(c context.Context, req *pb_subscription.LeavePublicationRequest) (*pb_subscription.Response, error) {
	data := &db.LeavePublicationRequest{
		UserId:        int(req.GetUserId()),
		PublicationId: int(req.GetPublicationId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.LeavePublication(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to leave publication: %v", err)
	}

	return &pb_subscription.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) ChangeSubscriberSubscription(c context.Context, req *pb_subscription.ChangeSubscriberSubscriptionRequest) (*pb_subscription.Response, error) {
	data := &db.ChangeSubscriberSubscriptionRequest{
		UserId:        int(req.GetUserId()),
		PublicationId: int(req.GetPublicationId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.ChangeSubscriberSubscription(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to change subscriber subscription: %v", err)
	}

	return &pb_subscription.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) ConfirmPublisherSubscription(c context.Context, req *pb_subscription.ConfirmPublisherSubscriptionRequest) (*pb_subscription.Response, error) {
	data := &db.ConfirmPublisherSubscriptionRequest{
		UserId: int(req.GetUserId()),
		PlanId: int(req.GetPlanId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.ConfirmPublisherSubscription(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to confirm publisher subscription: %v", err)
	}

	return &pb_subscription.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) RevokePublisherSubscription(c context.Context, req *pb_subscription.RevokePublisherSubscriptionRequest) (*pb_subscription.Response, error) {
	data := &db.RevokePublisherSubscriptionRequest{
		UserId: int(req.GetUserId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.RevokePublisherSubscription(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to revoke publisher subscription: %v", err)
	}

	return &pb_subscription.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) ChangePublisherSubscription(c context.Context, req *pb_subscription.ChangePublisherSubscriptionRequest) (*pb_subscription.Response, error) {
	data := &db.ChangePublisherSubscriptionRequest{
		UserId:        int(req.GetUserId()),
		ChangedPlanId: int(req.GetPlanId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	if err := l.db.ChangePublisherSubscription(data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to change publisher subscription: %v", err)
	}

	return &pb_subscription.Response{IsSuccess: true}, nil
}
