package api

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/w3bstream/block"
)

type jsonRpcReq struct {
	ID      uint64 `json:"id"             binding:"required"`
	Version string `json:"jsonrpc"        binding:"required"`
	Method  string `json:"method"         binding:"required"`
	Params  any    `json:"params"`
}

type jsonRpcRsp struct {
	ID      uint64 `json:"id"`
	Version string `json:"jsonrpc"`
	Error   string `json:"error"`
	Result  any    `json:"result"`
}

type blockTemplate struct {
	Meta          string `json:"meta"`
	PrevBlockHash string `json:"previousblockhash"`
	MerkleRoot    string `json:"merkleroot"`
	Difficulty    uint32 `json:"difficulty"`
	Ts            uint64 `json:"ts"`
	NonceRange    string `json:"noncerange"`
}

type submitBlockParam struct {
	Meta          string `json:"meta"`
	PrevBlockHash string `json:"previousblockhash"`
	MerkleRoot    string `json:"merkleroot"`
	Difficulty    uint32 `json:"difficulty"`
	Ts            uint64 `json:"ts"`
	Nonce         string `json:"nonce"`
}

func (p *submitBlockParam) toBlockHeader() (*block.Header, error) {
	meta, err := hex.DecodeString(p.Meta)
	if err != nil {
		return nil, err
	}
	prevBlockHash, err := hex.DecodeString(p.PrevBlockHash)
	if err != nil {
		return nil, err
	}
	merkleRoot, err := hex.DecodeString(p.MerkleRoot)
	if err != nil {
		return nil, err
	}
	nonce, err := hex.DecodeString(p.Nonce)
	if err != nil {
		return nil, err
	}
	h := &block.Header{
		Meta:       [4]byte{},
		PrevHash:   common.Hash{},
		MerkleRoot: [32]byte{},
		Nonce:      [8]byte{},
	}
	copy(h.Meta[:], meta)
	copy(h.PrevHash[:], prevBlockHash)
	copy(h.MerkleRoot[:], merkleRoot)
	copy(h.Nonce[:], nonce)
	h.Difficulty = p.Difficulty
	return h, nil
}
