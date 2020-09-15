package serialize

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPackByte(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackByte(bufWriter, byte(0x01))
	fmt.Println("pack byte:", bytesBuf.Bytes())
}

func TestPackInt8(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackInt8(bufWriter, int8(-1))
	fmt.Println("pack int8:", bytesBuf.Bytes())
}

func TestPackUint8(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackUint8(bufWriter, uint8(0x01))
	fmt.Println("pack uint8:", bytesBuf.Bytes())
}

func TestPackInt16(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackInt16(bufWriter, int16(-1))
	fmt.Println("pack int16:", bytesBuf.Bytes())
}

func TestPackUint16(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackUint16(bufWriter, uint16(0x0102))
	fmt.Println("pack uint16:", bytesBuf.Bytes())
}

func TestPackInt32(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackInt32(bufWriter, int32(-1))
	fmt.Println("pack int32:", bytesBuf.Bytes())
}

func TestPackUint32(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackUint32(bufWriter, uint32(0x01020304))
	fmt.Println("pack uint32:", bytesBuf.Bytes())
}

func TestPackInt64(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackInt64(bufWriter, int64(-1))
	fmt.Println("pack int64:", bytesBuf.Bytes())
}

func TestPackUint64(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackUint64(bufWriter, uint64(0x0102030405060708))
	fmt.Println("pack uint64:", bytesBuf.Bytes())
}

func TestPackFloat32(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackFloat32(bufWriter, float32(1.234567))
	fmt.Println("pack float32:", bytesBuf.Bytes())
}

func TestPackFloat64(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackFloat64(bufWriter, float64(1.23456789012345))
	fmt.Println("pack float64:", bytesBuf.Bytes())
}

func TestCompactSizeLen(t *testing.T) {
	fmt.Println(CompactSizeLen(252))
	fmt.Println(CompactSizeLen((2 << 15) - 1))
	fmt.Println(CompactSizeLen((2 << 31) - 1))
	fmt.Println(CompactSizeLen((2 << 63) - 1))
}

func TestPackCompactSize(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = PackCompactSize(bufWriter, 0x1f)
	bufWriter = io.Writer(bytesBuf)
	_ = PackCompactSize(bufWriter, 0x1fff)
	bufWriter = io.Writer(bytesBuf)
	_ = PackCompactSize(bufWriter, 0x1fffffff)
	bufWriter = io.Writer(bytesBuf)
	_ = PackCompactSize(bufWriter, 0x1fffffffffffffff)
	fmt.Println("pack compact:", bytesBuf.Bytes())
}

func TestUnPackByte(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{1})
	bufReader := io.Reader(bytesBuf)
	b, _ := UnPackByte(bufReader)
	fmt.Printf("unpack byte: 0x%02x\n", b)
}

func TestUnPackInt8(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{255})
	bufReader := io.Reader(bytesBuf)
	i8, _ := UnPackInt8(bufReader)
	fmt.Println("unpack int8:", i8)
}

func TestUnPackUint8(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{1})
	bufReader := io.Reader(bytesBuf)
	ui8, _ := UnPackUint8(bufReader)
	fmt.Printf("unpack uint8: 0x%02x\n", ui8)
}

func TestUnPackInt16(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{255, 255})
	bufReader := io.Reader(bytesBuf)
	i16, _ := UnPackInt16(bufReader)
	fmt.Println("unpack int16:", i16)
}

func TestUnPackUint16(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{2, 1})
	bufReader := io.Reader(bytesBuf)
	ui16, _ := UnPackUint16(bufReader)
	fmt.Printf("unpack uint16: 0x%04x\n", ui16)
}

func TestUnPackInt32(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{255, 255, 255, 255})
	bufReader := io.Reader(bytesBuf)
	i32, _ := UnPackInt32(bufReader)
	fmt.Println("unpack int32:", i32)
}

func TestUnPackUint32(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{4, 3, 2, 1})
	bufReader := io.Reader(bytesBuf)
	ui32, _ := UnPackUint32(bufReader)
	fmt.Printf("unpack uint32: 0x%08x\n", ui32)
}

func TestUnPackInt64(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{255, 255, 255, 255, 255, 255, 255, 255})
	bufReader := io.Reader(bytesBuf)
	i64, _ := UnPackInt64(bufReader)
	fmt.Println("unpack int64:", i64)
}

func TestUnPackUint64(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{8, 7, 6, 5, 4, 3, 2, 1})
	bufReader := io.Reader(bytesBuf)
	ui64, _ := UnPackUint64(bufReader)
	fmt.Printf("unpack uint64: 0x%016x\n", ui64)
}

func TestUnPackFloat32(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{75, 6, 158, 63})
	bufReader := io.Reader(bytesBuf)
	f32, _ := UnPackFloat32(bufReader)
	fmt.Println("unpack float32:", f32)
}

func TestUnPackFloat64(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{221, 89, 140, 66, 202, 192, 243, 63})
	bufReader := io.Reader(bytesBuf)
	f64, _ := UnPackFloat64(bufReader)
	fmt.Println("unpack float64:", f64)
}

func TestUnPackCompactSize(t *testing.T) {
	bytesBuf := bytes.NewBuffer([]byte{31, 253, 255, 31, 254, 255, 255, 255, 31, 255,
		255, 255, 255, 255, 255, 255, 255, 31})
	bufReader := io.Reader(bytesBuf)
	ui8, _ := UnPackCompactSize(bufReader)
	ui16, _ := UnPackCompactSize(bufReader)
	ui32, _ := UnPackCompactSize(bufReader)
	ui64, _ := UnPackCompactSize(bufReader)
	fmt.Printf("unpack compact: %x %x %x %x\n", ui8, ui16, ui32, ui64)
}
