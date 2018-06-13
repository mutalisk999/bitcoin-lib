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
