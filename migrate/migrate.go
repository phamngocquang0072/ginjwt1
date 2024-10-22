package main

import (
	"log"

	"github.com/phamngocquang0072/ginjwt1/initializers"
	"github.com/phamngocquang0072/ginjwt1/internal/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// Load .env file
	initializers.LoadEnv()
	DB = initializers.ConnectDB()
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	log.Fatal("Migrated!")
}

func main() {
}
