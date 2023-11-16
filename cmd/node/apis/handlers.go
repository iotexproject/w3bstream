package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/msg/messages"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
)

func (s *Server) handleRequest(c *gin.Context) {
	req := &msgReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, newErrResp(err))
		return
	}
	slog.Debug("received your message, handling")
	if err := s.msgHandler.Handle(&msg.Msg{
		ProjectID:      req.ProjectID,
		ProjectVersion: req.ProjectVersion,
		Data:           req.Data,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) queryByMessageID(c *gin.Context) {
	messageID := c.Param("messageID")

	slog.Debug("received message querying", "message_id", messageID)
	m, ok := messages.Query(messageID)
	if !ok {
		c.JSON(http.StatusNotFound, newErrResp(errors.Errorf("message id %s expired or not exists", messageID)))
		return
	}

	c.JSON(http.StatusOK, m)
}
