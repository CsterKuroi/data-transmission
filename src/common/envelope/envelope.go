package envelope

import (
	"uniswitch-agent/src/common/box"
	"uniswitch-agent/src/common/secretbox"
)

func Seal(msg string, secretKey string, sessionPub string) (string, string, string) {
	tempPub, tempPri, err := box.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	cipher := secretbox.Seal(secretKey, msg)
	encryptedSecretKey := box.Seal(secretKey, sessionPub, tempPri)
	return cipher, encryptedSecretKey, tempPub
}

func Open(cipher string, encryptedSecretKey string, tempPub string, sessionPri string) (string, bool) {
	decryptedSecretKey, ok := box.Open(encryptedSecretKey, tempPub, sessionPri)
	if !ok {
		return "", ok
	}
	plain, ok := secretbox.Open(decryptedSecretKey, cipher)
	return plain, ok
}
