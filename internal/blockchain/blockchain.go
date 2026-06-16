package blockchain

import (
	"fmt"

	"github.com/fazrilrama/bitcoin-go/internal/block"
	"github.com/fazrilrama/bitcoin-go/internal/transaction"
)

type Blockchain struct {
	Blocks []*block.Block
}

func NewBlockchain(address string) *Blockchain {
	coinbase := transaction.NewCoinbaseTX(address, "")
	return &Blockchain{[]*block.Block{block.Genesis(coinbase)}}
}

func (bc *Blockchain) AddBlock(txs []*transaction.Transaction) {
	prev := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.CreateBlock(txs, prev.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) AllTransactions() []*transaction.Transaction {
	var txs []*transaction.Transaction
	for _, b := range bc.Blocks {
		txs = append(txs, b.Transactions...)
	}
	return txs
}

func (bc *Blockchain) PrintChain() {
	for _, b := range bc.Blocks {
		fmt.Printf("\n--- Block ---\n")
		fmt.Printf("Timestamp:  %d\n", b.Timestamp)
		fmt.Printf("PrevHash:   %x\n", b.PrevHash)
		fmt.Printf("Hash:       %x\n", b.Hash)
		fmt.Printf("MerkleRoot: %x\n", b.MerkleRoot)
		fmt.Printf("Nonce:      %d\n", b.Nonce)
		fmt.Printf("Txs:        %d\n", len(b.Transactions))
	}
}
