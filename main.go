package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {

	myChain := Blockchain{}

	log.Println("creating genesis block...")
	genBlock, err := myChain.NewBlock(1, 100, true)
	if err != nil {
		log.Fatal("unable go create genesis block: ", err)
	}
	myChain.AddBlock(genBlock)
	log.Println("genesis block created, have fun!")

	http.Handle("/list", blockListHandler(myChain))
	http.Handle("/mine", mineHandler(myChain))
	http.Handle("/newTransaction", newTransactionHandler(myChain))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func blockListHandler(b Blockchain) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
			return
		}
		chainJSON, err := json.Marshal(b)
		if err != nil {
			log.Println("error with JSON: ", err)
		}
		log.Println("full chain: ", string(chainJSON))
		w.Write(chainJSON)
	})
}

func mineHandler(b Blockchain) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
			return
		}
		// start mining
		log.Println("starting to mine...")
		lastProof := b.LastBlock()
		newProof := b.ProofOfWork(lastProof.Header.Proof)
		log.Println("new proof: ", newProof)

		// add new block
		newBlock, err := b.NewBlock(b.LastBlock().Header.Index+1, newProof, false)
		if err != nil {
			log.Println(err)
		}
		b.AddBlock(newBlock)
		blockJSON, err := json.Marshal(newBlock)
		if err != nil {
			log.Println("error with JSON: ", err)
		}
		log.Println("new block: ", string(blockJSON))
		w.Write(blockJSON)
	})
}

func newTransactionHandler(b Blockchain) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
			return
		}
		to := r.URL.Query().Get("to")
		from := r.URL.Query().Get("from")
		amount := r.URL.Query().Get("amount")
		amountInt, err := strconv.ParseInt(amount, 10, 32)
		if err != nil {
			log.Println("error converting string -> int: ", err)
		}
		newTrans := b.NewTransaction(from, to, amountInt)
		transJSON, err := json.Marshal(newTrans)
		if err != nil {
			log.Println("error with JSON: ", err)
		}
		log.Println("new transaction: ", string(transJSON))
		w.Write(transJSON)
	})
}
