package script

import (
	"blob"
	"bytes"
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

func (s Script) PackToHex() (string, error) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	err := s.Pack(bufWriter)
	if err != nil {
		return "", err
	}
	Blob := new(blob.Byteblob)
	Blob.SetData(bytesBuf.Bytes())
	return Blob.GetHex(), nil
}

func (s *Script) UnPack(reader io.Reader) error {
	u64, err := serialize.UnPackCompactSize(reader)
	if err != nil {
		return err
	}
	dataRead := make([]byte, u64)
	_, err = reader.Read(dataRead[0:u64])
	if err != nil {
		return err
	}
	for _, c := range dataRead {
		s.data = append(s.data, c)
	}
	return nil
}

func (s *Script) UnPackFromHex(hexStr string) error {
	Blob := new(blob.Byteblob)
	err := Blob.SetHex(hexStr)
	if err != nil {
		return err
	}
	bytesBuf := bytes.NewBuffer(Blob.GetData())
	bufReader := io.Reader(bytesBuf)
	err = s.UnPack(bufReader)
	if err != nil {
		return err
	}
	return nil
}

func (s Script) GetScriptBytes() []byte {
	return s.data
}

type ScriptWitness struct {
	stack [][]byte
}

func (s ScriptWitness) Pack(writer io.Writer) error {
	err := serialize.PackCompactSize(writer, uint64(len(s.stack)))
	if err != nil {
		return err
	}
	for _, bytesSigWit := range s.stack {
		err := serialize.PackCompactSize(writer, uint64(len(bytesSigWit)))
		if err != nil {
			return err
		}
		_, err = writer.Write(bytesSigWit[0:])
		if err != nil {
			return err
		}
	}
	return nil
}

func (s ScriptWitness) PackToHex() (string, error) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	err := s.Pack(bufWriter)
	if err != nil {
		return "", err
	}
	Blob := new(blob.Byteblob)
	Blob.SetData(bytesBuf.Bytes())
	return Blob.GetHex(), nil
}

func (s *ScriptWitness) UnPack(reader io.Reader) error {
	stackLength, err := serialize.UnPackCompactSize(reader)
	if err != nil {
		return err
	}
	for i := 0; i < int(stackLength); i++ {
		u64, err := serialize.UnPackCompactSize(reader)
		if err != nil {
			return err
		}
		dataRead := make([]byte, u64)
		_, err = reader.Read(dataRead[0:u64])
		if err != nil {
			return err
		}
		s.stack = append(s.stack, dataRead)
	}
	return nil
}

func (s *ScriptWitness) UnPackFromHex(hexStr string) error {
	Blob := new(blob.Byteblob)
	err := Blob.SetHex(hexStr)
	if err != nil {
		return err
	}
	bytesBuf := bytes.NewBuffer(Blob.GetData())
	bufReader := io.Reader(bytesBuf)
	err = s.UnPack(bufReader)
	if err != nil {
		return err
	}
	return nil
}

func (s ScriptWitness) GetScriptWitnessBytes() [][]byte {
	return s.stack
}
