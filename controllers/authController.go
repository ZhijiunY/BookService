package controllers

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/zhijiunY/BookService/config"
	"github.com/zhijiunY/BookService/helper"
	"github.com/zhijiunY/BookService/models"
)

// // var validate = validator.New()
var (
	jwtSecret = []byte("jwtSecret")
)

//GET //////////////////////////////////////////////////////////////////

func IndexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", nil)
	}
}

func GetLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "Please Login",
		})
	}
}

func GetRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "register", gin.H{
			"title": "Register First",
		})
	}
}

func SuccRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "succRegister", gin.H{
			"title": "Success Register",
		})
	}
}

func SuccLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "succLogin", gin.H{
			"title": "Success Login",
		})
	}
}

func GetChgpsw() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "chgpsw", gin.H{
			"title": "Change Password",
		})
	}
}

func GetFgtpsw() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "fgtpsw", gin.H{
			"title": "Forget Password",
		})
	}
}

//POST ///////////////////////////////////

// post register
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		ctx.Done()

		var payload models.Register
		payload.Name = c.PostForm("name")
		payload.Email = c.PostForm("email")
		payload.Password = c.PostForm("password")
		payload.PasswordConfirm = c.PostForm("passwordConfirm")

		if err := c.ShouldBind(&payload); err != nil {
			log.Println("Error binding JSON:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		emailValid := helper.EmailValid(payload.Email)
		if !emailValid {
			log.Println("Error validating email")
			// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format!"})
			c.HTML(http.StatusSeeOther, "refail", gin.H{
				"title": "Invalid email format!",
			})
			return
		}

		passwordValid := helper.PswFmtValid(payload.Password)
		if !passwordValid {
			log.Println("passwordValid error")
			// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password !"})
			c.HTML(http.StatusSeeOther, "refail", gin.H{
				"title": "Invalid password format!",
			})
			return
		}

		if payload.Password != payload.PasswordConfirm {
			log.Println("PasswordConfirm is incorrect")
			// c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
			c.HTML(http.StatusSeeOther, "refail", gin.H{
				"title": "Passwords do not match!",
			})
			return
		}

		hashedPassword, err := helper.HashPassword(payload.Password)
		if err != nil {
			// c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
			log.Println("Error creating hashed password:", err)
			c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
			return
		}

		now := time.Now()
		user := models.User{
			ID:              uuid.New(),
			Name:            payload.Name,
			Email:           strings.ToLower(payload.Email),
			Password:        hashedPassword,
			PasswordConfirm: hashedPassword,
			CreatedAt:       now,
			UpdatedAt:       now,
		}

		// config.DB.Create(&user)
		result := config.DB.Create(&user)
		if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
			log.Println("Error creating user:", result.Error)
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
			return
		} else if result.Error != nil {
			log.Println("Error creating user:", result.Error)
			c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
			return
		}

		c.Redirect(http.StatusSeeOther, "/auth/succRegister")
	}
}

// post login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		ctx.Done()

		var payload models.Login

		payload.Email = c.PostForm("email")
		payload.Password = c.PostForm("password")
		// 解析 JSON 資料為 Login 結構
		if err := c.ShouldBind(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 在數據庫中查詢用户信息
		var u models.User
		// result := config.DB.First(&u, "email = ?", strings.ToLower(l.Email))
		result := config.DB.Find(&u, "email = ?", strings.ToLower(payload.Email))
		if result.Error != nil {
			log.Println("error getting email:", result.Error)
			// c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
			c.HTML(http.StatusSeeOther, "lofail", gin.H{
				"title": "Invalid email!",
			})
			return
		}

		// 验证密码是否正确
		if err := helper.CheckPassword(u.Password, payload.Password); err != nil {
			log.Println("error comparing password", err.Error())
			// c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Password"})
			c.HTML(http.StatusSeeOther, "lofail", gin.H{
				"title": "Invalid password!",
			})
			return
		}

		// helper.SessionSet(c, u.ID)

		// 生成 JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256,
			jwt.MapClaims{
				"password": u.Password,
				"name":     u.Name,
				"email":    u.Email,
			},
		)

		// 簽署 JWT 並獲取字符串格式
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			log.Println("error token", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
			return
		}

		u.Token = tokenString
		// newToken := models.User{

		// 	Token: tokenString,
		// }
		// config.DB.Create(&newToken)

		// c.JSON(http.StatusOK, gin.H{"login": "success"})

		// // 登入成功，返回 JWT
		// c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})

		// c.JSON(http.StatusOK, gin.H{"message": "login success"})
		c.Redirect(http.StatusSeeOther, "/auth/succLogin")

	}
}

// logout
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		ctx.Done()

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "accessToken",
			Value:    "",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
		})
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "refreshToken",
			Value:    "",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
		})

		helper.SessionClear(c)

		c.JSON(http.StatusOK, gin.H{"logout": "success"})

	}
}
