package utility

import (
	"testing"
	"fmt"
)

func TestDumpUint16ToBytes(t *testing.T) {
	u16 := uint16(0x0102)
	bytes := DumpUint16ToBytes(u16)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("u16 bytes[%d]=%02x\n", i, bytes[i])
	}
}

func TestDumpUint32ToBytes(t *testing.T) {
	u32 := uint32(0x01020304)
	bytes := DumpUint32ToBytes(u32)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("u32 bytes[%d]=%02x\n", i, bytes[i])
	}
}

func TestDumpUint64ToBytes(t *testing.T) {
	u64 := uint64(0x0102030405060708)
	bytes := DumpUint64ToBytes(u64)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("u64 bytes[%d]=%02x\n", i, bytes[i])
	}
}

func TestLoadUint16FromBytes(t *testing.T) {
	bytes := []byte{0x02, 0x01}
	u16 := LoadUint16FromBytes(bytes)
	fmt.Printf("u16: 0x%04x\n", u16)
}

func TestLoadUint32FromBytes(t *testing.T) {
	bytes := []byte{0x04, 0x03, 0x02, 0x01}
	u32 := LoadUint32FromBytes(bytes)
	fmt.Printf("u32: 0x%08x\n", u32)
}

func TestLoadUint64FromBytes(t *testing.T) {
	bytes := []byte{0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01}
	u64 := LoadUint64FromBytes(bytes)
	fmt.Printf("u64: 0x%016x\n", u64)
}

