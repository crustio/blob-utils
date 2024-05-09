package main

import (
	"context"
	"log"
	"os"

	example "github.com/crustio/blob-utils/v2"
	"github.com/ethereum/go-ethereum/crypto"

	bu "github.com/crustio/blob-utils"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "tx",
			Usage:  "send a blob transaction",
			Action: TxApp,
			Flags:  bu.TxFlags,
		},
		{
			Name:   "download",
			Usage:  "download blobs from the beacon net",
			Action: bu.DownloadApp,
			Flags:  bu.DownloadFlags,
		},
		{
			Name:   "proof",
			Usage:  "generate kzg proof for any input point by using jth blob polynomial",
			Action: ProofApp,
			Flags:  bu.ProofFlags,
		},
		{
			Name:   "tx-simple",
			Usage:  "generate blob from any input point by using jth blob polynomial",
			Action: TxSimpleApp,
			// Flags:  bu.BlobFlags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("App failed: %v", err)
	}
}

func TxSimpleApp(ctx *cli.Context) error {
	key, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	client, err := example.New("https://rpc-devnet2.ethda.io", key)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	log.Printf("start to send blobs...")
	hash, err := client.PostBlob(context.Background(), []byte("ethda"))
	if err != nil {
		log.Fatalf("Failed to post blob: %v", err)
	}
	log.Printf("blob hash: %v", hash)
	data, err := client.GetBlob(hash)
	if err != nil {
		log.Fatalf("Failed to get blob: %v", err)
	}
	log.Printf("blob data decode result: %s", string(data))
	return nil
}

func TxApp(cliCtx *cli.Context) error {

	return nil
}

func ProofApp(cliCtx *cli.Context) error {

	return nil
}
