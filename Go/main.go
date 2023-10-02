package main

import (
	"fmt"
	"go-blockchain/blockchain"
	"strconv"
)

func main() {

	chain := blockchain.CreateBlockchain()

	chain.AddNewBlock("FIRST BLOCK AFTER GENISIS")
	chain.AddNewBlock("SECOND BLOCK AFTER GENISIS")
	chain.AddNewBlock("THIRD BLOCK AFTER GENISIS")

	for _, block := range chain.Blocks {
		fmt.Printf("\nData in Block: %s\n", block.Data)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
	}
}
