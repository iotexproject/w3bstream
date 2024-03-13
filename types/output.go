package types

type Output string

const (
	OutputStdout           Output = "stdout"
	OutputEthereumContract Output = "ethereumContract"
	OutputSolanaProgram    Output = "solanaProgram"
	OutputTextile          Output = "textile"
)
