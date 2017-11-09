package envelope

import (
	"fmt"
	"io/ioutil"
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

func Test_myEnvelopeText(t *testing.T) {
	secretKey := secretbox.GenerateSecretKey()
	fmt.Println(secretKey)

	sessionPub, sessionPri, err := box.GenerateKeyPair()
	fmt.Println(sessionPub, sessionPri, err)

	f, err := ioutil.ReadFile("envelope.go")
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	msg := string(f)
	fmt.Println(msg)

	cipher, encryptedSecretKey, tempPub := Seal(msg, secretKey, sessionPub)
	fmt.Println(cipher, encryptedSecretKey, tempPub)

	plain, ok := Open(cipher, encryptedSecretKey, tempPub, sessionPri)
	fmt.Println(plain, ok)
}

func Test_myEnvelopePhoto(t *testing.T) {
	secretKey := secretbox.GenerateSecretKey()
	fmt.Println(secretKey)

	sessionPub, sessionPri, err := box.GenerateKeyPair()
	fmt.Println(sessionPub, sessionPri, err)

	f, err := ioutil.ReadFile("shan.jpg")
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	msg := string(f)
	fmt.Println(msg)

	cipher, encryptedSecretKey, tempPub := Seal(msg, secretKey, sessionPub)
	fmt.Println(cipher, encryptedSecretKey, tempPub)

	plain, ok := Open(cipher, encryptedSecretKey, tempPub, sessionPri)
	fmt.Println(plain, ok)

	fn := "shan_return.jpg"
	data := []byte(plain)
	ioutil.WriteFile(fn, data, 0664)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
}

// too big
//func Test_myEnvelopeVideo(t *testing.T) {
//	secretKey := secretbox.GenerateSecretKey()
//	fmt.Println(secretKey)
//
//	sessionPub, sessionPri, err := box.GenerateKeyPair()
//	fmt.Println(sessionPub, sessionPri, err)
//
//	f, err := ioutil.ReadFile("second.mp4")
//	if err != nil {
//		fmt.Printf("%s\n", err)
//		panic(err)
//	}
//	msg := string(f)
//	fmt.Println("msg")
//
//	cipher, encryptedSecretKey, tempPub := Seal(msg, secretKey, sessionPub)
//	fmt.Println(cipher, encryptedSecretKey, tempPub)
//
//	plain, _ := Open(cipher, encryptedSecretKey, tempPub, sessionPri)
//	//fmt.Println(plain, ok)
//
//	fn := "second_return.mp4"
//	data := []byte(plain)
//	fmt.Println(os.ModeAppend)
//	ioutil.WriteFile(fn, data, 0664)
//	if err != nil {
//		fmt.Printf("%s\n", err)
//		panic(err)
//	}
//}
