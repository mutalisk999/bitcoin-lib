package script

import (
	"testing"
	"blob"
	"fmt"
)

func TestSolverWitness(t *testing.T) {

}

func TestSolverP2PKH(t *testing.T) {

}

func TestSolverP2PK(t *testing.T) {

}

func TestSolverP2SH(t *testing.T) {

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
