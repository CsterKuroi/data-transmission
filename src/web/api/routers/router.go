package routers

import (
	"github.com/astaxie/beego"

	"uniswitch-agent/src/web/api/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Hello")
	beego.Router("/sendDataToAnotherAgent", &controllers.MainController{}, "post:SendDataToAnotherAgent")
}
