package blob

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestByteblob(t *testing.T) {
	blob := new(Byteblob)
	blob.SetHex("000000000000000004ec466ce4732fe6f1ed1cddc2ed4b328fff5224276e3f6f")
	fmt.Println(blob.GetData())
	fmt.Println(blob.GetHex())
	fmt.Println(blob.GetDataSize())

	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	blob.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())

	bytesBuf = bytes.NewBuffer(bytesBuf.Bytes())
	bufReader := io.Reader(bytesBuf)
	blob = new(Byteblob)
	blob.UnPack(bufReader)
	fmt.Println("blob data:", blob.GetData())
}
