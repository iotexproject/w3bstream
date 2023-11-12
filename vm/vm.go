package vm

import (
	"context"
	"log/slog"

	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/vm/server"
	"github.com/machinefi/w3bstream-mainnet/vm/types"
	"github.com/pkg/errors"
)

type Handler struct {
	endpoints             map[types.Type]string
	projectConfigFilePath string

	instanceMgr *server.Mgr
}

func (r *Handler) Handle(msg *msg.Msg) ([]byte, error) {
	// TODO get project bin data by real project info
	testdata := getTestData(r.projectConfigFilePath)

	endpoint, ok := r.endpoints[testdata.VMType]
	if !ok {
		return nil, errors.New("unsupported vm type")
	}

	ins, err := r.instanceMgr.Acquire(msg, endpoint, testdata.Code, testdata.CodeExpParam)
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

func NewHandler(risc0Endpoint, halo2Endpoint, projectConfigFilePath string) *Handler {
	return &Handler{
		endpoints: map[types.Type]string{
			types.Risc0: risc0Endpoint,
			types.Halo2: halo2Endpoint,
		},
		projectConfigFilePath: projectConfigFilePath,
		instanceMgr:           server.NewMgr(),
	}
}
