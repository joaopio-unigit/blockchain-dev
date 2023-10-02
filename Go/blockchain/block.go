package blockchain

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}

	//Run POW
	pow := NewProof(block)
	nonce, hash := pow.Run()
	//fmt.Printf("\nNONCE: %d\n", nonce)

	block.Nonce = nonce
	block.Hash = hash[:]

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
