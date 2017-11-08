package controllers

import (
	"uniswitch-agent/src/core/task"

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

func (m *MainController) SendDataToAnotherAgent() {
	param := m.Ctx.Input.RequestBody
	logs.Info("Receive data send task :", param)
	task.EnqueueTask(param)
}
