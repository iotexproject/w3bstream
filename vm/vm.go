package vm

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/vm/server"
)

type Type string

const (
	Risc0  Type = "risc0"
	Halo2  Type = "halo2"
	ZKwasm Type = "zkwasm"
)

type Handler struct {
	vmServerEndpoints map[Type]string
	instanceMgr       *server.Mgr
}

func (r *Handler) Handle(projectID uint64, vmtype Type, code string, expParam string, data [][]byte) ([]byte, error) {
	endpoint, ok := r.vmServerEndpoints[vmtype]
	if !ok {
		return nil, errors.New("unsupported vm type")
	}

	ins, err := r.instanceMgr.Acquire(projectID, endpoint, code, expParam)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get instance")
	}
	slog.Debug(fmt.Sprintf("acquire %s instance success", vmtype))
	defer r.instanceMgr.Release(projectID, ins)

	res, err := ins.Execute(context.Background(), projectID, data)
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
