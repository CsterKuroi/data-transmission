package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"uniswitch-agent/src/common/box"
	"uniswitch-agent/src/common/secretbox"
)

var (
	oid                 = secretbox.GenerateSecretKey()
	public, private, _  = box.GenerateKeyPair()
	tempPub, tempPri, _ = box.GenerateKeyPair()
	secret              = secretbox.GenerateSecretKey()
	address             = "http://127.0.0.1:8099/data"
	data                = "A pedestrian wades through the flooded road in Haikou, South China's Hainan province, Nov 14, 2017."
)

func TestMainController_Public(t *testing.T) {
	url := "http://127.0.0.1:8099/public"

	result := make(map[string]string)
	result["oid"] = oid
	result["temp"] = tempPub
	result["public"] = box.Seal(public, agentPub, tempPri)

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
}

func TestMainController_Private(t *testing.T) {
	url := "http://127.0.0.1:8099/private"

	result := make(map[string]string)
	result["oid"] = oid
	result["temp"] = tempPub
	result["private"] = box.Seal(private, agentPub, tempPri)

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
}

func TestMainController_Secret(t *testing.T) {
	url := "http://127.0.0.1:8099/secret"

	result := make(map[string]string)
	result["oid"] = oid
	result["temp"] = tempPub
	result["secret"] = box.Seal(secret, agentPub, tempPri)

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
}

func TestMainController_Address(t *testing.T) {
	url := "http://127.0.0.1:8099/address"

	result := make(map[string]string)
	result["oid"] = oid
	result["temp"] = tempPub
	result["address"] = box.Seal(address, agentPub, tempPri)

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
}

//func TestMainController_Data(t *testing.T) {
//	url := "http://127.0.0.1:8099/data"
//
//	result := make(map[string]string)
//	result["oid"] = oid
//	result["temp"] = tempPub
//	result["secret"] = box.Seal(secret, public, tempPri)
//	result["data"] = secretbox.Seal(secret, data)
//
//	jsonStr, _ := json.Marshal(result)
//	fmt.Println("json:", result)
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
//	if err != nil {
//		panic(err)
//	}
//	req.Header.Set("Content-Type", "application/json")
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	defer resp.Body.Close()
//	fmt.Println("response Status:", resp.Status)
//	fmt.Println("response Headers:", resp.Header)
//	body, _ := ioutil.ReadAll(resp.Body)
//	fmt.Println("response Body:", string(body))
//}

func TestMainController_DecryptData(t *testing.T) {
	url := "http://127.0.0.1:8099/decrypt"

	result := make(map[string]string)
	result["oid"] = oid

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
}

func TestMainController_DestroyData(t *testing.T) {
	url := "http://127.0.0.1:8099/destroy"

	result := make(map[string]string)
	result["oid"] = oid

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
}

func Test_var(t *testing.T) {
	fmt.Println(agentPub, agentPri)
	fmt.Println()
	fmt.Println(oid)
	fmt.Println(public, private)
	fmt.Println(tempPub, tempPri)
	fmt.Println(secret)
	fmt.Println(address)
	fmt.Println(data)
}
