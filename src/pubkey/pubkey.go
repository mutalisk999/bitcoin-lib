package pubkey

import (
	"errors"
	"io"
	"serialize"
	"utility"
)

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

func (p PubKey) GetPubKeyData() ([]byte, error) {
	isValid := ValidSize(p.data)
	if !isValid {
		return []byte{}, errors.New("PubKey::GetPubKeyData : invalid pubkey size")
	}
	return p.data, nil
}

func (p PubKey) Pack(writer io.Writer) error {
	err := serialize.PackCompactSize(writer, uint64(len(p.data)))
	if err != nil {
		return err
	}
	_, err = writer.Write(p.data[0:len(p.data)])
	if err != nil {
		return err
	}
	return nil
}

func (p *PubKey) UnPack(reader io.Reader) error {
	u64, err := serialize.UnPackCompactSize(reader)
	if err != nil {
		return err
	}
	dataRead := make([]byte, u64)
	_, err = reader.Read(dataRead[0:u64])
	if err != nil {
		return err
	}
	for _, c := range dataRead {
		p.data = append(p.data, c)
	}
	return nil
}

func (p *PubKey) CalcKeyIDBytes() ([]byte, error) {
	isValid := ValidSize(p.data)
	if !isValid {
		return []byte{}, errors.New("PubKey::CaclKeyIDBytes : invalid pubkey size")
	}
	return utility.Hash160(p.data), nil
}
