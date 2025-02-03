package grpc_server

import (
	"fmt"
	"log"
	"net"
	"sync"
	pb_content "tiny-letter/subscription/cmd/grpc/pb/content"
	pb_coordinator "tiny-letter/subscription/cmd/grpc/pb/coordinator"
	pb_notification "tiny-letter/subscription/cmd/grpc/pb/notification"
	"tiny-letter/subscription/pkg/db"
	content_handlers "tiny-letter/subscription/pkg/handlers/content"
	coordinator_handlers "tiny-letter/subscription/pkg/handlers/coordinator"
	notification_handlers "tiny-letter/subscription/pkg/handlers/notification"

	"google.golang.org/grpc"
)

func Serve(wg *sync.WaitGroup, db *db.Repository, addr string) {
	defer wg.Done()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb_content.RegisterContentListenerServer(s, content_handlers.GetContentHandlers(db))
	pb_coordinator.RegisterCoordinatorListenerServer(s, coordinator_handlers.GetCoordinatorHandlers(db))
	pb_notification.RegisterNotificationListenerServer(s, notification_handlers.GetNotificationHandlers(db))

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
