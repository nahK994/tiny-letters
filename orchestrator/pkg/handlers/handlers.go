package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandlePublisherSubscription(c *gin.Context) {}

func (h *Handler) HandlePublisherUnsubscription(c *gin.Context) {}

func (h *Handler) HandleChangePublisherPlan(c *gin.Context) {}

func (h *Handler) HandleJoinPublication(c *gin.Context) {}

func (h *Handler) HandleLeavePublication(c *gin.Context) {}

func (h *Handler) HandleChangeSubscriberPlan(c *gin.Context) {}
