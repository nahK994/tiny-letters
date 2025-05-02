package handlers

import (
	"encoding/json"
	"tiny-letter/email/pkg/app"
	"tiny-letter/email/pkg/models"
)

var mqConfig = app.GetConfig().MQ

type ConsumerHandlers struct {
}

func NewConsumerHandlers() *ConsumerHandlers {
	return &ConsumerHandlers{}
}

func (h *ConsumerHandlers) HandleMsg(msg []byte) error {
	var data models.MessageItem
	topic := mqConfig.Topic
	json.Unmarshal(msg, &data)

	switch data.Topic {
	case topic.JoinPublication:
		return h.handleSubscriberSubscription(msg)

	case topic.LeavePublication:
		return h.handleSubscriberUnsubscription(msg)

	case topic.SubscriberChangePlan:
		return h.handleChangeSubscriberSubscription(msg)

	case topic.PublisherSubscribe:
		return h.handlePublisherSubscription(msg)

	case topic.PublisherUnsubscribe:
		return h.handlePublisherUnsubscribe(msg)

	case topic.PublisherChangePlan:
		return h.handleChangePublisherSubscription(msg)

	case topic.PublishLetter:
		return h.handlePublishLetter(msg)
	}
	return nil
}

func (h *ConsumerHandlers) handleSubscriberSubscription(msg []byte) error {
	var data models.JoinPublication
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return err
	}
	// Handle join publication logic here
	return nil
}

func (h *ConsumerHandlers) handleSubscriberUnsubscription(msg []byte) error {
	var data models.LeavePublication
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return err
	}
	// Handle leave publication logic here
	return nil
}

func (h *ConsumerHandlers) handleChangeSubscriberSubscription(msg []byte) error {
	var data models.ChangeSubscriberSubscription
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return err
	}
	// Handle change subscriber subscription logic here
	return nil
}

func (h *ConsumerHandlers) handlePublisherSubscription(msg []byte) error {
	var data models.PublisherSubscription
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return err
	}
	// Handle publisher subscription logic here
	return nil
}

func (h *ConsumerHandlers) handlePublisherUnsubscribe(msg []byte) error {
	var data models.PublisherUnsubscription
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return err
	}
	// Handle publisher unsubscribe logic here
	return nil
}

func (h *ConsumerHandlers) handleChangePublisherSubscription(msg []byte) error {
	var data models.ChangePublisherSubscription
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return err
	}
	// Handle change publisher subscription logic here
	return nil
}

func (h *ConsumerHandlers) handlePublishLetter(msg []byte) error {
	var data models.PublishLetter
	err := json.Unmarshal(msg, &data)
	if err != nil {
		return err
	}
	// Handle publish letter logic here
	return nil
}
