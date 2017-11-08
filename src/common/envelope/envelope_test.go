package envelope

import (
	"fmt"
	"testing"

	"uniswitch-agent/src/common/box"
	"uniswitch-agent/src/common/secretbox"
)

// seal: msg, secretKey, sessionPub, ----> cipher, encryptedSecretKey, tempPub
// open: cipher, encryptedSecretKey, tempPub, sessionPri ----> plain, ok
func Test_myEnvelope(t *testing.T) {
	secretKey := secretbox.GenerateSecretKey()
	fmt.Println(secretKey)

	sessionPub, sessionPri, err := box.GenerateKeyPair()
	fmt.Println(sessionPub, sessionPri, err)

	msg := "f*ck envelope seal and open ?"
	fmt.Println(msg)

	cipher, encryptedSecretKey, tempPub := Seal(msg, secretKey, sessionPub)
	fmt.Println(cipher, encryptedSecretKey, tempPub)

	plain, ok := Open(cipher, encryptedSecretKey, tempPub, sessionPri)
	fmt.Println(plain, ok)
}

func Test_envelope(t *testing.T) {
	secretKey := secretbox.GenerateSecretKey()
	fmt.Println(secretKey)

	sessionPub, sessionPri, err := box.GenerateKeyPair()
	fmt.Println(sessionPub, sessionPri, err)

	tempPub, tempPri, err := box.GenerateKeyPair()
	fmt.Println(tempPub, tempPri, err)

	msg := "f*ck envelope"
	fmt.Println(msg)

	cipher := secretbox.Seal(secretKey, msg)
	fmt.Println(cipher)

	encryptedSecretKey := box.Seal(secretKey, sessionPub, tempPri)
	fmt.Println(encryptedSecretKey)

	decryptedSecretKey, ok := box.Open(encryptedSecretKey, tempPub, sessionPri)
	fmt.Println(decryptedSecretKey, ok)

	plain, ok := secretbox.Open(decryptedSecretKey, cipher)
	fmt.Println(plain, ok)
}
