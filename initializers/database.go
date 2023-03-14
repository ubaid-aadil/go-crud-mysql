package initializers

import (
	"golang-crud-mysql/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabse() {

	var err error

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		log.Fatal("Failed to connect to a  database")

	}
	DB.AutoMigrate(&models.Post{})
}
