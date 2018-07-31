package transaction

import (
	"bytes"
	"fmt"
	"github.com/mutalisk999/bitcoin-lib/src/blob"
	"io"
	"testing"
)

func TestTransaction(t *testing.T) {
	Blob := new(blob.Byteblob)
	Blob.SetHex("0200000001f940e14f9cc1b7fe004fe5e09714727eca5e3719c5c560ae1eb68ed03441f871010000006b483045022100effc237749e91ba46a38bec5685dca75c93c2a36434df8630ef52c872365228b022054b9ac11cd31f98485bab72a2eab725de462726730d13ffdb0c1e9a0be3584c901210274d7bd86feb8ac4833987cdd2adb05f14780840bebccb6b084060476cb84cffffeffffff0200562183000000001976a9146f24b7c21f4038710da905f62e73d363c1d186f088ac972c6200000000001976a914ca58ff72bb9168f49adb4a0c786e94639598208b88acf5070800")
	fmt.Println("byte blob:", Blob.GetData())

	bytesBuf := bytes.NewBuffer(Blob.GetData())
	bufReader := io.Reader(bytesBuf)
	trx := new(Transaction)
	trx.UnPack(bufReader)
	fmt.Println("trx version:", trx.Version)
	fmt.Println("trx locktime", trx.LockTime)
	fmt.Println("trx vin size:", len(trx.Vin))
	for i := 0; i < len(trx.Vin); i++ {
		fmt.Println("vin prevout:", trx.Vin[i].PrevOut)
		fmt.Println("vin scriptsig:", trx.Vin[i].ScriptSig)
		fmt.Println("vin sequence:", trx.Vin[i].Sequence)
		fmt.Println("vin scriptwitness:", trx.Vin[i].ScriptWitness)
	}
	fmt.Println("trx vout size:", len(trx.Vout))
	for i := 0; i < len(trx.Vout); i++ {
		fmt.Println("vout value", trx.Vout[i].Value)
		fmt.Println("vout scriptpubkey:", trx.Vout[i].ScriptPubKey)
	}
	bytesBuf = bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	trx.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())
}

func TestTransactionHex(t *testing.T) {
	trx := new(Transaction)
	trx.UnPackFromHex("0200000001f940e14f9cc1b7fe004fe5e09714727eca5e3719c5c560ae1eb68ed03441f871010000006b483045022100effc237749e91ba46a38bec5685dca75c93c2a36434df8630ef52c872365228b022054b9ac11cd31f98485bab72a2eab725de462726730d13ffdb0c1e9a0be3584c901210274d7bd86feb8ac4833987cdd2adb05f14780840bebccb6b084060476cb84cffffeffffff0200562183000000001976a9146f24b7c21f4038710da905f62e73d363c1d186f088ac972c6200000000001976a914ca58ff72bb9168f49adb4a0c786e94639598208b88acf5070800")
	fmt.Println("trx version:", trx.Version)
	fmt.Println("trx locktime", trx.LockTime)
	fmt.Println("trx vin size:", len(trx.Vin))
	for i := 0; i < len(trx.Vin); i++ {
		fmt.Println("vin prevout:", trx.Vin[i].PrevOut)
		fmt.Println("vin scriptsig:", trx.Vin[i].ScriptSig)
		fmt.Println("vin sequence:", trx.Vin[i].Sequence)
		fmt.Println("vin scriptwitness:", trx.Vin[i].ScriptWitness)
	}
	fmt.Println("trx vout size:", len(trx.Vout))
	for i := 0; i < len(trx.Vout); i++ {
		fmt.Println("vout value", trx.Vout[i].Value)
		fmt.Println("vout scriptpubkey:", trx.Vout[i].ScriptPubKey)
	}
	hexStr, _ := trx.PackToHex()
	fmt.Println("hex string:", hexStr)
}

func TestWitnessTransaction(t *testing.T) {
	Blob := new(blob.Byteblob)
	Blob.SetHex("010000000001018fd47668c7921a018be8e8e8b27eb88627de03c7d3bb8827fb71dc126a5040280100000000ffffffff02d06e57000000000017a914a000cf5d47102c48f493483de12c699a688789ad87a3a2970000000000220020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d0400483045022100e9575bf83561ba9c713497074062c59facb5793094f8718af48302b4ca95936d022033d144c70b022a6430ae44324b811263ee670fc40d387cfd3a74b4de4eb5732d01483045022100a85a12f9f349cb052c8190ef4cc9f31bfb8d6e9b8297984656d218307087b53f02200d3fb116429d45e7031aa8cbb101410179a2e66d6d7bca9ac54379e8e01b8009016952210375e00eb72e29da82b89367947f29ef34afb75e8654f6ea368e0acdfd92976b7c2103a1b26313f430c4b15bb1fdce663207659d8cac749a0e53d70eff01874496feff2103c96d495bfdd5ba4145e3e046fee45e84a8a48ad05bd8dbb395c011a32cf9f88053ae00000000")
	fmt.Println("byte blob:", Blob.GetData())

	bytesBuf := bytes.NewBuffer(Blob.GetData())
	bufReader := io.Reader(bytesBuf)
	trx := new(Transaction)
	trx.UnPack(bufReader)
	fmt.Println("trx version:", trx.Version)
	fmt.Println("trx locktime", trx.LockTime)
	fmt.Println("trx vin size:", len(trx.Vin))
	for i := 0; i < len(trx.Vin); i++ {
		fmt.Println("vin prevout:", trx.Vin[i].PrevOut)
		fmt.Println("vin scriptsig:", trx.Vin[i].ScriptSig)
		fmt.Println("vin sequence:", trx.Vin[i].Sequence)
		fmt.Println("vin scriptwitness:", trx.Vin[i].ScriptWitness)
	}
	fmt.Println("trx vout size:", len(trx.Vout))
	for i := 0; i < len(trx.Vout); i++ {
		fmt.Println("vout value", trx.Vout[i].Value)
		fmt.Println("vout scriptpubkey:", trx.Vout[i].ScriptPubKey)
	}
	bytesBuf = bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	trx.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())
}

func TestWitnessTransactionHex(t *testing.T) {
	trx := new(Transaction)
	trx.UnPackFromHex("010000000001018fd47668c7921a018be8e8e8b27eb88627de03c7d3bb8827fb71dc126a5040280100000000ffffffff02d06e57000000000017a914a000cf5d47102c48f493483de12c699a688789ad87a3a2970000000000220020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d0400483045022100e9575bf83561ba9c713497074062c59facb5793094f8718af48302b4ca95936d022033d144c70b022a6430ae44324b811263ee670fc40d387cfd3a74b4de4eb5732d01483045022100a85a12f9f349cb052c8190ef4cc9f31bfb8d6e9b8297984656d218307087b53f02200d3fb116429d45e7031aa8cbb101410179a2e66d6d7bca9ac54379e8e01b8009016952210375e00eb72e29da82b89367947f29ef34afb75e8654f6ea368e0acdfd92976b7c2103a1b26313f430c4b15bb1fdce663207659d8cac749a0e53d70eff01874496feff2103c96d495bfdd5ba4145e3e046fee45e84a8a48ad05bd8dbb395c011a32cf9f88053ae00000000")
	fmt.Println("trx version:", trx.Version)
	fmt.Println("trx locktime", trx.LockTime)
	fmt.Println("trx vin size:", len(trx.Vin))
	for i := 0; i < len(trx.Vin); i++ {
		fmt.Println("vin prevout:", trx.Vin[i].PrevOut)
		fmt.Println("vin scriptsig:", trx.Vin[i].ScriptSig)
		fmt.Println("vin sequence:", trx.Vin[i].Sequence)
		fmt.Println("vin scriptwitness:", trx.Vin[i].ScriptWitness)
	}
	fmt.Println("trx vout size:", len(trx.Vout))
	for i := 0; i < len(trx.Vout); i++ {
		fmt.Println("vout value", trx.Vout[i].Value)
		fmt.Println("vout scriptpubkey:", trx.Vout[i].ScriptPubKey)
	}
	hexStr, _ := trx.PackToHex()
	fmt.Println("hex string:", hexStr)
}

func TestCalcTrxId(t *testing.T) {
	trx := new(Transaction)
	trx.UnPackFromHex("01000000010123449fad6289dda5365a197fa822e320cdfc106bed243bf773cac64cfdb237050000006a473044022036be6403aeb4e0e6fd54720b328d9d81bea32fb79684da0288743668fb5ef3ee02202023a71ef7217061fb9b4f35a05143de71447032e5a35b39c3d14b3210bad10b0121032725846bb7bc2e47b7b5a50670d77c8268f4d7f3243bdcf1b22174a67faaf528feffffff0200213900000000001976a914659042e01e864e2f29641ea3a213c51a956d33c788ac288c0f00000000001976a9144fc238bcda3f884ff6ce8d9feeb89b50dfd3da8888ac2c480700")
	trxId, _ := trx.CalcTrxId()
	fmt.Println("trx id:", trxId)
}
