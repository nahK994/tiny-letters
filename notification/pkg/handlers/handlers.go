package handlers

import (
	"encoding/json"
	"tiny-letter/notification/cmd/grpc/client/subscription"
	"tiny-letter/notification/pkg/app"
	"tiny-letter/notification/pkg/models"
	mq_producer "tiny-letter/notification/pkg/mq/producer"
)

var mq = app.GetConfig().MQ

type Handler struct {
	producer *mq_producer.Producer
}

func NewHandler(producer *mq_producer.Producer) *Handler {
	return &Handler{
		producer: producer,
	}
}

func (h *Handler) HandleConfirmationMsg(msg []byte) error {
	var data models.ConfirmationMessage
	json.Unmarshal(msg, &data)

	switch data.Action {
	case mq.MsgAction.JoinPublication:
		return h.handleJoinPublication(msg)

	case mq.MsgAction.LeavePublication:
		return h.handleLeavePublication(msg)

	case mq.MsgAction.SubscriberChangePlan:
		return h.handleChangeSubscriberSubscription(msg)

	case mq.MsgAction.PublisherSubscribe:
		return h.handleConfirmPublisherSubscription(msg)

	case mq.MsgAction.PublisherUnsubscribe:
		return h.handleRevokePublisherSubscription(msg)

	case mq.MsgAction.PublisherChangePlan:
		return h.handleChangePublisherSubscription(msg)
	}
	return nil
}

func (h *Handler) HandlePublicationMsg(msgBytes []byte) error {
	var msg models.ConsumedContentMessage
	json.Unmarshal(msgBytes, &msg)

	var msgData models.ContentData
	json.Unmarshal(msg.Data, &msgData)

	subscriberIds, _ := subscription.GetContentSubscribers(msgData.ContentId)
	data, _ := json.Marshal(models.PublishedContentMessage{
		Content:       msgData.Content,
		SubscriberIds: subscriberIds,
	})
	return h.producer.Push(mq.Topic.PublicationEmail, data)
}

func (h *Handler) handleJoinPublication(msg []byte) error {
	return h.producer.Push(mq.Topic.ConfirmationEmail, msg)
}

func (h *Handler) handleLeavePublication(msg []byte) error {
	return h.producer.Push(mq.Topic.ConfirmationEmail, msg)
}

func (h *Handler) handleChangeSubscriberSubscription(msg []byte) error {
	return h.producer.Push(mq.Topic.ConfirmationEmail, msg)
}

func (h *Handler) handleConfirmPublisherSubscription(msg []byte) error {
	return h.producer.Push(mq.Topic.ConfirmationEmail, msg)
}

func (h *Handler) handleRevokePublisherSubscription(msg []byte) error {
	return h.producer.Push(mq.Topic.ConfirmationEmail, msg)
}

func (h *Handler) handleChangePublisherSubscription(msg []byte) error {
	return h.producer.Push(mq.Topic.ConfirmationEmail, msg)
}
