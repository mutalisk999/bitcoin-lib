package utility

import (
	"fmt"
	"testing"
)

func TestHexCharToNumber(t *testing.T) {
	var c byte
	for c = '0'; c <= '9'; c++ {
		fmt.Println(HexCharToNumber(c))
	}
	for c = 'A'; c <= 'F'; c++ {
		fmt.Println(HexCharToNumber(c))
	}
	for c = 'a'; c <= 'f'; c++ {
		fmt.Println(HexCharToNumber(c))
	}
}
