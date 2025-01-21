package coordinator_handlers

import (
	"context"
	pb_coordinator "tiny-letter/subscription/cmd/grpc/pb/coordinator"
	"tiny-letter/subscription/pkg/db"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CoordinatorListener struct {
	pb_coordinator.UnimplementedCoordinatorListenerServer
	db *db.Repository
}

func GetCoordinatorHandlers(db *db.Repository) *CoordinatorListener {
	return &CoordinatorListener{
		db: db,
	}
}

func (l *CoordinatorListener) JoinPublication(c context.Context, req *pb_coordinator.JoinPublicationRequest) (*pb_coordinator.JoinPublicationResponse, error) {
	data := &db.JoinPublicationRequest{
		UserId:        int(req.GetUserId()),
		PublicationId: int(req.GetPublicationId()),
		IsPremium:     req.GetIsPremium(),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	subscriptionId, err := l.db.JoinPublication(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to join publication: %v", err)
	}

	return &pb_coordinator.JoinPublicationResponse{
		SubscriptionId: int32(subscriptionId),
	}, nil
}
func (l *CoordinatorListener) RollbackJoinPublication(c context.Context, req *pb_coordinator.RollbackJoinPublicationRequest) (*pb_coordinator.Response, error) {
	data := &db.RollbackJoinPublicationRequest{
		SubscriptionId: int(req.GetSubscriptionId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	err := l.db.RollbackJoinPublication(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rollback join publication: %v", err)
	}

	return &pb_coordinator.Response{IsSuccess: true}, nil
}

func (l *CoordinatorListener) LeavePublication(c context.Context, req *pb_coordinator.LeavePublicationRequest) (*pb_coordinator.LeavePublicationResponse, error) {
	data := &db.LeavePublicationRequest{
		UserId:        int(req.GetUserId()),
		PublicationId: int(req.GetPublicationId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	isPremium, err := l.db.LeavePublication(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to leave publication: %v", err)
	}

	return &pb_coordinator.LeavePublicationResponse{
		IsPremium: bool(isPremium),
	}, nil
}

func (l *CoordinatorListener) RollbackLeavePublication(c context.Context, req *pb_coordinator.RollbackLeavePublicationRequest) (*pb_coordinator.Response, error) {
	data := &db.RollbackLeavePublicationRequest{
		UserId:        int(req.GetUserId()),
		PublicationId: int(req.GetPublicationId()),
		IsPremium:     req.GetIsPremium(),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	err := l.db.RollbackLeavePublication(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rollback leave publication: %v", err)
	}

	return &pb_coordinator.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) ChangeSubscriberSubscription(c context.Context, req *pb_coordinator.ChangeSubscriberSubscriptionRequest) (*pb_coordinator.ChangeSubscriberSubscriptionResponse, error) {
	data := &db.ChangeSubscriberSubscriptionRequest{
		UserId:        int(req.GetUserId()),
		PublicationId: int(req.GetPublicationId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	subscriptionId, err := l.db.ChangeSubscriberSubscription(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to change publication plan: %v", err)
	}

	return &pb_coordinator.ChangeSubscriberSubscriptionResponse{
		SubscriptionId: int32(subscriptionId),
	}, nil
}
func (l *CoordinatorListener) RollbackChangePublicationPlan(c context.Context, req *pb_coordinator.RollbackChangeSubscriberSubscriptionRequest) (*pb_coordinator.Response, error) {
	data := &db.RollbackChangeSubscriberSubscriptionRequest{
		SubscriptionId: int(req.GetSubscriptionId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	err := l.db.RollbackChangeSubscriberPlan(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rollback change publication plan: %v", err)
	}

	return &pb_coordinator.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) ConfirmPublisherSubscription(c context.Context, req *pb_coordinator.ConfirmPublisherSubscriptionRequest) (*pb_coordinator.ConfirmPublisherSubscriptionResponse, error) {
	data := &db.ConfirmPublisherSubscriptionRequest{
		UserId: int(req.GetUserId()),
		PlanId: int(req.GetPlanId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	subscriptionId, err := l.db.ConfirmPublisherSubscription(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to confirm publisher subscription: %v", err)
	}

	return &pb_coordinator.ConfirmPublisherSubscriptionResponse{
		SubscriptionId: int32(subscriptionId),
	}, nil
}
func (l *CoordinatorListener) RollbackConfirmPublisherSubscription(c context.Context, req *pb_coordinator.RollbackConfirmPublisherSubscriptionRequest) (*pb_coordinator.Response, error) {
	data := &db.RollbackConfirmPublisherSubscriptionRequest{
		SubscriptionId: int(req.GetSubscriptionId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	err := l.db.RollbackConfirmPublisherSubscription(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rollback confirm publisher subscription: %v", err)
	}

	return &pb_coordinator.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) RevokePublisherSubscription(c context.Context, req *pb_coordinator.RevokePublisherSubscriptionRequest) (*pb_coordinator.RevokePublisherSubscriptionResponse, error) {
	data := &db.RevokePublisherSubscriptionRequest{
		UserId: int(req.GetUserId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	oldPlanId, err := l.db.RevokePublisherSubscription(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to revoke publisher subscription: %v", err)
	}

	return &pb_coordinator.RevokePublisherSubscriptionResponse{
		UserId: int32(data.UserId),
		PlanId: int32(oldPlanId),
	}, nil
}
func (l *CoordinatorListener) RollbackRevokePublisherSubscription(c context.Context, req *pb_coordinator.RollbackRevokePublisherSubscriptionRequest) (*pb_coordinator.Response, error) {
	data := &db.RollbackRevokePublisherSubscriptionRequest{
		UserId: int(req.GetUserId()),
		PlanId: int(req.GetPlanId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	err := l.db.RollbackRevokePublisherSubscription(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rollback revoke publisher subscription: %v", err)
	}

	return &pb_coordinator.Response{IsSuccess: true}, nil
}
func (l *CoordinatorListener) ChangePublisherSubscription(c context.Context, req *pb_coordinator.ChangePublisherSubscriptionRequest) (*pb_coordinator.ChangePublisherSubscriptionResponse, error) {
	data := &db.ChangePublisherSubscriptionRequest{
		UserId:        int(req.GetUserId()),
		ChangedPlanId: int(req.GetPlanId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	subscriptionId, oldPlanId, err := l.db.ChangePublisherSubscription(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to change publisher subscription: %v", err)
	}

	return &pb_coordinator.ChangePublisherSubscriptionResponse{
		SubscriptionId: int32(subscriptionId),
		OldPlanId:      int32(oldPlanId),
	}, nil
}
func (l *CoordinatorListener) RollbackChangePublisherSubscription(c context.Context, req *pb_coordinator.RollbackChangePublisherSubscriptionRequest) (*pb_coordinator.Response, error) {
	data := &db.RollbackChangePublisherSubscriptionRequest{
		SubscriptionId: int(req.GetSubscriptionId()),
		OldPlanId:      int(req.GetOldPlanId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	err := l.db.RollbackChangePublisherSubscription(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rollback change publisher subscription: %v", err)
	}

	return &pb_coordinator.Response{IsSuccess: true}, nil
}
