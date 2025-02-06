package main

import (
	"log"
	grpc_client "tiny-letter/coordinator/cmd/grpc/client"
	rest_server "tiny-letter/coordinator/cmd/rest/server"
	"tiny-letter/coordinator/pkg/app"
	mq_producer "tiny-letter/coordinator/pkg/mq"
)

func main() {
	config := app.GetConfig()

	if err := grpc_client.IsGRPC_ClientAvailable(&config.GRPC); err != nil {
		log.Fatal(err.Error())
	}

	producer, err := mq_producer.NewProducer(&config.MQ)
	if err != nil {
		log.Fatalf("Failed to connect to MQ: %v", err)
	}
	rest_server.Serve(&config.REST, producer)
}
