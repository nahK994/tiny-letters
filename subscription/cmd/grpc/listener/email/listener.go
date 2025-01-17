package email

import (
	"fmt"
	"log"
	"net"
	"sync"
	pb_email "tiny-letter/subscription/cmd/grpc/pb/email"
	"tiny-letter/subscription/pkg/app"
	email_handlers "tiny-letter/subscription/pkg/handlers/email"

	"google.golang.org/grpc"
)

func Listen(wg *sync.WaitGroup, commConfig *app.GRPCConfig) {
	defer wg.Done()
	addr := fmt.Sprintf("%s:%d", commConfig.Domain, commConfig.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb_email.RegisterEmailListenerServer(s, email_handlers.GetEmailHandlers())

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
