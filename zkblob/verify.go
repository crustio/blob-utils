package zkblob

import (
	"context"
	"fmt"
	"math/big"

	"github.com/crustio/blob-utils/zkblob/contracts/zkblob"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlobVerify struct {
	l1Host string

	client *ethclient.Client
	zkBlob *zkblob.ZkBlob
}

func NewBlobVerify(l1Host string, zkBlobAddress common.Address) (*BlobVerify, error) {

	client, err := ethclient.DialContext(context.Background(), l1Host)
	if err != nil {
		return nil, fmt.Errorf("connect to l1: %w", err)
	}

	zkBlob, err := zkblob.NewZkBlob(zkBlobAddress, client)
	if err != nil {
		return nil, err
	}

	return &BlobVerify{client: client, zkBlob: zkBlob, l1Host: l1Host}, nil
}

func (v *BlobVerify) Verify(batchNumber int64, hash common.Hash, proof [][32]byte) (bool, error) {
	ok, err := v.zkBlob.VerifyBlob(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, big.NewInt(batchNumber), hash, proof)
	if err != nil {
		return false, err
	}

	return ok, nil
}
