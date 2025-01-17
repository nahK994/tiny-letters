package orchestrator_handlers

import (
	"context"
	pb_orchestrator "tiny-letter/subscription/cmd/grpc/pb/orchestrator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrchestratorListener struct {
	pb_orchestrator.UnimplementedOrchestratorListenerServer
}

func GetOrchestratorHandlers() *OrchestratorListener {
	return &OrchestratorListener{}
}

func (l *OrchestratorListener) JoinPublication(context.Context, *pb_orchestrator.JoinPublicationRequest) (*pb_orchestrator.JoinPublicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPublication not implemented")
}
func (l *OrchestratorListener) RollbackJoinPublication(context.Context, *pb_orchestrator.RollbackJoinPublicationRequest) (*pb_orchestrator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackJoinPublication not implemented")
}
func (l *OrchestratorListener) LeavePublication(context.Context, *pb_orchestrator.LeavePublicationRequest) (*pb_orchestrator.LeavePublicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeavePublication not implemented")
}
func (l *OrchestratorListener) RollbackLeavePublication(context.Context, *pb_orchestrator.RollbackLeavePublicationRequest) (*pb_orchestrator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackLeavePublication not implemented")
}
func (l *OrchestratorListener) ChangePublicationPlan(context.Context, *pb_orchestrator.ChangePublicationPlanRequest) (*pb_orchestrator.ChangePublicationPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublicationPlan not implemented")
}
func (l *OrchestratorListener) RollbackChangePublicationPlan(context.Context, *pb_orchestrator.RollbackChangePublicationPlanRequest) (*pb_orchestrator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackChangePublicationPlan not implemented")
}
func (l *OrchestratorListener) ConfirmPublisherSubscription(context.Context, *pb_orchestrator.ConfirmPublisherSubscriptionRequest) (*pb_orchestrator.ConfirmPublisherSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmPublisherSubscription not implemented")
}
func (l *OrchestratorListener) RollbackConfirmPublisherSubscription(context.Context, *pb_orchestrator.RollbackConfirmPublisherSubscriptionRequest) (*pb_orchestrator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackConfirmPublisherSubscription not implemented")
}
func (l *OrchestratorListener) RevokePublisherSubscription(context.Context, *pb_orchestrator.RevokePublisherSubscriptionRequest) (*pb_orchestrator.RevokePublisherSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokePublisherSubscription not implemented")
}
func (l *OrchestratorListener) RollbackRevokePublisherSubscription(context.Context, *pb_orchestrator.RollbackRevokePublisherSubscriptionRequest) (*pb_orchestrator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackRevokePublisherSubscription not implemented")
}
func (l *OrchestratorListener) ChangePublisherSubscription(context.Context, *pb_orchestrator.ChangePublisherSubscriptionRequest) (*pb_orchestrator.ChangePublisherSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublisherSubscription not implemented")
}
func (l *OrchestratorListener) RollbackChangePublisherSubscription(context.Context, *pb_orchestrator.RollbackChangePublisherSubscriptionRequest) (*pb_orchestrator.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackChangePublisherSubscription not implemented")
}
