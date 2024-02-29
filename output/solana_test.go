package output

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/types"
)

func TestNewSolanaProgram(t *testing.T) {
	require := require.New(t)

	chainEndpoint := "https://solana"
	programID := "0x5Ea91218CB1E329806a746E0816A8BD533637b42"
	secretKey := "b7255a24"
	stateAccountPK := "accountPK"

	t.Run("NewSolana", func(t *testing.T) {
		_, err := NewSolanaProgram(chainEndpoint, programID, secretKey, stateAccountPK)
		require.NoError(err)
	})

	t.Run("SecretKeyNil", func(t *testing.T) {
		_, err := NewSolanaProgram(chainEndpoint, programID, "", stateAccountPK)
		require.EqualError(err, "secretkey is empty")
	})
}

func TestSolanaOutput(t *testing.T) {
	require := require.New(t)

	chainEndpoint := "https://solana"
	programID := "0x5Ea91218CB1E329806a746E0816A8BD533637b42"
	secretKey := "fd6ac80f1b9886a6d157cd8e71f842a63c52ebd237cf48fba03ae587e197d511f0b2439ae6da236d26f17f56c68f05d48513cd99b33143fa0b1aec7838ce4276"
	stateAccountPK := "accountPK"

	task := &types.Task{
		ID: "",
		Messages: []*types.Message{{
			ID:             "id1",
			ProjectID:      uint64(0x1),
			ProjectVersion: "0.1",
			Data:           "data",
		}},
	}

	contract, err := NewSolanaProgram(chainEndpoint, programID, secretKey, stateAccountPK)
	require.NoError(err)

	t.Run("RpcFail", func(t *testing.T) {
		_, err = contract.Output(task, []byte("proof"))
		require.ErrorContains(err, "failed to get solana latest block hash")
	})
}
