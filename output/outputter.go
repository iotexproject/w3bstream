package output

import (
	"errors"

	"github.com/machinefi/w3bstream-mainnet/output/adapter"
)

type (
	// Outputter is the interface for outputting proofs
	Outputter interface {
		// Output outputs the proof
		Output(proof []byte) error
	}
)

// NewOutputter returns a new outputter based on the config
func NewOutputter(cfg Config) (out Outputter, err error) {
	switch cfg.Type {
	case Stdout:
		// TODO: implement
	case EthereumContract:
		out, err = adapter.NewEthereumContract(cfg.ChainEndpoint, cfg.SecretKey, cfg.ContractAddress)
	default:
		return nil, errors.New("invalid output type")
	}
	return
}
