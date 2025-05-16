package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	"tiny-letter/auth/pkg/app"
	"tiny-letter/auth/pkg/db"
	pb_auth_manager "tiny-letter/auth/pkg/grpc/pb/auth_manager"
	"tiny-letter/auth/pkg/handlers"

	"google.golang.org/grpc"
)

func ServeGRPC(wg *sync.WaitGroup, db *db.Repository, commConfig *app.CommConfig) {
	defer wg.Done()
	s := grpc.NewServer()
	pb_auth_manager.RegisterAuthManagerServer(s, handlers.GetCoordinatorHandlers(db))

	addr := fmt.Sprintf("%s:%d", commConfig.Domain, commConfig.Port)
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
