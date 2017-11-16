package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"uniswitch-agent/src/common/box"
	"uniswitch-agent/src/common/secretbox"
	"uniswitch-agent/src/core/task"
	"uniswitch-agent/src/db/redis"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	agentPub, agentPri = "BFhZXZmUWqd5B7YdF9xshvHuJkskcUvnx5zTXsB22Mrk", "AzsqoRz1p47uVMWdtaqc6xNniK9z149YLg6jT1ScHhr9"
)

type MainController struct {
	beego.Controller
}

func (m *MainController) Hello() {
	m.Ctx.WriteString("Hello,")
	m.Ctx.ResponseWriter.Write([]byte(" Welcome to Uni-Switch Agent!"))
}

func (m *MainController) Public() {
	var result map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &result)
	logs.Info("Api receive public", result)
	plain, ok := box.Open(result["cipher"], result["temp"], agentPri)
	logs.Info("public open box", plain, ok)
	redis.Store(result["oid"], "public", plain)
	//TODO submit
}

func (m *MainController) Private() {
	var result map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &result)
	logs.Info("Api receive private", result)
	plain, ok := box.Open(result["cipher"], result["temp"], agentPri)
	logs.Info("private open box", plain, ok)
	redis.Store(result["oid"], "private", plain)
	//TODO submit
}

func (m *MainController) Secret() {
	var result map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &result)
	logs.Info("Api receive secret", result)
	plain, ok := box.Open(result["cipher"], result["temp"], agentPri)
	logs.Info("secret open box", plain, ok)
	redis.Store(result["oid"], "secret", plain)
	//TODO submit
}

func sendData(oid, public, secret, address, data string) {
	tempPub, tempPri, _ := box.GenerateKeyPair()
	url := address

	result := make(map[string]string)
	result["oid"] = oid
	result["temp"] = tempPub
	result["secret"] = box.Seal(secret, public, tempPri)
	result["data"] = secretbox.Seal(secret, data)

	jsonStr, _ := json.Marshal(result)
	fmt.Println("json:", result)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	//TODO submit
}

func (m *MainController) Address() {
	var result map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &result)
	time.Sleep(time.Second)

	logs.Info("Api receive address", result)
	plain, ok := box.Open(result["cipher"], result["temp"], agentPri)
	logs.Info("address open box", plain, ok)
	redis.Store(result["oid"], "address", plain)

	//send data
	res, err := redis.Get(result["oid"], "public")
	public := string(res.([]byte))
	logs.Info("redis get public", public, err)

	res, err = redis.Get(result["oid"], "secret")
	secret := string(res.([]byte))
	logs.Info("redis get secret", secret, err)
	data := "A staff member in costume waits for visitors at a booth for Chinese Twitter-like Sina Weibo at the Global Mobile Internet Conference in Beijing, April 27, 2017."
	go sendData(result["oid"], public, secret, plain, data)
	//TODO submit
}

func (m *MainController) Data() {
	var result map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &result)
	logs.Info("Api receive data", result)
	time.Sleep(time.Second * 2)

	res, err := redis.Get(result["oid"], "private")
	private := string(res.([]byte))
	logs.Info("redis get private", private, err)
	secret, ok := box.Open(result["secret"], result["temp"], private)
	logs.Info("secret open box", secret, ok)
	//data, ok := secretbox.Open(secret, result["data"])
	//logs.Info("data open secretbox", data, ok)
	redis.Store(result["oid"], "edata", result["data"])
	//TODO submit
}

func (m *MainController) DecryptData() {
	var result map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &result)
	logs.Info("Api receive data", result)
	time.Sleep(time.Second * 3)

	res, err := redis.Get(result["oid"], "secret")
	secret := string(res.([]byte))
	logs.Info("redis get secret", secret, err)

	res, err = redis.Get(result["oid"], "edata")
	edata := string(res.([]byte))
	logs.Info("redis get edata", edata, err)

	data, ok := secretbox.Open(secret, edata)
	logs.Info("data open secretbox", data, ok)
	redis.Store(result["oid"], "data", data)
	//TODO submit
}

func (m *MainController) DestroyData() {
	var result map[string]string
	json.Unmarshal(m.Ctx.Input.RequestBody, &result)
	logs.Info("Api receive data", result)
	time.Sleep(time.Second * 4)

	res, err := redis.Delete(result["oid"])
	logs.Info("redis delete key", res, err)
	//TODO submit
}

func (m *MainController) SendDataToAnotherAgent() {
	param := m.Ctx.Input.RequestBody
	logs.Info("Receive data send task :", param)
	task.EnqueueTask(param)
}
