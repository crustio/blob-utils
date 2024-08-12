package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/crustio/blob-utils/zkblob"
)

var (
	l2NetworkFlag = &cli.StringFlag{
		Name:  "rpc",
		Usage: "custom L2 JSON-RPC endpoint",
		Value: "https://rpc-testnet.ethda.io",
	}

	l1NetworkFlag = &cli.StringFlag{
		Name:  "vrpc",
		Usage: "custom L1 JSON-RPC endpoint",
		Value: "https://rpc-sepolia.ethda.io",
	}

	zkblobAddressFlag = &cli.StringFlag{
		Name:  "zkblobAddress",
		Usage: "custom akblob contract address",
		Value: "0x306c92cb2127Dfbc588373c6aF466147b9b721Fc",
	}

	getRpc    string
	verifyRpc string
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "get",
			Usage:  "Get proof by hash",
			Action: getProof,
			Flags: []cli.Flag{
				l2NetworkFlag,
			},
			Before: beforeGetAction,
		},
		{
			Name:   "verify",
			Usage:  "Verify zkblob proof, <batchNum>,<hash>,<proof>",
			Action: verifyProof,
			Flags: []cli.Flag{
				l1NetworkFlag,
				zkblobAddressFlag,
			},
			Before: beforeVerifyAction,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("App failed: %v", err)
	}
}

func beforeVerifyAction(context *cli.Context) error {
	verifyRpc = context.String(l1NetworkFlag.Name)
	return nil
}

func beforeGetAction(context *cli.Context) error {
	getRpc = context.String(l2NetworkFlag.Name)
	return nil
}

func getProof(cliCtx *cli.Context) error {
	if cliCtx.NArg() == 0 {
		return errors.New("no hash provided")
	}

	hash := cliCtx.Args().First()

	fmt.Println("get proof from: ", getRpc)
	proof, err := zkblob.GetProofByHash(getRpc, common.HexToHash(hash))
	if err != nil {
		return err
	}

	fmt.Println(proof)

	return nil
}

func verifyProof(cliCtx *cli.Context) error {
	zkblobAddress := cliCtx.String("zkblobAddress")
	if zkblobAddress == "" {
		return errors.New("no zkblob contract address provided")
	}

	if cliCtx.NArg() < 3 {
		return errors.New("not enough arguments")
	}
	batchNumStr := cliCtx.Args().Get(0)
	batchNum, err := strconv.ParseUint(batchNumStr, 10, 64)
	if err != nil {
		return err
	}

	rest := cliCtx.Args().Tail()
	hash := rest[0]
	proofStrSlice := rest[1:]

	proof := make([][32]byte, len(proofStrSlice))
	for i, proofStr := range proofStrSlice {
		proof[i] = common.HexToHash(proofStr)
	}

	fmt.Println("------------------------------------------------------------------------")

	fmt.Printf("batchNum : %d\n", batchNum)
	fmt.Printf("hash     : %s\n", hash)
	fmt.Printf("proof    : [%s]\n", strings.Join(proofStrSlice, ","))

	fmt.Println("------------------------------------------------------------------------")

	bv, err := zkblob.NewBlobVerify(verifyRpc, common.HexToAddress(zkblobAddress))
	if err != nil {
		return err
	}

	ret, err := bv.Verify(int64(batchNum), common.HexToHash(hash), proof)
	if err != nil {
		return err
	}

	fmt.Println("verify result: ", ret)
	return nil
}
