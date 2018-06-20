package utility

import (
	"testing"
	"fmt"
)

func TestSha256(t *testing.T) {
	fmt.Println("sha256 calc", Sha256Hex([]byte{0x01}))
}
