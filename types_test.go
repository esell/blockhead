package main

import (
	"strings"
	"testing"
)

func TestHashBlock(t *testing.T) {
	staticHash := strings.ToLower("22E6E24D132B3E42A25F0571D7B3DD2232197EFE14085219D05420B439AA7B73")
	testBlockHeader := BlockHeader{Index: 1, Proof: 100, Timestamp: 1511295968402564000}
	testBlock := Block{Header: testBlockHeader}
	testBlock.HashBlock()
	if testBlock.Header.Hash != staticHash {
		t.Errorf("Block hash is %v, should be %v", testBlock.Header.Hash, staticHash)
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

func TestOverwrite(t *testing.T) {
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

	// Transaction #1
	tempTrans := testChain.NewTransaction("blahfrom", "blahto", 666)
	if tempTrans != 1 {
		t.Errorf("Transaction index is %v, should be %v", tempTrans, 1)
	}
	if len(testChain.CurrentTransactions) != 1 {
		t.Errorf("Transaction count on chain is %v, should be %v", len(testChain.CurrentTransactions), 1)
	}

	// mine the transaction
	lastProof := testChain.LastBlock()
	newProof := testChain.ProofOfWork(lastProof.Header.Proof)
	// add Block #1
	newBlock, err := testChain.NewBlock(testChain.LastBlock().Header.Index+1, newProof, false)
	if err != nil {
		t.Errorf("Error adding new block: %v", err)
	}
	testChain.AddBlock(newBlock)
	blockOneTrans := testChain.Blocks[1].Transactions[0].(Transaction)
	if blockOneTrans.Sender != "blahfrom" {
		t.Errorf("Transaction sender is %v, should be %v", blockOneTrans.Sender, "blahfrom")
	}
	if len(testChain.CurrentTransactions) != 0 {
		t.Errorf("Transaction count on chain is %v, should be %v", len(testChain.CurrentTransactions), 0)
	}

	// Transaction #2
	tempTransTwo := testChain.NewTransaction("blahfromtwo", "blahtotwo", 999)
	if tempTransTwo != 2 {
		t.Errorf("Transaction index is %v, should be %v", tempTransTwo, 2)
	}
	if len(testChain.CurrentTransactions) != 1 {
		t.Errorf("Transaction count on chain is %v, should be %v", len(testChain.CurrentTransactions), 1)
	}
	// mine the transaction

	lastProofTwo := testChain.LastBlock()
	newProofTwo := testChain.ProofOfWork(lastProofTwo.Header.Proof)
	// add Block #2
	newBlockTwo, err := testChain.NewBlock(testChain.LastBlock().Header.Index+1, newProofTwo, false)
	if err != nil {
		t.Errorf("Error adding new block: %v", err)
	}
	testChain.AddBlock(newBlockTwo)
	blockTwoTrans := testChain.Blocks[2].Transactions[0].(Transaction)
	if blockTwoTrans.Sender != "blahfromtwo" {
		t.Errorf("Transaction sender is %v, should be %v", blockTwoTrans.Sender, "blahfromtwo")
	}
	if len(testChain.CurrentTransactions) != 0 {
		t.Errorf("Transaction count on chain is %v, should be %v", len(testChain.CurrentTransactions), 0)
	}

	// did block #1 stay the same?
	blockOneTransNew := testChain.Blocks[1].Transactions[0].(Transaction)
	if blockOneTransNew.Sender != "blahfrom" {
		t.Errorf("Transaction sender is %v, should be %v", blockOneTransNew.Sender, "blahfrom")
	}

	if len(testChain.CurrentTransactions) != 0 {
		t.Errorf("Transaction count on chain is %v, should be %v", len(testChain.CurrentTransactions), 0)
	}

}
