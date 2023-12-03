package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/coordinator"
	"github.com/machinefi/sprout/types"
)

type errResp struct {
	Error string `json:"error,omitempty"`
}

func newErrResp(err error) *errResp {
	return &errResp{Error: err.Error()}
}

type handleMessageReq struct {
	ProjectID uint64 `json:"projectID"        binding:"required"`
	// TODO support project version
	ProjectVersion string `json:"projectVersion"   binding:"required"`
	Data           string `json:"data"             binding:"required"`
}

type handleMessageResp struct {
	MessageID string `json:"messageID"`
}

type stateLog struct {
	State   string    `json:"state"`
	Time    time.Time `json:"time"`
	Comment string    `json:"comment"`
}

type queryMessageStateLogResp struct {
	MessageID string      `json:"messageID"`
	States    []*stateLog `json:"states"`
}

type HttpServer struct {
	engine      *gin.Engine
	coordinator *coordinator.Coordinator
}

func NewHttpServer(c *coordinator.Coordinator) *HttpServer {
	s := &HttpServer{
		engine:      gin.Default(),
		coordinator: c,
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
	if err := s.coordinator.Save(&types.Message{
		MessageID: id,
		ProjectID: req.ProjectID,
		Data:      req.Data,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	slog.Debug("message was handled", "messageID", id)
	c.JSON(http.StatusOK, &handleMessageResp{MessageID: id})
}

func (s *HttpServer) queryStateLogByID(c *gin.Context) {
	messageID := c.Param("id")

	ls, err := s.coordinator.FetchStateLog(messageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	ss := []*stateLog{}
	for _, l := range ls {
		ss = append(ss, &stateLog{
			State:   l.State.String(),
			Time:    l.CreatedAt,
			Comment: l.Comment,
		})
	}

	slog.Debug("received message querying", "message_id", messageID)
	c.JSON(http.StatusOK, &queryMessageStateLogResp{MessageID: messageID, States: ss})
}
