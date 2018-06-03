package bigint

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestUint160(t *testing.T) {
	uint160 := new(Uint160)
	uint160.SetHex("6608a2bdf5a96d58e8ec18548eae3724dae59797")
	fmt.Println(uint160.GetData())
	fmt.Println(uint160.GetHex())
	fmt.Println(uint160.GetDataSize())

	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	uint160.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())

	bytesBuf = bytes.NewBuffer(bytesBuf.Bytes())
	bufReader := io.Reader(bytesBuf)
	uint160 = new(Uint160)
	uint160.UnPack(bufReader)
	fmt.Println("uint160 data:", uint160.GetData())
}
