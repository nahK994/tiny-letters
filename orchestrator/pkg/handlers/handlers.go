package handlers

import (
	"encoding/json"
	"net/http"
	"tiny-letter/orchestrator/pkg/app"
	"tiny-letter/orchestrator/pkg/grpc/client/auth"
	"tiny-letter/orchestrator/pkg/grpc/client/subscription"
	"tiny-letter/orchestrator/pkg/models"

	"github.com/gin-gonic/gin"
)

var mq = app.GetConfig().MQ

type producer interface {
	Push(topic string, val []byte) error
}

type Handler struct {
	producer producer
}

func NewHandler(producer producer) *Handler {
	return &Handler{
		producer: producer,
	}
}

func (h *Handler) pushToQueue(topic string, data json.RawMessage) {
	msg := models.MessageItem{
		Topic: topic,
		Data:  data,
	}
	msgBytes, _ := json.Marshal(msg)
	h.producer.Push(topic, msgBytes)
}

func (h *Handler) HandlerPublisherRegistration(c *gin.Context) {
	var req struct {
		UserId   int    `json:"user_id" binding:"required"`
		Email    string `json:"email" binding:"required"`
		PlanId   int    `json:"plan_id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := auth.CreatePublisher(req.Email, req.Password, req.PlanId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = subscription.CreateSubscriptionForPublisher(req.UserId, req.PlanId)
	if err != nil {
		auth.RollbackCreatePublisher(int(userId))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	msgData, _ := json.Marshal(
		req,
	)
	h.pushToQueue(mq.Topic.PublisherRegister, msgData)

	c.JSON(http.StatusOK, "Publisher registration confirmed")
}
