package handlers

import (
	"database/sql"
	"errors"
	"log"
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

func (h *Handler) Login(c *gin.Context) {
	var payload db.LoginRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	userInfo, err := h.repo.GetUserInfoByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, "email not found")
		} else {
			c.JSON(http.StatusInternalServerError, "something unexpected happened")
		}
		return
	}

	if !checkPasswordHash(payload.Password, userInfo.Password) {
		c.JSON(http.StatusUnauthorized, "wrong email or password")
		return
	}

	accessToken, err1 := generateJWT(&db.GenerateTokenRequest{
		Id:   userInfo.Id,
		Role: userInfo.Role,
	})
	if err1 != nil {
		log.Fatal(err1.Error())
		c.JSON(http.StatusInternalServerError, err1.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"access_token": accessToken,
	})
}

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
