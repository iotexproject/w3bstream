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
	risc0ServerAddr       string
	halo2ServerAddr       string
	projectConfigFilePath string

	instanceMgr *server.Mgr
}

func (r *Handler) Handle(msg *msg.Msg) ([]byte, error) {
	// TODO get project bin data by real project info
	testdata := getTestData(r.projectConfigFilePath)
	serverAddr := r.halo2ServerAddr
	if testdata.VMType == types.Risc0 {
		serverAddr = r.risc0ServerAddr
	}

	ins, err := r.instanceMgr.Acquire(msg, serverAddr, testdata.Code, testdata.CodeExpParam)
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
		risc0ServerAddr:       risc0ServerAddr,
		halo2ServerAddr:       halo2ServerAddr,
		projectConfigFilePath: projectConfigFilePath,
		instanceMgr:           server.NewMgr(),
	}
}
