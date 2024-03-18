package output

import (
	"encoding/hex"
	"log/slog"

	"github.com/machinefi/sprout/types"
)

type stdout struct{}

func (r *stdout) Type() types.Output {
	return types.OutputStdout
}

func (r *stdout) Output(task *types.Task, proof []byte) (string, error) {
	slog.Info("stdout", "proof", hex.EncodeToString(proof))
	return "", nil
}
