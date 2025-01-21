package grpc_server

import (
	"fmt"
	"log"
	"net"
	"sync"
	pb_auth "tiny-letter/auth/cmd/grpc/pb"
	"tiny-letter/auth/pkg/db"
	grpc_handlers "tiny-letter/auth/pkg/handlers/grpc"

	"google.golang.org/grpc"
)

func Serve(wg *sync.WaitGroup, db *db.Repository, addr string) {
	defer wg.Done()
	s := grpc.NewServer()
	pb_auth.RegisterNotifyAuthServer(s, grpc_handlers.GetCoordinatorHandlers(db))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
