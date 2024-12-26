package grpc_server

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	pb "tiny-letter-subscription/cmd/grpc/pb/user"
	"tiny-letter-subscription/pkg/app"

	"google.golang.org/grpc"
)

type grpc_server struct {
	pb.UnimplementedUserServer
}

func (s *grpc_server) Create(c context.Context, request *pb.Request) (*pb.Response, error) {
	fmt.Println(request.GetId(), request.GetRole(), request.GetSubscriptionType())
	return &pb.Response{
		IsSuccess: true,
	}, nil
}

func Serve(wg *sync.WaitGroup) {
	defer wg.Done()
	config := app.GetConfig()
	addr := fmt.Sprintf("%s:%s", config.App.GRPC.User.Domain, config.App.GRPC.User.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &grpc_server{})

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
