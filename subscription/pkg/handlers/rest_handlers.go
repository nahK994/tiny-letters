package handlers

import (
	"net/http"
	"tiny-letter/subscription/pkg/db"
	"tiny-letter/subscription/pkg/models"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

type REST_Handler struct {
	db *db.Repository
}

func GetREST_Handler(db *db.Repository) *REST_Handler {
	return &REST_Handler{
		db: db,
	}
}

func (l *REST_Handler) JoinPublication(c *gin.Context) {
	var data models.JoinPublicationRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(int(codes.InvalidArgument), gin.H{"error": "invalid argument"})
		return
	}

	if err := data.Validate(); err != nil {
		c.JSON(int(codes.InvalidArgument), gin.H{"error": "invalid argument"})
		return
	}

	subscriptionId, err := l.db.JoinPublication(&models.JoinPublicationRequest{
		UserId:        data.UserId,
		PublicationId: data.PublicationId,
		IsPremium:     data.IsPremium,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to join publication"})
		return
	}

	c.JSON(http.StatusOK, subscriptionId)
}

func (l *REST_Handler) LeavePublication(c *gin.Context) {
	var data models.LeavePublicationRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(int(codes.InvalidArgument), gin.H{"error": "invalid argument"})
		return
	}

	if err := data.Validate(); err != nil {
		c.JSON(int(codes.InvalidArgument), gin.H{"error": "invalid argument"})
		return
	}

	err := l.db.LeavePublication(&models.LeavePublicationRequest{
		UserId:        data.UserId,
		PublicationId: data.PublicationId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to leave publication"})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (l *REST_Handler) ChangeSubscriberSubscription(c *gin.Context) {
	var data models.ChangeSubscriberSubscriptionRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(int(codes.InvalidArgument), gin.H{"error": "invalid argument"})
		return
	}

	if err := data.Validate(); err != nil {
		c.JSON(int(codes.InvalidArgument), gin.H{"error": "invalid argument"})
		return
	}

	err := l.db.ChangeSubscriberSubscription(&models.ChangeSubscriberSubscriptionRequest{
		UserId:        data.UserId,
		PublicationId: data.PublicationId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to change plan"})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (l *REST_Handler) ChangePublisherSubscription(c *gin.Context) {
	var data models.ChangePublisherSubscriptionRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(int(codes.InvalidArgument), gin.H{"error": "invalid argument"})
		return
	}

	if err := data.Validate(); err != nil {
		c.JSON(int(codes.InvalidArgument), gin.H{"error": "invalid argument"})
		return
	}

	err := l.db.ChangePublisherSubscription(&models.ChangePublisherSubscriptionRequest{
		UserId:        data.UserId,
		ChangedPlanId: data.ChangedPlanId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to change plan"})
		return
	}

	c.JSON(http.StatusOK, "ok")
}
