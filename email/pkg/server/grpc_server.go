package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	"tiny-letter/email/pkg/app"
	"tiny-letter/email/pkg/db"
	pb_email_service "tiny-letter/email/pkg/grpc/pb/email_service"
	"tiny-letter/email/pkg/handlers"

	"google.golang.org/grpc"
)

func Serve(wg *sync.WaitGroup, db *db.Repository, commConfig *app.CommConfig) {
	defer wg.Done()
	s := grpc.NewServer()
	pb_email_service.RegisterEmailServiceServer(s, handlers.GetEmailHandlers(db))

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
