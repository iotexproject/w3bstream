package output

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/types"
)

func TestNewEthereum(t *testing.T) {
	require := require.New(t)

	chainEndpoint := "https://iotex"
	secretKey := "b7255a24"
	contractAddress := "0x5Ea91218CB1E329806a746E0816A8BD533637b42"
	contractAbiJSON := `[{"constant":false,"inputs":[{"internalType":"bytes","name":"proof","type":"bytes"}],"name":"setProof","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getProof","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"}]`
	contractMethod := "setProof"

	t.Run("NewEthereum", func(t *testing.T) {
		_, err := NewEthereum(chainEndpoint, secretKey, contractAddress, "", contractAbiJSON, contractMethod)
		require.NoError(err)
	})

	t.Run("AbiNil", func(t *testing.T) {
		_, err := NewEthereum(chainEndpoint, secretKey, contractAddress, "", "", contractMethod)
		require.EqualError(err, "EOF")
	})

	t.Run("SecretKeyNil", func(t *testing.T) {
		_, err := NewEthereum(chainEndpoint, "", contractAddress, "", contractAbiJSON, contractMethod)
		require.EqualError(err, "secretkey is empty")
	})

}

func TestEthOutput(t *testing.T) {
	require := require.New(t)

	chainEndpoint := "https://iotex"
	secretKey := "b7255a24"
	contractAddress := "0x5Ea91218CB1E329806a746E0816A8BD533637b42"
	contractAbiJSON := `[{"constant":false,"inputs":[{"internalType":"bytes","name":"proof","type":"bytes"}],"name":"setProof","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getProof","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"}]`
	contractMethod := "setProof"

	task := &types.Task{
		ID: "",
		Messages: []*types.Message{{
			ID:             "id1",
			ProjectID:      uint64(0x1),
			ProjectVersion: "0.1",
			Data:           "data",
		}},
	}

	t.Run("MissMethod", func(t *testing.T) {
		contractMissMethod := "setProof1"
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, "", contractAbiJSON, contractMissMethod)
		require.NoError(err)

		_, err = contract.Output(task, []byte("proof"))
		require.EqualError(err, "contract abi miss the contract method setProof1")
	})

	t.Run("MissParam", func(t *testing.T) {
		contractMissParamAbiJSON := `[{"inputs":[{"internalType":"address","name":"depinRC20Address","type":"address"},{"internalType":"uint256","name":"nonce","type":"uint256"},{"internalType":"address","name":"sender","type":"address"},{"internalType":"bytes","name":"proof","type":"bytes"}],"name":"mine","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"depinRC20","outputs":[{"internalType":"contract IDepinRC20","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
		contractMissParamMethod := "mine"
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, "", contractMissParamAbiJSON, contractMissParamMethod)
		require.NoError(err)

		_, err = contract.Output(task, []byte(hex.EncodeToString([]byte("this is proof"))))
		require.ErrorContains(err, "miss param")
	})

	t.Run("TransactionFailed", func(t *testing.T) {
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, "", contractAbiJSON, contractMethod)
		require.NoError(err)

		_, err = contract.Output(task, []byte(hex.EncodeToString([]byte("this is proof"))))
		require.ErrorContains(err, "transaction failed")
	})
}
