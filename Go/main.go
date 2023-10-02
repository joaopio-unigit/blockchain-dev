package main

import (
	"fmt"
	"go-blockchain/blockchain"
)

func main() {

	blockchain := blockchain.CreateBlockchain()

	blockchain.AddNewBlock("FIRST BLOCK AFTER GENISIS")
	blockchain.AddNewBlock("SECOND BLOCK AFTER GENISIS")

	for _, block := range blockchain.Blocks {
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
