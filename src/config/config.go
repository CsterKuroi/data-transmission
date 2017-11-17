package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"uniswitch-agent/src/common"
	"uniswitch-agent/src/common/box"
	"uniswitch-agent/src/common/sign"

	"github.com/astaxie/beego/logs"
)

type _Config struct {
	Sign    Keypair
	Encrypt Keypair
}

type Keypair struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
}

var Config _Config

func init() {
	_user, err := user.Current()
	if err != nil {
		logs.Error(err.Error())
	}
	fileName := _user.HomeDir + "/.agent"
	_, err = os.Open(fileName)
	if err != nil {
		logs.Info(err.Error())
		return
	}
	FileToConfig()
}

func FileToConfig() {
	_user, err := user.Current()
	if err != nil {
		logs.Error(err.Error())
	}
	fileName := _user.HomeDir + "/.agent"
	file, err := os.Open(fileName)
	if err != nil {
		logs.Error(err.Error())
		logs.Error("please create default config by 'agent configure' or 'go run main.go agent'")
		os.Exit(1)
	}
	_byte, err := ioutil.ReadAll(file)
	if err != nil {
		logs.Error(err.Error())
		logs.Error("please checkout your config file OR remove it", fileName)
		os.Exit(1)
	}
	err = json.Unmarshal(_byte, &Config)
	if err != nil {
		logs.Error(err.Error())
		logs.Error("please checkout your config file OR remove it", fileName)
		os.Exit(1)
	}
	logs.Debug(common.Serialize(Config))
}

func createNewConfig() _Config {
	var newConfig _Config
	//Sign keypair
	pub, priv := sign.GenerateKeypair()
	newConfig.Sign.PublicKey = pub
	newConfig.Sign.PrivateKey = priv
	//Encrypt keyring
	pub, priv, _ = box.GenerateKeyPair()
	newConfig.Encrypt.PublicKey = pub
	newConfig.Encrypt.PrivateKey = priv
	return newConfig
}

func ConfigToFile() {
	_user, err := user.Current()
	if err != nil {
		logs.Error(err.Error())
	}
	fileName := _user.HomeDir + "/.agent"
	_, err = os.Stat(fileName)
	if err == nil { //文件存在
		fmt.Println("Config file already exist, do you want to override it?")
		fmt.Println("Please input y(es) or n(o) ")
		inputReader := bufio.NewReader(os.Stdin)
		p := make([]byte, 10)
		inputReader.Read(p)
		if p[0] != []byte("y")[0] {
			fmt.Println("Give up to override it!")
			return
		}
	}
	configFile, err := os.Create(fileName)
	defer configFile.Close()
	if err != nil {
		logs.Error(err.Error())
	}

	newConfig := createNewConfig()
	str := common.SerializePretty(newConfig)
	_, err = configFile.Write([]byte(str + "\n"))
	if err != nil {
		logs.Error(err.Error())
	} else {
		fmt.Println("crate config file successful!\n", str)
	}
}
