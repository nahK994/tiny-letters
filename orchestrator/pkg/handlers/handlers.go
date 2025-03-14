package handlers

import (
	"encoding/json"
	"net/http"
	"tiny-letter/orchestrator/pkg/app"
	"tiny-letter/orchestrator/pkg/client/auth"
	"tiny-letter/orchestrator/pkg/client/subscription"
	"tiny-letter/orchestrator/pkg/models"

	"github.com/gin-gonic/gin"
)

var mq = app.GetConfig().MQ

type producer interface {
	Push(val []byte) error
}

type Handler struct {
	producer producer
}

func NewHandler(producer producer) *Handler {
	return &Handler{
		producer: producer,
	}
}

func (h *Handler) pushToQueue(action string, data json.RawMessage) {
	msg := models.ConfirmationMessage{
		Action: action,
		Data:   data,
	}
	msgBytes, _ := json.Marshal(msg)
	h.producer.Push(msgBytes)
}

func (h *Handler) HandleConfirmPublisherSubscription(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id" binding:"required"`
		PlanID int `json:"plan_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionId, err := subscription.ConfirmPublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = auth.ConfirmPublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		subscription.RollbackConfirmPublisherSubscription(int(subscriptionId))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	msgData, _ := json.Marshal(
		models.ConfirmPublisherSubscriptionData{
			UserId: req.UserID,
			PlanId: req.PlanID,
		},
	)
	h.pushToQueue(mq.MsgAction.PublisherSubscribe, msgData)

	c.JSON(http.StatusOK, "Publisher subscription confirmed")
}

func (h *Handler) HandleRevokePublisherSubscription(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	planId, err := subscription.RevokePublisherSubscription(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = auth.RevokePublisherSubscription(req.UserID, int(planId))
	if err != nil {
		subscription.RollbackRevokePublisherSubscription(req.UserID, int(planId))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	msgData, _ := json.Marshal(
		models.RevokePublisherSubscriptionData{
			UserId: req.UserID,
		},
	)
	h.pushToQueue(mq.MsgAction.PublisherUnsubscribe, msgData)

	c.JSON(http.StatusOK, "Publisher unsubscription confirmed")
}

func (h *Handler) HandleChangePublisherSubscription(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id" binding:"required"`
		PlanID int `json:"plan_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionID, oldPlanId, err := subscription.ChangePublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = auth.ChangePublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		subscription.RollbackChangePublisherSubscription(int(subscriptionID), int(oldPlanId))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	msgData, _ := json.Marshal(
		models.ChangePublisherSubscriptionData{
			UserId:    req.UserID,
			NewPlanId: req.PlanID,
			OldPlanId: int(oldPlanId),
		},
	)
	h.pushToQueue(mq.MsgAction.PublisherChangePlan, msgData)

	c.JSON(http.StatusOK, "Publisher plan change confirmed")
}

func (h *Handler) HandleJoinPublication(c *gin.Context) {
	var req struct {
		UserID        int  `json:"user_id" binding:"required"`
		PublicationID int  `json:"publication_id" binding:"required"`
		IsPremium     bool `json:"is_premium" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionID, err := subscription.JoinPublication(req.UserID, req.PublicationID, req.IsPremium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = auth.JoinPublication(req.UserID, req.PublicationID, req.IsPremium)
	if err != nil {
		subscription.RollbackJoinPublication(int(subscriptionID))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	msgData, _ := json.Marshal(
		models.JoinPublicationData{
			UserId:        req.UserID,
			PlanType:      req.IsPremium,
			PublicationId: req.PublicationID,
		},
	)
	h.pushToQueue(mq.MsgAction.JoinPublication, msgData)

	c.JSON(http.StatusOK, "Join publication confirmed")
}

func (h *Handler) HandleLeavePublication(c *gin.Context) {
	var req struct {
		UserID        int `json:"user_id" binding:"required"`
		PublicationID int `json:"publication_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isPremium, err := subscription.LeavePublication(req.UserID, req.PublicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = auth.LeavePublication(req.UserID, req.PublicationID)
	if err != nil {
		subscription.RollbackLeavePublication(req.UserID, req.PublicationID, bool(isPremium))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	msgData, _ := json.Marshal(
		models.LeavePublicationData{
			UserId:        req.UserID,
			PublicationId: req.PublicationID,
		},
	)
	h.pushToQueue(mq.MsgAction.LeavePublication, msgData)

	c.JSON(http.StatusOK, "Leave publication confirmed")
}

func (h *Handler) HandleChangeSubscriberSubscription(c *gin.Context) {
	var req struct {
		UserID        int `json:"user_id" binding:"required"`
		PublicationID int `json:"publication_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionID, err := subscription.ChangeSubscriberSubscription(req.UserID, req.PublicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = auth.ChangeSubscriberSubscription(req.UserID, req.PublicationID)
	if err != nil {
		subscription.RollbackChangeSubscriberSubscription(int(subscriptionID))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	msgData, _ := json.Marshal(
		models.ChangeSubscriberSubscriptionData{
			UserId:        req.UserID,
			PublicationId: req.PublicationID,
		},
	)
	h.pushToQueue(mq.MsgAction.SubscriberChangePlan, msgData)

	c.JSON(http.StatusOK, "Subscriber plan change confirmed")
}
