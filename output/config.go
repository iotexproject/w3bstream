package output

type (
	// Config is the configuration for the outputter
	Config struct {
		Type Type

		// for the ethereum & solana contract outputter
		ChainEndpoint   string
		ContractAddress string
		SecretKey       string
		// for the solana program outputter
		StateAccountPK string
	}

	// Type is the type of outputter
	Type string
)

// Types of outputters
const (
	Stdout           Type = "stdout"
	EthereumContract Type = "ethereumContract"
	SolanaProgram    Type = "solanaProgram"
)
