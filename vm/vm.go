package vm

import (
	"context"
	"encoding/hex"
	"fmt"
	"log/slog"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm/server"
)

type Handler struct {
	vmServerEndpoints map[types.VM]string
	instanceMgr       *server.Mgr
}

func (r *Handler) Handle(task *types.Task, vmtype types.VM, code string, expParam string) ([]byte, error) {
	if len(task.Data) == 0 {
		return nil, errors.New("empty task")
	}
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
	raw, err := hex.DecodeString(string(res))
	if err != nil {
		return nil, errors.Wrap(err, "hex decode proof failed")
	}
	return raw, nil
}

func NewHandler(vmServerEndpoints map[types.VM]string) *Handler {
	return &Handler{
		vmServerEndpoints: vmServerEndpoints,
		instanceMgr:       server.NewMgr(),
	}
}
