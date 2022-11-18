package conf

import (
	"os"
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var AppConfig Config

type Config struct {
	DbClient *gorm.DB
	JWTConfig
}

type JWTConfig struct {
	JWTSignatureKey      string
	JWTAdminSignatureKey string
	JWTExpirationTime    time.Duration
}

func init() {
	err := godotenv.Load() //Load .env file
	if err != nil {
		panic(err)
	}

	beego.BConfig.RunMode = os.Getenv("beego_runmode")
	logs.Info("beego.BConfig.RunMode", beego.BConfig.RunMode)

	AppConfig.JWTConfig.JWTSignatureKey = os.Getenv("jwt_signature_key")
	if AppConfig.JWTConfig.JWTSignatureKey == "" {
		panic("jwt_signature_key not set")
	}

	AppConfig.JWTConfig.JWTAdminSignatureKey = os.Getenv("jwt_admin_signature_key")
	if AppConfig.JWTConfig.JWTAdminSignatureKey == "" {
		panic("jwt_admin_signature_key not set")
	}

	jwtExpirationTimeStr := os.Getenv("jwt_expiration_time")
	jwtExpirationTime, _ := strconv.Atoi(jwtExpirationTimeStr)
	if jwtExpirationTime == 0 {
		jwtExpirationTime = 24 * 60 * 60
	}
	AppConfig.JWTConfig.JWTExpirationTime = time.Duration(jwtExpirationTime) * time.Second
}
