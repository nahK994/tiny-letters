package coordinator_handlers

import (
	"context"
	pb_coordinator "tiny-letter/subscription/cmd/grpc/pb/coordinator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CoordinatorListener struct {
	pb_coordinator.UnimplementedCoordinatorListenerServer
}

func GetOrchestratorHandlers() *CoordinatorListener {
	return &CoordinatorListener{}
}

func (l *CoordinatorListener) JoinPublication(context.Context, *pb_coordinator.JoinPublicationRequest) (*pb_coordinator.JoinPublicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPublication not implemented")
}
func (l *CoordinatorListener) RollbackJoinPublication(context.Context, *pb_coordinator.RollbackJoinPublicationRequest) (*pb_coordinator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackJoinPublication not implemented")
}
func (l *CoordinatorListener) LeavePublication(context.Context, *pb_coordinator.LeavePublicationRequest) (*pb_coordinator.LeavePublicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeavePublication not implemented")
}
func (l *CoordinatorListener) RollbackLeavePublication(context.Context, *pb_coordinator.RollbackLeavePublicationRequest) (*pb_coordinator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackLeavePublication not implemented")
}
func (l *CoordinatorListener) ChangePublicationPlan(context.Context, *pb_coordinator.ChangePublicationPlanRequest) (*pb_coordinator.ChangePublicationPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublicationPlan not implemented")
}
func (l *CoordinatorListener) RollbackChangePublicationPlan(context.Context, *pb_coordinator.RollbackChangePublicationPlanRequest) (*pb_coordinator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackChangePublicationPlan not implemented")
}
func (l *CoordinatorListener) ConfirmPublisherSubscription(context.Context, *pb_coordinator.ConfirmPublisherSubscriptionRequest) (*pb_coordinator.ConfirmPublisherSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmPublisherSubscription not implemented")
}
func (l *CoordinatorListener) RollbackConfirmPublisherSubscription(context.Context, *pb_coordinator.RollbackConfirmPublisherSubscriptionRequest) (*pb_coordinator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackConfirmPublisherSubscription not implemented")
}
func (l *CoordinatorListener) RevokePublisherSubscription(context.Context, *pb_coordinator.RevokePublisherSubscriptionRequest) (*pb_coordinator.RevokePublisherSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokePublisherSubscription not implemented")
}
func (l *CoordinatorListener) RollbackRevokePublisherSubscription(context.Context, *pb_coordinator.RollbackRevokePublisherSubscriptionRequest) (*pb_coordinator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackRevokePublisherSubscription not implemented")
}
func (l *CoordinatorListener) ChangePublisherSubscription(context.Context, *pb_coordinator.ChangePublisherSubscriptionRequest) (*pb_coordinator.ChangePublisherSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublisherSubscription not implemented")
}
func (l *CoordinatorListener) RollbackChangePublisherSubscription(context.Context, *pb_coordinator.RollbackChangePublisherSubscriptionRequest) (*pb_coordinator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackChangePublisherSubscription not implemented")
}
