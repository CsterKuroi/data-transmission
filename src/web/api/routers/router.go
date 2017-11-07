package routers

import (
	"github.com/astaxie/beego"
	"uniswitch-agent/src/web/api/controllers"
)

func init() {
	beego.Router("/hello", &controllers.MainController{})
}
