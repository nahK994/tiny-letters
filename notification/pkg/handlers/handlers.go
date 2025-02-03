package handlers

import mq_producer "tiny-letter/notification/pkg/mq/producer"

type Handler struct {
	producer *mq_producer.Producer
}

func NewHandler(producer *mq_producer.Producer) *Handler {
	return &Handler{
		producer: producer,
	}
}

func (h *Handler) HandleConfirmationMsg(msg []byte) {

}

func (h *Handler) HandlePublicationMsg(msg []byte) {

}
