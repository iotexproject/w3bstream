package output

import (
	"errors"
)

type (
	// Outputter is the interface for outputting proofs
	Outputter interface {
		// Output outputs the proof
		Output(proof []byte) error
	}
)

// NewOutputter returns a new outputter based on the config
func NewOutputter(cfg Config) (out Outputter, err error) {
	switch cfg.Type {
	case Stdout:
		// TODO: implement
	case EvmContract:
		// TODO: implement
	default:
		return nil, errors.New("invalid output type")
	}
	return
}
