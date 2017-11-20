package main

import (
	"encoding/json"
	"log"
)

func main() {

	myChain := Blockchain{}
	log.Println(myChain)

	log.Println("creating genesis block...")
	genBlock, err := myChain.NewBlock(1, 100, true)
	if err != nil {
		log.Println(err)
	}
	genBlockJSON, err := json.Marshal(genBlock)
	if err != nil {
		log.Println("error with JSON: ", err)
	}
	myChain.AddBlock(genBlock)
	log.Println("genesis block: ", string(genBlockJSON))
	returnChain(&myChain)
	returnLastBlock(&myChain)

	// create new transaction
	tempTrans := myChain.NewTransaction("blahfrom", "blahto", 666)
	log.Println("returned index: ", tempTrans)
	returnChain(&myChain)
	returnLastBlock(&myChain)

	for i := 0; i < 4; i++ {
		// start mining
		log.Println("starting to mine...")
		lastProof := myChain.LastBlock()
		newProof := myChain.ProofOfWork(lastProof.Proof)
		log.Println("new proof: ", newProof)

		// add new block
		myNewBlock, err := myChain.NewBlock(myChain.LastBlock().Index+1, newProof, false)
		if err != nil {
			log.Println(err)
			break
		}
		myChain.AddBlock(myNewBlock)
	}
	// print new chain
	returnChain(&myChain)
	returnLastBlock(&myChain)
}

func returnChain(chain *Blockchain) {
	log.Println()
	log.Println("****************************************")
	chainJSON, err := json.Marshal(chain)
	if err != nil {
		log.Println("error with JSON: ", err)
	}
	log.Println("full chain: ", string(chainJSON))
	log.Println("****************************************")
	log.Println()
}

func returnLastBlock(chain *Blockchain) {
	log.Println()
	log.Println("****************************************")
	chainJSON, err := json.Marshal(chain.LastBlock())
	if err != nil {
		log.Println("error with JSON: ", err)
	}
	log.Println("last block: ", string(chainJSON))
	log.Println("****************************************")
	log.Println()
}
