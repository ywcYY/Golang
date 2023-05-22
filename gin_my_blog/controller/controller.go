package controller

import (
	"gin_my_blog/dao"
	"gin_my_blog/model"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.RegisterUser(&user)
	c.Redirect(301, "/")
}

func GoRegisterUser(c *gin.Context) {
	c.HTML(200, "register.html", nil)

}
func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
