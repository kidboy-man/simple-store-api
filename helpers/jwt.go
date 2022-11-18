package helpers

import (
	"simple-store-api/conf"
	"simple-store-api/middlewares"
	"simple-store-api/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user *models.User) (result string, err error) {
	claims := middlewares.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "simple-store-api",
			ExpiresAt: time.Now().Add(conf.AppConfig.JWTConfig.JWTExpirationTime).Unix(),
		},
		UID:      user.ID,
		Username: user.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err = token.SignedString([]byte(conf.AppConfig.JWTConfig.JWTSignatureKey))
	return
}

func GenerateAdminToken(admin *models.Admin) (result string, err error) {
	claims := middlewares.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "simple-store-api",
			ExpiresAt: time.Now().Add(conf.AppConfig.JWTConfig.JWTExpirationTime).Unix(),
		},
		UID:    admin.ID,
		RoleID: admin.RoleID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err = token.SignedString([]byte(conf.AppConfig.JWTConfig.JWTAdminSignatureKey))
	return
}
