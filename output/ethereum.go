package output

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"

	"github.com/machinefi/sprout/types"
)

var (
	errSnarkProofDataMissingFieldSnark           = errors.New("missing field `Snark.snark`")
	errSnarkProofDataMissingFieldPostStateDigest = errors.New("missing field `Snark.post_state_digest`")
	errSnarkProofDataMissingFieldJournal         = errors.New("missing field `Snark.journal`")
	errMissingReceiverParam                      = errors.New("missing receiver param")
)

type ethereumContract struct {
	client          *ethclient.Client
	contractAddress common.Address
	receiverAddress string
	secretKey       *ecdsa.PrivateKey
	signer          ethtypes.Signer
	contractABI     abi.ABI
	contractMethod  abi.Method
}

func (e *ethereumContract) Output(task *types.Task, proof []byte) (string, error) {
	params := []interface{}{}
	for _, a := range e.contractMethod.Inputs {
		switch a.Name {
		case "proof", "_proof":
			params = append(params, proof)

		case "projectId", "_projectId":
			i := new(big.Int).SetUint64(task.ProjectID)
			params = append(params, i)

		case "receiver", "_receiver":
			if e.receiverAddress == "" {
				return "", errMissingReceiverParam
			}
			params = append(params, common.HexToAddress(e.receiverAddress))

		case "data_snark", "_data_snark":
			valueSeal := gjson.GetBytes(proof, "Snark.snark").String()
			if valueSeal == "" {
				return "", errSnarkProofDataMissingFieldSnark
			}
			valueDigest := gjson.GetBytes(proof, "Snark.post_state_digest").String()
			if valueDigest == "" {
				return "", errSnarkProofDataMissingFieldPostStateDigest
			}
			valueJournal := gjson.GetBytes(proof, "Snark.journal").String()
			if valueJournal == "" {
				return "", errSnarkProofDataMissingFieldJournal
			}

			abiBytes, err := abi.NewType("bytes", "", nil)
			if err != nil {
				return "", errors.Wrap(err, "new ethereum accounts abi pack failed")
			}
			args := abi.Arguments{
				{Type: abiBytes, Name: "proof_snark_seal"},
				{Type: abiBytes, Name: "proof_snark_post_state_digest"},
				{Type: abiBytes, Name: "proof_snark_journal"},
			}

			packed, err := args.Pack([]byte(valueSeal), []byte(valueDigest), []byte(valueJournal))
			if err != nil {
				return "", errors.Wrap(err, "ethereum accounts abi pack failed")
			}
			params = append(params, packed)

		default:
			value := gjson.GetBytes(task.Data[0], a.Name)
			param := value.String()
			if param == "" {
				return "", errors.Errorf("miss param %s for contract abi", a.Name)
			}
			switch a.Type.String() {
			case "address":
				params = append(params, common.HexToAddress(param))
			case "uint256":
				i := new(big.Int)
				i.SetString(strings.TrimPrefix(param, "0x"), 16)
				params = append(params, i)
			default:
				params = append(params, param)
			}
		}
	}
	calldata, err := e.contractABI.Pack(e.contractMethod.Name, params...)
	if err != nil {
		return "", errors.Wrap(err, "failed to pack by contract abi")
	}

	txHash, err := e.sendTX(context.Background(), calldata)
	if err != nil {
		return "", errors.Wrap(err, "failed to send transaction")
	}

	return txHash, nil
}

func (e *ethereumContract) sendTX(ctx context.Context, data []byte) (string, error) {
	sender := crypto.PubkeyToAddress(e.secretKey.PublicKey)

	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", errors.Wrap(err, "get suggest gas price failed")
	}

	nonce, err := e.client.PendingNonceAt(ctx, sender)
	if err != nil {
		return "", errors.Wrap(err, "get pending nonce failed")
	}

	msg := ethereum.CallMsg{
		From:     sender,
		To:       &e.contractAddress,
		GasPrice: gasPrice,
		Data:     data,
	}
	gasLimit, err := e.client.EstimateGas(ctx, msg)
	if err != nil {
		return "", errors.Wrap(err, "estimate gas failed")
	}

	tx := ethtypes.NewTx(
		&ethtypes.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      gasLimit,
			To:       &e.contractAddress,
			Data:     data,
		})

	signedTx, err := ethtypes.SignTx(tx, e.signer, e.secretKey)
	if err != nil {
		return "", errors.Wrap(err, "sign tx failed")
	}

	if err := e.client.SendTransaction(ctx, signedTx); err != nil {
		return "", errors.Wrap(err, "send transaction failed")
	}

	return signedTx.Hash().Hex(), nil
}

func newEthereum(cfg EthereumConfig, secretKey string) (*ethereumContract, error) {
	if cfg.ContractAddress == "" {
		return nil, errors.New("secret key is empty")
	}
	contractABI, err := abi.JSON(strings.NewReader(cfg.ContractAbiJSON))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode contract abi")
	}
	method, ok := contractABI.Methods[cfg.ContractMethod]
	if !ok {
		return nil, errors.New("the contract method not exist in abi")
	}
	cli, err := ethclient.Dial(cfg.ChainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "dial eth endpoint %s failed", cfg.ChainEndpoint)
	}
	chainid, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "get chain id failed")
	}

	return &ethereumContract{
		client:          cli,
		secretKey:       crypto.ToECDSAUnsafe(common.FromHex(secretKey)),
		signer:          ethtypes.NewLondonSigner(chainid),
		contractAddress: common.HexToAddress(cfg.ContractAddress),
		receiverAddress: cfg.ReceiverAddress,
		contractABI:     contractABI,
		contractMethod:  method,
	}, nil
}
