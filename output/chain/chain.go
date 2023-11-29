package chain

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Chain names
const (
	IoTeXTestNet  Name = "iotex-testnet"
	SolanaTestNet Name = "solana-testnet"
)

type (
	// Name is the name of the chain
	Name string

	// Chain is the chain configuration
	Chain struct {
		Name     Name   `json:"name"`
		ID       uint64 `json:"chainID"`
		Endpoint string `json:"endpoint"`
	}
)

var (
	supportedChains = []Name{
		IoTeXTestNet,
		SolanaTestNet,
	}
)

// ChainsFromJSON parses the chain config from json
func ChainsFromJSON(data []byte) (map[Name]Chain, error) {
	// parse the chains
	chains := make([]Chain, 0)
	if err := json.Unmarshal(data, &chains); err != nil {
		return nil, err
	}

	// convert to map
	chainMap := make(map[Name]Chain)
	for _, chain := range chains {
		chainMap[chain.Name] = chain
	}

	// check if all supported chains are present
	for _, name := range supportedChains {
		if _, ok := chainMap[name]; !ok {
			return nil, errors.Errorf("missing chain config:%s", name)
		}
	}

	return chainMap, nil
}
