package adapter

import (
	"context"
	"log/slog"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/machinefi/w3bstream-mainnet/output/chain/eth"
)

// contractAbiJSON is the ABI of the contract
// solidity interface: function submitProof(string memory _proof) external;
const (
	contractMethod  = "submitProof"
	contractAbiJSON = `[
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes",
				"name": "_proof",
				"type": "bytes"
			}
		],
		"name": "submitProof",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
)

// EthereumContract is the adapter for outputting proofs to an ethereum-compatible contract
type EthereumContract struct {
	chainEndpoint   string
	contractAddress string
	secretKey       string
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

// NewEthereumContract returns a new ethereum contract adapter
func NewEthereumContract(chainEndpoint, secretKey, contractAddress string) *EthereumContract {
	return &EthereumContract{
		chainEndpoint:   chainEndpoint,
		secretKey:       secretKey,
		contractAddress: contractAddress,
	}
}

// Output outputs the proof to the ethereum contract
func (e *EthereumContract) Output(proof []byte) error {
	// pack contract data
	data, err := contractABI.Pack(contractMethod, proof)
	if err != nil {
		return err
	}

	// send tx
	txHash, err := eth.SendTX(context.Background(), e.chainEndpoint, e.secretKey, e.contractAddress, data)
	if err != nil {
		return err
	}
	slog.Debug("ethereum contract output", "txHash", txHash)

	return nil
}
