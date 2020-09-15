package bigint

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestUint160(t *testing.T) {
	uint160 := new(Uint160)
	_ = uint160.SetHex("6608a2bdf5a96d58e8ec18548eae3724dae59797")
	fmt.Println(uint160.GetData())
	fmt.Println(uint160.GetHex())
	fmt.Println(uint160.GetDataSize())

	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = uint160.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())

	bytesBuf = bytes.NewBuffer(bytesBuf.Bytes())
	bufReader := io.Reader(bytesBuf)
	uint160 = new(Uint160)
	_ = uint160.UnPack(bufReader)
	fmt.Println("uint160 data:", uint160.GetData())
}

func TestUint160_2(t *testing.T) {
	uint160 := new(Uint160)
	_ = uint160.SetData([]byte{151, 151, 229, 218, 36, 55, 174, 142, 84, 24, 236, 232,
		88, 109, 169, 245, 189, 162, 8, 102})
	fmt.Println(uint160.GetData())
	fmt.Println(uint160.GetHex())
	fmt.Println(uint160.GetDataSize())

	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = uint160.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())

	bytesBuf = bytes.NewBuffer(bytesBuf.Bytes())
	bufReader := io.Reader(bytesBuf)
	uint160 = new(Uint160)
	_ = uint160.UnPack(bufReader)
	fmt.Println("uint160 data:", uint160.GetData())
}
