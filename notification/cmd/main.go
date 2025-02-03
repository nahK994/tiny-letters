package main

import (
	"log"
	"sync"
	"tiny-letter/notification/pkg/handlers"
	mq_consumer "tiny-letter/notification/pkg/mq/consumer"
	mq_producer "tiny-letter/notification/pkg/mq/producer"
)

func main() {
	producer, err := mq_producer.NewProducer()
	if err != nil {
		log.Fatal(err.Error())
	}
	handlers := handlers.NewHandler(producer)
	consumer, err := mq_consumer.NewConsumer(handlers)
	if err != nil {
		log.Fatal(err.Error())
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go consumer.StartConsuming(&wg)
	wg.Wait()
}
