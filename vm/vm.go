package vm

import (
	"context"
	"log/slog"

	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/vm/server"
	"github.com/pkg/errors"
)

type Handler struct {
	endpoints   map[Type]string
	instanceMgr *server.Mgr
}

func (r *Handler) Handle(msg *msg.Msg, vmtype Type, code []byte, expParam string) ([]byte, error) {
	// TODO get project bin data by real project info
	// code, expParam := data.GetTestData(r.projectConfigFilePath)

	endpoint, ok := r.endpoints[vmtype]
	if !ok {
		return nil, errors.New("unsupported vm type")
	}

	ins, err := r.instanceMgr.Acquire(msg, endpoint, code, expParam)
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

func NewHandler(risc0Endpoint, halo2Endpoint string) *Handler {
	return &Handler{
		endpoints: map[Type]string{
			Risc0: risc0Endpoint,
			Halo2: halo2Endpoint,
		},
		instanceMgr: server.NewMgr(),
	}
}
