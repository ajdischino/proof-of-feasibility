package main

import (
	"context"
	"log"

	"github.com/gagliardetto/solana-go/rpc"
)

// test
/*
run go mod tody to import needed dependencies (like solana-go)
*/
func main() {
	node := rpc.New("https://api.mainnet-beta.solana.com")

	blockHash, err := node.GetLatestBlockhash(context.Background(), rpc.CommitmentProcessed)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("current block hash:", blockHash.Value.Blockhash.String())
}
