package grpc_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	pb "tiny-letter/auth/cmd/grpc/pb/subscriber-auth"
	"tiny-letter/auth/pkg/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpc_server struct {
	pb.UnimplementedSubscriptionAuthServer
}

func (server *grpc_server) JoinPublication(context.Context, *pb.ManagePublicationSubscriptionRequest) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPublication not implemented")
}
func (server *grpc_server) LeavePublication(context.Context, *pb.ManagePublicationSubscriptionRequest) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeavePublication not implemented")
}
func (server *grpc_server) SubscribePublisherPlan(context.Context, *pb.PublisherSubscriptionRequest) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribePublisherPlan not implemented")
}
func (server *grpc_server) UnsubscribePublisherPlan(context.Context, *pb.PublisherUnsubscriptionRequest) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubscribePublisherPlan not implemented")
}
func (server *grpc_server) ChangePublicationSubscription(context.Context, *pb.ChangePublicationSubscriptionRequest) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublicationSubscription not implemented")
}
func (server *grpc_server) ChangePublisherSubscriptionPlan(context.Context, *pb.ChangePublisherPlanRequest) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublisherSubscriptionPlan not implemented")
}

func Serve(wg *sync.WaitGroup) {
	defer wg.Done()
	config := app.GetConfig()
	addr := fmt.Sprintf("%s:%s", config.App.GRPC.Subscriber_Auth.Domain, config.App.GRPC.Subscriber_Auth.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSubscriptionAuthServer(s, &grpc_server{})

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
