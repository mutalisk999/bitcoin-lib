package serialize

import "io"
import "utility"

func pack_data8(writer io.Writer, data8 uint8) error {
	_, err := writer.Write([]byte{data8})
	if err != nil {
		return err
	}
	return nil
}

func pack_data16(writer io.Writer, data16 uint16) error {
	bytes := utility.DumpUint16ToBytes(utility.ConvertUint16ToLittleEndian(data16))
	utility.Assert(len(bytes)==2, "incorrect bytes length, not 2")
	_, err := writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func pack_data32(writer io.Writer, data32 uint32) error {
	bytes := utility.DumpUint32ToBytes(utility.ConvertUint32ToLittleEndian(data32))
	utility.Assert(len(bytes)==4, "incorrect bytes length, not 4")
	_, err := writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func pack_data64(writer io.Writer, data64 uint64) error {
	bytes := utility.DumpUint64ToBytes(utility.ConvertUint64ToLittleEndian(data64))
	utility.Assert(len(bytes)==8, "incorrect bytes length, not 8")
	_, err := writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func unpack_data8(reader io.Reader) (uint8, error) {
	var bytes [1]byte
	_, err := reader.Read(bytes[0:1])
	if err != nil {
		return 0, err
	}
	return uint8(bytes[0]), nil
}

func unpack_data16(reader io.Reader) (uint16, error) {
	var bytes [2]byte
	_, err := reader.Read(bytes[0:2])
	if err != nil {
		return 0, err
	}
	u16 := utility.ConvertUint16FromLittleEndian(utility.LoadUint16FromBytes(bytes[0:2]))
	return u16, nil
}

func unpack_data32(reader io.Reader) (uint32, error) {
	var bytes [4]byte
	_, err := reader.Read(bytes[0:4])
	if err != nil {
		return 0, err
	}
	u32 := utility.ConvertUint32FromLittleEndian(utility.LoadUint32FromBytes(bytes[0:4]))
	return u32, nil
}

func unpack_data64(reader io.Reader) (uint64, error) {
	var bytes [8]byte
	_, err := reader.Read(bytes[0:8])
	if err != nil {
		return 0, err
	}
	u64 := utility.ConvertUint64FromLittleEndian(utility.LoadUint64FromBytes(bytes[0:8]))
	return u64, nil
}
