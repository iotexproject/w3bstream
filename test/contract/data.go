package contract

import (
	"bytes"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func BuildData(param []byte) ([]byte, error) {
	abiJson, err := os.ReadFile("test/contract/Store.abi")
	if err != nil {
		return nil, err
	}
	storeABI, err := abi.JSON(bytes.NewReader(abiJson))
	if err != nil {
		return nil, err
	}
	return storeABI.Pack("setProof", string(param))
}
