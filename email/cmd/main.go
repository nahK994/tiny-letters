package main

import (
	"log"
	"sync"
	"tiny-letter/email/pkg/app"
	"tiny-letter/email/pkg/db"
	"tiny-letter/email/pkg/handlers"
	"tiny-letter/email/pkg/mq"
	"tiny-letter/email/pkg/server"
)

func main() {
	config := app.GetConfig()

	db, err := db.Init(&config.DB)
	if err != nil {
		log.Fatal(err)
	}

	consumerHandlers := handlers.NewConsumerHandlers()
	mq, err := mq.NewConsumer(consumerHandlers, &config.MQ)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go mq.StartConsuming(&wg)
	go server.Serve(&wg, db, &config.GRPC)
	wg.Wait()
}
