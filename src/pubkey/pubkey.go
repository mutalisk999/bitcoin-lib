package pubkey

import "errors"

const (
	PUBLIC_KEY_SIZE            = 65
	COMPRESSED_PUBLIC_KEY_SIZE = 33
	SIGNATURE_SIZE             = 72
	COMPACT_SIGNATURE_SIZE     = 65
)

type PubKey struct {
	data []byte
}

func GetLen(c byte) int {
	if c == 0x2 || c == 0x3 {
		return COMPRESSED_PUBLIC_KEY_SIZE
	} else if c == 0x4 || c == 0x6 || c == 0x7 {
		return PUBLIC_KEY_SIZE
	}
	return 0
}

func ValidSize(pubKeyData []byte) bool {
	return len(pubKeyData) > 0 && GetLen(pubKeyData[0]) == len(pubKeyData)
}

func (p *PubKey) SetPubKeyData(pubKeyBytes []byte) error {
	isValid := ValidSize(pubKeyBytes)
	if !isValid {
		return errors.New("PubKey::SetPubKeyData : invalid pubkey size")
	}
	p.data = pubKeyBytes
	return nil
}

func (p PubKey) GetPubKeyData() []byte {
	return p.data
}
