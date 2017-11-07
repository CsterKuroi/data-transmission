package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (m *MainController) Post() {
	m.Ctx.ResponseWriter.Write([]byte("world!"))
	m.Ctx.WriteString("hello")
}
