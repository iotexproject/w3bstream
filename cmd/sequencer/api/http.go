package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/message"
	"github.com/machinefi/sprout/sequencer"
	"github.com/machinefi/sprout/tasks"
)

type errResp struct {
	Error string `json:"error,omitempty"`
}

func newErrResp(err error) *errResp {
	return &errResp{Error: err.Error()}
}

type handleMessageReq struct {
	ProjectID      uint64 `json:"projectID"        binding:"required"`
	ProjectVersion string `json:"projectVersion"   binding:"required"`
	Data           string `json:"data"             binding:"required"`
}

type handleMessageResp struct {
	MessageID string `json:"messageID"`
}

type HttpServer struct {
	engine *gin.Engine
	seq    *sequencer.Sequencer
}

func NewHttpServer(seq *sequencer.Sequencer) *HttpServer {
	s := &HttpServer{
		engine: gin.Default(),
		seq:    seq,
	}

	s.engine.POST("/message", s.handleMessage)
	s.engine.GET("/message/:id", s.queryByID)

	return s
}

// this func will block caller
func (s *HttpServer) Run(endpoint string) error {
	if err := s.engine.Run(endpoint); err != nil {
		return errors.Wrap(err, "start http server failed")
	}
	return nil
}

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
