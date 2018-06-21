package pubkey

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCalcKeyIDBytes(t *testing.T) {
	pubKey := new(PubKey)
	pubKeyBytes, _ := hex.DecodeString("0394988ba55556eec01cb6e7fe8fb4e377908a121cbda759bd786101cc813a4907")
	pubKey.SetPubKeyData(pubKeyBytes)
	keyIDBytes, _ := pubKey.CalcKeyIDBytes()
	fmt.Println(hex.EncodeToString(keyIDBytes))
}
