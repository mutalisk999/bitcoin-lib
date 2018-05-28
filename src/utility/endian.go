package utility

import "unsafe"

var isEndianInit = false
var isLittleEndian = false

func IsLittleEndian() bool {
	if isEndianInit {
		return isLittleEndian
	}

	var v uint16 = 0x0102
	uint16ptr := unsafe.Pointer(&v)
	lowUint8ptr := (*uint8)(unsafe.Pointer(uint16ptr))
	highUint8ptr := (*uint8)(unsafe.Pointer(uintptr(uint16ptr) + uintptr(1)))

	if *lowUint8ptr == 0x01 && *highUint8ptr == 0x02 {
		isLittleEndian = false
	} else if *lowUint8ptr == 0x02 && *highUint8ptr == 0x01 {
		isLittleEndian = true
	} else {
		panic("IsLittleEndian(): can not judge if platform is little endian or not")
	}

	isEndianInit = true
	return isLittleEndian
}

func IsBigEndian() bool {
	return !IsLittleEndian()
}

func ConvertUint16ToLittleEndian(data16 uint16) uint16 {
	if IsLittleEndian() {
		return data16
	} else {
		byte16 := (data16 & 0xFF00) >> 8
		byte8 := data16 & 0x00FF
		return uint16((byte8 << 8) | byte16)
	}
}

func ConvertUint16ToBigEndian(data16 uint16) uint16 {
	if IsBigEndian() {
		return data16
	} else {
		byte16 := (data16 & 0xFF00) >> 8
		byte8 := data16 & 0x00FF
		return uint16((byte8 << 8) | byte16)
	}
}

func ConvertUint32ToLittleEndian(data32 uint32) uint32 {
	if IsLittleEndian() {
		return data32
	} else {
		byte32 := (data32 & 0xFF000000) >> 24
		byte24 := (data32 & 0x00FF0000) >> 16
		byte16 := (data32 & 0x0000FF00) >> 8
		byte8 := data32 & 0x000000FF
		return uint32((byte8 << 24) | (byte16 << 16) | (byte24 << 8) | byte32)
	}
}

func ConvertUint32ToBigEndian(data32 uint32) uint32 {
	if IsBigEndian() {
		return data32
	} else {
		byte32 := (data32 & 0xFF000000) >> 24
		byte24 := (data32 & 0x00FF0000) >> 16
		byte16 := (data32 & 0x0000FF00) >> 8
		byte8 := data32 & 0x000000FF
		return uint32((byte8 << 24) | (byte16 << 16) | (byte24 << 8) | byte32)
	}
}

func ConvertUint64ToLittleEndian(data64 uint64) uint64 {
	if IsLittleEndian() {
		return data64
	} else {
		byte64 := (data64 & 0xFF00000000000000) >> 56
		byte56 := (data64 & 0x00FF000000000000) >> 48
		byte48 := (data64 & 0x0000FF0000000000) >> 40
		byte40 := (data64 & 0x000000FF00000000) >> 32
		byte32 := (data64 & 0x00000000FF000000) >> 24
		byte24 := (data64 & 0x0000000000FF0000) >> 16
		byte16 := (data64 & 0x000000000000FF00) >> 8
		byte8 := data64 & 0x00000000000000FF
		return uint64((byte8 << 56) | (byte16 << 48) | (byte24 << 40) | (byte32 << 32) |
			(byte40 << 24) | (byte48 << 16) | (byte56 << 8) | byte64)
	}
}

func ConvertUint64ToBigEndian(data64 uint64) uint64 {
	if IsBigEndian() {
		return data64
	} else {
		byte64 := (data64 & 0xFF00000000000000) >> 56
		byte56 := (data64 & 0x00FF000000000000) >> 48
		byte48 := (data64 & 0x0000FF0000000000) >> 40
		byte40 := (data64 & 0x000000FF00000000) >> 32
		byte32 := (data64 & 0x00000000FF000000) >> 24
		byte24 := (data64 & 0x0000000000FF0000) >> 16
		byte16 := (data64 & 0x000000000000FF00) >> 8
		byte8 := data64 & 0x00000000000000FF
		return uint64((byte8 << 56) | (byte16 << 48) | (byte24 << 40) | (byte32 << 32) |
			(byte40 << 24) | (byte48 << 16) | (byte56 << 8) | byte64)
	}
}

func ConvertUint16FromLittleEndian(data16 uint16) uint16 {
	if IsLittleEndian() {
		return data16
	} else {
		byte16 := (data16 & 0xFF00) >> 8
		byte8 := data16 & 0x00FF
		return uint16((byte8 << 8) | byte16)
	}
}

func ConvertUint16FromBigEndian(data16 uint16) uint16 {
	if IsBigEndian() {
		return data16
	} else {
		byte16 := (data16 & 0xFF00) >> 8
		byte8 := data16 & 0x00FF
		return uint16((byte8 << 8) | byte16)
	}
}

func ConvertUint32FromLittleEndian(data32 uint32) uint32 {
	if IsLittleEndian() {
		return data32
	} else {
		byte32 := (data32 & 0xFF000000) >> 24
		byte24 := (data32 & 0x00FF0000) >> 16
		byte16 := (data32 & 0x0000FF00) >> 8
		byte8 := data32 & 0x000000FF
		return uint32((byte8 << 24) | (byte16 << 16) | (byte24 << 8) | byte32)
	}
}

func ConvertUint32FromBigEndian(data32 uint32) uint32 {
	if IsBigEndian() {
		return data32
	} else {
		byte32 := (data32 & 0xFF000000) >> 24
		byte24 := (data32 & 0x00FF0000) >> 16
		byte16 := (data32 & 0x0000FF00) >> 8
		byte8 := data32 & 0x000000FF
		return uint32((byte8 << 24) | (byte16 << 16) | (byte24 << 8) | byte32)
	}
}

func ConvertUint64FromLittleEndian(data64 uint64) uint64 {
	if IsLittleEndian() {
		return data64
	} else {
		byte64 := (data64 & 0xFF00000000000000) >> 56
		byte56 := (data64 & 0x00FF000000000000) >> 48
		byte48 := (data64 & 0x0000FF0000000000) >> 40
		byte40 := (data64 & 0x000000FF00000000) >> 32
		byte32 := (data64 & 0x00000000FF000000) >> 24
		byte24 := (data64 & 0x0000000000FF0000) >> 16
		byte16 := (data64 & 0x000000000000FF00) >> 8
		byte8 := data64 & 0x00000000000000FF
		return uint64((byte8 << 56) | (byte16 << 48) | (byte24 << 40) | (byte32 << 32) |
			(byte40 << 24) | (byte48 << 16) | (byte56 << 8) | byte64)
	}
}

func ConvertUint64FromBigEndian(data64 uint64) uint64 {
	if IsBigEndian() {
		return data64
	} else {
		byte64 := (data64 & 0xFF00000000000000) >> 56
		byte56 := (data64 & 0x00FF000000000000) >> 48
		byte48 := (data64 & 0x0000FF0000000000) >> 40
		byte40 := (data64 & 0x000000FF00000000) >> 32
		byte32 := (data64 & 0x00000000FF000000) >> 24
		byte24 := (data64 & 0x0000000000FF0000) >> 16
		byte16 := (data64 & 0x000000000000FF00) >> 8
		byte8 := data64 & 0x00000000000000FF
		return uint64((byte8 << 56) | (byte16 << 48) | (byte24 << 40) | (byte32 << 32) |
			(byte40 << 24) | (byte48 << 16) | (byte56 << 8) | byte64)
	}
}
