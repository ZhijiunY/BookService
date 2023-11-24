package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhijiunY/BookService/config"
	"github.com/zhijiunY/BookService/models"
)

// @Summary Get User
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Router /users [get]
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}

		users := []models.User{}
		config.DB.Find(&users)
		c.JSON(200, &users)
	}
}

// @Summary Create User
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Router /users [post]
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}

		var user models.User
		c.BindJSON(&user)
		config.DB.Create(&user)
		c.JSON(200, &user)
	}
}

// @Summary Delete User
// @Description Delete a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param auth header string true "token"
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /{id} [delete]
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}

		var user models.User
		config.DB.Where("id=?", c.Param("id")).Delete(&user)
		c.JSON(200, &user)
	}
}

// @Summary Update User
// @Description Update a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param auth header string true "token"
// @Param id path int true "User ID"
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Router /{id} [put]
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}

		var user models.User
		config.DB.Where("id= ?", c.Param("id")).First(&user)
		c.BindJSON(&user)
		config.DB.Save(&user)
		c.JSON(200, &user)
	}
}
