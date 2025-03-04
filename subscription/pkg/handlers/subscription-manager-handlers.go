package handlers

import (
	"context"
	"tiny-letter/subscription/pkg/db"
	pb_subscription_manager "tiny-letter/subscription/pkg/grpc/pb/subscription_manager"
	"tiny-letter/subscription/pkg/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SubscriptionManagerHandler struct {
	pb_subscription_manager.UnimplementedSubscriptionManagerServer
	db *db.Repository
}

func GetSubscriptionManagerHandler(db *db.Repository) *SubscriptionManagerHandler {
	return &SubscriptionManagerHandler{
		db: db,
	}
}

func (l *SubscriptionManagerHandler) JoinPublication(c context.Context, req *pb_subscription_manager.JoinPublicationRequest) (*pb_subscription_manager.JoinPublicationResponse, error) {
	data := &models.JoinPublicationRequest{
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

	return &pb_subscription_manager.JoinPublicationResponse{
		SubscriptionId: int32(subscriptionId),
	}, nil
}
func (l *SubscriptionManagerHandler) RollbackJoinPublication(c context.Context, req *pb_subscription_manager.RollbackJoinPublicationRequest) (*pb_subscription_manager.Response, error) {
	data := &models.RollbackJoinPublicationRequest{
		SubscriptionId: int(req.GetSubscriptionId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	err := l.db.RollbackJoinPublication(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rollback join publication: %v", err)
	}

	return &pb_subscription_manager.Response{IsSuccess: true}, nil
}

func (l *SubscriptionManagerHandler) LeavePublication(c context.Context, req *pb_subscription_manager.LeavePublicationRequest) (*pb_subscription_manager.LeavePublicationResponse, error) {
	data := &models.LeavePublicationRequest{
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

	return &pb_subscription_manager.LeavePublicationResponse{
		IsPremium: bool(isPremium),
	}, nil
}

func (l *SubscriptionManagerHandler) RollbackLeavePublication(c context.Context, req *pb_subscription_manager.RollbackLeavePublicationRequest) (*pb_subscription_manager.Response, error) {
	data := &models.RollbackLeavePublicationRequest{
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

	return &pb_subscription_manager.Response{IsSuccess: true}, nil
}

func (l *SubscriptionManagerHandler) ChangeSubscriberSubscription(c context.Context, req *pb_subscription_manager.ChangeSubscriberSubscriptionRequest) (*pb_subscription_manager.ChangeSubscriberSubscriptionResponse, error) {
	data := &models.ChangeSubscriberSubscriptionRequest{
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

	return &pb_subscription_manager.ChangeSubscriberSubscriptionResponse{
		SubscriptionId: int32(subscriptionId),
	}, nil
}

func (l *SubscriptionManagerHandler) RollbackChangePublicationPlan(c context.Context, req *pb_subscription_manager.RollbackChangeSubscriberSubscriptionRequest) (*pb_subscription_manager.Response, error) {
	data := &models.RollbackChangeSubscriberSubscriptionRequest{
		SubscriptionId: int(req.GetSubscriptionId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	err := l.db.RollbackChangeSubscriberPlan(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rollback change publication plan: %v", err)
	}

	return &pb_subscription_manager.Response{IsSuccess: true}, nil
}

func (l *SubscriptionManagerHandler) ConfirmPublisherSubscription(c context.Context, req *pb_subscription_manager.ConfirmPublisherSubscriptionRequest) (*pb_subscription_manager.ConfirmPublisherSubscriptionResponse, error) {
	data := &models.ConfirmPublisherSubscriptionRequest{
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

	return &pb_subscription_manager.ConfirmPublisherSubscriptionResponse{
		SubscriptionId: int32(subscriptionId),
	}, nil
}

func (l *SubscriptionManagerHandler) RollbackConfirmPublisherSubscription(c context.Context, req *pb_subscription_manager.RollbackConfirmPublisherSubscriptionRequest) (*pb_subscription_manager.Response, error) {
	data := &models.RollbackConfirmPublisherSubscriptionRequest{
		SubscriptionId: int(req.GetSubscriptionId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	err := l.db.RollbackConfirmPublisherSubscription(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rollback confirm publisher subscription: %v", err)
	}

	return &pb_subscription_manager.Response{IsSuccess: true}, nil
}

func (l *SubscriptionManagerHandler) RevokePublisherSubscription(c context.Context, req *pb_subscription_manager.RevokePublisherSubscriptionRequest) (*pb_subscription_manager.RevokePublisherSubscriptionResponse, error) {
	data := &models.RevokePublisherSubscriptionRequest{
		UserId: int(req.GetUserId()),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}

	oldPlanId, err := l.db.RevokePublisherSubscription(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to revoke publisher subscription: %v", err)
	}

	return &pb_subscription_manager.RevokePublisherSubscriptionResponse{
		PlanId: int32(oldPlanId),
	}, nil
}

func (l *SubscriptionManagerHandler) RollbackRevokePublisherSubscription(c context.Context, req *pb_subscription_manager.RollbackRevokePublisherSubscriptionRequest) (*pb_subscription_manager.Response, error) {
	data := &models.RollbackRevokePublisherSubscriptionRequest{
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

	return &pb_subscription_manager.Response{IsSuccess: true}, nil
}

func (l *SubscriptionManagerHandler) ChangePublisherSubscription(c context.Context, req *pb_subscription_manager.ChangePublisherSubscriptionRequest) (*pb_subscription_manager.ChangePublisherSubscriptionResponse, error) {
	data := &models.ChangePublisherSubscriptionRequest{
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

	return &pb_subscription_manager.ChangePublisherSubscriptionResponse{
		SubscriptionId: int32(subscriptionId),
		OldPlanId:      int32(oldPlanId),
	}, nil
}

func (l *SubscriptionManagerHandler) RollbackChangePublisherSubscription(c context.Context, req *pb_subscription_manager.RollbackChangePublisherSubscriptionRequest) (*pb_subscription_manager.Response, error) {
	data := &models.RollbackChangePublisherSubscriptionRequest{
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

	return &pb_subscription_manager.Response{IsSuccess: true}, nil
}

func (n *SubscriptionManagerHandler) GetContentSubscribers(c context.Context, req *pb_subscription_manager.GetContentSubscribersRequest) (*pb_subscription_manager.GetContentSubscribersResponse, error) {
	data := &models.GetContentSubscribersRequest{
		PublicationId: int(req.PublicationId),
	}
	if err := data.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
	}
	subscriberIds, err := n.db.GetContentSubscribers(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get subscriberIds: %v", err)
	}
	return &pb_subscription_manager.GetContentSubscribersResponse{
		SubscriberIds: subscriberIds,
	}, nil
}
