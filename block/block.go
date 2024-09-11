package block

import (
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/btcsuite/btcd/blockchain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/iotexproject/w3bstream/task"
)

type Header struct {
	Meta       [4]byte
	PrevHash   common.Hash
	MerkleRoot [32]byte
	Difficulty uint32
	Nonce      [8]byte
}

type Block struct {
	Header Header
	Tasks  []*task.Task
}

func (h *Header) Hash() common.Hash {
	numberBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(numberBytes, h.Difficulty)
	return crypto.Keccak256Hash(h.Meta[:], h.PrevHash[:], h.MerkleRoot[:], numberBytes[:], h.Nonce[:])
}

func (h *Header) Sha256Hash() common.Hash {
	numberBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(numberBytes, h.Difficulty)
	data := []byte{}
	data = append(data, h.Meta[:]...)
	data = append(data, h.PrevHash[:]...)
	data = append(data, h.MerkleRoot[:]...)
	data = append(data, numberBytes[:]...)
	data = append(data, h.Nonce[:]...)
	return sha256.Sum256(data)
}

func (h *Header) IsValid() bool {
	target := blockchain.CompactToBig(h.Difficulty)
	firstHash := h.Sha256Hash()
	finalHash := sha256.Sum256(firstHash[:])
	hashInt := new(big.Int).SetBytes(finalHash[:])

	return hashInt.Cmp(target) == -1
}
