package output

import (
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/blocto/solana-go-sdk/client"
	soltypes "github.com/blocto/solana-go-sdk/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func patchSolanaProgramSendTX(p *Patches, txhash string, err error) *Patches {
	return p.ApplyPrivateMethod(&solanaProgram{}, "sendTX", func(*solanaProgram, []soltypes.Instruction) (string, error) {
		return txhash, err
	})
}

func Test_solanaProgram_Output(t *testing.T) {
	r := require.New(t)
	e1 := &solanaProgram{stateAccountPK: ""}
	e2 := &solanaProgram{stateAccountPK: "any"}

	t.Run("FailedToSendTX", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = patchSolanaProgramSendTX(p, "", errors.New(t.Name()))

		txHash, err := e1.Output(1, [][]byte{}, nil)
		r.Equal(txHash, "")
		r.ErrorContains(err, t.Name())

		txHash, err = e2.Output(1, [][]byte{}, nil)
		r.Equal(txHash, "")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		txHashRet := "anyTxHash"
		p := NewPatches()
		defer p.Reset()
		p = patchSolanaProgramSendTX(p, txHashRet, nil)

		txHash, err := e1.Output(1, [][]byte{}, nil)
		r.Equal(txHash, txHashRet)
		r.NoError(err)

		txHash, err = e2.Output(1, [][]byte{}, nil)
		r.Equal(txHash, txHashRet)
		r.NoError(err)
	})
}

func Test_solanaProgram_sendTX(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()

	contract := &solanaProgram{
		secretKey: "fd6ac80f1b9886a6d157cd8e71f842a63c52ebd237cf48fba03ae587e197d511f0b2439ae6da236d26f17f56c68f05d48513cd99b33143fa0b1aec7838ce4276",
	}
	ins := contract.packInstructions([]byte("proof"))

	t.Run("MissingInstructionData", func(t *testing.T) {
		p = p.ApplyFuncReturn(client.NewClient, &client.Client{})
		_, err := contract.sendTX(nil)
		r.EqualError(err, "missing instruction data")
	})

	t.Run("GetSolanaBlockFailed", func(t *testing.T) {
		p = p.ApplyMethodReturn(&client.Client{}, "GetLatestBlockhash", nil, errors.New(t.Name()))
		_, err := contract.sendTX(ins)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyMethodReturn(&client.Client{}, "GetLatestBlockhash", nil, nil)

	t.Run("BuildSolanaTxFailed", func(t *testing.T) {
		p = p.ApplyFuncReturn(soltypes.NewTransaction, nil, errors.New(t.Name()))
		_, err := contract.sendTX(ins)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(soltypes.NewTransaction, nil, nil)

	t.Run("SendSolanaTxFailed", func(t *testing.T) {
		p = p.ApplyMethodReturn(&client.Client{}, "SendTransaction", "", errors.New(t.Name()))
		_, err := contract.sendTX(ins)
		r.ErrorContains(err, t.Name())
	})

	t.Run("SendSolanaTxSuccess", func(t *testing.T) {
		p = p.ApplyMethodReturn(&client.Client{}, "SendTransaction", t.Name(), nil)

		hash, err := contract.sendTX(ins)
		r.NoError(err)
		r.Equal(t.Name(), hash)
	})
}
