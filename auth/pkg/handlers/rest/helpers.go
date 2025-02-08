package rest_handlers

import (
	"time"
	"tiny-letter/auth/pkg/app"
	"tiny-letter/auth/pkg/models"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(r *models.GenerateTokenRequest) (string, error) {
	appConfig := app.GetConfig().App
	now := time.Now()
	expTime := now.Add(time.Duration(appConfig.JWT_exp_minutes) * time.Minute)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   r.Id,
		"iss":   "TinyLetter",
		"exp":   expTime.Unix(),
		"iat":   now.Unix(),
		"roles": r.Role,
	})

	tokenString, err := claims.SignedString([]byte(appConfig.JWT_secret))
	return tokenString, err
}
