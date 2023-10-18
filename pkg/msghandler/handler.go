package msghandler

import (
	"github.com/machinefi/w3bstream-mainnet/pkg/msg"
	"github.com/machinefi/w3bstream-mainnet/pkg/util/mq"
	"github.com/machinefi/w3bstream-mainnet/pkg/util/mq/gochan"
)

type Handler struct {
	mq mq.MQ
}

func New() *Handler {
	return &Handler{
		mq: gochan.New(),
	}
}

func (r *Handler) Handle(msg *msg.Msg) error {
	return r.mq.Enqueue(msg)
}
