package bigint

import (
	"blob"
	"io"
	"utility"
)

type Uint256 struct {
	blob blob.Baseblob
}

func (u Uint256) isValidHex(hexStr string) bool {
	if hexStr[0] == '0' && hexStr[1] == 'x' {
		hexStr = hexStr[2:]
	}

	if len(hexStr) != 64 {
		return false
	}

	return true
}

func (u *Uint256) SetHex(hexStr string) {
	utility.Assert(u.isValidHex(hexStr), "invalid hex str")
	u.blob.SetHex(hexStr)
}

func (u Uint256) GetHex() string {
	return u.blob.GetHex()
}

func (u Uint256) GetData() []byte {
	return u.blob.GetData()
}

func (u Uint256) GetDataSize() int {
	return u.blob.GetDataSize()
}

func (u Uint256) Pack(writer io.Writer) error {
	err := u.blob.Pack(writer)
	if err != nil {
		return err
	}
	return nil
}

func (u *Uint256) UnPack(reader io.Reader) error {
	err := u.blob.UnPack(reader)
	if err != nil {
		return err
	}
	utility.Assert(u.blob.GetDataSize() == 32, "Uint256::UnPack: invalid size of Uint256")
	return nil
}
