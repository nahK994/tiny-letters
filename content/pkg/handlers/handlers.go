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
}

func (h *Handler) HandleCreatePost(c *gin.Context) {
}
