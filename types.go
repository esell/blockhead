package main

import (
	"crypto/sha256"
	"time"
)

type Blockchain struct {
	CurrentTransactions []Transaction
	Blocks              []Block
}

func (b *Blockchain) NewBlock(index, proof int64, isGenesis bool) Block {
	block := Block{}
	block.Index = index
	if !isGenesis {
		block.PreviousHash = HashBlock(b.Blocks[len(b.Blocks)-1])
	}
	block.Proof = proof
	block.Timestamp = time.Now().UnixNano()
	block.Transactions = b.CurrentTransactions

	return block
}

func (b *Blockchain) AddBlock(block Block) {
	b.Blocks = append(b.Blocks, block)
}

func (b *Blockchain) NewTransaction(from, to string, amount int64) int64 {
	transaction := Transaction{Recipient: to, Sender: from, Amount: amount}
	b.CurrentTransactions = append(b.CurrentTransactions, transaction)

	return b.LastBlock().Index
}

func (b *Blockchain) LastBlock() Block {
	return b.Blocks[len(b.Blocks)]
}

type Block struct {
	Index        int64         `json:"index"`
	PreviousHash string        `json:"previous_hash"`
	Proof        int64         `json:"proof"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Amount    int64  `json:"amount"`
	Recipient string `json:"recipient"`
	Sender    string `json:"sender"`
}

func HashBlock(block Block) string {
	hash := sha256.Sum256([]byte("hello world\n"))
	return string(hash[:32])
}
