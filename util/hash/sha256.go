package hash

import (
	"crypto/sha256"
	"encoding/binary"
)

func Sum256Uint64(number uint64) [sha256.Size]byte {
	numberBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(numberBytes, number)
	return sha256.Sum256(numberBytes)
}
