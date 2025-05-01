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

type REST_Handler struct {
	repo *db.Repository
}

func GetREST_Handlers(repo *db.Repository) *REST_Handler {
	return &REST_Handler{repo: repo}
}

func (h *REST_Handler) Login(c *gin.Context) {
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

func (h *REST_Handler) HandleSubscriberRegistration(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, "wrong user info format")
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	user_id, err := h.repo.CreateUser(
		&models.UserRegistration{
			Email:    req.Email,
			Password: hashedPassword,
			Role:     "subscriber",
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, user_id)
}
