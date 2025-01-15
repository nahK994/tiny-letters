package handler

import (
	"net/http"
	grpc_auth "tiny-letter/orchestrator/cmd/grpc/client/auth"
	grpc_subscription "tiny-letter/orchestrator/cmd/grpc/client/subscription"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// HandlePublisherSubscription handles publisher subscription with 2PC
func (h *Handler) HandlePublisherSubscription(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id" binding:"required"`
		PlanID int `json:"plan_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Phase 1: Prepare
	authSubscriptionID, err := grpc_auth.ConfirmPublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth service prepare failed"})
		return
	}

	subscriptionID, err := grpc_subscription.ConfirmPublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		// Rollback Auth
		grpc_auth.RollbackConfirmPublisherSubscription(authSubscriptionID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription service prepare failed"})
		return
	}

	// Phase 2: Commit
	c.JSON(http.StatusOK, gin.H{
		"message":         "Publisher subscription confirmed",
		"auth_id":         authSubscriptionID,
		"subscription_id": subscriptionID,
	})
}

// HandlePublisherUnsubscription handles publisher unsubscription with 2PC
func (h *Handler) HandlePublisherUnsubscription(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Phase 1: Prepare
	authSubscriptionID, err := grpc_auth.RevokePublisherSubscription(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth service prepare failed"})
		return
	}

	subscriptionID, err := grpc_subscription.RevokePublisherSubscription(req.UserID)
	if err != nil {
		// Rollback Auth
		grpc_auth.RollbackRevokePublisherSubscription(authSubscriptionID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription service prepare failed"})
		return
	}

	// Phase 2: Commit
	c.JSON(http.StatusOK, gin.H{
		"message":         "Publisher unsubscription confirmed",
		"auth_id":         authSubscriptionID,
		"subscription_id": subscriptionID,
	})
}

// HandleChangePublisherPlan handles plan changes for publishers with 2PC
func (h *Handler) HandleChangePublisherPlan(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id" binding:"required"`
		PlanID int `json:"plan_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Phase 1: Prepare
	authSubscriptionID, err := grpc_auth.ChangePublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth service prepare failed"})
		return
	}

	subscriptionID, err := grpc_subscription.ChangePublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		// Rollback Auth
		grpc_auth.RollbackChangePublisherSubscription(authSubscriptionID, req.PlanID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription service prepare failed"})
		return
	}

	// Phase 2: Commit
	c.JSON(http.StatusOK, gin.H{
		"message":         "Publisher plan change confirmed",
		"auth_id":         authSubscriptionID,
		"subscription_id": subscriptionID,
	})
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

	// Phase 1: Prepare
	authSubscriptionID, err := grpc_auth.JoinPublication(req.UserID, req.PublicationID, req.IsPremium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth service prepare failed"})
		return
	}

	subscriptionID, err := grpc_subscription.JoinPublication(req.UserID, req.PublicationID, req.IsPremium)
	if err != nil {
		// Rollback Auth
		grpc_auth.RollbackJoinPublication(authSubscriptionID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription service prepare failed"})
		return
	}

	// Phase 2: Commit
	c.JSON(http.StatusOK, gin.H{
		"message":         "Join publication confirmed",
		"auth_id":         authSubscriptionID,
		"subscription_id": subscriptionID,
	})
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

	// Phase 1: Prepare
	authSubscriptionID, err := grpc_auth.LeavePublication(req.UserID, req.PublicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth service prepare failed"})
		return
	}

	subscriptionID, err := grpc_subscription.LeavePublication(req.UserID, req.PublicationID)
	if err != nil {
		// Rollback Auth
		grpc_auth.RollbackLeavePublication(authSubscriptionID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription service prepare failed"})
		return
	}

	// Phase 2: Commit
	c.JSON(http.StatusOK, gin.H{
		"message":         "Leave publication confirmed",
		"auth_id":         authSubscriptionID,
		"subscription_id": subscriptionID,
	})
}

// HandleChangeSubscriberPlan handles plan changes for subscribers with 2PC
func (h *Handler) HandleChangeSubscriberPlan(c *gin.Context) {
	var req struct {
		UserID        int `json:"user_id" binding:"required"`
		PublicationID int `json:"publication_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Phase 1: Prepare
	authSubscriptionID, err := grpc_auth.ChangePublicationPlan(req.UserID, req.PublicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth service prepare failed"})
		return
	}

	subscriptionID, err := grpc_subscription.ChangePublicationPlan(req.UserID, req.PublicationID)
	if err != nil {
		// Rollback Auth
		grpc_auth.RollbackChangePublicationPlan(authSubscriptionID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription service prepare failed"})
		return
	}

	// Phase 2: Commit
	c.JSON(http.StatusOK, gin.H{
		"message":         "Subscriber plan change confirmed",
		"auth_id":         authSubscriptionID,
		"subscription_id": subscriptionID,
	})
}
