package output

type (
	// Config is the configuration for the outputter
	Config struct {
		Type Type
	}

	// Type is the type of outputter
	Type string
)

// Types of outputters
const (
	Stdout      Type = "stdout"
	EvmContract Type = "evmContract"
)
