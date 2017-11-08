package routers

import (
	"uniswitch-agent/src/web/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Hello")
	beego.Router("/sendDataToAnotherAgent", &controllers.MainController{}, "post:SendDataToAnotherAgent")
}
