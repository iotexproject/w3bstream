package adapter

import (
	"context"
	"log/slog"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/machinefi/w3bstream-mainnet/output/chain/eth"
)

// contractAbiJSON is the ABI of the contract
// solidity interface: function setProof(string memory _proof) external;
const (
	contractMethod  = "setProof"
	contractAbiJSON = `[
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_proof",
				"type": "string"
			}
		],
		"name": "setProof",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
)

// EvmContract is the adapter for outputting proofs to an EVM contract
type EvmContract struct {
	chainEndpoint   string
	contractAddress string
	sk              string
}

var (
	contractABI abi.ABI
)

func init() {
	var err error
	contractABI, err = abi.JSON(strings.NewReader(contractAbiJSON))
	if err != nil {
		panic(err)
	}
}

// NewEvmContract returns a new EVM contract adapter
func NewEvmContract(chainEndpoint, sk, contractAddress string) *EvmContract {
	return &EvmContract{
		chainEndpoint:   chainEndpoint,
		sk:              sk,
		contractAddress: contractAddress,
	}
}

// Output outputs the proof to the EVM contract
func (e *EvmContract) Output(proof []byte) error {
	// pack contract data
	data, err := contractABI.Pack(contractMethod, string(proof))
	if err != nil {
		return err
	}

	// send tx
	txHash, err := eth.SendTX(context.Background(), e.chainEndpoint, e.sk, e.contractAddress, data)
	if err != nil {
		return err
	}
	slog.Debug("evm contract output", "txHash", txHash)

	return nil
}
