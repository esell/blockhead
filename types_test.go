package main

import (
	"strings"
	"testing"
)

func TestHashBlock(t *testing.T) {
	staticHash := strings.ToLower("4C5DE129BB39267AADEE63781962D3ADA87D77DA2791E1AA1EA4F535D19BEF2A")
	testBlockHeader := BlockHeader{Index: 1, Proof: 100, Timestamp: 1511295968402564000}
	testBlock := Block{Header: testBlockHeader}
	testBlockHash := HashBlock(testBlock)
	if testBlockHash != staticHash {
		t.Errorf("Block hash is %v, should be %v", testBlockHash, staticHash)
	}
}

func TestGenesisNewBlock(t *testing.T) {
	testChain := Blockchain{}
	testBlock, err := testChain.NewBlock(1, 100, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if testBlock.Header.Index != 1 {
		t.Errorf("Block index is %v, should be %v\n", testBlock.Header.Index, 1)
	}
	if testBlock.Header.Proof != 100 {
		t.Errorf("Block proof is %v, should be %v\n", testBlock.Header.Proof, 100)
	}
}

func TestNewBlockNoTrans(t *testing.T) {
	testChain := Blockchain{}
	// this should fail since there are no transactions
	_, err := testChain.NewBlock(2, 200, false)
	if err == nil {
		t.Error("No error returned despite having zero transactions")
	}
}

func TestNewBlock(t *testing.T) {
	testChain := Blockchain{}
	// create genesis block, needed for testing
	testBlock, err := testChain.NewBlock(1, 100, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if testBlock.Header.Index != 1 {
		t.Errorf("Genesis Block index is %v, should be %v\n", testBlock.Header.Index, 1)
	}
	if testBlock.Header.Proof != 100 {
		t.Errorf("Genesis Block proof is %v, should be %v\n", testBlock.Header.Proof, 100)
	}

	// need to add this block in order to continue tests
	// assumes AddBlock() works :)
	testChain.AddBlock(testBlock)

	// add a dummy transaction, again, assuming NewTransaction() works
	tempTrans := testChain.NewTransaction("blahfrom", "blahto", 666)
	if tempTrans != 1 {
		t.Errorf("Transaction index is %v, should be %v", tempTrans, 1)
	}

	// add new block that we actually want to test
	// hard code index and proof for simplicity
	testBlock2, err := testChain.NewBlock(2, 200, false)
	if err != nil {
		t.Errorf("Error creating new block: %v", err)
	}
	if testBlock2.Header.Index != 2 {
		t.Errorf("Block index is %v, should be %v\n", testBlock2.Header.Index, 2)
	}
	if testBlock2.Header.Proof != 200 {
		t.Errorf("Block proof is %v, should be %v\n", testBlock2.Header.Proof, 200)
	}
	// merkle of raw block transactions
	//tree, _ := merkletree.NewTree(testChain.CurrentTransactions)
	//_ = tree.MerkleRoot()

}

func TestNewTransaction(t *testing.T) {
	testChain := Blockchain{}
	// create genesis block, needed for testing
	testBlock, err := testChain.NewBlock(1, 100, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if testBlock.Header.Index != 1 {
		t.Errorf("Genesis Block index is %v, should be %v\n", testBlock.Header.Index, 1)
	}
	if testBlock.Header.Proof != 100 {
		t.Errorf("Genesis Block proof is %v, should be %v\n", testBlock.Header.Proof, 100)
	}
	// need to add this block in order to continue tests
	// assumes AddBlock() works :)
	testChain.AddBlock(testBlock)

	tempTrans := testChain.NewTransaction("blahfrom", "blahto", 666)
	if tempTrans != 1 {
		t.Errorf("Transaction index is %v, should be %v", tempTrans, 1)
	}
	if len(testChain.CurrentTransactions) != 1 {
		t.Errorf("Transaction count on chain is %v, should be %v", len(testChain.CurrentTransactions), 1)
	}
}
