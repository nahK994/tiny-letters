package main

import (
	"log"
	"sync"
	grpc_server "tiny-letter/email/cmd/grpc/server"
	"tiny-letter/email/pkg/app"
	"tiny-letter/email/pkg/db"
	mq_handlers "tiny-letter/email/pkg/handlers/mq"
	mq_consumer "tiny-letter/email/pkg/mq"
)

func main() {
	config := app.GetConfig()

	db, err := db.Init(&config.DB)
	if err != nil {
		log.Fatal(err)
	}

	consumerHandlers := mq_handlers.New_ConsumerHandlers(db)
	mq, err := mq_consumer.NewConsumer(consumerHandlers, &config.MQ)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go mq.StartConsuming(&wg)
	go grpc_server.Serve(&wg, db, &config.GRPC)
	wg.Wait()
}
