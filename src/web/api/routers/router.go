package routers

import (
	"uniswitch-agent/src/web/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Hello")

	beego.Router("/secret", &controllers.MainController{}, "post:Secret")
	beego.Router("/public", &controllers.MainController{}, "post:Public")
	beego.Router("/private", &controllers.MainController{}, "post:Private")
	beego.Router("/address", &controllers.MainController{}, "post:Address")
	beego.Router("/data", &controllers.MainController{}, "post:Data")
	beego.Router("/decrypt", &controllers.MainController{}, "post:DecryptData")
	beego.Router("/destroy", &controllers.MainController{}, "post:DestroyData")

	beego.Router("/sendDataToAnotherAgent", &controllers.MainController{}, "post:SendDataToAnotherAgent")
}
