package routers

import (
	"uniswitch-agent/src/web/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Hello")

	beego.Router("/public", &controllers.MainController{}, "post:Public")
	beego.Router("/private", &controllers.MainController{}, "post:Private")
	beego.Router("/secret", &controllers.MainController{}, "post:Secret")
	beego.Router("/address", &controllers.MainController{}, "post:Address")
	beego.Router("/data", &controllers.MainController{}, "post:Data")

	beego.Router("/sendDataToAnotherAgent", &controllers.MainController{}, "post:SendDataToAnotherAgent")
}
