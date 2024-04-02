package output

import (
	"log/slog"

	"github.com/machinefi/sprout/types"
)

type stdout struct{}

func (r *stdout) Output(task *types.Task, proof []byte) (string, error) {
	slog.Info("stdout", "proof", string(proof))
	return "", nil
}

func newStdout() Output {
	return &stdout{}
}
