package handlers

import (
	"net/http"
	"tiny-letter/orchestrator/pkg/grpc/client/auth"
	"tiny-letter/orchestrator/pkg/grpc/client/email"
	"tiny-letter/orchestrator/pkg/grpc/client/subscription"
	"tiny-letter/orchestrator/pkg/models"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandlerPublisherRegistration(c *gin.Context) {
	var req struct {
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

	subscriptionId, err := subscription.CreateSubscriptionForPublisher(userId, req.PlanId)
	if err != nil {
		auth.RollbackCreatePublisher(int(userId))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	if err := email.OnboardUser(&models.OnboardUserData{
		UserId: int(userId),
		Email:  req.Email,
		Role:   "publisher",
	}); err != nil {
		auth.RollbackCreatePublisher(int(userId))
		subscription.RollbackCreateSubscriptionForPublisher(int(subscriptionId))
		c.JSON(http.StatusInternalServerError, "Email service prepare failed")
		return
	}

	c.JSON(http.StatusOK, "Publisher registration confirmed")
}

func (h *Handler) HandlerSubscriberRegistration(c *gin.Context) {
	var req struct {
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := auth.CreateSubscriber(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	if err := email.OnboardUser(&models.OnboardUserData{
		UserId: int(userId),
		Email:  req.Email,
		Role:   "subscriber",
	}); err != nil {
		auth.RollbackCreateSubscriber(int(userId))
		c.JSON(http.StatusInternalServerError, "Email service prepare failed")
		return
	}

	c.JSON(http.StatusOK, "Subscriber registration confirmed")
}
