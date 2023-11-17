package handler

import (
	"context"
	"log/slog"

	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/msg/messages"
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
	messages.New(msg)
	return r.mq.Enqueue(msg)
}

func (r *Handler) asyncHandle(m *msg.Msg) {
	slog.Debug("message popped", "message_id", m.ID)

	project, err := r.projectManager.Get(m.ProjectID)
	if err != nil {
		slog.Error("get project failed: ", err)
		messages.OnFailed(m.ID, err)
		return
	}

	messages.OnSubmitProving(m.ID)
	res, err := r.vmHandler.Handle(m, project.Config.VMType, project.Config.Code, project.Config.CodeExpParam)
	if err != nil {
		slog.Error("proof failed: ", err)
		messages.OnFailed(m.ID, err)
		return
	}
	slog.Debug("proof result", "proof_result", string(res))
	messages.OnProved(m.ID, string(res))

	if r.operatorPrivateKey == "" {
		info := "missing operator private key, will not write to chain"
		slog.Debug(info)
		messages.OnSucceeded(m.ID, info)
		return
	}

	data, err := contract.BuildData(res)
	if err != nil {
		slog.Error(err.Error())
		messages.OnFailed(m.ID, err)
		return
	}

	slog.Debug("writing proof to chain")

	messages.OnSubmitToBlockchain(m.ID)
	txHash, err := eth.SendTX(context.Background(), r.chainEndpoint, r.operatorPrivateKey, "0x6e30b42554DDA34bAFca9cB00Ec4B464f452a671", data)
	if err != nil {
		slog.Error(err.Error())
		messages.OnFailed(m.ID, err)
		return
	}
	messages.OnSucceeded(m.ID, txHash)
	slog.Debug("transaction hash", "tx_hash", txHash)
}
