package chain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChainsFromJSON(t *testing.T) {
	r := require.New(t)

	jsonData := []byte(`[
		{
			"name": "iotex-testnet",
			"chainID": 4690,
			"endpoint": "endpoint"
		},
		{
			"name": "solana-testnet",
			"endpoint": "endpointsolana"
		}
	]`)
	chains, err := ChainsFromJSON(jsonData)
	r.NoError(err)
	r.Len(chains, 2)
	r.EqualValues(chains[IoTeXTestNet].Name, IoTeXTestNet)
	r.EqualValues(chains[IoTeXTestNet].ID, uint64(4690))
	r.EqualValues(chains[IoTeXTestNet].Endpoint, "endpoint")
	r.EqualValues(chains[SolanaTestNet].Name, SolanaTestNet)
	r.EqualValues(chains[SolanaTestNet].Endpoint, "endpointsolana")
}
