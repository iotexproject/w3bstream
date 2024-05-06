package output

import (
	"log/slog"

	"github.com/machinefi/sprout/task"
)

type stdout struct{}

func (r *stdout) Output(task *task.Task, proof []byte) (string, error) {
	slog.Info("stdout", "proof", string(proof))
	return "", nil
}

func newStdout() *stdout {
	return &stdout{}
}
