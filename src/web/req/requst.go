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

func RegisterAgentToSwitch(url string, jsonData string) (str string) {
	req := httplib.Post(url)
	//req.Header("Content-Encoding", "utf-8")
	req.Header("Content-Type", "application/json; charset=utf-8")
	req.Body(jsonData)
	str, err := req.String()
	if err != nil {
		return ""
	}
	logs.Info(str)
	return str
}
