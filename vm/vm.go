package vm

import (
	"context"
	"log/slog"

	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/vm/instance/manager"
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
	slog.Debug("acquire risc0 instance success")
	defer r.instanceMgr.Release(msg.Key(), ins)

	res, err := ins.Execute(context.Background(), msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute instance")
	}
	slog.Debug("ask risc0 generate proof success, the proof is")
	slog.Debug(string(res))
	return res, nil
}

func NewHandler(risc0ServerAddr, halo2ServerAddr, projectConfigFilePath string) *Handler {
	return &Handler{
		manager.NewMgr(&manager.Config{
			Risc0ServerAddr:       risc0ServerAddr,
			Halo2ServerAddr:       halo2ServerAddr,
			ProjectConfigFilePath: projectConfigFilePath,
		}),
	}
}
