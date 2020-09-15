package blob

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestByteblob(t *testing.T) {
	blob := new(Byteblob)
	_ = blob.SetHex("000000000000000004ec466ce4732fe6f1ed1cddc2ed4b328fff5224276e3f6f")
	fmt.Println(blob.GetData())
	fmt.Println(blob.GetHex())
	fmt.Println(blob.GetDataSize())

	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = blob.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())

	bytesBuf = bytes.NewBuffer(bytesBuf.Bytes())
	bufReader := io.Reader(bytesBuf)
	blob = new(Byteblob)
	_ = blob.UnPack(bufReader)
	fmt.Println("blob data:", blob.GetData())
}

func TestByteblob_2(t *testing.T) {
	blob := new(Byteblob)
	blob.SetData([]byte{0, 0, 0, 0, 0, 0, 0, 0, 4, 236, 70, 108, 228, 115, 47, 230, 241, 237, 28,
		221, 194, 237, 75, 50, 143, 255, 82, 36, 39, 110, 63, 111})
	fmt.Println(blob.GetData())
	fmt.Println(blob.GetHex())
	fmt.Println(blob.GetDataSize())

	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = blob.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())

	bytesBuf = bytes.NewBuffer(bytesBuf.Bytes())
	bufReader := io.Reader(bytesBuf)
	blob = new(Byteblob)
	_ = blob.UnPack(bufReader)
	fmt.Println("blob data:", blob.GetData())
}
