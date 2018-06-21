package script

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSolverP2WPKH(t *testing.T) {
	// from tx 7dcdd326eeba8aef44f20913751c845a88a98859fcd5e188c0e27f6252004a10
	scriptBytes, _ := hex.DecodeString("0014edec29b622585acbc82c914710070831bfb2b648")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverP2WSH(t *testing.T) {
	// from tx 14ddab2379a23faf4e3ac26737827a59a0dd78f0c7730e1c925f389238ae9a3b
	scriptBytes, _ := hex.DecodeString("00202122f4719add322f4d727f48379f8a8ba36a40ec4473fd99a2fdcfd89a16e048")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverP2PKH(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("76a914a7092d2dc8778b56d4c352697081c687b451ab6d88ac")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverP2PK(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("21027d886b785ddf10f085ce709810750fe3dc938c8dbe40d96b07f1a7ab909cb05fac")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverP2SH(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("a914d550ecfc60d9f976de1f2a43bdf4e491b684cd6887")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverMultiSig(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("5121032487c2a32f7c8d57d2a93906a6457afd00697925b0e6e145d89af6d3bca330162102308673d16987eaa010e540901cc6fe3695e758c19f46ce604e174dac315e685a52ae")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverNullData(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("6a08456c657068656e74")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestExtractDestinationP2WPKH(t *testing.T) {
	// from tx 7dcdd326eeba8aef44f20913751c845a88a98859fcd5e188c0e27f6252004a10
	scriptBytes, _ := hex.DecodeString("0014edec29b622585acbc82c914710070831bfb2b648")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, destAddrs := ExtractDestination(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("destAddrs:", destAddrs)
}

func TestExtractDestinationP2WSH(t *testing.T) {
	// from tx 14ddab2379a23faf4e3ac26737827a59a0dd78f0c7730e1c925f389238ae9a3b
	scriptBytes, _ := hex.DecodeString("00202122f4719add322f4d727f48379f8a8ba36a40ec4473fd99a2fdcfd89a16e048")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, destAddrs := ExtractDestination(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("destAddrs:", destAddrs)
}

func TestExtractDestinationP2PKH(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("76a914a7092d2dc8778b56d4c352697081c687b451ab6d88ac")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, destAddrs := ExtractDestination(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("destAddrs:", destAddrs)
}

func TestExtractDestinationP2PK(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("21027d886b785ddf10f085ce709810750fe3dc938c8dbe40d96b07f1a7ab909cb05fac")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, destAddrs := ExtractDestination(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("destAddrs:", destAddrs)
}

func TestExtractDestinationP2SH(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("a914d550ecfc60d9f976de1f2a43bdf4e491b684cd6887")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, destAddrs := ExtractDestination(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", destAddrs)
}

func TestExtractDestinationMultiSig(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("5121032487c2a32f7c8d57d2a93906a6457afd00697925b0e6e145d89af6d3bca330162102308673d16987eaa010e540901cc6fe3695e758c19f46ce604e174dac315e685a52ae")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, destAddrs := ExtractDestination(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", destAddrs)
}

func TestExtractDestinationNullData(t *testing.T) {
	scriptBytes, _ := hex.DecodeString("6a08456c657068656e74")
	fmt.Println("script bytes:", scriptBytes)

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(scriptBytes)
	isSuccess, scriptType, destAddrs := ExtractDestination(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", destAddrs)
}
