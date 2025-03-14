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

func (h *Handler) HandleConfirmationMsg(msg []byte) error {
	var data models.ConfirmationMessage
	json.Unmarshal(msg, &data)

	switch data.Action {
	case mqConfig.MsgAction.JoinPublication:
		return h.handleJoinPublication(msg)

	case mqConfig.MsgAction.LeavePublication:
		return h.handleLeavePublication(msg)

	case mqConfig.MsgAction.SubscriberChangePlan:
		return h.handleChangeSubscriberSubscription(msg)

	case mqConfig.MsgAction.PublisherSubscribe:
		return h.handleConfirmPublisherSubscription(msg)

	case mqConfig.MsgAction.PublisherUnsubscribe:
		return h.handleRevokePublisherSubscription(msg)

	case mqConfig.MsgAction.PublisherChangePlan:
		return h.handleChangePublisherSubscription(msg)
	}
	return nil
}

func (h *Handler) HandlePublicationMsg(msgBytes []byte) error {
	var msg models.ConsumedContentMessage
	json.Unmarshal(msgBytes, &msg)

	var msgData models.ContentData
	json.Unmarshal(msg.Data, &msgData)

	subscriberIds, _ := h.grpc.GetContentSubscribers(msgData.ContentId)
	data, _ := json.Marshal(models.PublishedContentMessage{
		Content:       msgData.Content,
		SubscriberIds: subscriberIds,
	})
	return h.producer.Push(mqConfig.Topic.PublicationEmail, data)
}

func (h *Handler) handleJoinPublication(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.ConfirmationEmail, msg)
}

func (h *Handler) handleLeavePublication(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.ConfirmationEmail, msg)
}

func (h *Handler) handleChangeSubscriberSubscription(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.ConfirmationEmail, msg)
}

func (h *Handler) handleConfirmPublisherSubscription(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.ConfirmationEmail, msg)
}

func (h *Handler) handleRevokePublisherSubscription(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.ConfirmationEmail, msg)
}

func (h *Handler) handleChangePublisherSubscription(msg []byte) error {
	return h.producer.Push(mqConfig.Topic.ConfirmationEmail, msg)
}
