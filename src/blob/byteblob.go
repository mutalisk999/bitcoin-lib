package blob

import (
	"errors"
	"github.com/mutalisk999/bitcoin-lib/src/serialize"
	"github.com/mutalisk999/bitcoin-lib/src/utility"
	"io"
)

type Byteblob struct {
	data []byte
}

func (b *Byteblob) SetData(bytes []byte) {
	b.data = bytes
}

func (b *Byteblob) SetHex(hexStr string) error {
	if hexStr[0] == '0' && hexStr[1] == 'x' {
		hexStr = hexStr[2:]
	}
	if !utility.IsValidHex(hexStr) {
		return errors.New("invalid hex string")
	}
	blobLength := len(hexStr) / 2
	b.data = make([]byte, blobLength, blobLength)
	for i := 0; i < blobLength; i++ {
		num1, _ := utility.HexCharToNumber(hexStr[2*i])
		num2, _ := utility.HexCharToNumber(hexStr[2*i+1])
		//b.data = append(b.data, byte((num1<<4)|num2))
		b.data[i] = byte((num1 << 4) | num2)
	}
	return nil
}

func (b Byteblob) GetHex() string {
	var bytes []byte
	bytes = make([]byte, 2*len(b.data), 2*len(b.data))
	for i := 0; i < len(b.data); i++ {
		var h4b byte
		var l4b byte
		h4b, _ = utility.NumberToHexChar((b.data[i] & 0xf0) >> 4)
		l4b, _ = utility.NumberToHexChar(b.data[i] & 0x0f)
		bytes[2*i], bytes[2*i+1] = h4b, l4b
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
	dataRead := make([]byte, u64, u64)
	_, err = reader.Read(dataRead[0:u64])
	if err != nil {
		return err
	}
	b.data = dataRead
	return nil
}
