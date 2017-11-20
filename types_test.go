package main

import (
	"testing"
)

func TestGenesisNewBlock(t *testing.T) {
	testChain := Blockchain{}
	testBlock, err := testChain.NewBlock(1, 100, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if testBlock.Index != 1 {
		t.Errorf("Block index is %v, should be %v\n", testBlock.Index, 1)
	}
	if testBlock.Proof != 100 {
		t.Errorf("Block proof is %v, should be %v\n", testBlock.Proof, 100)
	}
	if len(testBlock.Transactions) > 0 {
		t.Errorf("Block transaction count is %v, should be %v\n", len(testBlock.Transactions), 0)
	}
}

func TestNewBlockNoTrans(t *testing.T) {
	testChain := Blockchain{}
	// this should fail since there are no transactions
	_, err := testChain.NewBlock(2, 200, false)
	if err == nil {
		t.Errorf("No error returned despite having zero transactions")
	}
}

func TestNewBlock(t *testing.T) {
	testChain := Blockchain{}
	// create genesis block, needed for testing
	testBlock, err := testChain.NewBlock(1, 100, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if testBlock.Index != 1 {
		t.Errorf("Genesis Block index is %v, should be %v\n", testBlock.Index, 1)
	}
	if testBlock.Proof != 100 {
		t.Errorf("Genesis Block proof is %v, should be %v\n", testBlock.Proof, 100)
	}
	if len(testBlock.Transactions) > 0 {
		t.Errorf("Genesis Block transaction count is %v, should be %v\n", len(testBlock.Transactions), 0)
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
		t.Errorf("Error creating new block: ", err)
	}
	if testBlock2.Index != 2 {
		t.Errorf("Block index is %v, should be %v\n", testBlock2.Index, 2)
	}
	if testBlock2.Proof != 200 {
		t.Errorf("Block proof is %v, should be %v\n", testBlock2.Proof, 200)
	}

	if len(testBlock2.Transactions) != 1 {
		t.Errorf("Block transaction count is %v, should be %v\n", len(testBlock2.Transactions), 1)
	}
}

func TestNewTransaction(t *testing.T) {
	testChain := Blockchain{}
	// create genesis block, needed for testing
	testBlock, err := testChain.NewBlock(1, 100, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if testBlock.Index != 1 {
		t.Errorf("Genesis Block index is %v, should be %v\n", testBlock.Index, 1)
	}
	if testBlock.Proof != 100 {
		t.Errorf("Genesis Block proof is %v, should be %v\n", testBlock.Proof, 100)
	}
	if len(testBlock.Transactions) > 0 {
		t.Errorf("Genesis Block transaction count is %v, should be %v\n", len(testBlock.Transactions), 0)
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
