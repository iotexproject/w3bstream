package output

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
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
	patches := NewPatches()

	chainEndpoint := "https://iotex"
	secretKey := "b7255a24"
	contractAddress := "0x5Ea91218CB1E329806a746E0816A8BD533637b42"
	receiverAddress := "0x5Ea91218CB1E329806a746E0816A8BD533637b42"
	contractAbiJSON := `[{"inputs":[],"name":"getJournal","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getPostStateDigest","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getProjectId","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getReceiver","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getSeal","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes","name":"_proof","type":"bytes"}],"name":"setProof","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"_projectId","type":"uint256"},{"internalType":"address","name":"_receiver","type":"address"},{"internalType":"bytes","name":"_data_snark","type":"bytes"}],"name":"submit","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	contractMethod := "submit"

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

	t.Run("MissReceiverAddress", func(t *testing.T) {
		//contractMissParamAbiJSON := `[{"inputs":[],"name":"getJournal","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getPostStateDigest","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getProjectId","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getReceiver","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getSeal","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes","name":"_proof","type":"bytes"}],"name":"setProof","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"_projectId","type":"uint256"},{"internalType":"address","name":"_receiver","type":"address"},{"internalType":"bytes","name":"_data_snark","type":"bytes"}],"name":"submit","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
		//contractMissParamMethod := "submit"
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, "", contractAbiJSON, contractMethod)
		require.NoError(err)

		_, err = contract.Output(task, []byte("this is proof"))
		require.ErrorContains(err, "miss param")
	})

	t.Run("GetSnarkFailed", func(t *testing.T) {
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, receiverAddress, contractAbiJSON, contractMethod)
		require.NoError(err)

		_, err = contract.Output(task, []byte("this is proof"))
		require.ErrorContains(err, "get Snark.snark failed")
	})

	t.Run("GetPostStateDigestFailed", func(t *testing.T) {
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, receiverAddress, contractAbiJSON, contractMethod)
		require.NoError(err)

		//task.Messages[0].Data = "{\"Snark\":{\"snark\":{\"a\":[[11,176,218,102,82,247],[19,201,71,203,]],\"b\":[[[37,238,237,46],[36,124,137]],[[5,237,77],[41,187,159]]],\"c\":[[31,108,130],[34,189,130]],\"public\":[[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,68],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,197],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5]]},\"post_state_digest\":[244,204,22,124,129,242],\"journal\":[82,0,0,0,73,32]}}"
		//task.Messages[0].Data = "{\"Snark\":{\"snark\":{\"a\":[[11,176,218,102,82,247],[19,201,71,203,]],\"b\":[[[37,238,237,46],[36,124,137]],[[5,237,77],[41,187,159]]],\"c\":[[31,108,130],[34,189,130]],\"public\":[[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,68],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,197],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5]]},\"journal\":[82,0,0,0,73,32]}}"
		proof := "{\"Snark\":{\"snark\":{\"a\":[[11,176,218,102,82,247],[19,201,71,203,]],\"b\":[[[37,238,237,46],[36,124,137]],[[5,237,77],[41,187,159]]],\"c\":[[31,108,130],[34,189,130]],\"public\":[[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,68],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,197],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5]]},\"journal\":[82,0,0,0,73,32]}}"
		_, err = contract.Output(task, []byte(proof))
		require.ErrorContains(err, "get Snark.post_state_digest failed")
	})

	t.Run("GetJournalFailed", func(t *testing.T) {
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, receiverAddress, contractAbiJSON, contractMethod)
		require.NoError(err)

		proof := "{\"Snark\":{\"snark\":{\"a\":[[11,176,218,102,82,247],[19,201,71,203,]],\"b\":[[[37,238,237,46],[36,124,137]],[[5,237,77],[41,187,159]]],\"c\":[[31,108,130],[34,189,130]],\"public\":[[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,68],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,197],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5]]},\"post_state_digest\":[244,204,22,124,129,242]}}"
		_, err = contract.Output(task, []byte(proof))
		require.ErrorContains(err, "get Snark.journal failed")
	})

	t.Run("NewAbiPackFailed", func(t *testing.T) {
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, receiverAddress, contractAbiJSON, contractMethod)
		require.NoError(err)

		proof := "{\"Snark\":{\"snark\":{\"a\":[[11,176,218,102,82,247],[19,201,71,203,]],\"b\":[[[37,238,237,46],[36,124,137]],[[5,237,77],[41,187,159]]],\"c\":[[31,108,130],[34,189,130]],\"public\":[[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,68],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,197],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5]]},\"post_state_digest\":[244,204,22,124,129,242],\"journal\":[82,0,0,0,73,32]}}"
		patches = patches.ApplyFuncReturn(abi.NewType, nil, errors.New(t.Name()))
		defer patches.Reset()

		_, err = contract.Output(task, []byte(proof))
		require.ErrorContains(err, t.Name())
	})

	t.Run("AbiPackFailed", func(t *testing.T) {
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, receiverAddress, contractAbiJSON, contractMethod)
		require.NoError(err)

		proof := "{\"Snark\":{\"snark\":{\"a\":[[11,176,218,102,82,247],[19,201,71,203,]],\"b\":[[[37,238,237,46],[36,124,137]],[[5,237,77],[41,187,159]]],\"c\":[[31,108,130],[34,189,130]],\"public\":[[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,68],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,197],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5]]},\"post_state_digest\":[244,204,22,124,129,242],\"journal\":[82,0,0,0,73,32]}}"
		patches = patches.ApplyFuncReturn(abi.NewType, nil, nil)
		patches = ethArgumentsPack(patches, errors.New(t.Name()))
		defer patches.Reset()

		_, err = contract.Output(task, []byte(proof))
		require.ErrorContains(err, "ethereum accounts abi pack failed")
	})

	t.Run("MissProofSnarkParam", func(t *testing.T) {
		contractSnarkAbiJSON := `[{"inputs":[],"name":"getJournal","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getPostStateDigest","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getProjectId","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getReceiver","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getSeal","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes","name":"_proof","type":"bytes"}],"name":"setProof","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"_projectId","type":"uint256"},{"internalType":"address","name":"_receiver","type":"address"},{"internalType":"bytes","name":"proof_snark_seal","type":"bytes"},{"internalType":"bytes","name":"proof_snark_post_state_digest","type":"bytes"},{"internalType":"bytes","name":"proof_snark_journal","type":"bytes"}],"name":"submit","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, receiverAddress, contractSnarkAbiJSON, contractMethod)
		require.NoError(err)

		_, err = contract.Output(task, []byte("this is proof"))
		require.ErrorContains(err, "miss param")
	})

	t.Run("MissParam", func(t *testing.T) {
		//contractMissParamAbiJSON := `[{"inputs":[{"internalType":"address","name":"depinRC20Address","type":"address"},{"internalType":"uint256","name":"nonce","type":"uint256"},{"internalType":"address","name":"sender","type":"address"},{"internalType":"bytes","name":"proof","type":"bytes"}],"name":"mine","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"depinRC20","outputs":[{"internalType":"contract IDepinRC20","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
		//contractMissParamMethod := "mine"
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, "", contractAbiJSON, contractMethod)
		require.NoError(err)

		_, err = contract.Output(task, []byte("this is proof"))
		require.ErrorContains(err, "miss param")
	})

	t.Run("ContractABIPackFailed", func(t *testing.T) {
		contractAbiJSON = `[{"constant":false,"inputs":[{"internalType":"bytes","name":"proof","type":"bytes"}],"name":"setProof","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getProof","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"}]`
		contractMethod = "setProof"
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, receiverAddress, contractAbiJSON, contractMethod)
		require.NoError(err)

		patches = ethABIPack(patches, errors.New(t.Name()))
		defer patches.Reset()

		_, err = contract.Output(task, []byte("this is proof"))
		require.ErrorContains(err, t.Name())
	})

	t.Run("TransactionFailed", func(t *testing.T) {
		contractAbiJSON = `[{"constant":false,"inputs":[{"internalType":"bytes","name":"proof","type":"bytes"}],"name":"setProof","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getProof","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"}]`
		contractMethod = "setProof"
		contract, err := NewEthereum(chainEndpoint, secretKey, contractAddress, receiverAddress, contractAbiJSON, contractMethod)
		require.NoError(err)

		_, err = contract.Output(task, []byte("this is proof"))
		require.ErrorContains(err, "transaction failed")
	})
}

func TestEthSendTX(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()

	chainEndpoint := "https://iotex"
	//chainEndpoint := "https://babel-api.testnet.iotex.io"
	secretKey := "b7255a24"
	contractAddress := "0x5Ea91218CB1E329806a746E0816A8BD533637b42"
	receiverAddress := "0x5Ea91218CB1E329806a746E0816A8BD533637b42"
	contractAbiJSON := `[{"constant":false,"inputs":[{"internalType":"bytes","name":"proof","type":"bytes"}],"name":"setProof","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getProof","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"}]`
	contractMethod := "setProof"
	contractABI, err := abi.JSON(strings.NewReader(contractAbiJSON))
	require.NoError(err)
	contract := &ethereumContract{
		chainEndpoint:   chainEndpoint,
		secretKey:       secretKey,
		contractAddress: contractAddress,
		receiverAddress: receiverAddress,
		contractABI:     contractABI,
		contractMethod:  contractMethod,
	}
	require.NoError(err)

	ctx := context.Background()
	data := []byte("this is proof")

	t.Run("DialEthFailed", func(t *testing.T) {
		patches.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))
		defer patches.Reset()

		_, err := contract.sendTX(ctx, chainEndpoint, secretKey, contractAddress, data)
		require.ErrorContains(err, t.Name())
	})

}

func ethArgumentsPack(p *Patches, err error) *Patches {
	var args *abi.Arguments
	return p.ApplyMethodFunc(
		reflect.TypeOf(args),
		"Pack",
		func(args ...interface{}) ([]byte, error) {
			return nil, err
		},
	)
}

func ethABIPack(p *Patches, err error) *Patches {
	var a *abi.ABI
	return p.ApplyMethodFunc(
		reflect.TypeOf(a),
		"Pack",
		func(name string, args ...interface{}) ([]byte, error) {
			return nil, err
		},
	)
}
