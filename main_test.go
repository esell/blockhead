package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestBlockListHandler(t *testing.T) {
	// create dummy chain
	myChain := Blockchain{}
	genBlock, err := myChain.NewBlock(1, 100, true)
	if err != nil {
		t.Fatal("unable go create genesis block: ", err)
	}
	myChain.AddBlock(genBlock)

	listHandler := blockListHandler(&myChain)
	// GET
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	listHandler.ServeHTTP(w, req)
	responseBody, _ := ioutil.ReadAll(w.Result().Body)
	var returnChain Blockchain
	json.Unmarshal(responseBody, &returnChain)
	if returnChain.Blocks[0].Header.Index != genBlock.Header.Index {
		t.Errorf("blockListHandler index returned %v, should be %v", returnChain.Blocks[0].Header.Index, genBlock.Header.Index)
	}
	if returnChain.Blocks[0].Header.Proof != genBlock.Header.Proof {
		t.Errorf("blockListHandler proof returned %v, should be %v", returnChain.Blocks[0].Header.Proof, genBlock.Header.Proof)
	}
}

func TestNewTransactionHandler(t *testing.T) {
	// create dummy chain
	myChain := Blockchain{}
	genBlock, err := myChain.NewBlock(1, 100, true)
	if err != nil {
		t.Fatal("unable go create genesis block: ", err)
	}
	myChain.AddBlock(genBlock)

	newTransHandler := newTransactionHandler(&myChain)
	// POST
	req, _ := http.NewRequest("POST", "", nil)
	req.PostForm = url.Values{"to": {"testto"}, "from": {"testfrom"}, "amount": {"123"}}

	w := httptest.NewRecorder()
	newTransHandler.ServeHTTP(w, req)
	if len(myChain.CurrentTransactions) != 1 {
		t.Errorf("newTransactionHandler returned %v, should be %v", len(myChain.CurrentTransactions), 1)
	}
	parsedTrans := myChain.CurrentTransactions[0].(Transaction)
	if parsedTrans.Recipient != "testto" {
		t.Errorf("newTransactionHandler transaction returned %v, should be %v", parsedTrans.Recipient, "testto")
	}
	if parsedTrans.Sender != "testfrom" {
		t.Errorf("newTransactionHandler transaction returned %v, should be %v", parsedTrans.Sender, "testfrom")
	}
	if parsedTrans.Amount != 123 {
		t.Errorf("newTransactionHandler transaction returned %v, should be %v", parsedTrans.Amount, 123)
	}
}
