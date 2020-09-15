package blob

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestBaseblob(t *testing.T) {
	blob := new(Baseblob)
	_ = blob.SetHex("000000000000000004ec466ce4732fe6f1ed1cddc2ed4b328fff5224276e3f6f")
	fmt.Println(blob.GetData())
	fmt.Println(blob.GetHex())
	fmt.Println(blob.GetDataSize())

	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = blob.Pack(bufWriter, 32)
	fmt.Println("byte buffer:", bytesBuf.Bytes())

	bytesBuf = bytes.NewBuffer(bytesBuf.Bytes())
	bufReader := io.Reader(bytesBuf)
	blob = new(Baseblob)
	_ = blob.UnPack(bufReader, 32)
	fmt.Println("blob data:", blob.GetData())
}

func TestBaseblob_2(t *testing.T) {
	blob := new(Baseblob)
	blob.SetData([]byte{111, 63, 110, 39, 36, 82, 255, 143, 50, 75, 237, 194, 221, 28, 237, 241,
		230, 47, 115, 228, 108, 70, 236, 4, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(blob.GetData())
	fmt.Println(blob.GetHex())
	fmt.Println(blob.GetDataSize())

	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = blob.Pack(bufWriter, 32)
	fmt.Println("byte buffer:", bytesBuf.Bytes())

	bytesBuf = bytes.NewBuffer(bytesBuf.Bytes())
	bufReader := io.Reader(bytesBuf)
	blob = new(Baseblob)
	_ = blob.UnPack(bufReader, 32)
	fmt.Println("blob data:", blob.GetData())
}
