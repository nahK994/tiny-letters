package mq

import (
	"encoding/json"
	"fmt"
	"tiny-letter/content/pkg/app"
	"tiny-letter/content/pkg/models"

	"github.com/IBM/sarama"
)

type MQ struct {
	producer sarama.SyncProducer
}

func NewProducer(config *app.MQ_config) (*MQ, error) {
	producer, err := connectProducer(config)
	if err != nil {
		return nil, err
	}
	defer producer.Close()
	return &MQ{producer: producer}, nil
}

func connectProducer(config *app.MQ_config) (sarama.SyncProducer, error) {
	broker := fmt.Sprintf("%s:%d", config.Domain, config.Port)

	mqConfig := sarama.NewConfig()
	mqConfig.Producer.Return.Successes = config.Producer.IsProducerReturnSuccess
	mqConfig.Producer.RequiredAcks = sarama.WaitForAll
	mqConfig.Producer.Retry.Max = config.Producer.NumberOfRetry

	return sarama.NewSyncProducer([]string{broker}, mqConfig)
}

func (p *MQ) PushToQueue(topic string, data json.RawMessage) {
	msg := models.MessageItem{
		Topic: topic,
		Data:  data,
	}
	msgBytes, _ := json.Marshal(msg)

	p.push(topic, msgBytes)
}

func (p *MQ) push(topic string, val []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(val),
	}
	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
