package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/machinefi/w3bstream-mainnet/msg"
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
