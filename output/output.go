package output

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/types"
)

var (
	errInvalidEthereumSecretKey    = errors.New("invalid ethereum secret key")
	errInvalidEthereumABI          = errors.New("invalid ethereum abi")
	errInvalidEthereumMethod       = errors.New("invalid ethereum method")
	errInvalidEthereumOutputConfig = errors.New("invalid ethereum config")
	errInvalidSolanaSecretKey      = errors.New("invalid solana secret key  ")
	errInvalidSolanaOutputConfig   = errors.New("invalid solana config")
	errInvalidTextileSecretKey     = errors.New("invalid textile secret key ")
	errInvalidTextileOutputConfig  = errors.New("invalid textile config")
)

type Output interface {
	Output(task *types.Task, proof []byte) (string, error)
	Type() types.Output
}

type Config struct {
	Type     types.Output    `json:"type"`
	Ethereum *ethereumConfig `json:"ethereum,omitempty"`
	Solana   *solanaConfig   `json:"solana,omitempty"`
	Textile  *textileConfig  `json:"textile,omitempty"`

	prvKeyECDSA   string
	prvKeyED25519 string
}

func (c *Config) SetPrivateKey(skECDSA, skED25519 string) *Config {
	c.prvKeyED25519 = skED25519
	c.prvKeyECDSA = skECDSA
	return c
}

func (c *Config) Output() (Output, error) {
	switch c.Type {
	default:
		return &stdout{}, nil
	case types.OutputEthereumContract:
		cc := c.Ethereum
		if cc == nil {
			return nil, errInvalidEthereumOutputConfig
		}
		cc.sk = c.prvKeyECDSA
		if len(cc.sk) == 0 {
			return nil, errInvalidEthereumSecretKey
		}
		contractABI, err := abi.JSON(strings.NewReader(cc.ContractAbiJSON))
		if err != nil {
			return nil, errInvalidEthereumABI
		}
		contractMethod, ok := contractABI.Methods[cc.ContractMethod]
		if !ok {
			return nil, errInvalidEthereumMethod
		}

		return &ethereumContract{
			chainEndpoint:   cc.ChainEndpoint,
			secretKey:       cc.sk,
			contractAddress: cc.ContractAddress,
			receiverAddress: cc.ReceiverAddress,
			contractABI:     contractABI,
			contractMethod:  contractMethod,
		}, nil
	case types.OutputSolanaProgram:
		cc := c.Solana
		if cc == nil {
			return nil, errInvalidSolanaOutputConfig
		}
		cc.sk = c.prvKeyED25519
		if len(cc.sk) == 0 {
			return nil, errInvalidSolanaSecretKey
		}
		return &solanaProgram{
			endpoint:       cc.ChainEndpoint,
			programID:      cc.ProgramID,
			secretKey:      cc.sk,
			stateAccountPK: cc.StateAccountPK,
		}, nil
	case types.OutputTextile:
		cc := c.Textile
		if cc == nil {
			return nil, errInvalidTextileOutputConfig
		}
		cc.sk = c.prvKeyECDSA
		if len(cc.sk) == 0 {
			return nil, errInvalidTextileSecretKey
		}

		pk := crypto.ToECDSAUnsafe(common.FromHex(cc.sk))
		return &textileDB{
			endpoint:  fmt.Sprintf("https://basin.tableland.xyz/vaults/%s/events", cc.VaultID),
			secretKey: pk,
		}, nil
	}
}

type ethereumConfig struct {
	ChainEndpoint   string `json:"chainEndpoint"`
	ContractAddress string `json:"contractAddress"`
	ReceiverAddress string `json:"receiverAddress,omitempty"`
	ContractMethod  string `json:"contractMethod"`
	ContractAbiJSON string `json:"contractAbiJSON"`

	sk string
}

type solanaConfig struct {
	ChainEndpoint  string `json:"chainEndpoint"`
	ProgramID      string `json:"programID"`
	StateAccountPK string `json:"stateAccountPK"`

	sk string
}

type textileConfig struct {
	VaultID string `json:"vaultID"`

	sk string
}
