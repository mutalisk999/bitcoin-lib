package bigint

import "utility"

type Uint256 struct {
	blob utility.Baseblob
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

func (u* Uint256) SetHex(hexStr string) {
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