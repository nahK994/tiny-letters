package handlers

import (
	"encoding/json"
	"tiny-letter/notification/pkg/app"
	"tiny-letter/notification/pkg/models"
)

var mqConfig = app.GetConfig().MQ

type producer interface {
	Push(topic string, message []byte) error
}

type grpcClient interface {
	GetContentSubscribers(publicationId int) ([]int32, error)
}

type Handler struct {
	producer producer
	grpc     grpcClient
}

func NewHandler(producer producer, grpcClient grpcClient) *Handler {
	return &Handler{
		producer: producer,
		grpc:     grpcClient,
	}
}

func (h *Handler) HandleMsg(msg []byte) error {
	var data models.MessageItem
	topic := mqConfig.Topic
	json.Unmarshal(msg, &data)

	switch data.Topic {
	case topic.SubscriberRegister:
		return h.handleSubscriberRegistration(msg)

	case topic.JoinPublication:
		return h.handleJoinPublication(msg)

	case topic.LeavePublication:
		return h.handleLeavePublication(msg)

	case topic.SubscriberChangePlan:
		return h.handleChangeSubscriberSubscription(msg)

	case topic.PublisherSubscribe:
		return h.handlePublisherSubscription(msg)

	case topic.PublisherUnsubscribe:
		return h.handleRevokePublisherSubscription(msg)

	case topic.PublisherChangePlan:
		return h.handleChangePublisherSubscription(msg)

	case topic.PublishLetter:
		return h.handlePublishLetter(msg)
	}
	return nil
}

func (h *Handler) handleJoinPublication(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.JoinPublication, msg)
}

func (h *Handler) handleLeavePublication(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.LeavePublication, msg)
}

func (h *Handler) handleChangeSubscriberSubscription(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.SubscriberChangePlan, msg)
}

func (h *Handler) handlePublisherSubscription(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.PublisherChangePlan, msg)
}

func (h *Handler) handleRevokePublisherSubscription(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.PublisherUnsubscribe, msg)
}

func (h *Handler) handleChangePublisherSubscription(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.PublisherChangePlan, msg)
}

func (h *Handler) handlePublisherRegistration(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.PublisherRegister, msg)
}

func (h *Handler) handleSubscriberRegistration(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.SubscriberRegister, msg)
}

func (h *Handler) handlePublishLetter(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.PublishLetter, msg)
}
