package vm

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm/server"
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
	instanceMgr       *server.Mgr
}

func (r *Handler) Handle(task *types.Task, vmtype Type, code string, expParam string) ([]byte, error) {
	endpoint, ok := r.vmServerEndpoints[vmtype]
	if !ok {
		return nil, errors.New("unsupported vm type")
	}

	ins, err := r.instanceMgr.Acquire(task.ProjectID, endpoint, code, expParam)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get instance")
	}
	slog.Debug(fmt.Sprintf("acquire %s instance success", vmtype))
	defer r.instanceMgr.Release(task.ProjectID, ins)

	res, err := ins.Execute(context.Background(), task)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute instance")
	}
	return res, nil
}

func NewHandler(vmServerEndpoints map[Type]string) *Handler {
	return &Handler{
		vmServerEndpoints: vmServerEndpoints,
		instanceMgr:       server.NewMgr(),
	}
}
