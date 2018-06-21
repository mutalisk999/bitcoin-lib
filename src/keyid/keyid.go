package keyid

import (
	"base58"
	"bech32"
	"errors"
	"utility"
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

func (k KeyID) GetKeyIDData() ([]byte, error) {
	isValid := ValidSize(k.data)
	if !isValid {
		return []byte{}, errors.New("KeyID::GetKeyIDData: invalid keyid size")
	}
	return k.data, nil
}

func (k KeyID) ToBase58Address(version byte) (string, error) {
	isValid := ValidSize(k.data)
	if !isValid {
		return "", errors.New("KeyID::ToBase58Address: invalid keyid size")
	}
	payLoad := []byte{version}
	payLoad = []byte(string(payLoad) + string(k.data))
	checkSum := utility.Sha256(utility.Sha256(payLoad))[0:4]
	payLoad = []byte(string(payLoad) + string(checkSum))
	return base58.Encode(payLoad), nil
}

func (k KeyID) ToBech32AddressP2WPKH(hrp string) (string, error) {
	isValid := ValidSize(k.data)
	if !isValid {
		return "", errors.New("KeyID::ToBech32AddressP2WPKH: invalid keyid size")
	}
	payLoad := []byte{0x0, KEY_ID_SIZE}
	payLoad = []byte(string(payLoad) + string(k.data))
	return bech32.SegWitAddressEncode(hrp, payLoad)
}

func ToBech32AddressP2WSH(hrp string, p2wshBytes []byte) (string, error) {
	if len(p2wshBytes) != 0x20 {
		return "", errors.New("ToBech32AddressP2WSH: invalid p2wshBytes size")
	}
	payLoad := []byte{0x0, 0x20}
	payLoad = []byte(string(payLoad) + string(p2wshBytes))
	return bech32.SegWitAddressEncode(hrp, payLoad)
}
