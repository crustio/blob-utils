package blobutils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	"log"
	"math/big"
	"net/http"
	"time"
)

type Client struct {
	chainId *big.Int
	rpcUrl  string
	privKey *ecdsa.PrivateKey
	client  *ethclient.Client
}

func New(
	rpcUrl string,
	privKey *ecdsa.PrivateKey) (*Client, error) {
	if rpcUrl == "" {
		return nil, fmt.Errorf("empty ethda rpc url")
	}

	if privKey == nil {
		return nil, fmt.Errorf("empty private key")
	}

	client, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("connect to ethda: %w", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get chain id: %w", err)
	}

	return &Client{
		chainId: chainId,
		rpcUrl:  rpcUrl,
		privKey: privKey,
		client:  client,
	}, nil
}

func (cli *Client) PostBlob(ctx context.Context, data []byte) (common.Hash, error) {
	nonce, err := cli.client.NonceAt(ctx, crypto.PubkeyToAddress(cli.privKey.PublicKey), big.NewInt(-1))
	if err != nil {
		return common.Hash{}, err
	}

	var gasPrice256 *uint256.Int
	val, err := cli.client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	var nok bool
	gasPrice256, nok = uint256.FromBig(val)
	if nok {
		return common.Hash{}, fmt.Errorf("gas price is too high! got %v", val.String())
	}

	priorityGasPrice256 := gasPrice256
	maxFeePerBlobGas := "3000000000"
	maxFeePerBlobGas256, err := DecodeUint256String(maxFeePerBlobGas)
	if err != nil {
		return common.Hash{}, err
	}

	blobs, commitments, proofs, versionedHashes, err := EncodeBlobs(data)
	if err != nil {
		return common.Hash{}, err
	}

	value256, _ := uint256.FromBig(big.NewInt(10000000000000000))

	to := common.HexToAddress("0x")

	signer := types.NewEIP155Signer(cli.chainId)

	gasLimit := uint64(21000)
	legacyTx := types.NewTransaction(uint64(nonce), to, value256.ToBig(), gasLimit, gasPrice256.ToBig(), nil)
	legacySignedTx, err := types.SignTx(legacyTx, signer, cli.privKey)

	v, r, s := legacySignedTx.RawSignatureValues()

	blobTx := &types.BlobTx{
		ChainID:    uint256.MustFromBig(cli.chainId),
		Nonce:      uint64(nonce),
		GasTipCap:  priorityGasPrice256,
		GasFeeCap:  gasPrice256,
		Gas:        gasLimit,
		To:         to,
		Value:      value256,
		BlobFeeCap: maxFeePerBlobGas256,
		BlobHashes: versionedHashes,
		Sidecar:    &types.BlobTxSidecar{Blobs: blobs, Commitments: commitments, Proofs: proofs},
		V:          uint256.NewInt(v.Uint64() - cli.chainId.Uint64()*2 - 35),
		R:          uint256.MustFromBig(r),
		S:          uint256.MustFromBig(s),
	}
	tx := types.NewTx(blobTx)

	err = cli.client.SendTransaction(ctx, tx)
	if err != nil {
		return common.Hash{}, err
	} else {
		log.Printf("successfully sent transaction. txhash=%v", legacySignedTx.Hash())
	}

	//var receipt *types.Receipt
	for {
		_, err = cli.client.TransactionReceipt(ctx, legacySignedTx.Hash())
		if err == ethereum.NotFound {
			time.Sleep(1 * time.Second)
		} else if err != nil {
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				break
			}
		} else {
			break
		}
	}

	return legacySignedTx.Hash(), nil
}

type rpcTransaction struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Result  *BlobResult `json:"result"`
}

type BlobTxSidecar struct {
	Blobs []string `json:"blobs"` // Blobs needed by the blob pool
}

type BlobResult struct {
	BlobHashes []common.Hash  `json:"blobVersionedHashes"`
	Sidecar    *BlobTxSidecar `json:"sidecar"`
}

func (cli *Client) GetBlob(hash common.Hash) ([]byte, error) {
	rj, err := json.Marshal(map[string]interface{}{
		"method":  "eth_getTransactionByHash",
		"id":      "1",
		"jsonrpc": "2.0",
		"params":  []string{hash.Hex()},
	})

	resp, err := http.Post(cli.rpcUrl, "application/json", bytes.NewBuffer(rj))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var rpcTx *rpcTransaction
	if err := json.NewDecoder(resp.Body).Decode(&rpcTx); err != nil {
		return nil, fmt.Errorf("decode response, %w", err)
	}
	blobs := rpcTx.Result.Sidecar.Blobs

	var r []byte
	for _, blob := range blobs {
		b := DecodeBlob(common.Hex2Bytes(blob))
		r = append(r, b...)
	}

	return r, nil
}
