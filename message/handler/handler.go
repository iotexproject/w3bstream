package handler

import (
	"context"
	"log/slog"

	"github.com/machinefi/sprout/message"
	"github.com/machinefi/sprout/output/chain/eth"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/tasks"
	"github.com/machinefi/sprout/test/contract"
	"github.com/machinefi/sprout/util/mq"
	"github.com/machinefi/sprout/util/mq/gochan"
	"github.com/machinefi/sprout/vm"
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

func (r *Handler) Handle(msg *message.Message) error {
	slog.Debug("push message into sequencer")
	tasks.New(msg)
	return r.mq.Enqueue(msg)
}

func (r *Handler) asyncHandle(m *message.Message) {
	slog.Debug("message popped", "message_id", m.ID)

	project, err := r.projectManager.Get(m.ProjectID)
	if err != nil {
		slog.Error("get project failed:", "error", err)
		tasks.OnFailed(m.ID, err)
		return
	}

	tasks.OnSubmitProving(m.ID)
	res, err := r.vmHandler.Handle(m, project.Config.VMType, project.Config.Code, project.Config.CodeExpParam)
	if err != nil {
		slog.Error("proof failed:", "error", err)
		tasks.OnFailed(m.ID, err)
		return
	}
	slog.Debug("proof result", "proof_result", string(res))
	tasks.OnProved(m.ID, string(res))

	if r.operatorPrivateKey == "" {
		info := "missing operator private key, will not write to chain"
		slog.Debug(info)
		tasks.OnSucceeded(m.ID, info)
		return
	}

	data, err := contract.BuildData(res)
	if err != nil {
		slog.Error(err.Error())
		tasks.OnFailed(m.ID, err)
		return
	}

	slog.Debug("writing proof to chain")

	tasks.OnSubmitToBlockchain(m.ID)
	txHash, err := eth.SendTX(context.Background(), r.chainEndpoint, r.operatorPrivateKey, "0x6e30b42554DDA34bAFca9cB00Ec4B464f452a671", data)
	if err != nil {
		slog.Error(err.Error())
		tasks.OnFailed(m.ID, err)
		return
	}
	tasks.OnSucceeded(m.ID, txHash)
	slog.Debug("transaction hash", "tx_hash", txHash)
}
