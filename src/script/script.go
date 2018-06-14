package script

import (
	"blob"
	"bytes"
	"io"
	"pubkey"
	"serialize"
	"utility"
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

func (s Script) GetScriptLength() int {
	return len(s.data)
}

func (s *Script) SetScriptBytes(scriptBytes []byte) {
	s.data = scriptBytes
}

func DecodeOPN(opCode byte) int {
	if opCode == OP_0 {
		return 0
	}
	utility.Assert(opCode >= OP_1 && opCode <= OP_16, "invalid opCode")
	return int(opCode - (OP_1 - 1))
}

func (s Script) IsPayToScriptHash() bool {
	return len(s.data) == 23 && s.data[0] == OP_HASH160 && s.data[1] == 0x14 && s.data[22] == OP_EQUAL
}

func (s Script) IsPayToPubKey() bool {
	return ((s.GetScriptLength() == (1+1+pubkey.PUBLIC_KEY_SIZE) && s.GetScriptBytes()[0] == pubkey.PUBLIC_KEY_SIZE) ||
		(s.GetScriptLength() == (1+1+pubkey.COMPRESSED_PUBLIC_KEY_SIZE) && s.GetScriptBytes()[0] == pubkey.COMPRESSED_PUBLIC_KEY_SIZE)) &&
		s.GetScriptBytes()[s.GetScriptLength()-1] == OP_CHECKSIG
}

func (s Script) IsPayToPubKeyHash() bool {
	return s.GetScriptLength() == 25 && s.GetScriptBytes()[0] == OP_DUP && s.GetScriptBytes()[1] == OP_HASH160 &&
		s.GetScriptBytes()[2] == 0x14 && s.GetScriptBytes()[23] == OP_EQUALVERIFY && s.GetScriptBytes()[24] == OP_CHECKSIG
}

func (s Script) IsMultiSig() bool {
	if len(s.GetScriptBytes()) < 3+1+pubkey.COMPRESSED_PUBLIC_KEY_SIZE {
		return false
	}

	// op_smallinteger
	opSmallInt1 := s.GetScriptBytes()[0]
	opSmallInt2 := s.GetScriptBytes()[s.GetScriptLength()-2]
	if opSmallInt1 != OP_0 && (opSmallInt1 < OP_1 || opSmallInt1 > OP_16) {
		return false
	}
	if opSmallInt2 != OP_0 && (opSmallInt2 < OP_1 || opSmallInt2 > OP_16) {
		return false
	}
	// n <= m
	if opSmallInt1 > opSmallInt2 {
		return false
	}

	return true
}

func (s Script) IsPayToWitnessScriptHash() bool {
	return len(s.data) == 34 && s.data[0] == OP_0 && s.data[1] == 0x20
}

func (s Script) IsWitnessProgram() (bool, int, []byte) {
	if len(s.data) < 4 || len(s.data) > 42 {
		return false, 0, []byte{}
	}
	if s.data[0] != OP_0 && (s.data[0] < OP_1 || s.data[0] > OP_16) {
		return false, 0, []byte{}
	}
	if int(s.data[1]+2) == len(s.data) {
		version := DecodeOPN(s.data[0])
		program := s.data[2:]
		return true, version, program
	}
	return false, 0, []byte{}
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

func (s *ScriptWitness) SetScriptWitnessBytes(witnessBytes [][]byte) {
	s.stack = witnessBytes
}
