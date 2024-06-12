package vm

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/task"
)

type Type string

const (
	Risc0  Type = "risc0"
	Halo2  Type = "halo2"
	ZKwasm Type = "zkwasm"
	Wasm   Type = "wasm"
)

type Handler struct {
	vmServerEndpoints map[Type]string
}

func (r *Handler) Handle(task *task.Task, vmtype Type, code string, expParam string) ([]byte, error) {
	endpoint, ok := r.vmServerEndpoints[vmtype]
	if !ok {
		return nil, errors.New("unsupported vm type")
	}

	ins, err := newInstance(context.Background(), task.ProjectID, endpoint, code, expParam)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new instance")
	}
	defer ins.release()
	slog.Debug("acquire vm instance success", "vm_type", vmtype)

	res, err := ins.execute(context.Background(), task)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute instance")
	}
	return res, nil
}

func NewHandler(vmServerEndpoints map[Type]string) *Handler {
	return &Handler{
		vmServerEndpoints: vmServerEndpoints,
	}
}
