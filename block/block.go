package block

import (
	"bytes"
	"crypto/sha256"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/iotexproject/w3bstream/task"
)

type Header struct {
	Meta       [4]byte
	PrevHash   common.Hash
	MerkleRoot [32]byte
	Difficulty [4]byte
	Nonce      [8]byte
}

type Block struct {
	Header Header
	Tasks  []*task.Task
}

func (h *Header) Sha256Hash() common.Hash {
	var buf bytes.Buffer
	buf.Write(h.Meta[:])
	buf.Write(h.PrevHash[:])
	buf.Write(h.MerkleRoot[:])
	buf.Write(h.Difficulty[:])
	buf.Write(h.Nonce[:])

	return sha256.Sum256(buf.Bytes())
}

func (h *Header) IsValid() bool {
	hash := h.Sha256Hash().Bytes()
	hashInt := new(big.Int).SetBytes(hash[:4])
	difficulty := new(big.Int).SetBytes(h.Difficulty[:])

	return hashInt.Cmp(difficulty) == -1
}
