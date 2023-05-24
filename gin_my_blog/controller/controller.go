package controller

import (
	"fmt"
	"gin_my_blog/dao"
	"gin_my_blog/model"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"strconv"
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
func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)
	if u.Username == "" {
		c.HTML(200, "login.html", "用户不存在")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登陆成功")
			c.Redirect(301, "/")
		}
	}
}

//博客列表

func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

//添加博客

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(302, "/post_index")
}

//跳转到添加博客

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}
func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(p.Content))

	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}
