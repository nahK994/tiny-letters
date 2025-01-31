package handlers

import (
	"tiny-letter/content/pkg/db"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *db.Repository
}

func GetHandler(db *db.Repository) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) HandleCreatePublication(c *gin.Context) {
	var req db.CreatePublicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	id, err := h.DB.CreatePublication(req.Name, req.PublisherID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, id)
}

func (h *Handler) HandleCreatePost(c *gin.Context) {
	var req db.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	id, err := h.DB.CreatePost(req.Title, req.Content, req.PublicationID, req.IsPremium)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, id)
}
