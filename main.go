package main

import (
	"encoding/json"
	"fmt"
	"os"

	"uniswitch-agent/src/common"
	"uniswitch-agent/src/config"
	"uniswitch-agent/src/db/redis"
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
	os.Args = append(os.Args, "start")
	if len(os.Args) == 1 {
		fmt.Println("cmd:\n" +
			"  init     : init sign and encrypt\n" +
			"  register : register to uniswitch\n" +
			"  start    : start uniswitch-agent\n" +
			"  stop     : stop uniswitch-agent \n" +
			"\n")
	} else {
		cmd := os.Args[1]
		if cmd == "init" {
			initKeys()
		} else if cmd == "register" {
			config.FileToConfig()
			register()
		} else if cmd == "start" {
			config.FileToConfig()
			start()
		} else if cmd == "stop" {
			stop()
		} else {
			fmt.Println("Agent only support these commands : init | register | start | stop. Please try again!")
		}
	}
}

func initKeys() {
	config.ConfigToFile()
}

func register() {
	logs.Info("Add Agent")
	res, err := req.AddAgent(uniSwitchHost, "")
	var registerResult map[string]interface{}
	err = json.Unmarshal([]byte(res), &registerResult)
	if err != nil {
		panic(err)
	}
	if registerResult["code"].(float64) != 200 {
		panic("添加Agent失败 ")
	}
	redis.Store(config.Config.Encrypt.PublicKey, "agentId", registerResult["result"].(string))
	if err != nil {
		panic(err)
	}
	var name, pwdF, pwdS string
	fmt.Print("Please input your name : ")
	fmt.Scanln(&name)
	fmt.Print("Please input your password : ")
	fmt.Scanln(&pwdF)
	fmt.Print("Please input your password again : ")
	fmt.Scanln(&pwdS)
	if len(pwdF) < 6 {
		fmt.Println("Password must contain at least 6 characters!")
	} else if pwdF != pwdS {
		fmt.Println("Passwords do not match!")
	} else {
		result := make(map[string]string)
		result["id"] = registerResult["result"].(string)
		result["name"] = name
		result["password"] = pwdF
		result["signPubkey"] = config.Config.Sign.PublicKey
		result["encryptPubkey"] = config.Config.Encrypt.PublicKey
		result["address"] = beego.AppConfig.String("agent.host")
		logs.Info("register Agent")
		res, err := req.RegisterAgent(uniSwitchHost, common.Serialize(result))
		logs.Debug(res, err)
		if err != nil {
			panic(err)
		}
	}
}

func login() {
	var name, pwdF string
	fmt.Print("Please input your name : ")
	fmt.Scanln(&name)
	fmt.Print("Please input your password : ")
	fmt.Scanln(&pwdF)
	result := make(map[string]string)
	result["name"] = name
	result["password"] = pwdF
	logs.Info("login Agent")
	res, err := req.AgentLogin(uniSwitchHost, common.Serialize(result))
	logs.Info(res, err)
	if err != nil {
		panic(err)
	}
	var loginResult map[string]interface{}
	err = json.Unmarshal([]byte(res), &loginResult)
	if err != nil {
		panic(err)
	}
	if loginResult["code"].(float64) != 200 {
		panic("用户名或密码有误 ")
	}
	redis.Store(config.Config.Encrypt.PublicKey, "token", loginResult["result"].(string))
	if err != nil {
		panic(err)
	}
}

func logout() {
	result := make(map[string]string)
	res, err := redis.Get(config.Config.Encrypt.PublicKey, "token")
	token := string(res.([]byte))
	logs.Debug(res, err)
	if err != nil {
		panic(err)
	}
	result["token"] = token
	logs.Info("logout Agent")
	res, err = req.AgentLogout(uniSwitchHost, common.Serialize(result))
	logs.Debug(res, err)
	if err != nil {
		panic(err)
	}
}

func heartbeat() {
	//TODO per hour
	result := make(map[string]string)
	res, err := redis.Get(config.Config.Encrypt.PublicKey, "token")
	token := string(res.([]byte))
	logs.Debug(res, err)
	if err != nil {
		panic(err)
	}
	result["token"] = token
	res, err = redis.Get(config.Config.Encrypt.PublicKey, "agentId")
	agentId := string(res.([]byte))
	logs.Debug(res, err)
	if err != nil {
		panic(err)
	}
	result["id"] = agentId
	logs.Info("heartbeat Agent")
	res, err = req.UploadAgentStatus(uniSwitchHost, common.Serialize(result))
	logs.Debug(res, err)
	if err != nil {
		panic(err)
	}
}

func start() {
	logInit()

	logs.Info("login with switch")
	login()

	logs.Info("heartbeat with switch")
	heartbeat()

	logs.Info("beego start run")
	beego.Run()
}

func stop() {
	logout()
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
