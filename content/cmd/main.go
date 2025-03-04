package main

import (
	"log"
	"tiny-letter/content/pkg/app"
	"tiny-letter/content/pkg/db"
	mq_producer "tiny-letter/content/pkg/mq"
	"tiny-letter/content/pkg/server"
)

func main() {
	config := app.GetConfig()

	db, err := db.Init(&config.DB)
	if err != nil {
		log.Fatal(err)
	}

	producer, err := mq_producer.NewProducer(&config.MQ)
	if err != nil {
		log.Fatalf("Failed to connect to MQ: %v", err)
	}

	server.Serve(db, &config.REST, producer)
}
