package handlers

import (
	"encoding/json"
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
	case constants.SubscriberSubscribe:
		var info models.JoinPublicationData
		json.Unmarshal(data.Data, &info)
		return h.handleJoinPublication(msg)

	case constants.SubscriberUnsubscribe:
		var info models.LeavePublicationData
		json.Unmarshal(data.Data, &info)
		return h.handleLeavePublication(msg)

	case constants.SubscriberChangePlan:
		var info models.ChangeSubscriberSubscriptionData
		json.Unmarshal(data.Data, &info)
		return h.handleChangeSubscriberSubscription(msg)

	case constants.PublisherSubscribe:
		var info models.ConfirmPublisherSubscriptionData
		json.Unmarshal(data.Data, &info)
		return h.handleConfirmPublisherSubscription(msg)

	case constants.PublisherUnsubscribe:
		var info models.RevokePublisherSubscriptionData
		json.Unmarshal(data.Data, &info)
		return h.handleRevokePublisherSubscription(msg)

	case constants.PublisherChangePlan:
		var info models.ChangePublisherSubscriptionData
		json.Unmarshal(data.Data, &info)
		return h.handleChangePublisherSubscription(msg)
	}
	return nil
}

func (h *Handler) HandlePublicationMsg(msg []byte) {

}
