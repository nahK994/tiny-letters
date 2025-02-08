package main

import (
	"log"
	grpc_client "tiny-letter/content/cmd/grpc/client"
	rest_server "tiny-letter/content/cmd/rest/server"
	"tiny-letter/content/pkg/app"
	"tiny-letter/content/pkg/db"
	mq_producer "tiny-letter/content/pkg/mq"
)

func main() {
	config := app.GetConfig()

	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := grpc_client.IsGRPC_ClientAvailable(&config.GRPC); err != nil {
		log.Fatal(err)
	}

	producer, err := mq_producer.NewProducer(&config.MQ)
	if err != nil {
		log.Fatalf("Failed to connect to MQ: %v", err)
	}

	rest_server.Serve(db, &config.REST, producer)
}
