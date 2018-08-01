package bigint

import (
	"bytes"
	"errors"
	"github.com/mutalisk999/bitcoin-lib/src/blob"
	"github.com/mutalisk999/bitcoin-lib/src/utility"
	"io"
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

func (u *Uint256) SetHex(hexStr string) error {
	if !u.isValidHex(hexStr) {
		return errors.New("invalid hex str")
	}
	return u.blob.SetHex(hexStr)
}

func (u *Uint256) SetData(bytes []byte) error {
	if len(bytes) != 32 {
		return errors.New("invalid bytes")
	}
	u.blob.SetData(bytes)
	return nil
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
	utility.Assert(u.GetDataSize() == 32, "Uint256::Pack : invalid data size")
	err := u.blob.Pack(writer, 32)
	if err != nil {
		return err
	}
	return nil
}

func (u *Uint256) UnPack(reader io.Reader) error {
	err := u.blob.UnPack(reader, 32)
	if err != nil {
		return err
	}
	if u.blob.GetDataSize() != 32 {
		return errors.New("Uint256::UnPack: invalid size of Uint256")
	}
	return nil
}

func IsUint256Equal(l *Uint256, r *Uint256) bool {
	dataLeft := l.GetData()
	dataRight := r.GetData()
	return bytes.Equal(dataLeft, dataRight)
}
