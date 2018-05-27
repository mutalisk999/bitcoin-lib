package serialize

import "io"
import (
	"utility"
	"unsafe"
)

func packU8(writer io.Writer, data8 uint8) error {
	_, err := writer.Write([]byte{data8})
	if err != nil {
		return err
	}
	return nil
}

func packU16(writer io.Writer, data16 uint16) error {
	bytes := utility.DumpUint16ToBytes(utility.ConvertUint16ToLittleEndian(data16))
	utility.Assert(len(bytes)==2, "incorrect bytes length, not 2")
	_, err := writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func packU32(writer io.Writer, data32 uint32) error {
	bytes := utility.DumpUint32ToBytes(utility.ConvertUint32ToLittleEndian(data32))
	utility.Assert(len(bytes)==4, "incorrect bytes length, not 4")
	_, err := writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func packU64(writer io.Writer, data64 uint64) error {
	bytes := utility.DumpUint64ToBytes(utility.ConvertUint64ToLittleEndian(data64))
	utility.Assert(len(bytes)==8, "incorrect bytes length, not 8")
	_, err := writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func packF32(writer io.Writer, data32 float32) error {
	data32ptr := uintptr(unsafe.Pointer(&data32))
	byte8 := *(*byte)(unsafe.Pointer(data32ptr))
	byte16 := *(*byte)(unsafe.Pointer(data32ptr + uintptr(1)))
	byte24 := *(*byte)(unsafe.Pointer(data32ptr + uintptr(2)))
	byte32 := *(*byte)(unsafe.Pointer(data32ptr + uintptr(3)))
	_, err := writer.Write([]byte{byte8, byte16, byte24, byte32})
	if err != nil {
		return err
	}
	return nil
}

func packF64(writer io.Writer, data64 float64) error {
	data64ptr := uintptr(unsafe.Pointer(&data64))
	byte8 := *(*byte)(unsafe.Pointer(data64ptr))
	byte16 := *(*byte)(unsafe.Pointer(data64ptr + uintptr(1)))
	byte24 := *(*byte)(unsafe.Pointer(data64ptr + uintptr(2)))
	byte32 := *(*byte)(unsafe.Pointer(data64ptr + uintptr(3)))
	byte40 := *(*byte)(unsafe.Pointer(data64ptr + uintptr(4)))
	byte48 := *(*byte)(unsafe.Pointer(data64ptr + uintptr(5)))
	byte56 := *(*byte)(unsafe.Pointer(data64ptr + uintptr(6)))
	byte64 := *(*byte)(unsafe.Pointer(data64ptr + uintptr(7)))
	_, err := writer.Write([]byte{byte8, byte16, byte24, byte32, byte40, byte48, byte56, byte64})
	if err != nil {
		return err
	}
	return nil
}

func unpackU8(reader io.Reader) (uint8, error) {
	var bytes [1]byte
	_, err := reader.Read(bytes[0:1])
	if err != nil {
		return 0, err
	}
	return uint8(bytes[0]), nil
}

func unpackU16(reader io.Reader) (uint16, error) {
	var bytes [2]byte
	_, err := reader.Read(bytes[0:2])
	if err != nil {
		return 0, err
	}
	u16 := utility.ConvertUint16FromLittleEndian(utility.LoadUint16FromBytes(bytes[0:2]))
	return u16, nil
}

func unpackU32(reader io.Reader) (uint32, error) {
	var bytes [4]byte
	_, err := reader.Read(bytes[0:4])
	if err != nil {
		return 0, err
	}
	u32 := utility.ConvertUint32FromLittleEndian(utility.LoadUint32FromBytes(bytes[0:4]))
	return u32, nil
}

func unpackU64(reader io.Reader) (uint64, error) {
	var bytes [8]byte
	_, err := reader.Read(bytes[0:8])
	if err != nil {
		return 0, err
	}
	u64 := utility.ConvertUint64FromLittleEndian(utility.LoadUint64FromBytes(bytes[0:8]))
	return u64, nil
}

func unpackF32(reader io.Reader) (float32, error) {
	var bytes [4]byte
	_, err := reader.Read(bytes[0:4])
	if err != nil {
		return 0, err
	}
	var f32 float32
	byte32ptr := uintptr(unsafe.Pointer(&f32))
	*(*byte)(unsafe.Pointer(byte32ptr)) = bytes[0]
	*(*byte)(unsafe.Pointer(byte32ptr + uintptr(1))) = bytes[1]
	*(*byte)(unsafe.Pointer(byte32ptr + uintptr(2))) = bytes[2]
	*(*byte)(unsafe.Pointer(byte32ptr + uintptr(3))) = bytes[3]
	return f32, nil
}

func unpackF64(reader io.Reader) (float64, error) {
	var bytes [8]byte
	_, err := reader.Read(bytes[0:8])
	if err != nil {
		return 0, err
	}
	var f64 float64
	byte64ptr := uintptr(unsafe.Pointer(&f64))
	*(*byte)(unsafe.Pointer(byte64ptr)) = bytes[0]
	*(*byte)(unsafe.Pointer(byte64ptr + uintptr(1))) = bytes[1]
	*(*byte)(unsafe.Pointer(byte64ptr + uintptr(2))) = bytes[2]
	*(*byte)(unsafe.Pointer(byte64ptr + uintptr(3))) = bytes[3]
	*(*byte)(unsafe.Pointer(byte64ptr + uintptr(4))) = bytes[4]
	*(*byte)(unsafe.Pointer(byte64ptr + uintptr(5))) = bytes[5]
	*(*byte)(unsafe.Pointer(byte64ptr + uintptr(6))) = bytes[6]
	*(*byte)(unsafe.Pointer(byte64ptr + uintptr(7))) = bytes[7]
	return f64, nil
}

func PackByte(writer io.Writer, b byte) error {
	return packU8(writer, uint8(b))
}

func PackInt8(writer io.Writer, i8 int8) error {
	return packU8(writer, uint8(i8))
}

func PackUint8(writer io.Writer, ui8 uint8) error {
	return packU8(writer, uint8(ui8))
}

func PackInt16(writer io.Writer, i16 int16) error {
	return packU16(writer, uint16(i16))
}

func PackUint16(writer io.Writer, ui16 uint16) error {
	return packU16(writer, uint16(ui16))
}

func PackInt32(writer io.Writer, i32 int32) error {
	return packU32(writer, uint32(i32))
}

func PackUint32(writer io.Writer, ui32 uint32) error {
	return packU32(writer, uint32(ui32))
}

func PackInt64(writer io.Writer, i64 int64) error {
	return packU64(writer, uint64(i64))
}

func PackUint64(writer io.Writer, ui64 uint64) error {
	return packU64(writer, uint64(ui64))
}

func PackFloat32(writer io.Writer, f32 float32) error {
	return packF32(writer, f32)
}

func PackFloat64(writer io.Writer, f64 float64) error {
	return packF64(writer, f64)
}

func UnPackByte(reader io.Reader) (byte, error) {
	ui8, err := unpackU8(reader)
	return byte(ui8), err
}

func UnPackInt8(reader io.Reader) (int8, error) {
	ui8, err := unpackU8(reader)
	return int8(ui8), err
}

func UnPackUint8(reader io.Reader) (uint8, error) {
	return unpackU8(reader)
}

func UnPackInt16(reader io.Reader) (int16, error) {
	ui16, err := unpackU16(reader)
	return int16(ui16), err
}

func UnPackUint16(reader io.Reader) (uint16, error) {
	return unpackU16(reader)
}

func UnPackInt32(reader io.Reader) (int32, error) {
	ui32, err := unpackU32(reader)
	return int32(ui32), err
}

func UnPackUint32(reader io.Reader) (uint32, error) {
	return unpackU32(reader)
}

func UnPackInt64(reader io.Reader) (int64, error) {
	ui64, err := unpackU64(reader)
	return int64(ui64), err
}

func UnPackUint64(reader io.Reader) (uint64, error) {
	return unpackU64(reader)
}

func UnPackFloat32(reader io.Reader) (float32, error) {
	return unpackF32(reader)
}

func UnPackFloat64(reader io.Reader) (float64, error) {
	return unpackF64(reader)
}
