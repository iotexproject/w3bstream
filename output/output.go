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
		// TODO: update parameters of newSolanaProgram
		solConf := conf.Solana
		return newSolanaProgram(solConf.ChainEndpoint, solConf.ProgramID, privateKeyED25519, solConf.StateAccountPK)
	case Textile:
		textileConf := conf.Textile
		return newTextileDBAdapter(textileConf.VaultID, privateKeyECDSA)
	default:
		return newStdout(), nil
	}
}
