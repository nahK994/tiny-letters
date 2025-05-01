package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	"tiny-letter/subscription/pkg/app"
	"tiny-letter/subscription/pkg/db"
	pb_subscription_manager "tiny-letter/subscription/pkg/grpc/pb/subscription_manager"
	"tiny-letter/subscription/pkg/handlers"

	"google.golang.org/grpc"
)

func ServeGRPC(wg *sync.WaitGroup, db *db.Repository, config *app.CommConfig) {
	addr := fmt.Sprintf("%s:%d", config.Domain, config.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb_subscription_manager.RegisterSubscriptionManagerServer(s, handlers.GetGRPC_Handler(db))

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
