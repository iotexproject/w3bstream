package output

import "github.com/machinefi/sprout/types"

type Type string

const (
	Stdout           Type = "stdout"
	EthereumContract Type = "ethereumContract"
	SolanaProgram    Type = "solanaProgram"
	Textile          Type = "textile"
)

type Config struct {
	Type     Type           `json:"type"`
	Ethereum EthereumConfig `json:"ethereum"`
	Solana   SolanaConfig   `json:"solana"`
	Textile  TextileConfig  `json:"textile"`
}

type EthereumConfig struct {
	ChainEndpoint   string `json:"chainEndpoint"`
	ContractAddress string `json:"contractAddress"`
	ReceiverAddress string `json:"receiverAddress,omitempty"`
	ContractMethod  string `json:"contractMethod"`
	ContractAbiJSON string `json:"contractAbiJSON"`
}

type SolanaConfig struct {
	ChainEndpoint  string `json:"chainEndpoint"`
	ProgramID      string `json:"programID"`
	StateAccountPK string `json:"stateAccountPK"`
}

type TextileConfig struct {
	VaultID string `json:"vaultID"`
}

type Output interface {
	Output(task *types.Task, proof []byte) (string, error)
}

func New(conf *Config, privateKeyECDSA, privateKeyED25519 string) (Output, error) {
	switch conf.Type {
	case EthereumContract:
		return newEthereum(conf.Ethereum, privateKeyECDSA)
	case SolanaProgram:
		return newSolanaProgram(conf.Solana, privateKeyED25519)
	case Textile:
		return newTextileDBAdapter(conf.Textile, privateKeyECDSA)
	default:
		return newStdout(), nil
	}
}
