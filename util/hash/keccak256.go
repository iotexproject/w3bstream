package hash

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Keccak256Uint64(number uint64) common.Hash {
	numberBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(numberBytes, number)
	return crypto.Keccak256Hash(numberBytes)
}
