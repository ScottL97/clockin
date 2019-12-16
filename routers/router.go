package routers

import (
	"beegoapp/clockin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/clockin", &controllers.ClockinController{})
	beego.Router("/clockin/login", &controllers.LoginController{})
	beego.Router("/clockin/result", &controllers.ResultController{})
}
