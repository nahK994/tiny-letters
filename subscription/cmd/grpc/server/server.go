package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	pb_content "tiny-letter/subscription/cmd/grpc/pb/content"
	pb_coordinator "tiny-letter/subscription/cmd/grpc/pb/coordinator"
	"tiny-letter/subscription/pkg/app"
	content_handlers "tiny-letter/subscription/pkg/handlers/content"
	coordinator_handlers "tiny-letter/subscription/pkg/handlers/coordinator"

	"google.golang.org/grpc"
)

func Serve(wg *sync.WaitGroup) {
	defer wg.Done()
	config := app.GetConfig()
	addr := fmt.Sprintf("%s:%d", config.App.Domain, config.App.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb_content.RegisterContentListenerServer(s, content_handlers.GetContentHandlers())
	pb_coordinator.RegisterCoordinatorListenerServer(s, coordinator_handlers.GetCoordinatorHandlers())

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
