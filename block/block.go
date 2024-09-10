package block

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

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

func (h *Header) Hash() common.Hash {
	return crypto.Keccak256Hash(h.Meta[:], h.PrevHash[:], h.MerkleRoot[:], h.Difficulty[:], h.Nonce[:])
}
