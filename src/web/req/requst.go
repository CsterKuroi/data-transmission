package req

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

func ReqGet() {
	req := httplib.Get("http://beego.me/")
	str, err := req.String()
	if err != nil {
		logs.Error(err)
	}
	fmt.Println(str)
}
