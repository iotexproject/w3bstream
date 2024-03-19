package output

import (
	"encoding/hex"
	"log/slog"
)

type stdout struct{}

func (r *stdout) Output(projectID uint64, taskData [][]byte, proof []byte) (string, error) {
	slog.Info("stdout", "proof", hex.EncodeToString(proof))
	return "", nil
}

func newStdout() Output {
	return &stdout{}
}
