package bigint

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestUint256(t *testing.T) {
	uint256 := new(Uint256)
	_ = uint256.SetHex("000000000000000004ec466ce4732fe6f1ed1cddc2ed4b328fff5224276e3f6f")
	fmt.Println(uint256.GetData())
	fmt.Println(uint256.GetHex())
	fmt.Println(uint256.GetDataSize())

	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = uint256.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())

	bytesBuf = bytes.NewBuffer(bytesBuf.Bytes())
	bufReader := io.Reader(bytesBuf)
	uint256 = new(Uint256)
	_ = uint256.UnPack(bufReader)
	fmt.Println("uint256 data:", uint256.GetData())
}

func TestUint256_2(t *testing.T) {
	uint256 := new(Uint256)
	_ = uint256.SetData([]byte{111, 63, 110, 39, 36, 82, 255, 143, 50, 75, 237, 194, 221, 28, 237,
		241, 230, 47, 115, 228, 108, 70, 236, 4, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(uint256.GetData())
	fmt.Println(uint256.GetHex())
	fmt.Println(uint256.GetDataSize())

	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = uint256.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())

	bytesBuf = bytes.NewBuffer(bytesBuf.Bytes())
	bufReader := io.Reader(bytesBuf)
	uint256 = new(Uint256)
	_ = uint256.UnPack(bufReader)
	fmt.Println("uint256 data:", uint256.GetData())
}
