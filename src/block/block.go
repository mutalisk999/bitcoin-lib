package block

import "bigint"

type BlockHeader struct {
	Version        int32
	HashPrevBlock  bigint.Uint256
	HashMerkleRoot bigint.Uint256
	Time           uint32
	Bits           uint32
	Nonce          uint32
}

type Block struct {
	Header BlockHeader
}
