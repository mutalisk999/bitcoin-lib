package utility

import (
	"testing"
	"fmt"
)

func TestIsLittleEndian(t *testing.T) {
	b := IsLittleEndian()
	fmt.Println("is little endian:", b)
}
