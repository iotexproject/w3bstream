package output

import "github.com/machinefi/sprout/output/chain"

type (
	// Config is the configuration for the outputter
	Config struct {
		Type            Type
		ChainName       chain.Name
		ContractAddress string
		SecretKey       string
		StateAccountPK  string
	}
)

// NewStdoutConfig creates a stdout config
func NewStdoutConfig() Config {
	return Config{
		Type: Stdout,
	}
}

// NewEthereumContractConfig creates an ethereum contract config
func NewEthereumContractConfig(chainName chain.Name, contractAddress, secretKey string) Config {
	return Config{
		Type:            EthereumContract,
		ChainName:       chainName,
		ContractAddress: contractAddress,
		SecretKey:       secretKey,
	}
}

// NewSolanaProgramConfig creates a solana program config
func NewSolanaProgramConfig(chainName chain.Name, programID, secretKey, stateAccountPK string) Config {
	return Config{
		Type:            SolanaProgram,
		ChainName:       chainName,
		ContractAddress: programID,
		SecretKey:       secretKey,
		StateAccountPK:  stateAccountPK,
	}
}
