package blob

import (
	"errors"
	"io"
	"serialize"
	"strings"
	"utility"
)

type Byteblob struct {
	data []byte
}

func (b Byteblob) isValidHex(hexStr string) bool {
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

func (b *Byteblob) SetData(bytes []byte) {
	b.data = bytes
}

func (b *Byteblob) SetHex(hexStr string) error {
	if hexStr[0] == '0' && hexStr[1] == 'x' {
		hexStr = hexStr[2:]
	}
	if !b.isValidHex(hexStr) {
		return errors.New("invalid hex string")
	}
	blobLength := len(hexStr) / 2
	for i := 0; i < blobLength; i++ {
		num1, _ := utility.HexCharToNumber(hexStr[2*i])
		num2, _ := utility.HexCharToNumber(hexStr[2*i+1])
		b.data = append(b.data, byte((num1<<4)|num2))
	}
	return nil
}

func (b Byteblob) GetHex() string {
	var bytes []byte
	for _, c := range b.data {
		var h4b byte
		var l4b byte
		h4b, _ = utility.NumberToHexChar((c & 0xf0) >> 4)
		l4b, _ = utility.NumberToHexChar(c & 0x0f)
		bytes = append(bytes, h4b, l4b)
	}
	return string(bytes)
}

func (b Byteblob) GetData() []byte {
	return b.data
}

func (b Byteblob) GetDataSize() int {
	return len(b.data)
}

func (b Byteblob) Pack(writer io.Writer) error {
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

func (b *Byteblob) UnPack(reader io.Reader) error {
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
