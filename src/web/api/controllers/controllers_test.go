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
	oid                = secretbox.GenerateSecretKey()
	public, private, _ = box.GenerateKeyPair()
	secret             = secretbox.GenerateSecretKey()
	address            = "localhost:8099"
	data               = "e_data"
)

func TestMainController_Public(t *testing.T) {
	url := "http://127.0.0.1:8099/public"

	pub := make(map[string]string)
	pub["public"] = public
	pub["oid"] = oid

	jsonStr, _ := json.Marshal(pub)
	fmt.Println("json:", jsonStr)
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

	pri := make(map[string]string)
	pri["private"] = private
	pri["oid"] = oid

	jsonStr, _ := json.Marshal(pri)
	fmt.Println("json:", jsonStr)
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

	sec := make(map[string]string)
	sec["secret"] = secret
	sec["oid"] = oid

	jsonStr, _ := json.Marshal(sec)
	fmt.Println("json:", jsonStr)
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

	add := make(map[string]string)
	add["address"] = address
	add["oid"] = oid

	jsonStr, _ := json.Marshal(add)
	fmt.Println("json:", jsonStr)
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

func TestMainController_Data(t *testing.T) {
	url := "http://127.0.0.1:8099/data"

	da := make(map[string]string)
	da["data"] = data
	da["oid"] = oid

	jsonStr, _ := json.Marshal(da)
	fmt.Println("json:", jsonStr)
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
	fmt.Println(oid)
	fmt.Println(public, private)
	fmt.Println(secret)
	fmt.Println(address)
	fmt.Println(data)
}
