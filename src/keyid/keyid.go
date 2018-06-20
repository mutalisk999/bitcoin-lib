package keyid

import (
	"errors"
	"utility"
	"base58"
	"bech32"
)

const (
	KEY_ID_SIZE = 0x14
)

type KeyID struct {
	data []byte
}

func ValidSize(keyIDData []byte) bool {
	return len(keyIDData) == KEY_ID_SIZE
}

func (k *KeyID) SetKeyIDData(keyIDBytes []byte) error {
	isValid := ValidSize(keyIDBytes)
	if !isValid {
		return errors.New("KeyID::SetKeyIDData: invalid keyid size")
	}
	k.data = keyIDBytes
	return nil
}

func (k KeyID) GetKeyIDData() []byte {
	return k.data
}

func (k KeyID) ToBase58Address(version byte) string {
	payLoad := []byte{version}
	payLoad = []byte(string(payLoad) + string(k.GetKeyIDData()))
	checkSum := utility.Sha256(utility.Sha256(payLoad))[0:4]
	payLoad = []byte(string(payLoad) + string(checkSum))
	return base58.Encode(payLoad)
}

func (k KeyID) ToBech32AddressP2WPKH(hrp string) (string, error) {
	payLoad := []byte{0x0, KEY_ID_SIZE}
	payLoad = []byte(string(payLoad) + string(k.GetKeyIDData()))
	return bech32.SegWitAddressEncode(hrp, payLoad)
}