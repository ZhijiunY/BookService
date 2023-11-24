package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zhijiunY/BookService/config"
	_ "github.com/zhijiunY/BookService/docs"
	"gorm.io/gorm"

	"github.com/zhijiunY/BookService/routes"
)

// @title Book service
// @version 1.0
// @description book service
// @schemes http,https
// @host localhost:8000
// @BasePath /
// @contact.name nathan.lu
// @contact.url https://tedmax100.github.io/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()
	router.Use(gin.Logger())
	routes.SetupRoutes(router, &gorm.DB{})
	config.Connect()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + port)
}
