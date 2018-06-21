package utility

import (
	"crypto/sha256"
	"encoding/hex"
	"ripemd160"
)

func Sha256(bytes []byte) []byte {
	hash := sha256.New()
	hash.Write(bytes)
	return hash.Sum(nil)
}

func Sha256Hex(bytes []byte) string {
	return hex.EncodeToString(Sha256(bytes))
}

func Ripemd160(bytes []byte) []byte {
	hash := ripemd160.New()
	hash.Write(bytes)
	return hash.Sum(nil)
}

func Ripemd160Hex(bytes []byte) string {
	return hex.EncodeToString(Ripemd160(bytes))
}

func Hash160(bytes []byte) []byte {
	return Ripemd160(Sha256(bytes))
}

func Hash160Hex(bytes []byte) string {
	return hex.EncodeToString(Hash160(bytes))
}
