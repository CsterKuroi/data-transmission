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

func AddAgent(url string, jsonData string) (str string, err error) {
	req := httplib.Post("http://"+url + "/uniswitchAgent/addAgent")
	//req.Header("Content-Encoding", "utf-8")
	req.Header("Content-Type", "application/json; charset=utf-8")
	req.Body(jsonData)
	str, err = req.String()
	if err != nil {
		logs.Info(str)
	}
	return str, err
}

func RegisterAgent(url string, jsonData string) (str string, err error) {
	req := httplib.Post("http://"+url + "/uniswitchAgent/registerAgent")
	//req.Header("Content-Encoding", "utf-8")
	req.Header("Content-Type", "application/json; charset=utf-8")
	req.Body(jsonData)
	str, err = req.String()
	if err != nil {
		logs.Info(str)
	}
	return str, err
}

func AgentLogin(url string, jsonData string) (str string, err error) {
	req := httplib.Post("http://"+url + "/uniswitchAgent/agentLogin")
	//req.Header("Content-Encoding", "utf-8")
	req.Header("Content-Type", "application/json; charset=utf-8")
	req.Body(jsonData)
	str, err = req.String()
	if err != nil {
		logs.Info(str)
	}
	return str, err
}

func AgentLogout(url string, jsonData string) (str string, err error) {
	req := httplib.Post("http://"+url + "/uniswitchAgent/agentLogout")
	//req.Header("Content-Encoding", "utf-8")
	req.Header("Content-Type", "application/json; charset=utf-8")
	req.Body(jsonData)
	str, err = req.String()
	if err != nil {
		logs.Info(str)
	}
	return str, err
}

func UploadAgentStatus(url string, jsonData string) (str string, err error) {
	req := httplib.Post("http://"+url + "/uniswitchAgent/uploadAgentStatus")
	//req.Header("Content-Encoding", "utf-8")
	req.Header("Content-Type", "application/json; charset=utf-8")
	req.Body(jsonData)
	str, err = req.String()
	if err != nil {
		logs.Info(str)
	}
	return str, err
}