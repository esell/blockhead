package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
