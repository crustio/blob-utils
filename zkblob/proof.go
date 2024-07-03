package zkblob

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

type ProofInfo struct {
	BatchNumber uint64        `json:"batch_number"`
	Proof       []common.Hash `json:"proof"`
}

type ProofResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      string     `json:"id"`
	Result  *ProofInfo `json:"result"`
}

func GetProofByHash(l2Host string, hash common.Hash) (*ProofInfo, error) {
	rj, err := json.Marshal(map[string]interface{}{
		"method":  "ethda_getProofByHash",
		"id":      "1",
		"jsonrpc": "2.0",
		"params":  []string{hash.Hex()},
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(l2Host, "application/json", bytes.NewBuffer(rj))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	proof := &ProofResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&proof); err != nil {
		return nil, err
	}

	return proof.Result, nil
}
