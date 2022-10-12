package conf

import (
	"log"
	"os"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var AppConfig Config

type Config struct {
	DbClient *gorm.DB
}

func init() {
	err := godotenv.Load() //Load .env file
	if err != nil {
		panic(err)
	}

	beego.BConfig.RunMode = os.Getenv("beego_runmode")
	log.Println("beego.BConfig.RunMode", beego.BConfig.RunMode)
}
