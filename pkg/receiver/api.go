package receiver

import (
	"github.com/gin-gonic/gin"
	"github.com/machinefi/w3bstream-mainnet/pkg/mq"
)

type MsgReceiver struct {
	router *gin.Engine
}

func New(q mq.MQ) *MsgReceiver {
	router := gin.Default()
	handler := &handler{q}
	router.POST("/message", handler.Receive)
	return &MsgReceiver{
		router: router,
	}
}

// this method will block the calling goroutine indefinitely unless an error happens.
func (r *MsgReceiver) Run(addr string) error {
	return r.router.Run(addr)
}
