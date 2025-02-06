package main

import (
	"log"
	"sync"
	grpc_client "tiny-letter/notification/cmd/grpc/client"
	"tiny-letter/notification/pkg/app"
	"tiny-letter/notification/pkg/handlers"
	mq_consumer "tiny-letter/notification/pkg/mq/consumer"
	mq_producer "tiny-letter/notification/pkg/mq/producer"
)

func main() {
	config := app.GetConfig()
	if err := grpc_client.IsGRPC_ClientAvailable(&config.GRPC); err != nil {
		log.Fatal(err.Error())
	}

	producer, err := mq_producer.NewProducer(&config.MQ)
	if err != nil {
		log.Fatal(err.Error())
	}
	handlers := handlers.NewHandler(producer)
	consumer, err := mq_consumer.NewConsumer(handlers, &config.MQ)
	if err != nil {
		log.Fatal(err.Error())
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go consumer.StartConsuming(&wg)
	wg.Wait()
}
