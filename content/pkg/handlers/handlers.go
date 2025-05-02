package handlers

import (
	"encoding/json"
	"tiny-letter/content/pkg/app"
	"tiny-letter/content/pkg/db"
	"tiny-letter/content/pkg/models"
	"tiny-letter/content/pkg/mq"

	"github.com/gin-gonic/gin"
)

type grpcClient interface {
	GetContentSubscribers(publicationId int, isContentPremium bool) ([]int32, error)
}

type Handler struct {
	db   *db.Repository
	mq   *mq.MQ
	grpc grpcClient
}

func GetHandler(db *db.Repository, mq *mq.MQ, grpcClient grpcClient) *Handler {
	return &Handler{
		db:   db,
		mq:   mq,
		grpc: grpcClient,
	}
}

func (h *Handler) HandleCreatePublication(c *gin.Context) {
	var req models.Publication
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	id, err := h.db.CreatePublication(req.Name, req.PublisherID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, id)
}

func (h *Handler) HandleCreatePost(c *gin.Context) {
	var req models.Post
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	id, err := h.db.CreatePost(req.Title, req.Content, req.PublicationID, req.IsPremium)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, id)
}

func (h *Handler) HandlePublishContent(c *gin.Context) {
	var req struct {
		ContentId int `json:"content_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	post, err := h.db.GetContentInfo(req.ContentId)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if post == nil {
		c.JSON(404, "Post not found")
		return
	}

	subscriberIds, err := h.grpc.GetContentSubscribers(post.PublicationID, post.IsPremium)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	if err := h.db.MarkPostAsPublished(req.ContentId); err != nil {
		c.JSON(500, err.Error())
		return
	}

	msgData, _ := json.Marshal(models.PublishContentData{
		ContentId:     req.ContentId,
		Content:       post.Content,
		Title:         post.Title,
		SubscriberIds: subscriberIds,
	})

	h.mq.PushToQueue(app.GetConfig().MQ.Topic.PublishLetter, msgData)

	c.JSON(200, "Content published successfully")
}
