package output

import (
	"log/slog"

	"github.com/iotexproject/w3bstream/task"
)

type stdout struct{}

func (r *stdout) Output(proverID uint64, task *task.Task, proof []byte) (string, error) {
	slog.Info("stdout", "proof", string(proof))
	return "", nil
}

func newStdout() *stdout {
	return &stdout{}
}
