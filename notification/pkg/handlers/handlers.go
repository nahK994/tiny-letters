package handlers

import (
	"encoding/json"
	"tiny-letter/notification/cmd/grpc/client/subscription"
	"tiny-letter/notification/pkg/constants"
	"tiny-letter/notification/pkg/models"
	mq_producer "tiny-letter/notification/pkg/mq/producer"
)

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
	case constants.JoinPublication:
		return h.handleJoinPublication(msg)

	case constants.LeavePublication:
		return h.handleLeavePublication(msg)

	case constants.SubscriberChangePlan:
		return h.handleChangeSubscriberSubscription(msg)

	case constants.PublisherSubscribe:
		return h.handleConfirmPublisherSubscription(msg)

	case constants.PublisherUnsubscribe:
		return h.handleRevokePublisherSubscription(msg)

	case constants.PublisherChangePlan:
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
	return h.producer.Push(constants.PublicationEmail, data)
}

func (h *Handler) handleJoinPublication(msg []byte) error {
	return h.producer.Push(constants.ConfirmationEmail, msg)
}

func (h *Handler) handleLeavePublication(msg []byte) error {
	return h.producer.Push(constants.ConfirmationEmail, msg)
}

func (h *Handler) handleChangeSubscriberSubscription(msg []byte) error {
	return h.producer.Push(constants.ConfirmationEmail, msg)
}

func (h *Handler) handleConfirmPublisherSubscription(msg []byte) error {
	return h.producer.Push(constants.ConfirmationEmail, msg)
}

func (h *Handler) handleRevokePublisherSubscription(msg []byte) error {
	return h.producer.Push(constants.ConfirmationEmail, msg)
}

func (h *Handler) handleChangePublisherSubscription(msg []byte) error {
	return h.producer.Push(constants.ConfirmationEmail, msg)
}
