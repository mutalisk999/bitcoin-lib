package utility

import (
	"fmt"
	"testing"
)

func TestIsLittleEndian(t *testing.T) {
	b := IsLittleEndian()
	fmt.Println("is little endian:", b)
}

func TestIsBigEndian(t *testing.T) {
	b := IsBigEndian()
	fmt.Println("is big endian:", b)
}

func TestConvertUint16ToLittleEndian(t *testing.T) {
	u16 := uint16(0x0102)
	u16 = ConvertUint16ToLittleEndian(u16)
	fmt.Printf("uint16 to little endian: 0x%04x\n", u16)
}

func TestConvertUint16ToBigEndian(t *testing.T) {
	u16 := uint16(0x0102)
	u16 = ConvertUint16ToBigEndian(u16)
	fmt.Printf("uint16 to big endian: 0x%04x\n", u16)
}

func TestConvertUint32ToLittleEndian(t *testing.T) {
	u32 := uint32(0x01020304)
	u32 = ConvertUint32ToLittleEndian(u32)
	fmt.Printf("uint32 to little endian: 0x%08x\n", u32)
}

func TestConvertUint32ToBigEndian(t *testing.T) {
	u32 := uint32(0x01020304)
	u32 = ConvertUint32ToBigEndian(u32)
	fmt.Printf("uint32 to big endian: 0x%08x\n", u32)
}

func TestConvertUint64ToLittleEndian(t *testing.T) {
	u64 := uint64(0x0102030405060708)
	u64 = ConvertUint64ToLittleEndian(u64)
	fmt.Printf("uint64 to little endian: 0x%016x\n", u64)
}

func TestConvertUint64ToBigEndian(t *testing.T) {
	u64 := uint64(0x0102030405060708)
	u64 = ConvertUint64ToBigEndian(u64)
	fmt.Printf("uint64 to big endian: 0x%016x\n", u64)
}

func TestConvertUint16FromLittleEndian(t *testing.T) {
	u16 := uint16(0x0102)
	u16 = ConvertUint16FromLittleEndian(u16)
	fmt.Printf("uint16 from little endian: 0x%04x\n", u16)
}

func TestConvertUint16FromBigEndian(t *testing.T) {
	u16 := uint16(0x0102)
	u16 = ConvertUint16FromBigEndian(u16)
	fmt.Printf("uint16 from big endian: 0x%04x\n", u16)
}

func TestConvertUint32FromLittleEndian(t *testing.T) {
	u32 := uint32(0x01020304)
	u32 = ConvertUint32FromLittleEndian(u32)
	fmt.Printf("uint32 from little endian: 0x%08x\n", u32)
}

func TestConvertUint32FromBigEndian(t *testing.T) {
	u32 := uint32(0x01020304)
	u32 = ConvertUint32FromBigEndian(u32)
	fmt.Printf("uint32 from big endian: 0x%08x\n", u32)
}

func TestConvertUint64FromLittleEndian(t *testing.T) {
	u64 := uint64(0x0102030405060708)
	u64 = ConvertUint64FromLittleEndian(u64)
	fmt.Printf("uint64 from little endian: 0x%016x\n", u64)
}

func TestConvertUint64FromBigEndian(t *testing.T) {
	u64 := uint64(0x0102030405060708)
	u64 = ConvertUint64FromBigEndian(u64)
	fmt.Printf("uint64 from big endian: 0x%016x\n", u64)
}
