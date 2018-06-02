package script

import (
	"io"
	"serialize"
)

type Script struct {
	data []byte
}

func (s Script) Pack(writer io.Writer) error {
	err := serialize.PackCompactSize(writer, uint64(len(s.data)))
	if err != nil {
		return err
	}
	_, err = writer.Write(s.data[0:len(s.data)])
	if err != nil {
		return err
	}
	return nil
}

func (s* Script) UnPack(reader io.Reader) error {
	u64, err := serialize.UnPackCompactSize(reader)
	if err != nil {
		return err
	}
	dataRead := make([]byte,u64)
	_, err = reader.Read(dataRead[0:u64])
	if err != nil {
		return err
	}
	for _, c:= range dataRead {
		s.data = append(s.data, c)
	}
	return nil
}

func (s Script) GetScriptBytes() []byte {
	return s.data
}

