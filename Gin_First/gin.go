package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `uri:"username"`
	Password string `uri:"password"`
}

func TestGetBind(c *gin.Context) {
	var user User
	err := c.ShouldBindUri(&user)
	if err != nil {
		log.Fatal(err)
	}
	c.String(200, "%v", user)
}

func main() {
	e := gin.Default()
	// http://localhost:8080/testGetBind/ywc/123
	e.GET("/testGetBind/:username/:password", TestGetBind)
	e.Run()
}
