package main

import (
	"log"
	"sync"
	"tiny-letter/notification/pkg/app"
	"tiny-letter/notification/pkg/handlers"
	"tiny-letter/notification/pkg/mq"
)

func main() {
	config := app.GetConfig()

	producer, err := mq.NewProducer(&config.MQ)
	if err != nil {
		log.Fatal(err.Error())
	}
	handlers := handlers.NewHandler(producer)

	consumer, err := mq.NewConsumer(handlers, &config.MQ)
	if err != nil {
		log.Fatal(err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go consumer.StartConsuming(&wg)
	wg.Wait()
}
