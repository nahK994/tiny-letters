package mq

import (
	"fmt"
	"tiny-letter/notification/pkg/app"

	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer(config *app.MQ_config) (*Producer, error) {
	producer, err := connectProducer(config)
	if err != nil {
		return nil, err
	}
	defer producer.Close()
	return &Producer{producer: producer}, nil
}

func connectProducer(config *app.MQ_config) (sarama.SyncProducer, error) {
	broker := fmt.Sprintf("%s:%d", config.Domain, config.Port)

	mqConfig := sarama.NewConfig()
	mqConfig.Producer.Return.Successes = config.Producer.IsProducerReturnSuccess
	mqConfig.Producer.RequiredAcks = sarama.WaitForAll
	mqConfig.Producer.Retry.Max = config.Producer.NumberOfRetry

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
