package vm

import (
	"context"
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

func (r *Handler) Handle(msgs []*types.Message, vmtype types.VM, code string, expParam string) ([]byte, error) {
	if len(msgs) == 0 {
		return nil, errors.New("missing messages")
	}
	endpoint, ok := r.vmServerEndpoints[vmtype]
	if !ok {
		return nil, errors.New("unsupported vm type")
	}

	ins, err := r.instanceMgr.Acquire(msgs[0].ProjectID, endpoint, code, expParam)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get instance")
	}
	slog.Debug(fmt.Sprintf("acquire %s instance success", vmtype))
	defer r.instanceMgr.Release(msgs[0].ProjectID, ins)

	res, err := ins.Execute(context.Background(), msgs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute instance")
	}
	return res, nil
}

func NewHandler(vmServerEndpoints map[types.VM]string) *Handler {
	return &Handler{
		vmServerEndpoints: vmServerEndpoints,
		instanceMgr:       server.NewMgr(),
	}
}
