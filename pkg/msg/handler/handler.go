package handler

import (
	"context"
	"log/slog"

	"github.com/machinefi/w3bstream-mainnet/pkg/msg"
	"github.com/machinefi/w3bstream-mainnet/pkg/output/chain/eth"
	"github.com/machinefi/w3bstream-mainnet/pkg/util/mq"
	"github.com/machinefi/w3bstream-mainnet/pkg/util/mq/gochan"
	"github.com/machinefi/w3bstream-mainnet/pkg/vm"
	"github.com/machinefi/w3bstream-mainnet/test/contract"
)

type Handler struct {
	mq                 mq.MQ
	vmHandler          *vm.Handler
	chainEndpoint      string
	operatorPrivateKey string
}

func New(vmHandler *vm.Handler, chainEndpoint, operatorPrivateKey string) *Handler {
	q := gochan.New()
	h := &Handler{
		mq:                 q,
		vmHandler:          vmHandler,
		chainEndpoint:      chainEndpoint,
		operatorPrivateKey: operatorPrivateKey,
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
	res, err := r.vmHandler.Handle(m)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	data, err := contract.BuildData(res)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Debug("writing proof to chain")
	txHash, err := eth.SendTX(context.Background(), r.chainEndpoint, r.operatorPrivateKey, "0x190Cc9af23504ac5Dc461376C1e2319bc3B9cD29", data)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Debug("writing proof to chain success, the transaction hash is")
	slog.Debug(txHash)
}
