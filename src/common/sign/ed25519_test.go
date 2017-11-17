package sign

import (
	"fmt"
	"testing"
)

func TestGenerateKeypair(t *testing.T) {
	fmt.Println(GenerateKeypair())
}

func TestSign(t *testing.T) {
	_, pri := GenerateKeypair()
	msg := "a"
	fmt.Println(Sign(pri, msg))
}

func TestVerify(t *testing.T) {
	pub, pri := GenerateKeypair()
	msg := "a"
	sig := Sign(pri, msg)
	fmt.Println(Verify(pub, msg, sig))
}
