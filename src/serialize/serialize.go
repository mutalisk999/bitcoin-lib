package serialize

import "io"
import (
	"utility"
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
	utility.Assert(len(bytes) == 2, "incorrect bytes length, not 2")
	_, err := writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func packU32(writer io.Writer, data32 uint32) error {
	bytes := utility.DumpUint32ToBytes(utility.ConvertUint32ToLittleEndian(data32))
	utility.Assert(len(bytes) == 4, "incorrect bytes length, not 4")
	_, err := writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func packU64(writer io.Writer, data64 uint64) error {
	bytes := utility.DumpUint64ToBytes(utility.ConvertUint64ToLittleEndian(data64))
	utility.Assert(len(bytes) == 8, "incorrect bytes length, not 8")
	_, err := writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func packF32(writer io.Writer, data32 float32) error {
	bytes := utility.DumpFloat32ToBytes(data32)
	utility.Assert(len(bytes) == 4, "incorrect bytes length, not 4")
	_, err := writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func packF64(writer io.Writer, data64 float64) error {
	bytes := utility.DumpFloat64ToBytes(data64)
	utility.Assert(len(bytes) == 8, "incorrect bytes length, not 8")
	_, err := writer.Write(bytes)
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
	f32 := utility.LoadFloat32FromBytes(bytes[0:4])
	return f32, nil
}

func unpackF64(reader io.Reader) (float64, error) {
	var bytes [8]byte
	_, err := reader.Read(bytes[0:8])
	if err != nil {
		return 0, err
	}
	f64 := utility.LoadFloat64FromBytes(bytes[0:8])
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
