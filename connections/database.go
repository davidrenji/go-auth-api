package connections

import (
	"log"
	"os"

	"github.com/davidrenji/go-bootcamp-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Database Connection
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION_STRING")), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database!")
	} else {
		DB.AutoMigrate(&models.User{})
	}
}
