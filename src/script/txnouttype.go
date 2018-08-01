package script

const (
	TX_NONSTANDARD = iota
	// 'standard' transaction types:
	TX_PUBKEY
	TX_PUBKEYHASH
	TX_SCRIPTHASH
	TX_MULTISIG
	TX_NULL_DATA //!< unspendable OP_RETURN script that carries data
	TX_WITNESS_V0_SCRIPTHASH
	TX_WITNESS_V0_KEYHASH
	TX_WITNESS_UNKNOWN //!< Only for Witness versions not already defined above
)

func IsSingleAddress(scriptType int) bool {
	if scriptType == TX_PUBKEY || scriptType == TX_PUBKEYHASH || scriptType == TX_SCRIPTHASH ||
		scriptType == TX_WITNESS_V0_SCRIPTHASH || scriptType == TX_WITNESS_V0_KEYHASH {
		return true
	}
	return false
}

func IsMultiAddress(scriptType int) bool {
	if scriptType == TX_MULTISIG {
		return true
	}
	return false
}

func IsNoneAddress(scriptType int) bool {
	if IsSingleAddress(scriptType) {
		return false
	} else if IsMultiAddress(scriptType) {
		return false
	} else {
		return true
	}
	return false
}
