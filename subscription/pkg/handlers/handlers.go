package handlers

import (
	"tiny-letter-subscription/pkg/db"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *db.Repository
}

func NewHandler(repo *db.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) HandlerSubscribePublisher(c *gin.Context) {}

func (h *Handler) HandlerSubscribePublication(c *gin.Context) {}

func (h *Handler) HandlerChangeSubscriberSubscriptionPlan(c *gin.Context) {}

func (h *Handler) HandlerChangePublisherSubscriptionPlan(c *gin.Context) {}
