package req

import (
	"fmt"
	"testing"

	"uniswitch-agent/src/common"
)

var url = "http://172.17.7.110:8091"

func TestReqGet(t *testing.T) {
	ReqGet()
}

func TestAddAgent(t *testing.T) {
	res, err := AddAgent(url, "")
	fmt.Println(res, "\n", err)
}

func TestRegisterAgent(t *testing.T) {
	result := make(map[string]string)
	result["id"] = "23"
	result["name"] = "lee"
	result["password"] = "123456"
	result["signPubkey"] = "aaa"
	result["encryptPubkey"] = "bbb"
	result["address"] = "172.17.7.112:8099"

	res, err := RegisterAgent(url, common.Serialize(result))
	fmt.Println(res, "\n", err)
}

func TestAgentLogin(t *testing.T) {
	result := make(map[string]string)
	result["name"] = "lee"
	result["password"] = "123456"

	res, err := AgentLogin(url, common.Serialize(result))
	fmt.Println(res, "\n", err)
}

func TestAgentLogout(t *testing.T) {
	result := make(map[string]string)
	result["token"] = "7F612ad48c9c534266f010E57C926cbB59022B6a84f2d383"

	res, err := AgentLogout(url, common.Serialize(result))
	fmt.Println(res, "\n", err)
}

func TestUploadAgentStatus(t *testing.T) {
	result := make(map[string]string)
	result["id"] = "23"
	result["token"] = "dC70C6f374F054f9594Ec1F979334f84ba5976caD7aD30C1"

	res, err := UploadAgentStatus(url, common.Serialize(result))
	fmt.Println(res, "\n", err)
}