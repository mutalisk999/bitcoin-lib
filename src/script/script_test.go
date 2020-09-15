package script

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"testing"
)

func TestScript(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("1976a914cb02eb6ea281ef1fa194f55c530f6a080cfce99288ac")
	fmt.Println("script bytes:", scriptBytes)
	bytesBuf := bytes.NewBuffer(scriptBytes)
	bufReader := io.Reader(bytesBuf)
	script := new(Script)
	_ = script.UnPack(bufReader)
	fmt.Println("script data:", script.GetScriptBytes())

	bytesBuf = bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = script.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())
}

func TestScriptHex(t *testing.T) {
	script := new(Script)
	_ = script.UnPackFromHex("1976a914cb02eb6ea281ef1fa194f55c530f6a080cfce99288ac")
	fmt.Println("script data:", script.GetScriptBytes())

	hexStr, _ := script.PackToHex()
	fmt.Println("hex string:", hexStr)
}

func TestScriptWitness(t *testing.T) {
	scriptWitnessBytes, _ := hex.DecodeString("0400483045022100e9575bf83561ba9c713497074062c59facb5793094f8718af48302b4ca95936d022033d144c70b022a6430ae44324b811263ee670fc40d387cfd3a74b4de4eb5732d01483045022100a85a12f9f349cb052c8190ef4cc9f31bfb8d6e9b8297984656d218307087b53f02200d3fb116429d45e7031aa8cbb101410179a2e66d6d7bca9ac54379e8e01b8009016952210375e00eb72e29da82b89367947f29ef34afb75e8654f6ea368e0acdfd92976b7c2103a1b26313f430c4b15bb1fdce663207659d8cac749a0e53d70eff01874496feff2103c96d495bfdd5ba4145e3e046fee45e84a8a48ad05bd8dbb395c011a32cf9f88053ae")
	fmt.Println("script witness bytes:", scriptWitnessBytes)

	bytesBuf := bytes.NewBuffer(scriptWitnessBytes)
	bufReader := io.Reader(bytesBuf)
	script := new(ScriptWitness)
	_ = script.UnPack(bufReader)
	fmt.Println("script witness stack:", script.GetScriptWitnessBytes())

	bytesBuf = bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	_ = script.Pack(bufWriter)
	fmt.Println("byte buffer:", bytesBuf.Bytes())
}

func TestScriptWitnessHex(t *testing.T) {
	script := new(ScriptWitness)
	_ = script.UnPackFromHex("0400483045022100e9575bf83561ba9c713497074062c59facb5793094f8718af48302b4ca95936d022033d144c70b022a6430ae44324b811263ee670fc40d387cfd3a74b4de4eb5732d01483045022100a85a12f9f349cb052c8190ef4cc9f31bfb8d6e9b8297984656d218307087b53f02200d3fb116429d45e7031aa8cbb101410179a2e66d6d7bca9ac54379e8e01b8009016952210375e00eb72e29da82b89367947f29ef34afb75e8654f6ea368e0acdfd92976b7c2103a1b26313f430c4b15bb1fdce663207659d8cac749a0e53d70eff01874496feff2103c96d495bfdd5ba4145e3e046fee45e84a8a48ad05bd8dbb395c011a32cf9f88053ae")
	fmt.Println("script witness stack:", script.GetScriptWitnessBytes())

	hexStr, _ := script.PackToHex()
	fmt.Println("hex string:", hexStr)
}
