package output

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	r := require.New(t)

	t.Run("Default", func(t *testing.T) {
		c := &Config{}
		o, err := New(c, "", "")
		r.NoError(err)
		_, ok := o.(*stdout)
		r.True(ok)
	})
	t.Run("Stdout", func(t *testing.T) {
		c := &Config{
			Type: Stdout,
		}
		o, err := New(c, "", "")
		r.NoError(err)
		_, ok := o.(*stdout)
		r.True(ok)
	})
	t.Run("Ethereum", func(t *testing.T) {
		c := &Config{
			Type: EthereumContract,
			Ethereum: EthereumConfig{
				ContractAbiJSON: `[{"constant":true,"inputs":[],"name":"getProof","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"bytes","name":"_proof","type":"bytes"}],"name":"setProof","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`,
				ContractMethod:  "getProof",
			},
		}
		o, err := New(c, "c47bbade736b0f82788aa6eaa06140cdf41a544707edef944299642e0d708cab", "")
		r.NoError(err)
		_, ok := o.(*ethereumContract)
		r.True(ok)
	})
	t.Run("Solana", func(t *testing.T) {
		c := &Config{
			Type:   SolanaProgram,
			Solana: SolanaConfig{},
		}
		o, err := New(c, "", "308edd7fca562182adbffaa59264a138d9e04f9f3adbda2c80ef1ca71b7dcfa4")
		r.NoError(err)
		_, ok := o.(*solanaProgram)
		r.True(ok)
	})
}
