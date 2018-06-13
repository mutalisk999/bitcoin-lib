package script

import "pubkey"

func Solver(scriptPubKey Script) (bool, int, [][]byte) {
	// p2sh
	if scriptPubKey.IsPayToScriptHash() {
		scriptHash := scriptPubKey.GetScriptBytes()[2:22]
		return true, TX_SCRIPTHASH, [][]byte{scriptHash}
	}

	// witness
	isWitness, witnessVersion, witnessProgram := scriptPubKey.IsWitnessProgram()
	if isWitness {
		if witnessVersion == 0 && len(witnessProgram) == WITNESS_V0_KEYHASH_SIZE {
			return true, TX_WITNESS_V0_KEYHASH, [][]byte{witnessProgram}
		} else if witnessVersion == 0 && len(witnessProgram) == WITNESS_V0_SCRIPTHASH_SIZE {
			return true, TX_WITNESS_V0_SCRIPTHASH, [][]byte{witnessProgram}
		} else if witnessVersion != 0 {
			return true, TX_WITNESS_UNKNOWN, [][]byte{{byte(witnessVersion)}, witnessProgram}
		}
		return false, TX_NONSTANDARD, [][]byte{}
	}

	// op_return
	// and not to judge if the script only contains PUSHDATA opcode and something else related
	if len(scriptPubKey.GetScriptBytes()) >= 1 && scriptPubKey.GetScriptBytes()[0] == OP_RETURN {
		return true, TX_NULL_DATA, [][]byte{}
	}

	// p2pkh

	// p2pk
	if len(scriptPubKey.GetScriptBytes()) >= 1 &&
		(scriptPubKey.GetScriptBytes()[0] == pubkey.PUBLIC_KEY_SIZE || scriptPubKey.GetScriptBytes()[0] == pubkey.COMPRESSED_PUBLIC_KEY_SIZE) {
		pubkeyBytes := scriptPubKey.GetScriptBytes()[1:]
		pubkey := new(pubkey.PubKey)
		err := pubkey.SetPubKeyData(pubkeyBytes)
		if err == nil {
			return true, TX_PUBKEY, [][]byte{pubkeyBytes}
		} else {
			return false, TX_NONSTANDARD, [][]byte{}
		}
	}

	// ms

	return false, TX_NONSTANDARD, [][]byte{}
}
