package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*	mysqluser, _ := beego.AppConfig.String("mysqluser")
		print(mysqluser)*/

	c.Data["Website"] = "ywc.vip"
	c.Data["Name"] = "ywc的golang"
	c.TplName = "index.tpl"

	c.Ctx.Output.Body([]byte("hello world"))
	params := c.Ctx.Input.Param(":id")
	fmt.Println(params)
	getString := c.GetString("username")
	c.Ctx.Output.Body([]byte(getString))

}
