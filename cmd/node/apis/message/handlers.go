package message

import (
	"github.com/gin-gonic/gin"
	"github.com/machinefi/w3bstream-mainnet/msg"
	msghandler "github.com/machinefi/w3bstream-mainnet/msg/handler"
	"log/slog"
	"net/http"
)

func handleRequest(c *gin.Context) {
	req := HandleReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, &HandleErrRsp{err.Error()})
		return
	}
	slog.Debug("received your message, handling")
	if err := msghandler.DefaultHandler.Handle(&msg.Msg{
		ProjectID:      req.ProjectID,
		ProjectVersion: req.ProjectVersion,
		Data:           req.Data,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, &HandleErrRsp{err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func Register(eng *gin.Engine) {
	base := "/message"
	eng.POST(base, handleRequest)
}
