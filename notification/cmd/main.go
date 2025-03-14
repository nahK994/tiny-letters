package main

import (
	"log"
	"sync"
	"tiny-letter/notification/pkg/app"
	"tiny-letter/notification/pkg/client"
	"tiny-letter/notification/pkg/handlers"
	"tiny-letter/notification/pkg/mq"
)

func main() {
	config := app.GetConfig()
	client, err := client.ConnectSubscriptionClient(&config.GRPC)
	if err != nil {
		log.Fatal(err.Error())
	}

	producer, err := mq.NewProducer(&config.MQ)
	if err != nil {
		log.Fatal(err.Error())
	}
	handlers := handlers.NewHandler(producer, client)

	consumer, err := mq.NewConsumer(handlers, &config.MQ)
	if err != nil {
		log.Fatal(err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go consumer.StartConsuming(&wg)
	wg.Wait()
}
