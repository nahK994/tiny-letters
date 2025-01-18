package coordinator

import (
	"fmt"
	"log"
	"net"
	"sync"
	pb_coordinator "tiny-letter/subscription/cmd/grpc/pb/coordinator"
	"tiny-letter/subscription/pkg/app"
	coordinator_handlers "tiny-letter/subscription/pkg/handlers/coordinator"

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
	pb_coordinator.RegisterCoordinatorListenerServer(s, coordinator_handlers.GetOrchestratorHandlers())

	fmt.Println("Starting server...")
	fmt.Printf("Hosting server on: %s\n", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
