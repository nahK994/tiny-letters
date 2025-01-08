package main

import (
	"log"
	grpc_client "tiny-letter/orchestrator/cmd/grpc/client"
	rest_server "tiny-letter/orchestrator/cmd/rest/server"
)

func main() {
	if err := grpc_client.InitAuthClient(); err != nil {
		grpc_client.ShutdownAuthClient()
		log.Fatal(err.Error())
	}

	if err := grpc_client.InitSubscriptionClient(); err != nil {
		grpc_client.ShutdownSubscriptionClient()
		log.Fatal(err.Error())
	}

	rest_server.Serve()
}
