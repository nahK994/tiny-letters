package main

import (
	"log"
	"tiny-letter/orchestrator/pkg/app"
	"tiny-letter/orchestrator/pkg/grpc/client"
	"tiny-letter/orchestrator/pkg/mq"
	"tiny-letter/orchestrator/pkg/server"
)

func main() {
	config := app.GetConfig()

	if err := client.ConnectGRPC(&config.GRPC); err != nil {
		log.Fatal(err.Error())
	}

	producer, err := mq.NewProducer(&config.MQ)
	if err != nil {
		log.Fatalf("Failed to connect to MQ: %v", err)
	}

	server.Serve(&config.REST, producer)
}
