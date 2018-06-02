package script

import (
	"testing"
	"utility"
	"io"
	"bytes"
	"fmt"
)

func TestScript(t *testing.T) {
	blob := new(utility.Byteblob)
	blob.SetHex("1976a914cb02eb6ea281ef1fa194f55c530f6a080cfce99288ac")
	fmt.Println("byte blob:", blob.GetData())

	bytesBuf := bytes.NewBuffer(blob.GetData())
	bufReader := io.Reader(bytesBuf)
	script := new(Script)
	script.UnPack(bufReader)
	fmt.Println("script data:", script.GetScriptBytes())

	bytesBuf = bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	script.Pack(bufWriter)
	fmt.Println("byte blob:", bytesBuf.Bytes())
}
