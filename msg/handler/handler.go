package handler

import (
	"context"
	"log/slog"

	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/output/chain/eth"
	"github.com/machinefi/w3bstream-mainnet/project"
	"github.com/machinefi/w3bstream-mainnet/test/contract"
	"github.com/machinefi/w3bstream-mainnet/util/mq"
	"github.com/machinefi/w3bstream-mainnet/util/mq/gochan"
	"github.com/machinefi/w3bstream-mainnet/vm"
)

type Handler struct {
	mq                 mq.MQ
	vmHandler          *vm.Handler
	projectManager     *project.Manager
	chainEndpoint      string
	operatorPrivateKey string
}

func New(vmHandler *vm.Handler, projectManager *project.Manager, chainEndpoint, operatorPrivateKey string) *Handler {
	q := gochan.New()
	h := &Handler{
		mq:                 q,
		vmHandler:          vmHandler,
		chainEndpoint:      chainEndpoint,
		operatorPrivateKey: operatorPrivateKey,
		projectManager:     projectManager,
	}
	go q.Watch(h.asyncHandle)
	return h
}

func (r *Handler) Handle(msg *msg.Msg) error {
	slog.Debug("push message into sequencer")
	return r.mq.Enqueue(msg)
}

func (r *Handler) asyncHandle(m *msg.Msg) {
	slog.Debug("message popped by proofer")
	project, err := r.projectManager.Get(m.ProjectID)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	res, err := r.vmHandler.Handle(m, project.Config.VMType, project.Config.Code, project.Config.CodeExpParam)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Debug("vm generate proof success, the proof is")
	slog.Debug(string(res))

	data, err := contract.BuildData(res)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Debug("writing proof to chain")
	txHash, err := eth.SendTX(context.Background(), r.chainEndpoint, r.operatorPrivateKey, "0xbc3c770272a8d274ba75ce2a104df397f7ca793e", data)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Debug("writing proof to chain success, the transaction hash is")
	slog.Debug(txHash)
}
