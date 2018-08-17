package utility

import (
	"errors"
	"strings"
)

func HexCharToNumber(charIn byte) (int8, error) {
	if charIn >= 0x30 && charIn <= 0x39 {
		return int8(charIn - 0x30), nil
	} else if charIn >= 0x41 && charIn <= 0x46 {
		return int8(charIn - 0x37), nil
	} else if charIn >= 0x61 && charIn <= 0x66 {
		return int8(charIn - 0x57), nil
	}
	return 0, errors.New("invalid hex char")
}

func NumberToHexChar(number uint8) (byte, error) {
	if number <= 9 {
		return byte(number + 0x30), nil
	} else if number >= 10 && number <= 15 {
		return byte(number + 0x57), nil
	}
	return 0, errors.New("invalid number")
}

func IsValidHex(hexStr string) bool {
	if hexStr[0] == '0' && hexStr[1] == 'x' {
		hexStr = hexStr[2:]
	}
	if len(hexStr)%2 == 1 {
		return false
	}
	hexStr = strings.ToLower(hexStr)
	for _, c := range []byte(hexStr) {
		if !((c >= 0x30 && c <= 0x39) || (c >= 0x61 && c <= 0x66)) {
			return false
		}
	}
	return true
}
