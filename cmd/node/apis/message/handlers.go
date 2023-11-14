package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/machinefi/w3bstream-mainnet/msg"
	msghandler "github.com/machinefi/w3bstream-mainnet/msg/handler"
	"github.com/machinefi/w3bstream-mainnet/msg/messages"
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
	messageID := uuid.NewString()
	if err := msghandler.DefaultHandler.Handle(&msg.Msg{
		ID:             messageID,
		ProjectID:      req.ProjectID,
		ProjectVersion: req.ProjectVersion,
		Data:           req.Data,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, &HandleErrRsp{err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, &HandleRsp{MessageID: messageID})
}

func queryByMessageID(c *gin.Context) {
	messageID := c.Param("messageID")

	slog.Debug("received message querying", "message_id", messageID)
	m, ok := messages.Query(messageID)
	if !ok {
		c.JSON(http.StatusNotFound, HandleErrRsp{
			Error: fmt.Sprintf("message id %s expired or not exists", messageID),
		})
		return
	}

	c.JSON(http.StatusOK, m)
}

func Register(eng *gin.Engine) {
	eng.POST("/message", handleRequest)
	eng.GET("/message/:messageID", queryByMessageID)
}
