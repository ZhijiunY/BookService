package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	controller "github.com/zhijiunY/BookService/controllers"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	// router.LoadHTMLGlob("templates/*/*.tmpl")
	render := multitemplate.NewRenderer()

	render.AddFromFiles("index", "templates/public/base.html", "templates/public/index.html")
	render.AddFromFiles("login", "templates/public/base.html", "templates/auth/login.html")
	render.AddFromFiles("register", "templates/public/base.html", "templates/auth/register.html")
	render.AddFromFiles("succRegister", "templates/public/base.html", "templates/auth/succRegister.html")
	render.AddFromFiles("succLogin", "templates/public/base.html", "templates/auth/succLogin.html")
	render.AddFromFiles("refail", "templates/public/base.html", "templates/auth/refail.html")
	render.AddFromFiles("lofail", "templates/public/base.html", "templates/auth/lofail.html")

	router.HTMLRender = render

	router.GET("/", controller.IndexPage())

	auth := router.Group("auth")
	{
		// auth.GET("refresh", controller.Refresh())
		auth.GET("/getlogin", controller.GetLogin())
		auth.GET("/getregister", controller.GetRegister())
		auth.GET("/chgpsw", controller.GetChgpsw())
		auth.GET("fgtpsw", controller.GetFgtpsw())
		auth.GET("/succRegister", controller.SuccRegister())
		auth.GET("/succLogin", controller.SuccLogin())
		auth.GET("/login", controller.Login())

		auth.POST("/register", controller.Register())
		auth.POST("/login", controller.Login())
		auth.POST("/googleOAuth", controller.GoogleOAuth())
		auth.POST("/gitHubOAuth", controller.GitHubOAuth())

		auth.DELETE("/logout", controller.Logout())

	}

	users := router.Group("users")
	{
		users.GET("/", controller.GetUsers())
		users.POST("/", controller.CreateUser())
		users.PUT("/:id", controller.UpdateUser())
		users.DELETE("/:id", controller.DeleteUser())
	}

	book := router.Group("book")
	{
		book.GET("/", controller.Getbook())
		book.POST("/", controller.Createbook())
		book.PUT("/:id", controller.Updatebook())
		book.DELETE("/:id", controller.Deletebook())
	}

}
