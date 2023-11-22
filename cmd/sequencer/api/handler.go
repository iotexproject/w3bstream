package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/message"
	"github.com/machinefi/sprout/tasks"
)

func (s *HttpServer) handleMessage(c *gin.Context) {
	req := &handleMessageReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, newErrResp(err))
		return
	}

	id := uuid.NewString()
	slog.Debug("received your message, handling")
	if err := s.seq.Save(&message.Message{
		ID:             id,
		ProjectID:      req.ProjectID,
		ProjectVersion: req.ProjectVersion,
		Data:           req.Data,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	slog.Debug("message was handled", "messageID", id)
	c.JSON(http.StatusOK, &handleMessageResp{MessageID: id})
}

func (s *HttpServer) queryByID(c *gin.Context) {
	messageID := c.Param("id")

	slog.Debug("received message querying", "message_id", messageID)
	m, ok := tasks.Query(messageID)
	if !ok {
		c.JSON(http.StatusNotFound, newErrResp(errors.Errorf("message [%s] was expired or not exists", messageID)))
		return
	}

	c.JSON(http.StatusOK, m)
}
