package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"tiny-letter/auth/pkg/db"
	"tiny-letter/auth/pkg/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *db.Repository
}

func GetREST_Handlers(repo *db.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Login(c *gin.Context) {
	var payload models.LoginRequest
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

	accessToken, err1 := generateJWT(&models.GenerateTokenRequest{
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

func (h *Handler) HandleUserRegistration(c *gin.Context) {
	var userRequest models.CreateUserRequest
	if err := c.ShouldBindJSON(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, "wrong user info format")
		return
	}

	hashedPassword, err := hashPassword(userRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	user_id, err := h.repo.CreateUser(
		&models.CreateUserRequest{
			Name:     userRequest.Name,
			Email:    userRequest.Email,
			Password: hashedPassword,
			Role:     userRequest.Role,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, user_id)
}
