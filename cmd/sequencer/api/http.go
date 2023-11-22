package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/message"
	"github.com/machinefi/sprout/sequencer"
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

type stateLog struct {
	State       string    `json:"state"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
}

type queryMessageStateLogResp struct {
	MessageID string      `json:"messageID"`
	States    []*stateLog `json:"states"`
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
	s.engine.GET("/message/:id", s.queryStateLogByID)

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

func (s *HttpServer) queryStateLogByID(c *gin.Context) {
	messageID := c.Param("id")

	ls, err := s.seq.FetchStateLog(messageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	ss := []*stateLog{}
	for _, l := range ls {
		ss = append(ss, &stateLog{
			State:       l.State.String(),
			Time:        l.CreatedAt,
			Description: l.Comment,
		})
	}

	slog.Debug("received message querying", "message_id", messageID)
	c.JSON(http.StatusOK, &queryMessageStateLogResp{MessageID: messageID, States: ss})
}
