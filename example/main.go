package main

import (
	"context"
	"log"

	butils "github.com/crustio/blob-utils"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	rpcURL = "https://rpc-testnet.ethda.io" // change to your ethda rpc url
)

func main() {
	key, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80") // change to your private key
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	client, err := butils.New(rpcURL, key)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// post blob tx
	log.Printf("start to send blobs...")
	hash, err := client.PostBlob(context.Background(), []byte("ethda"))
	if err != nil {
		log.Fatalf("Failed to post blob: %v", err)
	}
	log.Printf("blob hash: %v", hash)

	// get blob tx
	data, err := client.GetBlob(hash)
	if err != nil {
		log.Fatalf("Failed to get blob: %v", err)
	}

	log.Printf("blob data decode result: %s", string(data))
}
