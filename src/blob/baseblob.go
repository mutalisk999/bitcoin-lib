package blob

import (
	"fmt"
	"io"
	"serialize"
	"strings"
	"utility"
)

type Baseblob struct {
	data []byte
}

func DataReverse(dataIn []byte) []byte {
	var dataRet []byte
	for i := len(dataIn) - 1; i >= 0; i-- {
		dataRet = append(dataRet, dataIn[i])
	}
	return dataRet
}

func (b Baseblob) isValidHex(hexStr string) bool {
	if hexStr[0] == '0' && hexStr[1] == 'x' {
		hexStr = hexStr[2:]
	}

	if len(hexStr)%2 == 1 {
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

	utility.Assert(b.isValidHex(hexStr), "invalid hex string")
	blobLength := len(hexStr) / 2
	for i := blobLength - 1; i >= 0; i-- {
		num1, _ := utility.HexCharToNumber(hexStr[2*i])
		num2, _ := utility.HexCharToNumber(hexStr[2*i+1])
		b.data = append(b.data, byte((num1<<4)|num2))
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

func (b Baseblob) Pack(writer io.Writer) error {
	err := serialize.PackCompactSize(writer, uint64(len(b.data)))
	if err != nil {
		return err
	}
	_, err = writer.Write(b.data[0:len(b.data)])
	if err != nil {
		return err
	}
	return nil
}

func (b *Baseblob) UnPack(reader io.Reader) error {
	u64, err := serialize.UnPackCompactSize(reader)
	if err != nil {
		return err
	}
	dataRead := make([]byte, u64)
	_, err = reader.Read(dataRead[0:u64])
	if err != nil {
		return err
	}
	for _, c := range dataRead {
		b.data = append(b.data, c)
	}
	return nil
}
