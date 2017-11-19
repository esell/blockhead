package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
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

	// empty out transactions
	b.CurrentTransactions = b.CurrentTransactions[:0]
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

func (b *Blockchain) ProofOfWork(lastProof int64) int64 {
	log.Println("checking proof of work...")
	proof := int64(0)
	for !b.ValidProof(lastProof, proof) {
		proof += 1
	}
	return proof

}

func (b *Blockchain) ValidProof(lastProof, newProof int64) bool {
	concatString := string(lastProof) + string(newProof)
	hash := sha256.New()
	hash.Write([]byte(concatString))
	proofTestReadable := fmt.Sprintf("%x", hash.Sum(nil))
	if string(proofTestReadable[:4]) == "0000" {
		return true
	}
	return false
}

func (b *Blockchain) LastBlock() Block {
	return b.Blocks[len(b.Blocks)-1]
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
	blockJSON, err := json.Marshal(block)
	if err != nil {
		log.Println("error with JSON: ", err)
	}
	hash := sha256.New()
	hash.Write([]byte(blockJSON))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
