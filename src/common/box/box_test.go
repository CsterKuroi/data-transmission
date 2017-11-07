package box

import (
	cryptoRand "crypto/rand"
	"fmt"
	"io"
	"testing"

	"golang.org/x/crypto/nacl/box"
)

func Test_myBox(t *testing.T) {
	pub1, pri1, err := GenerateKeyPair()
	fmt.Println(pub1, pri1, err)

	pub2, pri2, err := GenerateKeyPair()
	fmt.Println(pub2, pri2, err)

	msg := "f*ck it again"
	fmt.Println(msg)

	cipher := Seal(msg, pub2, pri1)
	fmt.Println(cipher)

	plain, ok := Open(cipher, pub1, pri2)
	fmt.Println(plain, ok)
}

func Test_box(t *testing.T) {
	senderPublicKey, senderPrivateKey, err := box.GenerateKey(cryptoRand.Reader)
	if err != nil {
		panic(err)
	}

	recipientPublicKey, recipientPrivateKey, err := box.GenerateKey(cryptoRand.Reader)
	if err != nil {
		panic(err)
	}

	// You must use a different nonce for each message you encrypt with the
	// same key. Since the nonce here is 192 bits long, a random value
	// provides a sufficiently small probability of repeats.
	var nonce [24]byte
	if _, err := io.ReadFull(cryptoRand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	msg := []byte("Alas, poor Yorick! I knew him, Horatio")
	// This encrypts msg and appends the result to the nonce.
	encrypted := box.Seal(nonce[:], msg, &nonce, recipientPublicKey, senderPrivateKey)

	// The recipient can decrypt the message using their private key and the
	// sender's public key. When you decrypt, you must use the same nonce you
	// used to encrypt the message. One way to achieve this is to store the
	// nonce alongside the encrypted message. Above, we stored the nonce in the
	// first 24 bytes of the encrypted text.
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])
	decrypted, ok := box.Open(nil, encrypted[24:], &decryptNonce, senderPublicKey, recipientPrivateKey)
	if !ok {
		panic("decryption error")
	}
	fmt.Println(string(decrypted))
}
