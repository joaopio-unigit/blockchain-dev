package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) GetHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	newHash := sha256.Sum256(info)
	b.Hash = newHash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.GetHash()
	return block
}

func (chain *Blockchain) AddNewBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func CreateGenesisBlock() *Block {
	return CreateBlock("GENESIS", []byte{})
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesisBlock()}}
}
