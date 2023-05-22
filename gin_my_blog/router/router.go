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
	e.GET("/", controller.Index)
	e.Run()
}
