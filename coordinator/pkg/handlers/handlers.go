package handler

import (
	"encoding/json"
	"net/http"
	"tiny-letter/coordinator/cmd/grpc/client/auth"
	"tiny-letter/coordinator/cmd/grpc/client/subscription"
	"tiny-letter/coordinator/pkg/constants"
	"tiny-letter/coordinator/pkg/models"
	mq_producer "tiny-letter/coordinator/pkg/mq"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	producer *mq_producer.Producer
}

func NewHandler(producer *mq_producer.Producer) *Handler {
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

// HandlePublisherSubscription handles publisher subscription with 2PC
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
	h.pushToQueue(constants.PublisherSubscribe, msgData)

	c.JSON(http.StatusOK, "Publisher subscription confirmed")
}

// HandlePublisherUnsubscription handles publisher unsubscription with 2PC
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
	h.pushToQueue(constants.PublisherUnsubscribe, msgData)

	c.JSON(http.StatusOK, "Publisher unsubscription confirmed")
}

// HandleChangePublisherPlan handles plan changes for publishers with 2PC
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
	h.pushToQueue(constants.PublisherChangePlan, msgData)

	c.JSON(http.StatusOK, "Publisher plan change confirmed")
}

// HandleJoinPublication handles joining a publication with 2PC
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
	h.pushToQueue(constants.SubscriberSubscribe, msgData)

	c.JSON(http.StatusOK, "Join publication confirmed")
}

// HandleLeavePublication handles leaving a publication with 2PC
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
	h.pushToQueue(constants.SubscriberUnsubscribe, msgData)

	c.JSON(http.StatusOK, "Leave publication confirmed")
}

// HandleChangeSubscriberPlan handles plan changes for subscribers with 2PC
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
	h.pushToQueue(constants.SubscriberChangePlan, msgData)

	c.JSON(http.StatusOK, "Subscriber plan change confirmed")
}
