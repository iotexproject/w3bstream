package output

import (
	"testing"

	"github.com/blocto/solana-go-sdk/client"
	soltypes "github.com/blocto/solana-go-sdk/types"
	. "github.com/bytedance/mockey"
	"github.com/machinefi/sprout/types"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
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

	PatchConvey("SendTXFailed", t, func() {
		Mock((*solanaProgram).sendTX).Return("", errors.New(t.Name())).Build()
		_, err = contract.Output(task, []byte("proof"))
		So(err.Error(), ShouldEqual, t.Name())
	})

	PatchConvey("SendTXOk", t, func() {
		Mock((*solanaProgram).sendTX).Return(t.Name(), nil).Build()
		txHash, err := contract.Output(task, []byte("proof"))
		So(err, ShouldBeEmpty)
		So(txHash, ShouldEqual, t.Name())
	})
}

func TestSolanaSendTX(t *testing.T) {

	chainEndpoint := "https://solana"
	programID := "0x5Ea91218CB1E329806a746E0816A8BD533637b42"
	secretKey := "fd6ac80f1b9886a6d157cd8e71f842a63c52ebd237cf48fba03ae587e197d511f0b2439ae6da236d26f17f56c68f05d48513cd99b33143fa0b1aec7838ce4276"
	stateAccountPK := "accountPK"

	contract := &solanaProgram{
		endpoint:       chainEndpoint,
		programID:      programID,
		secretKey:      secretKey,
		stateAccountPK: stateAccountPK,
	}

	PatchConvey("MissingInstructionData", t, func() {
		Mock(client.NewClient).Return(&client.Client{}).Build()
		_, err := contract.sendTX(chainEndpoint, secretKey, nil)
		So(err.Error(), ShouldEqual, "missing instruction data")
	})

	PatchConvey("GetSolanaBlockFailed", t, func() {
		Mock(client.NewClient).Return(&client.Client{})
		Mock((*client.Client).GetLatestBlockhash).Return(nil, errors.New(t.Name())).Build()
		_, err := contract.sendTX(chainEndpoint, secretKey, contract.packInstructions([]byte("proof")))
		So(err.Error(), ShouldContainSubstring, t.Name())
	})

	PatchConvey("BuildSolanaTxFailed", t, func() {
		Mock(client.NewClient).Return(&client.Client{})
		Mock((*client.Client).GetLatestBlockhash).Return(nil, nil).Build()
		Mock(soltypes.NewTransaction).Return(nil, errors.New(t.Name())).Build()
		_, err := contract.sendTX(chainEndpoint, secretKey, contract.packInstructions([]byte("proof")))
		So(err.Error(), ShouldContainSubstring, t.Name())
	})

	PatchConvey("SendSolanaTxFailed", t, func() {
		Mock(client.NewClient).Return(&client.Client{})
		Mock((*client.Client).GetLatestBlockhash).Return(nil, nil).Build()
		Mock(soltypes.NewTransaction).Return(nil, nil).Build()
		Mock((*client.Client).SendTransaction).Return("", errors.New(t.Name())).Build()
		_, err := contract.sendTX(chainEndpoint, secretKey, contract.packInstructions([]byte("proof")))
		So(err.Error(), ShouldContainSubstring, t.Name())
	})

	PatchConvey("SendSolanaTxOk", t, func() {
		Mock(client.NewClient).Return(&client.Client{})
		Mock((*client.Client).GetLatestBlockhash).Return(nil, nil).Build()
		Mock(soltypes.NewTransaction).Return(nil, nil).Build()
		Mock((*client.Client).SendTransaction).Return(t.Name(), nil).Build()
		hash, err := contract.sendTX(chainEndpoint, secretKey, contract.packInstructions([]byte("proof")))
		So(err, ShouldBeEmpty)
		So(hash, ShouldEqual, hash)
	})
}
