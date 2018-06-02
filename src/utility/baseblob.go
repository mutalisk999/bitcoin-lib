package utility

import (
	"strings"
	"fmt"
	"errors"
)

type Baseblob struct {
	data	[]byte
}

func DataReverse(dataIn []byte) []byte {
	var dataRet []byte
	for i := len(dataIn)-1; i >= 0; i-- {
		dataRet = append(dataRet, dataIn[i])
	}
	return dataRet
}

func HexCharToNumber(charIn byte) (int8, error) {
	if charIn >= 0x30 && charIn <= 0x39 {
		return int8(charIn - 0x30), nil
	} else if charIn >= 0x41 && charIn <= 0x46 {
		return int8(charIn - 0x37), nil
	} else if charIn >= 0x61 && charIn <= 0x66 {
		return int8(charIn - 0x57), nil
	}
	return 0, errors.New("invalid hex char")
}

func (b Baseblob) isValidHex(hexStr string) bool {
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

func (b *Baseblob) SetHex(hexStr string) {
	if hexStr[0] == '0' && hexStr[1] == 'x' {
		hexStr = hexStr[2:]
	}

	Assert(b.isValidHex(hexStr), "invalid hex string")
	blobLength := len(hexStr) / 2
	for i := blobLength-1; i >= 0; i-- {
		num1, _ := HexCharToNumber(hexStr[2*i])
		num2, _ := HexCharToNumber(hexStr[2*i+1])
		b.data = append(b.data, byte((num1 << 4) | num2))
	}
}

func (b Baseblob) GetHex() string {
	var stringRet string
	dataRet := DataReverse(b.data)
	for _, c := range dataRet {
		stringRet += fmt.Sprintf("%02x", c)
	}
	return stringRet
}

func (b Baseblob) GetData() []byte {
	return b.data
}

func (b Baseblob) GetDataSize() int {
	return len(b.data)
}
