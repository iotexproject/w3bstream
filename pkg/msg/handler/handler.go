package handler

import (
	"log/slog"

	"github.com/machinefi/w3bstream-mainnet/pkg/msg"
	"github.com/machinefi/w3bstream-mainnet/pkg/util/mq"
	"github.com/machinefi/w3bstream-mainnet/pkg/util/mq/gochan"
	"github.com/machinefi/w3bstream-mainnet/pkg/vm"
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
	}
	slog.Info(string(res))
}
