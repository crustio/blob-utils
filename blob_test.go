package blobutils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSignBatchHash(t *testing.T) {
	key, err := crypto.HexToECDSA("f26d6aca18e0c75ac948a262d4b9435a8173515f84c258d1f90d171143039024")
	require.Nil(t, err)
	cli, err := New("https://rpc-devnet2.ethda.io", key)
	require.Nil(t, err)

	res, err := cli.SignBatchHash(common.HexToHash("76b56f65beab1cfe55242eb97eebfe2f32aacd957f202122835026bffb3ba282"))
	require.Nil(t, err)
	fmt.Println(common.BytesToHash(res))
}

func TestPostBlobReturnSign(t *testing.T) {
	key, err := crypto.HexToECDSA("f26d6aca18e0c75ac948a262d4b9435a8173515f84c258d1f90d171143039024")
	require.Nil(t, err)
	cli, err := New("https://rpc-devnet2.ethda.io", key)
	require.Nil(t, err)

	sig, hash, err := cli.PostBlobReturnSign(context.Background(), []byte("ethda"))
	require.Nil(t, err)
	fmt.Println(common.BytesToHash(sig))
	fmt.Println(hash.Hex())
}
