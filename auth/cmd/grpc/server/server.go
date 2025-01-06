package grpc_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	pb_auth "tiny-letter/auth/cmd/grpc/pb"
	"tiny-letter/auth/pkg/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpc_server struct {
	pb_auth.UnimplementedNotificationListenerServer
}

func (server *grpc_server) PublisherAction(context.Context, *pb_auth.PublisherActionRequest) (*pb_auth.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublisherAction not implemented")
}
func (server *grpc_server) SubscriberAction(context.Context, *pb_auth.SubscriberActionRequest) (*pb_auth.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscriberAction not implemented")
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
	pb_auth.RegisterNotificationListenerServer(s, &grpc_server{})

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
