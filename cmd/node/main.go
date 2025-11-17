package main

import (
    "fmt"
    "os"

    "bitcoin-go/internal/blockchain"
)

func main() {
    bc := blockchain.NewBlockchain()

    switch os.Args[1] {
    case "add":
        data := os.Args[2]
        bc.AddBlock(data)
        fmt.Println("Block added!")
    case "print":
        bc.PrintChain()
    default:
        fmt.Println("Commands:")
        fmt.Println(" add <data>")
        fmt.Println(" print")
    }
}