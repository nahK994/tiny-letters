package handlers

import (
	"encoding/json"
	"tiny-letter/email/pkg/app"
	"tiny-letter/email/pkg/models"
)

type ConsumerHandlers struct {
}

func NewConsumerHandlers() *ConsumerHandlers {
	return &ConsumerHandlers{}
}

func (h *ConsumerHandlers) HandleConfirmationMsg(msg []byte) error {
	var data models.ConfirmationMessage
	json.Unmarshal(msg, &data)
	msgAction := app.GetConfig().MQ.MsgAction

	switch data.Action {
	case msgAction.JoinPublication:
		return h.handleJoinPublication(msg)

	case msgAction.LeavePublication:
		return h.handleLeavePublication(msg)

	case msgAction.SubscriberChangePlan:
		return h.handleChangeSubscriberSubscription(msg)

	case msgAction.PublisherSubscribe:
		return h.handleConfirmPublisherSubscription(msg)

	case msgAction.PublisherUnsubscribe:
		return h.handleRevokePublisherSubscription(msg)

	case msgAction.PublisherChangePlan:
		return h.handleChangePublisherSubscription(msg)

	}
	return nil
}

func (h *ConsumerHandlers) HandlePublicationMsg(msg []byte) error {
	return nil
}

func (h *ConsumerHandlers) handleJoinPublication(msg []byte) error {
	return nil
}

func (h *ConsumerHandlers) handleLeavePublication(msg []byte) error {
	return nil
}

func (h *ConsumerHandlers) handleChangeSubscriberSubscription(msg []byte) error {
	return nil
}

func (h *ConsumerHandlers) handleConfirmPublisherSubscription(msg []byte) error {
	return nil
}

func (h *ConsumerHandlers) handleRevokePublisherSubscription(msg []byte) error {
	return nil
}

func (h *ConsumerHandlers) handleChangePublisherSubscription(msg []byte) error {
	return nil
}
