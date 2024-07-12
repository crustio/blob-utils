package main

import (
	"context"
	"log"

	butils "github.com/crustio/blob-utils"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	l2rpcURL = "https://rpc-testnet.ethda.io" // change to your ethda rpc l2rpcURL
)

func main() {
	key, err := crypto.HexToECDSA("privateKey") // change to your private key
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	client, err := butils.New(l2rpcURL, key)
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
