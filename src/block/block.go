package block

import (
	"bigint"
	"transaction"
	"io"
	"serialize"
	"bytes"
	"blob"
)

type BlockHeader struct {
	Version        int32
	HashPrevBlock  bigint.Uint256
	HashMerkleRoot bigint.Uint256
	Time           uint32
	Bits           uint32
	Nonce          uint32
}

func (b BlockHeader) Pack(writer io.Writer) error {
	var err error
	err = serialize.PackInt32(writer, b.Version)
	if err != nil {
		return err
	}
	err = b.HashPrevBlock.Pack(writer)
	if err != nil {
		return err
	}
	err = b.HashMerkleRoot.Pack(writer)
	if err != nil {
		return err
	}
	err = serialize.PackUint32(writer, b.Time)
	if err != nil {
		return err
	}
	err = serialize.PackUint32(writer, b.Bits)
	if err != nil {
		return err
	}
	err = serialize.PackUint32(writer, b.Nonce)
	if err != nil {
		return err
	}
	return nil
}

func (b *BlockHeader) UnPack(reader io.Reader) error {
	var err error
	b.Version, err = serialize.UnPackInt32(reader)
	if err != nil {
		return err
	}
	err = b.HashPrevBlock.UnPack(reader)
	if err != nil {
		return err
	}
	err = b.HashMerkleRoot.UnPack(reader)
	if err != nil {
		return err
	}
	b.Time, err = serialize.UnPackUint32(reader)
	if err != nil {
		return err
	}
	b.Bits, err = serialize.UnPackUint32(reader)
	if err != nil {
		return err
	}
	b.Nonce, err = serialize.UnPackUint32(reader)
	if err != nil {
		return err
	}
	return nil
}

type Block struct {
	Header BlockHeader
	Vtx    []transaction.Transaction
}

func (b Block) Pack(writer io.Writer) error {
	var err error
	err = b.Header.Pack(writer)
	if err != nil {
		return err
	}
	err = serialize.PackCompactSize(writer, uint64(len(b.Vtx)))
	if err != nil {
		return err
	}
	for _, tx := range b.Vtx {
		err = tx.Pack(writer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b Block) PackToHex() (string, error) {
	bytesBuf := bytes.NewBuffer([]byte{})
	bufWriter := io.Writer(bytesBuf)
	err := b.Pack(bufWriter)
	if err != nil {
		return "", err
	}
	Blob := new(blob.Byteblob)
	Blob.SetData(bytesBuf.Bytes())
	return Blob.GetHex(), nil
}

func (b *Block) UnPack(reader io.Reader) error {
	var err error
	var txCount uint64
	err = b.Header.UnPack(reader)
	if err != nil {
		return err
	}
	txCount, err = serialize.UnPackCompactSize(reader)
	if err != nil {
		return err
	}
	for i:=0; i<int(txCount); i++ {
		var tx transaction.Transaction
		err = tx.UnPack(reader)
		if err != nil {
			return err
		}
		b.Vtx = append(b.Vtx, tx)
	}
	return nil
}

func (b *Block) UnPackFromHex(hexStr string) error {
	Blob := new(blob.Byteblob)
	err := Blob.SetHex(hexStr)
	if err != nil {
		return err
	}
	bytesBuf := bytes.NewBuffer(Blob.GetData())
	bufReader := io.Reader(bytesBuf)
	err = b.UnPack(bufReader)
	if err != nil {
		return err
	}
	return nil
}