package controllers

import (
	"encoding/json"

	"uniswitch-agent/src/core/task"
	"uniswitch-agent/src/db/redis"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (m *MainController) Hello() {
	m.Ctx.WriteString("Hello,")
	m.Ctx.ResponseWriter.Write([]byte(" Welcome to Uni-Switch Agent!"))
}

func (m *MainController) Public() {
	var public map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &public)
	redis.Store(public["oid"],"public",public["public"])
	logs.Info("Api receive Public", public)
}

func (m *MainController) Private() {
	var private map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &private)
	redis.Store(private["oid"],"private",private["private"])
	logs.Info("Api receive Private", private)
}

func (m *MainController) Secret() {
	var secret map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &secret)
	redis.Store(secret["oid"],"secret",secret["secret"])
	logs.Info("Api receive Secret", secret)
}

func (m *MainController) Address() {
	var address map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &address)
	redis.Store(address["oid"],"address",address["address"])
	logs.Info("Api receive Address", address)
}

func (m *MainController) Data() {
	var data map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &data)
	redis.Store(data["oid"],"data",data["data"])
	logs.Info("Api receive Data", data)
}

func (m *MainController) SendDataToAnotherAgent() {
	param := m.Ctx.Input.RequestBody
	logs.Info("Receive data send task :", param)
	task.EnqueueTask(param)
}
