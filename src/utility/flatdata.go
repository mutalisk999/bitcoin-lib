package utility

import "unsafe"

func DumpUint16ToBytes(data16 uint16) []byte {
	uint16ptr := unsafe.Pointer(&data16)
	byte8 := *(*byte)(unsafe.Pointer(uint16ptr))
	byte16 := *(*byte)(unsafe.Pointer(uintptr(uint16ptr) + uintptr(1)))
	return []byte{byte8, byte16}
}

func DumpUint32ToBytes(data32 uint32) []byte {
	uint32ptr := unsafe.Pointer(&data32)
	byte8 := *(*byte)(unsafe.Pointer(uint32ptr))
	byte16 := *(*byte)(unsafe.Pointer(uintptr(uint32ptr) + uintptr(1)))
	byte24 := *(*byte)(unsafe.Pointer(uintptr(uint32ptr) + uintptr(2)))
	byte32 := *(*byte)(unsafe.Pointer(uintptr(uint32ptr) + uintptr(3)))
	return []byte{byte8, byte16, byte24, byte32}
}

func DumpUint64ToBytes(data64 uint64) []byte {
	uint64ptr := unsafe.Pointer(&data64)
	byte8 := *(*byte)(unsafe.Pointer(uint64ptr))
	byte16 := *(*byte)(unsafe.Pointer(uintptr(uint64ptr) + uintptr(1)))
	byte24 := *(*byte)(unsafe.Pointer(uintptr(uint64ptr) + uintptr(2)))
	byte32 := *(*byte)(unsafe.Pointer(uintptr(uint64ptr) + uintptr(3)))
	byte40 := *(*byte)(unsafe.Pointer(uintptr(uint64ptr) + uintptr(4)))
	byte48 := *(*byte)(unsafe.Pointer(uintptr(uint64ptr) + uintptr(5)))
	byte56 := *(*byte)(unsafe.Pointer(uintptr(uint64ptr) + uintptr(6)))
	byte64 := *(*byte)(unsafe.Pointer(uintptr(uint64ptr) + uintptr(7)))
	return []byte{byte8, byte16, byte24, byte32, byte40, byte48, byte56, byte64}
}

func DumpFloat32ToBytes(data32 float32) []byte {
	float32ptr := unsafe.Pointer(&data32)
	byte8 := *(*byte)(unsafe.Pointer(float32ptr))
	byte16 := *(*byte)(unsafe.Pointer(uintptr(float32ptr) + uintptr(1)))
	byte24 := *(*byte)(unsafe.Pointer(uintptr(float32ptr) + uintptr(2)))
	byte32 := *(*byte)(unsafe.Pointer(uintptr(float32ptr) + uintptr(3)))
	return []byte{byte8, byte16, byte24, byte32}
}

func DumpFloat64ToBytes(data64 float64) []byte {
	float64ptr := unsafe.Pointer(&data64)
	byte8 := *(*byte)(unsafe.Pointer(float64ptr))
	byte16 := *(*byte)(unsafe.Pointer(uintptr(float64ptr) + uintptr(1)))
	byte24 := *(*byte)(unsafe.Pointer(uintptr(float64ptr) + uintptr(2)))
	byte32 := *(*byte)(unsafe.Pointer(uintptr(float64ptr) + uintptr(3)))
	byte40 := *(*byte)(unsafe.Pointer(uintptr(float64ptr) + uintptr(4)))
	byte48 := *(*byte)(unsafe.Pointer(uintptr(float64ptr) + uintptr(5)))
	byte56 := *(*byte)(unsafe.Pointer(uintptr(float64ptr) + uintptr(6)))
	byte64 := *(*byte)(unsafe.Pointer(uintptr(float64ptr) + uintptr(7)))
	return []byte{byte8, byte16, byte24, byte32, byte40, byte48, byte56, byte64}
}

func LoadUint16FromBytes(bytes []byte) uint16 {
	Assert(len(bytes) == 2, "incorrect bytes length, not 2")
	byte8 := uint16(bytes[0])
	byte16 := uint16(bytes[1])
	return uint16((byte16 << 8) | byte8)
}

func LoadUint32FromBytes(bytes []byte) uint32 {
	Assert(len(bytes) == 4, "incorrect bytes length, not 4")
	byte8 := uint32(bytes[0])
	byte16 := uint32(bytes[1])
	byte24 := uint32(bytes[2])
	byte32 := uint32(bytes[3])
	return uint32((byte32 << 24) | (byte24 << 16) | (byte16 << 8) | byte8)
}

func LoadUint64FromBytes(bytes []byte) uint64 {
	Assert(len(bytes) == 8, "incorrect bytes length, not 8")
	byte8 := uint64(bytes[0])
	byte16 := uint64(bytes[1])
	byte24 := uint64(bytes[2])
	byte32 := uint64(bytes[3])
	byte40 := uint64(bytes[4])
	byte48 := uint64(bytes[5])
	byte56 := uint64(bytes[6])
	byte64 := uint64(bytes[7])
	return uint64((byte64 << 56) | (byte56 << 48) | (byte48 << 40) | (byte40 << 32) |
		(byte32 << 24) | (byte24 << 16) | (byte16 << 8) | byte8)
}

func LoadFloat32FromBytes(bytes []byte) float32 {
	Assert(len(bytes) == 4, "incorrect bytes length, not 4")
	var data32 float32
	data32ptr := unsafe.Pointer(&data32)
	*(*byte)(unsafe.Pointer(data32ptr)) = bytes[0]
	*(*byte)(unsafe.Pointer(uintptr(data32ptr) + uintptr(1))) = bytes[1]
	*(*byte)(unsafe.Pointer(uintptr(data32ptr) + uintptr(2))) = bytes[2]
	*(*byte)(unsafe.Pointer(uintptr(data32ptr) + uintptr(3))) = bytes[3]
	return data32
}

func LoadFloat64FromBytes(bytes []byte) float64 {
	Assert(len(bytes) == 8, "incorrect bytes length, not 8")
	var data64 float64
	data64ptr := unsafe.Pointer(&data64)
	*(*byte)(unsafe.Pointer(data64ptr)) = bytes[0]
	*(*byte)(unsafe.Pointer(uintptr(data64ptr) + uintptr(1))) = bytes[1]
	*(*byte)(unsafe.Pointer(uintptr(data64ptr) + uintptr(2))) = bytes[2]
	*(*byte)(unsafe.Pointer(uintptr(data64ptr) + uintptr(3))) = bytes[3]
	*(*byte)(unsafe.Pointer(uintptr(data64ptr) + uintptr(4))) = bytes[4]
	*(*byte)(unsafe.Pointer(uintptr(data64ptr) + uintptr(5))) = bytes[5]
	*(*byte)(unsafe.Pointer(uintptr(data64ptr) + uintptr(6))) = bytes[6]
	*(*byte)(unsafe.Pointer(uintptr(data64ptr) + uintptr(7))) = bytes[7]
	return data64
}
