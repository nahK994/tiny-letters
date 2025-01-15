package main

import (
	"fmt"
	"log"
	grpc_client "tiny-letter/orchestrator/cmd/grpc/client"
	rest_server "tiny-letter/orchestrator/cmd/rest/server"
	"tiny-letter/orchestrator/pkg/app"
)

func main() {
	config := app.GetConfig()
	addr := fmt.Sprintf("%s:%d", config.Domain, config.Port)

	if err := grpc_client.IsGRPC_ClientAvailable(addr); err != nil {
		log.Fatal(err.Error())
	}
	rest_server.Serve(addr)
}
