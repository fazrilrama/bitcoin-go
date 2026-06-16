package block

import (
    "time"

    "github.com/fazrilrama/bitcoin-go/internal/merkle"
    "github.com/fazrilrama/bitcoin-go/internal/transaction"
)

type Block struct {
    Timestamp   int64
    Transactions []*transaction.Transaction
    PrevHash    []byte
    Hash        []byte
    Nonce       int
    MerkleRoot  []byte
}

func CreateBlock(txs []*transaction.Transaction, prevHash []byte) *Block {
    block := &Block{
        Timestamp:    time.Now().Unix(),
        Transactions: txs,
        PrevHash:     prevHash,
        MerkleRoot:   merkleRoot(txs),
    }

    pow := NewProof(block)
    nonce, hash := pow.Run()
    block.Hash = hash
    block.Nonce = nonce

    return block
}

func Genesis(coinbase *transaction.Transaction) *Block {
    return CreateBlock([]*transaction.Transaction{coinbase}, []byte{})
}

func merkleRoot(txs []*transaction.Transaction) []byte {
    var txHashes [][]byte
    for _, tx := range txs {
        txHashes = append(txHashes, tx.ID)
    }
    return merkle.BuildMerkleRoot(txHashes)
}
