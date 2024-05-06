package blobutils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
	"log"
	"math/big"
	"testing"
	"time"
)

func makeBlob(siz int) []byte {
	b := make([]byte, siz)
	for i := range b {
		b[i] = byte(i)
	}
	return b
}

func TestBlobTx(t *testing.T) {
	var cid int64 = 1001
	addr := "https://rpc-devnet2.ethda.io"
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, addr)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	key, err := crypto.HexToECDSA("")
	require.Nil(t, err)

	nonce, err := client.NonceAt(ctx, crypto.PubkeyToAddress(key.PublicKey), big.NewInt(-1))
	if err != nil {
		log.Fatalf("Error getting nonce: %v", err)
	}
	//nonce := int64(40)

	var gasPrice256 *uint256.Int
	val, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("Error getting suggested gas price: %v", err)
	}
	var nok bool
	gasPrice256, nok = uint256.FromBig(val)
	if nok {
		log.Fatalf("gas price is too high! got %v", val.String())
	}

	priorityGasPrice256 := gasPrice256
	maxFeePerBlobGas := "3000000000"
	maxFeePerBlobGas256, err := DecodeUint256String(maxFeePerBlobGas)
	require.Nil(t, err)

	var a []byte
	for i := 0; i < 10; i++ {
		a = append(a, '2')
	}
	blobs, commitments, proofs, versionedHashes, err := EncodeBlobs(a)
	if err != nil {
		log.Fatalf("failed to compute commitments: %v", err)
	}

	value256, _ := uint256.FromBig(big.NewInt(1000000000000000000))

	to := common.HexToAddress("0x")
	chainId := big.NewInt(cid)

	signer := types.NewEIP155Signer(chainId)

	gasLimit := uint64(21000)
	legacyTx := types.NewTransaction(uint64(nonce), to, value256.ToBig(), gasLimit, gasPrice256.ToBig(), nil)
	legacySignedTx, err := types.SignTx(legacyTx, signer, key)
	fmt.Println("legacy sign hash: ", signer.Hash(legacyTx).Hex())
	require.Nil(t, err)

	v, r, s := legacySignedTx.RawSignatureValues()
	from, err := types.Sender(signer, legacySignedTx)
	require.Nil(t, err)
	fmt.Println("from: ", from.Hex())
	fmt.Println("legacy hash: ", legacySignedTx.Hash().Hex())

	ff, err := json.Marshal(legacySignedTx)
	require.Nil(t, err)
	fmt.Printf("%+v, %s\n", string(ff), from)

	blobTx := &types.BlobTx{
		ChainID:    uint256.MustFromBig(chainId),
		Nonce:      uint64(nonce),
		GasTipCap:  priorityGasPrice256,
		GasFeeCap:  gasPrice256,
		Gas:        gasLimit,
		To:         to,
		Value:      value256,
		BlobFeeCap: maxFeePerBlobGas256,
		BlobHashes: versionedHashes,
		Sidecar:    &types.BlobTxSidecar{Blobs: blobs, Commitments: commitments, Proofs: proofs},
		V:          uint256.NewInt(v.Uint64() - uint64(cid)*2 - 35),
		R:          uint256.MustFromBig(r),
		S:          uint256.MustFromBig(s),
	}
	tx := types.NewTx(blobTx)

	fmt.Println("blob tx hash:", tx.Hash())
	fmt.Println("blob tx type:", tx.Type())
	fmt.Println(tx.RawSignatureValues())
	//fff, err := tx.MarshalJSON()
	//require.Nil(t, err)
	//fmt.Printf("send: %+v\n", string(fff))

	err = client.SendTransaction(context.Background(), tx)

	if err != nil {
		log.Fatalf("failed to send transaction: %v", err)
	} else {
		log.Printf("successfully sent transaction. txhash=%v", legacySignedTx.Hash())
	}

	//var receipt *types.Receipt
	for {
		_, err = client.TransactionReceipt(context.Background(), legacySignedTx.Hash())
		if err == ethereum.NotFound {
			time.Sleep(1 * time.Second)
			fmt.Println(222222)
		} else if err != nil {
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				// TODO: ignore other errors for now. Some clients are treating the blobGasUsed as big.Int rather than uint64
				break
			}
		} else {
			break
		}
	}

	log.Printf("Transaction included. nonce=%d hash=%v", nonce, legacySignedTx.Hash())
}

func TestDataBlob(t *testing.T) {
	addr := "https://rpc-devnet.ethda.io"
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, addr)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	key, err := crypto.HexToECDSA("")
	require.Nil(t, err)

	pendingNonce, err := client.NonceAt(ctx, crypto.PubkeyToAddress(key.PublicKey), big.NewInt(-1))
	if err != nil {
		log.Fatalf("Error getting nonce: %v", err)
	}
	fmt.Println(pendingNonce)
	nonce := int64(40)

	var gasPrice256 *uint256.Int
	val, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("Error getting suggested gas price: %v", err)
	}
	var nok bool
	gasPrice256, nok = uint256.FromBig(val)
	if nok {
		log.Fatalf("gas price is too high! got %v", val.String())
	}

	require.Nil(t, err)

	var a []byte
	for i := 0; i < 10; i++ {
		a = append(a, '2')
	}
	//blobs, _, _, _, err := EncodeBlobs(a)
	//if err != nil {
	//	log.Fatalf("failed to compute commitments: %v", err)
	//}

	value256, _ := uint256.FromBig(big.NewInt(0))

	to := common.HexToAddress("0x")
	chainId := big.NewInt(177)

	signer := types.NewEIP155Signer(chainId)

	gasLimit := uint64(21000)
	legacyTx := types.NewTransaction(uint64(nonce), to, value256.ToBig(), gasLimit, gasPrice256.ToBig(), []byte("2"))
	legacySignedTx, err := types.SignTx(legacyTx, signer, key)
	require.Nil(t, err)

	from, err := types.Sender(signer, legacySignedTx)
	require.Nil(t, err)

	ff, err := json.Marshal(legacySignedTx)
	require.Nil(t, err)
	fmt.Printf("%+v, %s\n", string(ff), from)

	err = client.SendTransaction(context.Background(), legacyTx)

	if err != nil {
		log.Fatalf("failed to send transaction: %v", err)
	} else {
		log.Printf("successfully sent transaction. txhash=%v", legacySignedTx.Hash())
	}

	//var receipt *types.Receipt
	for {
		_, err = client.TransactionReceipt(context.Background(), legacySignedTx.Hash())
		if err == ethereum.NotFound {
			time.Sleep(1 * time.Second)
			fmt.Println(222222)
		} else if err != nil {
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				// TODO: ignore other errors for now. Some clients are treating the blobGasUsed as big.Int rather than uint64
				break
			}
		} else {
			break
		}
	}

	log.Printf("Transaction included. nonce=%d hash=%v", nonce, legacySignedTx.Hash())
}

/*
func TestBlobCodec(t *testing.T) {
	blobs := [][]byte{
		makeBlob(0),
		makeBlob(5),
		makeBlob(95),
		makeBlob(params.FieldElementsPerBlob * 31),
	}
	for i, blob := range blobs {
		encB := EncodeBlobs(blob)
		if len(encB) != 1 {
			t.Fatalf("(%d) expected 1 blob, got %d", i, len(encB))
		}
		enc := encB[0]
		dec := DecodeBlob(enc[:])
		if len(dec) != len(blob) {
			t.Fatalf("(%d) mismatched lengths: expected %d, got %d", i, len(blob), len(dec))
		}
		if !bytes.Equal(blob, dec) {
			t.Fatalf("(%d) expected %x, got %x", i, blob, dec)
		}
	}
}

func TestBlobsCodec(t *testing.T) {
	blob := makeBlob(params.FieldElementsPerBlob*32 + 10)
	encB := EncodeBlobs(blob)
	if len(encB) != 2 {
		t.Fatal("expected 2 blobs, got", len(encB))
	}
	dec1 := DecodeBlob(encB[0][:])
	dec2 := DecodeBlob(encB[1][:])
	dec := append(dec1, dec2...)
	if len(dec) != len(blob) {
		t.Fatalf("mismatched lengths: expected %d, got %d", len(blob), len(dec))
	}
	if !bytes.Equal(blob, dec) {
		t.Fatalf("expected %x, got %x", blob, dec)
	}
}
*/

func TestEIP155Signing(t *testing.T) {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)

	to := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	fmt.Println(addr.Hex())
	signer := types.NewEIP155Signer(big.NewInt(18))
	tx, err := types.SignTx(types.NewTransaction(0, to, new(big.Int), 0, new(big.Int), nil), signer, key)
	if err != nil {
		t.Fatal(err)
	}
	dd, err := tx.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", string(dd))

	from, err := types.Sender(signer, tx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(from)
	if from != addr {
		t.Errorf("exected from and address to be equal. Got %x want %x", from, addr)
	}
}

func TestGetBlock(t *testing.T) {
	addr := "https://rpc-devnet.ethda.io"
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, addr)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	block, err := client.BlockByNumber(context.Background(), big.NewInt(43))
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	jj, err := json.Marshal(block)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	fmt.Println(string(jj))

	fmt.Printf("%s", block.Hash().Hex())
}
