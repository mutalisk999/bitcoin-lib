package script

import (
	"blob"
	"fmt"
	"testing"
)

func TestSolverP2WPKH(t *testing.T) {
	Blob := new(blob.Byteblob)
	Blob.SetHex("0014c13948dfd851ad7d9c68f150bc4c93e252df8026")
	fmt.Println("byte blob:", Blob.GetData())

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(Blob.GetData())
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverP2WSH(t *testing.T) {
	Blob := new(blob.Byteblob)
	Blob.SetHex("0020558068a61b5569b9f5991ad8c674f5b8ad40125df81faa7d47d8df0c725514b4")
	fmt.Println("byte blob:", Blob.GetData())

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(Blob.GetData())
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverP2PKH(t *testing.T) {
	Blob := new(blob.Byteblob)
	Blob.SetHex("76a914a7092d2dc8778b56d4c352697081c687b451ab6d88ac")
	fmt.Println("byte blob:", Blob.GetData())

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(Blob.GetData())
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverP2PK(t *testing.T) {
	Blob := new(blob.Byteblob)
	Blob.SetHex("21027d886b785ddf10f085ce709810750fe3dc938c8dbe40d96b07f1a7ab909cb05fac")
	fmt.Println("byte blob:", Blob.GetData())

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(Blob.GetData())
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverP2SH(t *testing.T) {
	Blob := new(blob.Byteblob)
	Blob.SetHex("a914d550ecfc60d9f976de1f2a43bdf4e491b684cd6887")
	fmt.Println("byte blob:", Blob.GetData())

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(Blob.GetData())
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverMultiSig(t *testing.T) {
	Blob := new(blob.Byteblob)
	Blob.SetHex("5121032487c2a32f7c8d57d2a93906a6457afd00697925b0e6e145d89af6d3bca330162102308673d16987eaa010e540901cc6fe3695e758c19f46ce604e174dac315e685a52ae")
	fmt.Println("byte blob:", Blob.GetData())

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(Blob.GetData())
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}

func TestSolverNullData(t *testing.T) {
	Blob := new(blob.Byteblob)
	Blob.SetHex("6a08456c657068656e74")
	fmt.Println("byte blob:", Blob.GetData())

	scriptPubKey := new(Script)
	scriptPubKey.SetScriptBytes(Blob.GetData())
	isSuccess, scriptType, solution := Solver(*scriptPubKey)
	fmt.Println("isSuccess", isSuccess)
	fmt.Println("scriptType:", scriptType)
	fmt.Println("solution:", solution)
}
