package envelope

import (
	"fmt"
	"testing"

	"uniswitch-agent/src/common/box"
	"uniswitch-agent/src/common/secretbox"
)

// en: secretKey,sessionPub,[tempPir],msg
// de: encryptedSecretKey,sessionPri,[tempPub],cipher
func Test_myEnvelope(t *testing.T) {
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
