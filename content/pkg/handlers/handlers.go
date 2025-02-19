package handlers

import (
	"encoding/json"
	"tiny-letter/content/pkg/app"
	"tiny-letter/content/pkg/db"
	"tiny-letter/content/pkg/models"
	mq_producer "tiny-letter/content/pkg/mq"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db       *db.Repository
	producer *mq_producer.Producer
}

func GetHandler(db *db.Repository, producer *mq_producer.Producer) *Handler {
	return &Handler{
		db:       db,
		producer: producer,
	}
}

func (h *Handler) pushToQueue(action string, data json.RawMessage) {
	msg := models.PublishContentMessage{
		Action: action,
		Data:   data,
	}
	msgBytes, _ := json.Marshal(msg)
	h.producer.Push(msgBytes)
}

func (h *Handler) HandleCreatePublication(c *gin.Context) {
	var req models.CreatePublicationRequest
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
	var req models.CreatePostRequest
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

	msgData, _ := json.Marshal(models.ContentData{
		ContentId: req.PublicationID,
		Content:   req.Content,
	})

	mq_action := app.GetConfig().MQ.MsgAction.PublishContent
	h.pushToQueue(mq_action, msgData)

	c.JSON(200, id)
}
