package config

import (
	"log"

	"github.com/zhijiunY/BookService/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "postgres://postgres:password@localhost:5432/book_service"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to postgres")
	}

	log.Println("Connected to PostgreSQL database")
	db.AutoMigrate(
		&models.User{}, &models.Book{}, &models.Register{},
		&models.Login{},
	)
	DB = db
}
