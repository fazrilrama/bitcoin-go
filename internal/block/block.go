package block

import (
    "bytes"
    "crypto/sha256"
    "time"
)

type Block struct {
    Timestamp    int64
    Data         []byte
    PrevHash     []byte
    Hash         []byte
    Nonce        int
}

func CreateBlock(data string, prevHash []byte) *Block {
    block := &Block{
        Timestamp: time.Now().Unix(),
        Data:      []byte(data),
        PrevHash:  prevHash,
    }

    pow := NewProof(block)
    nonce, hash := pow.Run()

    block.Hash = hash
    block.Nonce = nonce

    return block
}

func Genesis() *Block {
    return CreateBlock("Genesis Block", []byte{})
}

func (b *Block) HashBlock() []byte {
    info := bytes.Join(
        [][]byte{
            b.PrevHash,
            b.Data,
            []byte(string(b.Timestamp)),
        },
        []byte{},
    )

    hash := sha256.Sum256(info)
    return hash[:]
}
