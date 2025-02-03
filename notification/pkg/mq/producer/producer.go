package mq_producer

import (
	"fmt"
	"tiny-letter/notification/pkg/app"

	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer() (*Producer, error) {
	producer, err := connectProducer()
	if err != nil {
		return nil, err
	}
	defer producer.Close()
	return &Producer{producer: producer}, nil
}

func connectProducer() (sarama.SyncProducer, error) {
	appConfig := app.GetConfig()
	broker := fmt.Sprintf("%s:%d", appConfig.MQ.Domain, appConfig.MQ.Port)

	mqConfig := sarama.NewConfig()
	mqConfig.Producer.Return.Successes = appConfig.Producer.IsProducerReturnSuccess
	mqConfig.Producer.RequiredAcks = sarama.WaitForAll
	mqConfig.Producer.Retry.Max = appConfig.Producer.NumberOfRetry

	return sarama.NewSyncProducer([]string{broker}, mqConfig)
}

func (p *Producer) Push(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
