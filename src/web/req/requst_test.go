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
	result["id"] = "25"
	result["name"] = "leehh"
	result["password"] = "123456"
	result["signPubkey"] = "aaa"
	result["encryptPubkey"] = "bbb"
	result["address"] = "172.17.7.112:8099"

	res, err := RegisterAgent(url, common.Serialize(result))
	fmt.Println(res, "\n", err)
}

func TestAgentLogin(t *testing.T) {
	result := make(map[string]string)
	result["name"] = "leehh"
	result["password"] = "123456"

	res, err := AgentLogin(url, common.Serialize(result))
	fmt.Println(res, "\n", err)
}

func TestAgentLogout(t *testing.T) {
	result := make(map[string]string)
	result["token"] = "8778C1e258B6a6d91851da3e2C3b5092e050dDffEf70f897"

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