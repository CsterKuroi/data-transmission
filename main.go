package main

import (
	"fmt"
	"os"

	"uniswitch-agent/src/core/task"
	_ "uniswitch-agent/src/web/api/routers"
	"uniswitch-agent/src/web/req"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	uniSwitchHost = beego.AppConfig.String("uniswitch.host")
	registerUrl   = beego.AppConfig.String("registerUrl")
	checkPwdUrl   = beego.AppConfig.String("checkPwdUrl")
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("cmd:\n" +
			"  register : register to uniswitch\n" +
			"  start    : start uniswitch-agent\n" +
			"  stop     : stop uniswitch-agent\n" +
			"\n")
	} else {
		cmd := os.Args[1]
		if cmd == "register" {
			register()
		} else if cmd == "start" {
			start()
		} else if cmd == "stop" {
			stop()
		} else {
			fmt.Println("Agent only support these commands : register | start | stop. Please try again!")
		}
	}
}

func register() {
	logs.Info("Register Agent")

	var pwdF, pwdS string
	fmt.Print("Please input your password : ")
	fmt.Scanln(&pwdF)
	fmt.Print("Please input your password again : ")
	fmt.Scanln(&pwdS)
	if len(pwdF) < 6 {
		fmt.Println("Password must contain at least 6 characters!")
	} else if pwdF != pwdS {
		fmt.Println("Passwords do not match!")
	} else {
		url := uniSwitchHost + registerUrl
		//TODO get agent param
		registerParam := ""
		res, err := req.RegisterAgentToSwitch(url, registerParam)
		//TODO deal with result
		logs.Info("res:", res)
		logs.Info("err:", err)
		fmt.Println("Register success! Please wait for the activation.")
	}
}

func start() {
	logInit()
	var pwdF string
	fmt.Print("Please input your password : ")
	fmt.Scanln(&pwdF)
	//TODO check pwd
	hashPwd := pwdF
	url := uniSwitchHost + checkPwdUrl
	_, err := req.CheckAgentPwdInSwitch(url, hashPwd)
	if err != nil {
		fmt.Println("Check password failed. Please try again!")
	} else {
		//TODO update login status by res

	}
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
