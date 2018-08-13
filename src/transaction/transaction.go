package transaction

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/mutalisk999/bitcoin-lib/src/bigint"
	"github.com/mutalisk999/bitcoin-lib/src/blob"
	"github.com/mutalisk999/bitcoin-lib/src/script"
	"github.com/mutalisk999/bitcoin-lib/src/serialize"
	"github.com/mutalisk999/bitcoin-lib/src/utility"
	"io"
	"strings"
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

type OutPointPrintAble struct {
	Hash string
	N    uint32
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
	return nil
}

type TxInPrintAble struct {
	PrevOut       OutPointPrintAble
	ScriptSig     string
	Sequence      uint32
	ScriptWitness []string
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

type TxOutPrintAble struct {
	Value        int64
	ScriptPubKey string
	Address      string
	ScriptType   string
}

type Transaction struct {
	Vin      []TxIn
	Vout     []TxOut
	Version  int32
	LockTime uint32
}

func (t Transaction) HasWitness() bool {
	for _, v := range t.Vin {
		if len(v.ScriptWitness.GetScriptWitnessBytes()) != 0 {
			return true
		}
	}
	return false
}

func (t Transaction) packVin(writer io.Writer, vin *[]TxIn) error {
	err := serialize.PackCompactSize(writer, uint64(len(*vin)))
	if err != nil {
		return err
	}
	for _, v := range *vin {
		err = v.Pack(writer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t Transaction) packVout(writer io.Writer, vout *[]TxOut) error {
	err := serialize.PackCompactSize(writer, uint64(len(*vout)))
	if err != nil {
		return err
	}
	for _, v := range *vout {
		err = v.Pack(writer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t Transaction) Pack(writer io.Writer) error {
	err := serialize.PackInt32(writer, t.Version)
	if err != nil {
		return err
	}
	var flags uint8 = 0
	if t.HasWitness() {
		flags = 1
	}
	if flags == 1 {
		// pack vinDummy and flags
		var vinDummy []TxIn
		err = t.packVin(writer, &vinDummy)
		if err != nil {
			return err
		}
		err = serialize.PackUint8(writer, flags)
		if err != nil {
			return err
		}
	}
	// pack Vin
	err = t.packVin(writer, &t.Vin)
	if err != nil {
		return err
	}
	// pack Vout
	err = t.packVout(writer, &t.Vout)
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

func (t Transaction) PackToHex() (string, error) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	err := t.Pack(bufWriter)
	if err != nil {
		return "", err
	}
	Blob := new(blob.Byteblob)
	Blob.SetData(bytesBuf.Bytes())
	return Blob.GetHex(), nil
}

func (t Transaction) PackNoWitness(writer io.Writer) error {
	err := serialize.PackInt32(writer, t.Version)
	if err != nil {
		return err
	}
	// pack Vin
	err = t.packVin(writer, &t.Vin)
	if err != nil {
		return err
	}
	// pack Vout
	err = t.packVout(writer, &t.Vout)
	if err != nil {
		return err
	}
	err = serialize.PackUint32(writer, t.LockTime)
	if err != nil {
		return err
	}
	return nil
}

func (t Transaction) CalcTrxId() (bigint.Uint256, error) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	err := t.PackNoWitness(bufWriter)
	if err != nil {
		return bigint.Uint256{}, err
	}
	bytesHash := utility.Sha256(utility.Sha256(bytesBuf.Bytes()))
	ui256 := new(bigint.Uint256)
	ui256.SetData(bytesHash)
	return *ui256, nil
}

func (t *Transaction) unpackVin(reader io.Reader) (*[]TxIn, error) {
	var vin []TxIn
	ui64, err := serialize.UnPackCompactSize(reader)
	if err != nil {
		return nil, err
	}
	vin = make([]TxIn, 0, ui64)
	for i := 0; i < int(ui64); i++ {
		var v TxIn
		err = v.UnPack(reader)
		if err != nil {
			return nil, err
		}
		vin = append(vin, v)
	}
	return &vin, nil
}

func (t *Transaction) unpackVout(reader io.Reader) (*[]TxOut, error) {
	var vout []TxOut
	ui64, err := serialize.UnPackCompactSize(reader)
	if err != nil {
		return nil, err
	}
	vout = make([]TxOut, 0, ui64)
	for i := 0; i < int(ui64); i++ {
		var v TxOut
		err = v.UnPack(reader)
		if err != nil {
			return nil, err
		}
		vout = append(vout, v)
	}
	return &vout, nil
}

func (t *Transaction) UnPack(reader io.Reader) error {
	var err error
	var flags uint8 = 0
	var vin *[]TxIn
	var vout *[]TxOut
	t.Version, err = serialize.UnPackInt32(reader)
	if err != nil {
		return err
	}
	// unpack Vin
	vin, err = t.unpackVin(reader)
	if err != nil {
		return err
	}
	t.Vin = *vin
	if len(*vin) == 0 { // witness
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
			t.Vin = *vin
			// unpack Vout
			vout, err = t.unpackVout(reader)
			if err != nil {
				return err
			}
			t.Vout = *vout
		}
	} else { // not witness
		// unpack Vout
		vout, err = t.unpackVout(reader)
		if err != nil {
			return err
		}
		t.Vout = *vout
	}
	if (flags & 1) == 1 {
		flags = flags ^ 1
		// unpack ScriptWitness
		for i := 0; i < len(t.Vin); i++ {
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

func (t *Transaction) UnPackFromHex(hexStr string) error {
	Blob := new(blob.Byteblob)
	err := Blob.SetHex(hexStr)
	if err != nil {
		return err
	}
	bytesBuf := bytes.NewBuffer(Blob.GetData())
	bufReader := io.Reader(bytesBuf)
	err = t.UnPack(bufReader)
	if err != nil {
		return err
	}
	return nil
}

type TrxPrintAble struct {
	Vin      []TxInPrintAble
	Vout     []TxOutPrintAble
	Version  int32
	LockTime uint32
}

func (t *Transaction) GetTrxPrintAble() TrxPrintAble {
	trxPrintAble := new(TrxPrintAble)
	for _, vin := range t.Vin {
		vinPrintAble := new(TxInPrintAble)
		vinPrintAble.PrevOut.Hash = vin.PrevOut.Hash.GetHex()
		vinPrintAble.PrevOut.N = vin.PrevOut.N
		if vin.PrevOut.Hash.GetHex() == "0000000000000000000000000000000000000000000000000000000000000000" {
			vinPrintAble.PrevOut.Hash = ""
		}
		vinPrintAble.ScriptSig = hex.EncodeToString(vin.ScriptSig.GetScriptBytes())
		vinPrintAble.Sequence = vin.Sequence
		vinPrintAble.ScriptWitness = []string{}
		for _, scriptWitness := range vin.ScriptWitness.GetScriptWitnessBytes() {
			vinPrintAble.ScriptWitness = append(vinPrintAble.ScriptWitness, hex.EncodeToString(scriptWitness))
		}
		trxPrintAble.Vin = append(trxPrintAble.Vin, *vinPrintAble)
	}
	for _, vout := range t.Vout {
		voutPrintAble := new(TxOutPrintAble)
		voutPrintAble.Value = vout.Value
		voutPrintAble.ScriptPubKey = hex.EncodeToString(vout.ScriptPubKey.GetScriptBytes())
		isSucc, scriptType, addresses := script.ExtractDestination(vout.ScriptPubKey)
		var addrStr string
		if isSucc {
			addrStr = ""
			if script.IsSingleAddress(scriptType) {
				addrStr = addresses[0]
			} else if script.IsMultiAddress(scriptType) {
				addrStr = strings.Join(addresses, ",")
			}
		}
		voutPrintAble.Address = addrStr
		voutPrintAble.ScriptType = script.GetScriptTypeStr(scriptType)
		trxPrintAble.Vout = append(trxPrintAble.Vout, *voutPrintAble)
	}
	trxPrintAble.Version = t.Version
	trxPrintAble.LockTime = t.LockTime
	return *trxPrintAble
}
