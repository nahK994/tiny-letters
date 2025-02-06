package grpc_server

import (
	"fmt"
	"log"
	"net"
	"sync"
	pb_publication_authorization "tiny-letter/subscription/cmd/grpc/pb/publication_authorization"
	pb_subscription_manager "tiny-letter/subscription/cmd/grpc/pb/subscription_manager"
	"tiny-letter/subscription/pkg/app"
	"tiny-letter/subscription/pkg/db"
	content_handlers "tiny-letter/subscription/pkg/handlers/content"
	coordinator_handlers "tiny-letter/subscription/pkg/handlers/coordinator"
	mq_producer "tiny-letter/subscription/pkg/mq/producer"

	"google.golang.org/grpc"
)

func Serve(wg *sync.WaitGroup, db *db.Repository, config *app.CommConfig, producer *mq_producer.Producer) {
	defer wg.Done()

	addr := fmt.Sprintf("%s:%d", config.Domain, config.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb_publication_authorization.RegisterPublicationAuthorizationServer(s, content_handlers.GetContentHandlers(db))
	pb_subscription_manager.RegisterSubscriptionManagerServer(s, coordinator_handlers.GetCoordinatorHandlers(db, producer))

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
