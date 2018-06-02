package utility

import (
	"testing"
	"fmt"
)

func TestByteblob(t *testing.T) {
	blob := new(Byteblob)
	blob.SetHex("000000000000000004ec466ce4732fe6f1ed1cddc2ed4b328fff5224276e3f6f")
	fmt.Println(blob.GetData())
	fmt.Println(blob.GetHex())
	fmt.Println(blob.GetDataSize())
}
