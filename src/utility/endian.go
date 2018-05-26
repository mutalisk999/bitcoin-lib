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
	} else if *lowUint8ptr == 0x02 && *highUint8ptr== 0x01 {
		isLittleEndian = true
	} else {
		panic("IsLittleEndian() can not judge if platform is little endian or not")
	}

	isEndianInit = true
	return isLittleEndian
}

