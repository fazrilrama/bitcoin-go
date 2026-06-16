package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fazrilrama/bitcoin-go/internal/blockchain"
	"github.com/fazrilrama/bitcoin-go/internal/transaction"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "createblockchain":
		address := os.Args[2]
		bc := blockchain.NewBlockchain(address)
		bc.PrintChain()
		fmt.Println("Blockchain created, reward sent to:", address)

	case "add":
		address := os.Args[2]
		bc := blockchain.NewBlockchain(address)
		tx := transaction.NewCoinbaseTX(address, os.Args[3])
		bc.AddBlock([]*transaction.Transaction{tx})
		fmt.Println("Block added!")

	case "print":
		bc := blockchain.NewBlockchain("genesis")
		bc.PrintChain()

	case "wallet":
		w := transaction.NewWallet()
		fmt.Printf("PublicKey:  %x\n", w.PublicKey)

	case "balance":
		address := os.Args[2]
		bc := blockchain.NewBlockchain(address)
		utxos := transaction.FindUTXO(address, bc.AllTransactions())
		balance := 0
		for _, out := range utxos {
			balance += out.Value
		}
		fmt.Printf("Balance of %s: %d\n", address, balance)

	case "send":
		from, to, amount := os.Args[2], os.Args[3], mustAtoi(os.Args[4])
		w := transaction.NewWallet()
		bc := blockchain.NewBlockchain(from)
		allTxs := bc.AllTransactions()

		acc, spendable := transaction.FindSpendableOutputs(from, amount, allTxs)
		if acc < amount {
			fmt.Println("Not enough funds")
			return
		}

		var inputs []transaction.TXInput
		for txID, outs := range spendable {
			for _, outIdx := range outs {
				inputs = append(inputs, transaction.TXInput{
					Txid:      []byte(txID),
					OutIndex:  outIdx,
					PubKey:    w.PublicKey,
				})
			}
		}
		outputs := []transaction.TXOutput{{Value: amount, PubKeyHash: []byte(to)}}
		if acc > amount {
			outputs = append(outputs, transaction.TXOutput{Value: acc - amount, PubKeyHash: []byte(from)})
		}

		tx := &transaction.Transaction{Vin: inputs, Vout: outputs}
		tx.SetID()
		tx.Sign(w.PrivateKey)

		bc.AddBlock([]*transaction.Transaction{tx})
		fmt.Printf("Sent %d from %s to %s\n", amount, from, to)

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Commands:")
	fmt.Println("  createblockchain <address>")
	fmt.Println("  add <address> <data>")
	fmt.Println("  print")
	fmt.Println("  wallet")
	fmt.Println("  balance <address>")
	fmt.Println("  send <from> <to> <amount>")
}

func mustAtoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
