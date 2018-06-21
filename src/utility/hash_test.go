package utility

import (
	"fmt"
	"testing"
)

func TestSha256(t *testing.T) {
	fmt.Println("sha256 calc", Sha256Hex([]byte{0x01}))
}

func TestRipemd160(t *testing.T) {
	fmt.Println("ripemd160 calc", Ripemd160Hex([]byte{0x01}))
}

func TestHash160(t *testing.T) {
	fmt.Println("hash160 calc", Hash160Hex([]byte{0x01}))
}
