package output

import (
	"testing"

	"github.com/blocto/solana-go-sdk/client"
	soltypes "github.com/blocto/solana-go-sdk/types"
	"github.com/bytedance/mockey"
	"github.com/machinefi/sprout/types"
	"github.com/pkg/errors"
	"github.com/smartystreets/goconvey/convey"
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

	t.Run("SendTXFailed", func(t *testing.T) {
		mockey.PatchConvey("SendTXFailed", t, func() {
			mockey.Mock((*solanaProgram).sendTX).Return("", errors.New(t.Name())).Build()
			_, err = contract.Output(task, []byte("proof"))
			convey.So(err.Error(), convey.ShouldEqual, t.Name())
		})
	})

	t.Run("SendTXOk", func(t *testing.T) {
		mockey.PatchConvey("SendTXOk", t, func() {
			mockey.Mock((*solanaProgram).sendTX).Return(t.Name(), nil).Build()
			txHash, err := contract.Output(task, []byte("proof"))
			convey.So(err, convey.ShouldBeEmpty)
			convey.So(txHash, convey.ShouldEqual, t.Name())
		})
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

	t.Run("MissingInstructionData", func(t *testing.T) {
		mockey.PatchConvey("MissingInstructionData", t, func() {
			mockey.Mock(client.NewClient).Return(&client.Client{}).Build()
			_, err := contract.sendTX(chainEndpoint, secretKey, nil)
			convey.So(err.Error(), convey.ShouldEqual, "missing instruction data")
		})
	})

	t.Run("GetSolanaBlockFailed", func(t *testing.T) {
		mockey.PatchConvey("GetSolanaBlockFailed", t, func() {
			mockey.Mock(client.NewClient).Return(&client.Client{})
			mockey.Mock((*client.Client).GetLatestBlockhash).Return(nil, errors.New(t.Name())).Build()
			_, err := contract.sendTX(chainEndpoint, secretKey, contract.packInstructions([]byte("proof")))
			convey.So(err.Error(), convey.ShouldContainSubstring, t.Name())
		})
	})

	t.Run("BuildSolanaTxFailed", func(t *testing.T) {
		mockey.PatchConvey("BuildSolanaTxFailed", t, func() {
			mockey.Mock(client.NewClient).Return(&client.Client{})
			mockey.Mock((*client.Client).GetLatestBlockhash).Return(nil, nil).Build()
			mockey.Mock(soltypes.NewTransaction).Return(nil, errors.New(t.Name())).Build()
			_, err := contract.sendTX(chainEndpoint, secretKey, contract.packInstructions([]byte("proof")))
			convey.So(err.Error(), convey.ShouldContainSubstring, t.Name())
		})
	})

	t.Run("SendSolanaTxFailed", func(t *testing.T) {
		mockey.PatchConvey("SendSolanaTxFailed", t, func() {
			mockey.Mock(client.NewClient).Return(&client.Client{})
			mockey.Mock((*client.Client).GetLatestBlockhash).Return(nil, nil).Build()
			mockey.Mock(soltypes.NewTransaction).Return(nil, nil).Build()
			mockey.Mock((*client.Client).SendTransaction).Return("", errors.New(t.Name())).Build()
			_, err := contract.sendTX(chainEndpoint, secretKey, contract.packInstructions([]byte("proof")))
			convey.So(err.Error(), convey.ShouldContainSubstring, t.Name())
		})
	})

	t.Run("SendSolanaTxOk", func(t *testing.T) {
		mockey.PatchConvey("SendSolanaTxOk", t, func() {
			mockey.Mock(client.NewClient).Return(&client.Client{})
			mockey.Mock((*client.Client).GetLatestBlockhash).Return(nil, nil).Build()
			mockey.Mock(soltypes.NewTransaction).Return(nil, nil).Build()
			mockey.Mock((*client.Client).SendTransaction).Return(t.Name(), nil).Build()
			hash, err := contract.sendTX(chainEndpoint, secretKey, contract.packInstructions([]byte("proof")))
			convey.So(err, convey.ShouldBeEmpty)
			convey.So(hash, convey.ShouldEqual, hash)
		})
	})
}
