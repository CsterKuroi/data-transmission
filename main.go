package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"os"
	"uniswitch-agent/src/core/task"
	_ "uniswitch-agent/src/web/api/routers"
	"uniswitch-agent/src/web/req"
)

var (
	uniSwitchHost = beego.AppConfig.String("uniswitch.host")
	registerUrl   = beego.AppConfig.String("registerUrl")
)

func main() {
	cmd := os.Args[0]
	logs.Info("operation : %s\n", cmd)
	//cmd: register start stop
	if cmd == "register" {
		register()
	} else if cmd == "start" {
		start()
	} else if cmd == "stop" {
		stop()
	}
}

func register() {
	logs.Info("register agent")
	//TODO pwd
	url := "" + uniSwitchHost + registerUrl

	//TODO get agent param
	registerParam := ""
	req.RegisterAgentToSwitch(url, registerParam)
}

func start() {
	logInit()
	//TODO update login status

	go task.DequeueTask()
	logs.Info("beego start run")
	beego.Run()
}

func stop() {

}

func logInit() {
	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	logs.SetLogFuncCall(true)
	//如果你的应用自己封装了调用 log 包,那么需要设置 SetLogFuncCallDepth,默认是 2,
	// 也就是直接调用的层级,如果你封装了多层,那么需要根据自己的需求进行调整.
	// logs 里面修改的话,此处请勿重复设置!
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	beego.BConfig.Log.AccessLogs = true

	//如果不想在控制台输出log相关的，可以打开下面设置
	//if u want not output to console, open following line!
	//beego.BeeLogger.DelLogger("console")

	// order 顺序必须按照
	// 1. logs.SetLevel(level)
	// 2. logs.SetLogger(logs.AdapterMultiFile, log_config_str)
	logs.SetLevel(logs.LevelDebug)
	//logs.SetLevel(logs.LevelInfo)
}
