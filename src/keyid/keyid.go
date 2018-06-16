package keyid

import "errors"

const (
	KEY_ID_SIZE = 20
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