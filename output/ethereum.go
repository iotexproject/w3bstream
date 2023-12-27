package output

import (
	"context"
	"log/slog"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/types"
)

// // contractAbiJSON is the ABI of the contract
// // solidity interface: function submitProof(string memory _proof) external;
// const (
// 	contractMethod  = "submitProof"
// 	contractAbiJSON = `[
// 	{
// 		"constant": false,
// 		"inputs": [
// 			{
// 				"internalType": "bytes",
// 				"name": "_proof",
// 				"type": "bytes"
// 			}
// 		],
// 		"name": "submitProof",
// 		"outputs": [],
// 		"payable": false,
// 		"stateMutability": "nonpayable",
// 		"type": "function"
// 	}
// ]`
// )

type ethereumContract struct {
	chainEndpoint   string
	contractAddress string
	secretKey       string
	contractABI     abi.ABI
	contractMethod  string
}

func (e *ethereumContract) Output(task *types.Task, proof []byte) (string, error) {
	slog.Debug("outputing to ethereum contract", "chain endpoint", e.chainEndpoint)
	data, err := e.contractABI.Pack(e.contractMethod, proof)
	if err != nil {
		return "", err
	}

	txHash, err := e.sendTX(context.Background(), e.chainEndpoint, e.secretKey, e.contractAddress, data)
	if err != nil {
		return "", err
	}
	slog.Debug("output success", "txHash", txHash)

	return txHash, nil
}

func (e *ethereumContract) sendTX(ctx context.Context, endpoint, privateKey, toStr string, data []byte) (string, error) {
	cli, err := ethclient.Dial(endpoint)
	if err != nil {
		return "", errors.Wrapf(err, "dial eth endpoint %s failed", endpoint)
	}

	pk := crypto.ToECDSAUnsafe(common.FromHex(privateKey))
	sender := crypto.PubkeyToAddress(pk.PublicKey)
	to := common.HexToAddress(toStr)

	gasPrice, err := cli.SuggestGasPrice(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get suggest gas price failed")
	}

	chainid, err := cli.ChainID(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get chain id failed")
	}

	nonce, err := cli.PendingNonceAt(ctx, sender)
	if err != nil {
		return "", errors.Wrap(err, "get pending nonce failed")
	}

	msg := ethereum.CallMsg{
		From:     sender,
		To:       &to,
		GasPrice: gasPrice,
		Data:     data,
	}
	gasLimit, err := cli.EstimateGas(ctx, msg)
	if err != nil {
		return "", errors.Wrap(err, "estimate gas failed")
	}

	tx := ethtypes.NewTx(
		&ethtypes.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      gasLimit,
			To:       &to,
			Data:     data,
		})

	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewLondonSigner(chainid), pk)
	if err != nil {
		return "", errors.Wrap(err, "sign tx failed")
	}

	if err = cli.SendTransaction(ctx, signedTx); err != nil {
		return "", errors.Wrap(err, "send transaction failed")
	}

	return signedTx.Hash().Hex(), nil
}

func NewEthereum(chainEndpoint, secretKey, contractAddress, contractAbiJSON, contractMethod string) (Output, error) {
	contractABI, err := abi.JSON(strings.NewReader(contractAbiJSON))
	if err != nil {
		return nil, err
	}
	if len(secretKey) == 0 {
		return nil, errors.New("secretkey is empty")
	}
	return &ethereumContract{
		chainEndpoint:   chainEndpoint,
		secretKey:       secretKey,
		contractAddress: contractAddress,
		contractABI:     contractABI,
		contractMethod:  contractMethod,
	}, nil
}
