package bigint

import (
	"testing"
	"fmt"
)

func TestUint160(t *testing.T) {
	uint160 := new(Uint160)
	uint160.SetHex("6608a2bdf5a96d58e8ec18548eae3724dae59797")
	fmt.Println(uint160.GetData())
	fmt.Println(uint160.GetHex())
	fmt.Println(uint160.GetDataSize())
}

