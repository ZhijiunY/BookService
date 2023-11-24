package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhijiunY/BookService/config"
	"github.com/zhijiunY/BookService/models"
)

func Getbook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}

		book := []models.Book{}
		config.DB.Find(&book)
		c.JSON(200, &book)
	}
}

func Createbook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}

		var book models.Book
		c.BindJSON(&book)
		config.DB.Create(&book)
		c.JSON(200, &book)
	}
}

func Deletebook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}

		var book models.Book
		config.DB.Where("id=?", c.Param("id")).Delete(&book)
		c.JSON(200, &book)

	}
}

func Updatebook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}

		var book models.Book
		config.DB.Where("id=?", c.Param("id")).First(&book)
		c.BindJSON(&book)
		config.DB.Save(&book)
		c.JSON(200, &book)

	}
}
