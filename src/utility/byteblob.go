package utility

import (
	"strings"
	"fmt"
)

type Byteblob struct {
	data	[]byte
}

func (b Byteblob) isValidHex(hexStr string) bool {
	if hexStr[0] == '0' && hexStr[1] == 'x' {
		hexStr = hexStr[2:]
	}

	if len(hexStr) % 2 == 1 {
		return false
	}

	hexStr = strings.ToLower(hexStr)
	for _, c := range []byte(hexStr) {
		if !((c >= 0x30 && c <= 0x39) || (c >= 0x61 && c <= 0x66)) {
			return false
		}
	}

	return true
}

func (b *Byteblob) SetHex(hexStr string) {
	if hexStr[0] == '0' && hexStr[1] == 'x' {
		hexStr = hexStr[2:]
	}

	Assert(b.isValidHex(hexStr), "invalid hex string")
	blobLength := len(hexStr) / 2
	for i := 0; i < blobLength; i++ {
		num1, _ := HexCharToNumber(hexStr[2*i])
		num2, _ := HexCharToNumber(hexStr[2*i+1])
		b.data = append(b.data, byte((num1 << 4) | num2))
	}
}

func (b Byteblob) GetHex() string {
	var stringRet string
	for _, c := range b.data {
		stringRet += fmt.Sprintf("%02x", c)
	}
	return stringRet
}

func (b Byteblob) GetData() []byte {
	return b.data
}

func (b Byteblob) GetDataSize() int {
	return len(b.data)
}
