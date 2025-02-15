package mq_handlers

import (
	"encoding/json"
	"tiny-letter/email/pkg/constants"
	"tiny-letter/email/pkg/models"
)

type MQ_ConsumerHandlers struct {
}

func New_ConsumerHandlers() *MQ_ConsumerHandlers {
	return &MQ_ConsumerHandlers{}
}

func (h *MQ_ConsumerHandlers) HandleConfirmationMsg(msg []byte) error {
	var data models.ConfirmationMessage
	json.Unmarshal(msg, &data)

	switch data.Action {
	case constants.SubscriberSubscribe:
		return h.handleJoinPublication(msg)

	case constants.SubscriberUnsubscribe:
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

func (h *MQ_ConsumerHandlers) HandlePublicationMsg(msg []byte) error {
	return nil
}

func (h *MQ_ConsumerHandlers) handleJoinPublication(msg []byte) error {
	return nil
}

func (h *MQ_ConsumerHandlers) handleLeavePublication(msg []byte) error {
	return nil
}

func (h *MQ_ConsumerHandlers) handleChangeSubscriberSubscription(msg []byte) error {
	return nil
}

func (h *MQ_ConsumerHandlers) handleConfirmPublisherSubscription(msg []byte) error {
	return nil
}

func (h *MQ_ConsumerHandlers) handleRevokePublisherSubscription(msg []byte) error {
	return nil
}

func (h *MQ_ConsumerHandlers) handleChangePublisherSubscription(msg []byte) error {
	return nil
}
