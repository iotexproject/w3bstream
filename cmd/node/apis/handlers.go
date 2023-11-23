package apis

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/proto"
	"github.com/machinefi/sprout/tasks"
)

func (s *Server) handleRequest(c *gin.Context) {
	req := &HandleReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, newErrResp(err))
		return
	}
	taskID := uuid.NewString()
	slog.Debug("received your message, handling")
	if err := s.msgHandler.Handle(&proto.Message{
		MessageID: taskID,
		ProjectID: req.ProjectID,
		Data:      req.Data,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}
	slog.Debug("message was handled", "taskID", taskID)
	c.JSON(http.StatusOK, &HandleRsp{TaskID: taskID})
}

func (s *Server) queryByTaskID(c *gin.Context) {
	taskID := c.Param("taskID")

	slog.Debug("received task querying", "task_id", taskID)
	m, ok := tasks.Query(taskID)
	if !ok {
		c.JSON(http.StatusNotFound, newErrResp(errors.Errorf("task [%s] was expired or not exists", taskID)))
		return
	}

	c.JSON(http.StatusOK, m)
}
