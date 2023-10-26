package handler

import (
	"context"
	"log/slog"

	"github.com/machinefi/w3bstream-mainnet/pkg/msg"
	"github.com/machinefi/w3bstream-mainnet/pkg/output/chain/eth"
	"github.com/machinefi/w3bstream-mainnet/pkg/util/mq"
	"github.com/machinefi/w3bstream-mainnet/pkg/util/mq/gochan"
	"github.com/machinefi/w3bstream-mainnet/pkg/vm"
	"github.com/machinefi/w3bstream-mainnet/test/sol"
)

type Handler struct {
	mq        mq.MQ
	vmHandler *vm.Handler
}

func New(vmHandler *vm.Handler) *Handler {
	q := gochan.New()
	h := &Handler{
		mq:        q,
		vmHandler: vmHandler,
	}
	go q.Watch(h.asyncHandle)
	return h
}

func (r *Handler) Handle(msg *msg.Msg) error {
	return r.mq.Enqueue(msg)
}

func (r *Handler) asyncHandle(m *msg.Msg) {
	res, err := r.vmHandler.Handle(m)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	data, err := sol.BuildData(res)
	txHash, err := eth.SendTX(context.Background(), "https://babel-api.testnet.iotex.io", "1dd02a766b13a3dd1d7bcaf928449ce0d1b56815caf96800f8e772326eb03012", "0x190Cc9af23504ac5Dc461376C1e2319bc3B9cD29", data)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(txHash)
}
