package sign

import (
	"bytes"

	"github.com/astaxie/beego/logs"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ed25519"
)

//encrypt
func GenerateKeypair(seed ...string) (pub string, priv string) {
	var publicKeyBytes, privateKeyBytes []byte
	var err error
	if len(seed) >= 1 {
		publicKeyBytes, privateKeyBytes, err = ed25519.GenerateKey(bytes.NewReader(base58.Decode(seed[0])))
	} else {
		publicKeyBytes, privateKeyBytes, err = ed25519.GenerateKey(nil)
	}
	if err != nil {
		logs.Error(err.Error())
	}
	publicKeyBase58 := base58.Encode(publicKeyBytes)
	privateKeyBase58 := base58.Encode(privateKeyBytes[0:32])
	return publicKeyBase58, privateKeyBase58
}

func Sign(priv string, msg string) string {
	pub, _ := GenerateKeypair(priv)
	privByte := base58.Decode(priv)
	pubByte := base58.Decode(pub)
	privateKey := make([]byte, 64)
	copy(privateKey[:32], privByte)
	copy(privateKey[32:], pubByte)
	sigByte := ed25519.Sign(privateKey, []byte(msg))
	return base58.Encode(sigByte)
}

func Verify(pub string, msg string, sig string) bool {
	pubByte := base58.Decode(pub)
	publicKey := make([]byte, 32)
	copy(publicKey, pubByte)
	return ed25519.Verify(publicKey, []byte(msg), base58.Decode(sig))
}
