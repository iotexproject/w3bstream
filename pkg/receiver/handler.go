package receiver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machinefi/w3bstream-mainnet/pkg/mq"
)

type handler struct {
	q mq.MQ
}

type receiveMsgReq struct {
	Data string `json:"data"        binding:"required"`
}

func (h *handler) Receive(c *gin.Context) {
	var req receiveMsgReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, newErrResp(err))
		return
	}
	msg := &mq.Msg{
		Data: []byte(req.Data),
	}
	if err := h.q.Enqueue(msg); err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	c.Status(http.StatusOK)
}
