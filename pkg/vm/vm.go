package vm

import (
	"context"

	"github.com/machinefi/w3bstream-mainnet/pkg/msg"
	"github.com/machinefi/w3bstream-mainnet/pkg/vm/instance/manager"
	"github.com/pkg/errors"
)

type Handler struct {
	instanceMgr *manager.Mgr
}

func (r *Handler) Handle(msg *msg.Msg) ([]byte, error) {
	ins, err := r.instanceMgr.Acquire(msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get instance")
	}
	defer r.instanceMgr.Release(msg.Key(), ins)

	res, err := ins.Execute(context.Background(), msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute instance")
	}
	return res, nil
}

func NewHandler(risc0ServerAddr string) *Handler {
	return &Handler{
		manager.NewMgr(&manager.Config{
			Risc0ServerAddr: risc0ServerAddr,
		}),
	}
}
