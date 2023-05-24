package router

import (
	"gin_my_blog/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.POST("/register", controller.RegisterUser)
	e.GET("/register", controller.GoRegisterUser)

	e.POST("/login", controller.Login)
	e.GET("/login", controller.GoLogin)
	e.GET("/", controller.Index)

	//操作博客
	e.GET("/post_index", controller.GetPostIndex)
	e.POST("/post", controller.AddPost)
	e.GET("/post", controller.GoAddPost)
	e.GET("/post_detail", controller.PostDetail)
	e.Run()
}
