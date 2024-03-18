package output

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/types"
)

func TestConfig_Output(t *testing.T) {
	r := require.New(t)

	t.Run("NewOutputByType", func(t *testing.T) {
		t.Run("Default", func(t *testing.T) {
			o, err := (&Config{}).Output()
			r.Equal(o.Type(), types.OutputStdout)
			r.NoError(err)
		})

		t.Run("Ethereum", func(t *testing.T) {
			t.Run("InvalidEthereumOutputConfig", func(t *testing.T) {
				o, err := (&Config{
					Type: types.OutputEthereumContract,
				}).Output()
				r.Nil(o)
				r.Equal(err, errInvalidEthereumOutputConfig)
			})
			t.Run("InvalidEthereumSecretKey", func(t *testing.T) {
				c := &Config{
					Type:     types.OutputEthereumContract,
					Ethereum: &ethereumConfig{},
				}
				o, err := c.Output()
				r.Nil(o)
				r.Equal(err, errInvalidEthereumSecretKey)
			})
			t.Run("InvalidEthereumABI", func(t *testing.T) {
				c := &Config{
					Type:     types.OutputEthereumContract,
					Ethereum: &ethereumConfig{},
				}
				c.SetPrivateKey("any", "any")
				o, err := c.Output()
				r.Nil(o)
				r.Equal(err, errInvalidEthereumABI)
			})
			t.Run("InvalidEthereumMethod", func(t *testing.T) {
				c := &Config{
					Type: types.OutputEthereumContract,
					Ethereum: &ethereumConfig{
						ContractAbiJSON: `[{"inputs": [],"name": "otherMethod","outputs": [],"type": "function"}]`,
						ContractMethod:  "expectedMethod",
					},
				}
				c.SetPrivateKey("any", "any")
				o, err := c.Output()
				r.Nil(o)
				r.Equal(err, errInvalidEthereumMethod)
			})
			t.Run("Success", func(t *testing.T) {
				c := &Config{
					Type: types.OutputEthereumContract,
					Ethereum: &ethereumConfig{
						ContractAbiJSON: `[{"inputs": [],"name": "expectedMethod","outputs": [],"type": "function"}]`,
						ContractMethod:  "expectedMethod",
					},
				}
				c.SetPrivateKey("any", "any")
				o, err := c.Output()
				r.Equal(o.Type(), types.OutputEthereumContract)
				r.NoError(err)
			})
		})

		t.Run("Solana", func(t *testing.T) {
			t.Run("InvalidSolanaConfig", func(t *testing.T) {
				o, err := (&Config{Type: types.OutputSolanaProgram}).Output()
				r.Nil(o)
				r.Equal(err, errInvalidSolanaOutputConfig)
			})

			t.Run("InvalidSolanaSecretKey", func(t *testing.T) {
				c := &Config{Type: types.OutputSolanaProgram, Solana: &solanaConfig{}}
				o, err := c.Output()
				r.Nil(o)
				r.Equal(err, errInvalidSolanaSecretKey)
			})

			t.Run("Success", func(t *testing.T) {
				c := &Config{Type: types.OutputSolanaProgram, Solana: &solanaConfig{}}
				c.SetPrivateKey("any", "any")
				o, err := c.Output()
				r.Equal(o.Type(), types.OutputSolanaProgram)
				r.NoError(err)
			})
		})

		t.Run("Textile", func(t *testing.T) {
			t.Run("InvalidTextileConfig", func(t *testing.T) {
				o, err := (&Config{Type: types.OutputTextile}).Output()
				r.Nil(o)
				r.Equal(err, errInvalidTextileOutputConfig)
			})

			t.Run("InvalidTextileSecretKey", func(t *testing.T) {
				c := &Config{Type: types.OutputTextile, Textile: &textileConfig{}}
				o, err := c.Output()
				r.Nil(o)
				r.Equal(err, errInvalidTextileSecretKey)
			})

			t.Run("Success", func(t *testing.T) {
				c := &Config{Type: types.OutputTextile, Textile: &textileConfig{}}
				c.SetPrivateKey("any", "any")
				o, err := c.Output()
				r.Equal(o.Type(), types.OutputTextile)
				r.NoError(err)
			})
		})
	})
}
