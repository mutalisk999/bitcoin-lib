package transaction

import (
	"bigint"
	"io"
	"script"
	"serialize"
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
	for _, txin := range t.Vin {
		if len(txin.ScriptWitness.GetScriptWitnessBytes()) != 0 {
			return true
		}
	}
	return false
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
		err = serialize.PackCompactSize(writer, uint64(len(vinDummy)))
		if err != nil {
			return err
		}
		err = serialize.PackUint8(writer, flags)
		if err != nil {
			return err
		}
	}
	// pack Vin
	err = serialize.PackCompactSize(writer, uint64(len(t.Vin)))
	if err != nil {
		return err
	}
	for _, vin := range t.Vin {
		err = vin.Pack(writer)
		if err != nil {
			return err
		}
	}
	// pack Vout
	err = serialize.PackCompactSize(writer, uint64(len(t.Vout)))
	if err != nil {
		return err
	}
	for _, vout := range t.Vout {
		err = vout.Pack(writer)
		if err != nil {
			return err
		}
	}
	if flags == 1 {
		// pack ScriptWitness
		for _, vin := range t.Vin {
			err = vin.ScriptWitness.Pack(writer)
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

func (t Transaction) PackToHex(witness bool) (error, string) {

}

func (t *Transaction) UnPack(reader io.Reader, witness bool) error {

}

func (t *Transaction) UnPackFromHex(hexStr string) (error) {

}
