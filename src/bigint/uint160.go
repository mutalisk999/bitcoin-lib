package bigint

import (
	"blob"
	"io"
	"utility"
)

type Uint160 struct {
	blob blob.Baseblob
}

func (u Uint160) isValidHex(hexStr string) bool {
	if hexStr[0] == '0' && hexStr[1] == 'x' {
		hexStr = hexStr[2:]
	}

	if len(hexStr) != 40 {
		return false
	}

	return true
}

func (u *Uint160) SetHex(hexStr string) {
	utility.Assert(u.isValidHex(hexStr), "invalid hex str")
	u.blob.SetHex(hexStr)
}

func (u Uint160) GetHex() string {
	return u.blob.GetHex()
}

func (u Uint160) GetData() []byte {
	return u.blob.GetData()
}

func (u Uint160) GetDataSize() int {
	return u.blob.GetDataSize()
}

func (b Uint160) Pack(writer io.Writer) error {
	err := b.blob.Pack(writer)
	if err != nil {
		return err
	}
	return nil
}

func (b *Uint160) UnPack(reader io.Reader) error {
	err := b.blob.UnPack(reader)
	if err != nil {
		return err
	}
	utility.Assert(b.blob.GetDataSize() == 20, "Uint160::UnPack: invalid size of Uint160")
	return nil
}
