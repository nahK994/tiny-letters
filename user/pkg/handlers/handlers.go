package handlers

import (
	"net/http"
	"tiny-letter-user/pkg/db"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *db.Repository
}

func NewHandler(repo *db.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Login(c *gin.Context) {}

func (h *Handler) HandleRegisterSubscriber(c *gin.Context) {
	var userRequest struct {
		name     string
		email    string
		password string
		role     string
	}
	if err := c.ShouldBindJSON(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, "wrong user info format")
		return
	}

	if err := h.repo.CreateUser(
		&db.CreateUserRequest{
			Name:     userRequest.name,
			Email:    userRequest.email,
			Password: userRequest.password,
			Roles:    []string{userRequest.role},
		},
	); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
}

func (h *Handler) HandleRegisterPublisher(c *gin.Context) {}
