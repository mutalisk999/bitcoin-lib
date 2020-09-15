package pubkey

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCalcKeyIDBytes(t *testing.T) {
	pubKey := new(PubKey)
	pubKeyBytes, _ := hex.DecodeString("0394988ba55556eec01cb6e7fe8fb4e377908a121cbda759bd786101cc813a4907")
	_ = pubKey.SetPubKeyData(pubKeyBytes)
	keyIDBytes, _ := pubKey.CalcKeyIDBytes()
	fmt.Println(hex.EncodeToString(keyIDBytes))
}

func TestGetUnCompressPubKey(t *testing.T) {
	pubKeyBytes, _ := hex.DecodeString("1bea5be697868ee1edb75e22df72aca99ed0de01e141f41996e50fec89dad96030d8ccad50d875a84daefc0b03856a4ce10b571d8609631f3378fbe0daa251ae")
	pubKey, _ := GetUnCompressPubKey(pubKeyBytes)
	pubKeyUnCompressBytes, _ := pubKey.GetPubKeyData()
	fmt.Println(hex.EncodeToString(pubKeyUnCompressBytes))
}

func TestGetCompressPubKey(t *testing.T) {
	pubKeyBytes, _ := hex.DecodeString("1bea5be697868ee1edb75e22df72aca99ed0de01e141f41996e50fec89dad96030d8ccad50d875a84daefc0b03856a4ce10b571d8609631f3378fbe0daa251ae")
	pubKey, _ := GetCompressPubKey(pubKeyBytes)
	pubKeyCompressBytes, _ := pubKey.GetPubKeyData()
	fmt.Println(hex.EncodeToString(pubKeyCompressBytes))
}
