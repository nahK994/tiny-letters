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

func (h *Handler) HandleRegister(c *gin.Context) {
	var userRequest struct {
		name             string
		email            string
		password         string
		role             string
		subscriptionType int
	}
	if err := c.ShouldBindJSON(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, "wrong user info format")
		return
	}

	hashedPassword, err := hashPassword(userRequest.password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	if err := h.repo.CreateUser(
		&db.CreateUserRequest{
			Name:             userRequest.name,
			Email:            userRequest.email,
			Password:         hashedPassword,
			Role:             userRequest.role,
			SubscriptionType: userRequest.subscriptionType,
		},
	); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
}
