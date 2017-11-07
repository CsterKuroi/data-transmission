package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"uniswitch-agent/src/core/task"
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
	logs.Info("Recive data send task :", param)
	task.EnqueueTask(param)
}
