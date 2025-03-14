package server

import (
	"fmt"
	"log"
	"net"
	"tiny-letter/subscription/pkg/app"
	"tiny-letter/subscription/pkg/db"
	pb_publication_authorization "tiny-letter/subscription/pkg/grpc/pb/publication_authorization"
	pb_subscription_manager "tiny-letter/subscription/pkg/grpc/pb/subscription_manager"
	"tiny-letter/subscription/pkg/handlers"

	"google.golang.org/grpc"
)

func Serve(db *db.Repository, config *app.CommConfig) {
	addr := fmt.Sprintf("%s:%d", config.Domain, config.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb_publication_authorization.RegisterPublicationAuthorizationServer(s, handlers.GetPublicationAuthorizationHandler(db))
	pb_subscription_manager.RegisterSubscriptionManagerServer(s, handlers.GetSubscriptionManagerHandler(db))

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
