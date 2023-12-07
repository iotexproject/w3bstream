package adapter

import (
	"encoding/hex"
	"log/slog"
)

// Stdout is an output adapter that writes the proof to stdout.
type Stdout struct{}

// NewStdout returns a new Stdout output adapter.
func NewStdout() *Stdout {
	return &Stdout{}
}

// Output writes the proof to stdout.
func (r *Stdout) Output(proof []byte) (Result, error) {
	slog.Info("stdout", "proof", hex.EncodeToString(proof))
	return "", nil
}
