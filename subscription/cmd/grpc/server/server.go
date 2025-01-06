package grpc_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	pb_subscription "tiny-letter/subscription/cmd/grpc/pb"
	"tiny-letter/subscription/pkg/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpc_server struct {
	pb_subscription.UnimplementedNotificationListenerServer
}

func (server *grpc_server) JoinPublication(context.Context, *pb_subscription.JoinPublicationRequest) (*pb_subscription.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPublication not implemented")
}
func (server *grpc_server) LeavePublication(context.Context, *pb_subscription.LeavePublicationRequest) (*pb_subscription.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeavePublication not implemented")
}
func (server *grpc_server) ChangePublicationPlan(context.Context, *pb_subscription.ChangePublicationPlanRequest) (*pb_subscription.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublicationPlan not implemented")
}
func (server *grpc_server) SubscribePublisherPlan(context.Context, *pb_subscription.SubscribePublisherRequest) (*pb_subscription.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribePublisherPlan not implemented")
}
func (server *grpc_server) UnsubscribePublisherPlan(context.Context, *pb_subscription.UnsubscribePublisherRequest) (*pb_subscription.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubscribePublisherPlan not implemented")
}
func (server *grpc_server) ChangePublisherPlan(context.Context, *pb_subscription.ChangePublisherPlanRequest) (*pb_subscription.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublisherPlan not implemented")
}

func Serve(wg *sync.WaitGroup) {
	defer wg.Done()
	config := app.GetConfig()
	addr := fmt.Sprintf("%s:%d", config.App.GRPC.Subscriber_Auth.Domain, config.App.GRPC.Subscriber_Auth.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb_subscription.RegisterNotificationListenerServer(s, &grpc_server{})

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
