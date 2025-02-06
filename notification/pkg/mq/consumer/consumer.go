package mq_consumer

import (
	"fmt"
	"sync"
	"tiny-letter/notification/pkg/app"
	"tiny-letter/notification/pkg/handlers"

	"github.com/IBM/sarama"
)

type (
	ConfirmationConsumer sarama.PartitionConsumer
	PublicationConsumer  sarama.PartitionConsumer
)

type Consumer struct {
	confirmationConsumer sarama.PartitionConsumer
	publicationConsumer  sarama.PartitionConsumer
	handler              *handlers.Handler
}

func NewConsumer(handler *handlers.Handler, config *app.MQ_config) (*Consumer, error) {
	confirmationConsumer, publicationConsumer, err := ConnectConsumer(config)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		confirmationConsumer: confirmationConsumer,
		publicationConsumer:  publicationConsumer,
		handler:              handler,
	}, nil
}

func ConnectConsumer(config *app.MQ_config) (ConfirmationConsumer, PublicationConsumer, error) {
	broker := fmt.Sprintf("%s:%d", config.Domain, config.Port)
	mqConfig := sarama.NewConfig()
	mqConfig.Consumer.Return.Errors = config.Consumer.IsConsumerReturnError
	worker, err := sarama.NewConsumer([]string{broker}, mqConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start worker")
	}

	confirmationConsumer, err := worker.ConsumePartition(config.Consumer.ConfirmationTopic, 0, sarama.OffsetOldest)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start confirmation consumer: %w", err)
	}
	publicationConsumer, err := worker.ConsumePartition(config.Consumer.PublicationTopic, 0, sarama.OffsetOldest)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start publication consumer: %w", err)
	}
	return confirmationConsumer, publicationConsumer, nil
}

func (consumer *Consumer) StartConsuming(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case err := <-consumer.confirmationConsumer.Errors():
			fmt.Println(err)
		case msg := <-consumer.confirmationConsumer.Messages():
			consumer.handler.HandleConfirmationMsg(msg.Value)
		case err := <-consumer.publicationConsumer.Errors():
			fmt.Println(err)
		case msg := <-consumer.publicationConsumer.Messages():
			consumer.handler.HandlePublicationMsg(msg.Value)
		}
	}
}
