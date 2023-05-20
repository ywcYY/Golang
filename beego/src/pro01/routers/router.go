package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"pro01/controllers"
)

func init() {
	beego.Router("/user", &controllers.UserController{}) //固定路由
	/*	beego.Router("/api/?:id", &controllers.MainController{})
	 */
	/*	beego.Router("/api/food", &controllers.MainController{}, "get:ListFood")
		beego.Router("/api/food", &controllers.MainController{}, "post:CreateFood")
	*/
}
