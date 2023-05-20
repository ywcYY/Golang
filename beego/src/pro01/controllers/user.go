package controllers

import beego "github.com/beego/beego/v2/server/web"

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	getString := c.GetString("username")
	c.Ctx.Output.Body([]byte(getString))

}
