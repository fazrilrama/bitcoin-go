package blockchain

import (
    "fmt"

    "bitcoin-go/internal/block"
)

type Blockchain struct {
    Blocks []*block.Block
}

func NewBlockchain() *Blockchain {
    return &Blockchain{[]*block.Block{block.Genesis()}}
}

func (bc *Blockchain) AddBlock(data string) {
    prev := bc.Blocks[len(bc.Blocks)-1]
    newBlock := block.CreateBlock(data, prev.Hash)
    bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) PrintChain() {
    for _, b := range bc.Blocks {
        fmt.Printf("\n--- Block ---\n")
        fmt.Printf("Timestamp: %d\n", b.Timestamp)
        fmt.Printf("Data: %s\n", b.Data)
        fmt.Printf("PrevHash: %x\n", b.PrevHash)
        fmt.Printf("Hash: %x\n", b.Hash)
        fmt.Printf("Nonce: %d\n", b.Nonce)
    }
}
