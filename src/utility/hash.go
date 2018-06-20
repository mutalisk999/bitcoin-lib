package utility

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(bytes []byte) []byte {
	hash := sha256.New()
	hash.Write(bytes)
	return hash.Sum(nil)
}

func Sha256Hex(bytes []byte) string {
	return hex.EncodeToString(Sha256(bytes))
}