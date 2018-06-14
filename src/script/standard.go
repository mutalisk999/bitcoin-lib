package script

import (
	"errors"
	"pubkey"
)

func loadPubKeys(pubKeysBytes []byte) (error, [][]byte) {
	var pubKeys [][]byte
	indexPubKeyStart := 0
	indexPubKeyEnd := 0

	for {
		pubKeySize := int(pubKeysBytes[indexPubKeyStart])
		if pubKeySize != pubkey.COMPRESSED_PUBLIC_KEY_SIZE && pubKeySize != pubkey.PUBLIC_KEY_SIZE {
			return errors.New("loadPubKeys: invalid pubkey size"), [][]byte{}
		}
		indexPubKeyEnd = indexPubKeyStart + 1 + pubKeySize
		if indexPubKeyEnd > len(pubKeysBytes) {
			return errors.New("loadPubKeys: invalid pubKeysBytes"), [][]byte{}
		}

		pubKeyBytes := pubKeysBytes[indexPubKeyStart+1 : indexPubKeyEnd]
		pubKeyNew := new(pubkey.PubKey)
		err := pubKeyNew.SetPubKeyData(pubKeyBytes)
		if err != nil {
			return err, [][]byte{}
		}
		pubKeys = append(pubKeys, pubKeyNew.GetPubKeyData())

		if indexPubKeyEnd == len(pubKeysBytes) {
			break
		}
		indexPubKeyStart = indexPubKeyEnd
	}
	return nil, pubKeys
}

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
	if scriptPubKey.IsPayToPubKeyHash() {
		return true, TX_PUBKEYHASH, [][]byte{scriptPubKey.GetScriptBytes()[3:23]}
	}

	// p2pk
	if scriptPubKey.IsPayToPubKey() {
		pubKeyBytes := scriptPubKey.GetScriptBytes()[1:]
		newPubKey := new(pubkey.PubKey)
		err := newPubKey.SetPubKeyData(pubKeyBytes)
		if err == nil {
			return true, TX_PUBKEY, [][]byte{pubKeyBytes}
		} else {
			return false, TX_NONSTANDARD, [][]byte{}
		}
	}

	// multisig
	if scriptPubKey.IsMultiSig() {
		_ = DecodeOPN(scriptPubKey.GetScriptBytes()[0])
		m := DecodeOPN(scriptPubKey.GetScriptBytes()[len(scriptPubKey.GetScriptBytes())-2])
		pubKeysBytes := scriptPubKey.GetScriptBytes()[1 : len(scriptPubKey.GetScriptBytes())-2]
		err, pubKeys := loadPubKeys(pubKeysBytes)
		if err != nil {
			return false, TX_NONSTANDARD, [][]byte{}
		} else {
			if m != len(pubKeys) {
				return false, TX_NONSTANDARD, [][]byte{}
			} else {
				return true, TX_MULTISIG, pubKeys
			}
		}
	}

	return false, TX_NONSTANDARD, [][]byte{}
}
