package transaction

import (
	"bigint"
	"io"
	"script"
	"serialize"
	"errors"
	"bytes"
	"blob"
)

type OutPoint struct {
	Hash bigint.Uint256
	N    uint32
}

func (o OutPoint) Pack(writer io.Writer) error {
	err := o.Hash.Pack(writer)
	if err != nil {
		return err
	}
	err = serialize.PackUint32(writer, o.N)
	if err != nil {
		return err
	}
	return nil
}

func (o *OutPoint) UnPack(reader io.Reader) error {
	err := o.Hash.UnPack(reader)
	if err != nil {
		return err
	}
	o.N, err = serialize.UnPackUint32(reader)
	if err != nil {
		return err
	}
	return nil
}

type TxIn struct {
	PrevOut       OutPoint
	ScriptSig     script.Script
	Sequence      uint32
	ScriptWitness script.ScriptWitness
}

func (t TxIn) Pack(writer io.Writer) error {
	err := t.PrevOut.Pack(writer)
	if err != nil {
		return err
	}
	err = t.ScriptSig.Pack(writer)
	if err != nil {
		return err
	}
	err = serialize.PackUint32(writer, t.Sequence)
	if err != nil {
		return err
	}
	err = t.ScriptWitness.Pack(writer)
	if err != nil {
		return err
	}
	return nil
}

func (t *TxIn) UnPack(reader io.Reader) error {
	err := t.PrevOut.UnPack(reader)
	if err != nil {
		return err
	}
	err = t.ScriptSig.UnPack(reader)
	if err != nil {
		return err
	}
	t.Sequence, err = serialize.UnPackUint32(reader)
	if err != nil {
		return err
	}
	err = t.ScriptWitness.UnPack(reader)
	if err != nil {
		return err
	}
	return nil
}

type TxOut struct {
	Value        int64
	ScriptPubKey script.Script
}

func (t TxOut) Pack(writer io.Writer) error {
	err := serialize.PackInt64(writer, t.Value)
	if err != nil {
		return err
	}
	err = t.ScriptPubKey.Pack(writer)
	if err != nil {
		return err
	}
	return nil
}

func (t *TxOut) UnPack(reader io.Reader) error {
	var err error
	t.Value, err = serialize.UnPackInt64(reader)
	if err != nil {
		return err
	}
	err = t.ScriptPubKey.UnPack(reader)
	if err != nil {
		return err
	}
	return nil
}

type Transaction struct {
	Vin			[]TxIn
	Vout		[]TxOut
	Version 	int32
	LockTime	uint32
}

func (t Transaction) HasWitness() bool {
	for _, v := range t.Vin {
		if len(v.ScriptWitness.GetScriptWitnessBytes()) != 0 {
			return true
		}
	}
	return false
}

func (t Transaction) packVin(writer io.Writer, vin []TxIn) error {
	err := serialize.PackCompactSize(writer, uint64(len(vin)))
	if err != nil {
		return err
	}
	for _, v := range vin {
		err = v.Pack(writer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t Transaction) packVout(writer io.Writer, vout []TxOut) error {
	err := serialize.PackCompactSize(writer, uint64(len(vout)))
	if err != nil {
		return err
	}
	for _, v := range vout {
		err = v.Pack(writer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t Transaction) Pack(writer io.Writer, witness bool) error {
	t.Version = 2
	err := serialize.PackInt32(writer, t.Version)
	if err != nil {
		return err
	}
	var flags uint8 = 0
	if witness && t.HasWitness() {
			flags = 1
	}
	if flags == 1 {
		// pack vinDummy and flags
		var vinDummy []TxIn
		err = t.packVin(writer, vinDummy)
		if err != nil {
			return err
		}
		err = serialize.PackUint8(writer, flags)
		if err != nil {
			return err
		}
	}
	// pack Vin
	err = t.packVin(writer, t.Vin)
	if err != nil {
		return err
	}
	// pack Vout
	err = t.packVout(writer, t.Vout)
	if err != nil {
		return err
	}
	if flags == 1 {
		// pack ScriptWitness
		for _, v := range t.Vin {
			err = v.ScriptWitness.Pack(writer)
			if err != nil {
				return err
			}
		}
	}
	err = serialize.PackUint32(writer, t.LockTime)
	if err != nil {
		return err
	}
	return nil
}

func (t Transaction) PackToHex(witness bool) (string, error) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	err := t.Pack(bufWriter, witness)
	if err != nil {
		return "", err
	}
	blob := new(blob.Byteblob)
	blob.SetData(bytesBuf.Bytes())
	return blob.GetHex(), nil
}

func (t *Transaction) unpackVin(reader io.Reader) ([]TxIn, error) {
	var vin []TxIn
	ui64, err := serialize.UnPackCompactSize(reader)
	if err != nil {
		return nil, err
	}
	for i:= 0; i < int(ui64); i++ {
		var v TxIn
		err = v.UnPack(reader)
		if err != nil {
			return nil,err
		}
		vin = append(vin, v)
	}
	return vin, nil
}

func (t *Transaction) unpackVout(reader io.Reader) ([]TxOut, error) {
	var vout []TxOut
	ui64, err := serialize.UnPackCompactSize(reader)
	if err != nil {
		return nil, err
	}
	for i:= 0; i < int(ui64); i++ {
		var v TxOut
		err = v.UnPack(reader)
		if err != nil {
			return nil,err
		}
		vout = append(vout, v)
	}
	return vout, nil
}

func (t *Transaction) UnPack(reader io.Reader, witness bool) error {
	var err error
	var flags uint8 = 0
	var vin []TxIn
	var vout []TxOut
	t.Version, err = serialize.UnPackInt32(reader)
	if err != nil {
		return err
	}
	// unpack Vin
	vin, err = t.unpackVin(reader)
	if err != nil {
		return err
	}
	t.Vin = vin
	if len(vin) == 0 && witness {   // witness
		flags, err = serialize.UnPackUint8(reader)
		if err != nil {
			return err
		}
		if flags != 0 {
			// unpack Vin
			vin, err = t.unpackVin(reader)
			if err != nil {
				return err
			}
			t.Vin = vin
			// unpack Vout
			vout, err = t.unpackVout(reader)
			if err != nil {
				return err
			}
			t.Vout = vout
		}
	} else {   // not witness
		// unpack Vout
		vout, err = t.unpackVout(reader)
		if err != nil {
			return err
		}
		t.Vout = vout
	}
	if ((flags & 1) == 1) && witness {
		flags = flags ^ 1
		// unpack ScriptWitness
		for i:=0; i < len(t.Vin); i++ {
			err = t.Vin[i].ScriptWitness.UnPack(reader)
			if err != nil {
				return err
			}
		}
	}
	if flags != 0 {
		return errors.New("Transaction::Unpack: Unknown transaction option data")
	}
	t.LockTime, err = serialize.UnPackUint32(reader)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) UnPackFromHex(hexStr string, witness bool) (error) {
	blob := new(blob.Byteblob)
	err := blob.SetHex(hexStr)
	if err != nil {
		return err
	}
	bytesBuf := bytes.NewBuffer(blob.GetData())
	bufReader := io.Reader(bytesBuf)
	err = t.UnPack(bufReader, witness)
	if err != nil {
		return err
	}
	return nil
}
