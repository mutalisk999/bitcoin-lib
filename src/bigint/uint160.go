package bigint

import (
	"blob"
	"errors"
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

func (u *Uint160) SetHex(hexStr string) error {
	if !u.isValidHex(hexStr) {
		return errors.New("invalid hex str")
	}
	return u.blob.SetHex(hexStr)
}

func (u *Uint160) SetData(bytes []byte) error {
	if len(bytes) != 20 {
		return errors.New("invalid bytes")
	}
	u.blob.SetData(bytes)
	return nil
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

func (u Uint160) Pack(writer io.Writer) error {
	utility.Assert(u.GetDataSize() == 20, "Uint160::Pack : invalid data size")
	err := u.blob.Pack(writer, 20)
	if err != nil {
		return err
	}
	return nil
}

func (u *Uint160) UnPack(reader io.Reader) error {
	err := u.blob.UnPack(reader, 20)
	if err != nil {
		return err
	}
	if u.blob.GetDataSize() != 20 {
		return errors.New("Uint160::UnPack: invalid size of Uint160")
	}
	return nil
}
