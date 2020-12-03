package keyid

import (
	"errors"
	"github.com/mutalisk999/bitcoin-lib/src/base58"
	"github.com/mutalisk999/bitcoin-lib/src/bech32"
	"github.com/mutalisk999/bitcoin-lib/src/serialize"
	"github.com/mutalisk999/bitcoin-lib/src/utility"
	"io"
)

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

func (k KeyID) GetKeyIDData() ([]byte, error) {
	isValid := ValidSize(k.data)
	if !isValid {
		return []byte{}, errors.New("KeyID::GetKeyIDData: invalid keyid size")
	}
	return k.data, nil
}

func (k KeyID) Pack(writer io.Writer) error {
	err := serialize.PackCompactSize(writer, uint64(len(k.data)))
	if err != nil {
		return err
	}
	_, err = writer.Write(k.data[0:len(k.data)])
	if err != nil {
		return err
	}
	return nil
}

func (k *KeyID) UnPack(reader io.Reader) error {
	u64, err := serialize.UnPackCompactSize(reader)
	if err != nil {
		return err
	}
	dataRead := make([]byte, u64, u64)
	_, err = reader.Read(dataRead[0:u64])
	if err != nil {
		return err
	}
	k.data = dataRead
	return nil
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
