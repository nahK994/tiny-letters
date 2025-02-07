package handlers

import "tiny-letter/notification/pkg/constants"

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
