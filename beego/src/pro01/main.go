package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "pro01/routers"
)

func main() {
	beego.Run()
}
