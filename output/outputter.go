package output

import (
	"github.com/machinefi/sprout/output/adapter"
	"github.com/machinefi/sprout/output/chain"
	"github.com/pkg/errors"
)

// Types of outputters
const (
	Stdout           Type = "stdout"
	EthereumContract Type = "ethereumContract"
	SolanaProgram    Type = "solanaProgram"
)

type (
	// Type is the type of outputter
	Type string

	// Outputter is the interface for outputting proofs
	Outputter interface {
		// Output outputs the proof
		Output(proof []byte) error
	}

	// Factory is the factory for creating outputters
	Factory struct {
		chains map[chain.Name]chain.Chain
	}
)

// NewFactory creates a new outputter factory
func NewFactory(chainConfig []byte) (*Factory, error) {
	chains, err := chain.ChainsFromJSON(chainConfig)
	if err != nil {
		return nil, err
	}
	return &Factory{
		chains: chains,
	}, nil
}

// NewOutputter returns a new outputter based on the config
func (f *Factory) NewOutputter(cfg Config) (out Outputter, err error) {
	switch cfg.Type {
	case Stdout:
		out = adapter.NewStdout()
	case EthereumContract:
		chain, ok := f.chains[cfg.ChainName]
		if !ok {
			return nil, errors.Errorf("invalid chain name: %s", cfg.ChainName)
		}
		out, err = adapter.NewEthereumContract(chain.Endpoint, cfg.SecretKey, cfg.ContractAddress)
	case SolanaProgram:
		chain, ok := f.chains[cfg.ChainName]
		if !ok {
			return nil, errors.Errorf("invalid chain name: %s", cfg.ChainName)
		}
		out = adapter.NewSolanaProgram(chain.Endpoint, cfg.ContractAddress, cfg.SecretKey, cfg.StateAccountPK)
	default:
		return nil, errors.New("invalid output type")
	}
	return
}
