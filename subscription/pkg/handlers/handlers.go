package handlers

import (
	"net/http"
	"tiny-letter/subscription/pkg/db"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *db.Repository
}

func NewHandler(repo *db.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) HandlerSubscribePublisher(c *gin.Context) {
	var req db.PublisherSubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.SubscribePublisherPlan(req); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, "Publisher subscribed successfully")
}

func (h *Handler) HandlerJoinPublication(c *gin.Context) {
	var req db.ManagePublicationSubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.JoinPublication(req); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, "Subscriber subscribed successfully")
}

func (h *Handler) HandlerChangePublicationSubscriptionPlan(c *gin.Context) {
	var req db.ChangePublisherPlanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.ChangePublisherSubscriptionPlan(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Subscriber subscription plan changed successfully")
}

func (h *Handler) HandlerChangePublisherSubscriptionPlan(c *gin.Context) {
	var req db.ChangePublisherPlanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.ChangePublisherSubscriptionPlan(req); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, "Publisher subscription plan changed successfully")
}

func (h *Handler) HandleLeavePublication(c *gin.Context) {
	var req db.ManagePublicationSubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.LeavePublication(req); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, "Subscriber unsubscribed successfully")
}

func (h *Handler) HandleUnsubscribePublisher(c *gin.Context) {
	var req db.UnsubscribePublisherRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.UnsubscriptionPublisherPlan(req); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, "Publisher unsubscribed successfully")
}
