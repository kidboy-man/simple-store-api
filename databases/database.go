package database

import (
	"fmt"
	"log"
	"os"
	"simple-store-api/conf"
	"simple-store-api/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {

	err := godotenv.Load() //Load .env file
	if err != nil {
		panic(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s port=%s, user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) // Build connection string
	log.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conf.AppConfig.DbClient = conn.Debug()                  // show query and runtime
	conf.AppConfig.DbClient.AutoMigrate(&models.Metadata{}) // Database migration
	conf.AppConfig.DbClient.AutoMigrate(&models.Category{}) // Database migration
	conf.AppConfig.DbClient.AutoMigrate(&models.User{})     // Database migration
	conf.AppConfig.DbClient.AutoMigrate(&models.Product{})  // Database migration
	conf.AppConfig.DbClient.AutoMigrate(&models.Variant{})  // Database migration
	conf.AppConfig.DbClient.AutoMigrate(&models.Price{})    // Database migration

}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return conf.AppConfig.DbClient
}
