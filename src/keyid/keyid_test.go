package keyid

import (
	"testing"
	"fmt"
)

func TestKeyID_ToBase58Address(t *testing.T) {
	keyID := new(KeyID)
	keyID.SetKeyIDData([]byte{0xe4,0x26,0xe1,0xe6,0x44,0x7a,0x59,0x04,0x10,0xdf,
		0x6b,0x15,0x25,0xe9,0xcb,0x2c,0xed,0x0b,0xb3,0x8b})
	fmt.Println("base58 address:", keyID.ToBase58Address(0))
}

func TestKeyID_ToBech32AddressP2WPKH(t *testing.T) {
	keyID := new(KeyID)
	keyID.SetKeyIDData([]byte{0xed,0xec,0x29,0xb6,0x22,0x58,0x5a,0xcb,0xc8,0x2c,
		0x91,0x47,0x10,0x07,0x08,0x31,0xbf,0xb2,0xb6,0x48})
	addrStr, _ := keyID.ToBech32AddressP2WPKH("bc")
	fmt.Println("bech32 address:", addrStr)
}
