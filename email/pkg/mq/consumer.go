package mq

import (
	"fmt"
	"log"
	"sync"
	"tiny-letter/email/pkg/app"

	"github.com/IBM/sarama"
)

type ConsumptionHandler interface {
	HandleMsg(msg []byte) error
}

type Consumer struct {
	consumers []sarama.PartitionConsumer
	handler   ConsumptionHandler
}

func NewConsumer(handler ConsumptionHandler, config *app.MQ_config) (*Consumer, error) {
	consumers, err := ConnectConsumers(config)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumers: consumers,
		handler:   handler,
	}, nil
}

func ConnectConsumers(config *app.MQ_config) ([]sarama.PartitionConsumer, error) {
	broker := fmt.Sprintf("%s:%d", config.Domain, config.Port)
	mqConfig := sarama.NewConfig()
	mqConfig.Consumer.Return.Errors = config.Consumer.IsConsumerReturnError

	worker, err := sarama.NewConsumer([]string{broker}, mqConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to start worker: %w", err)
	}

	consumers := []sarama.PartitionConsumer{}

	topics := []string{
		config.Topic.JoinPublication,
		config.Topic.LeavePublication,
		config.Topic.SubscriberChangePlan,
		config.Topic.PublisherSubscribe,
		config.Topic.PublisherUnsubscribe,
		config.Topic.PublisherChangePlan,
		config.Topic.PublishLetter,
	}

	for _, topic := range topics {
		consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
		if err != nil {
			return nil, fmt.Errorf("failed to start consumer for topic %s: %w", topic, err)
		}
		consumers = append(consumers, consumer)
	}

	return consumers, nil
}

func (consumer *Consumer) StartConsuming(wg *sync.WaitGroup) {
	defer wg.Done()

	var internalWg sync.WaitGroup

	for _, c := range consumer.consumers {
		internalWg.Add(1)
		go func(c sarama.PartitionConsumer) {
			defer internalWg.Done()

			for {
				select {
				case err := <-c.Errors():
					fmt.Println("Consumer error:", err)
				case msg := <-c.Messages():
					err := consumer.handler.HandleMsg(msg.Value)
					if err != nil {
						log.Printf("Error in handling message: %s", err.Error())
					}
				}
			}
		}(c)
	}

	internalWg.Wait()
}
