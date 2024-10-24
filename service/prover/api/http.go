package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/service/prover/db"
)

type ErrResp struct {
	Error string `json:"error,omitempty"`
}

func NewErrResp(err error) *ErrResp {
	return &ErrResp{Error: err.Error()}
}

type QueryTaskReq struct {
	ProjectID uint64 `json:"projectID"                  binding:"required"`
	TaskID    string `json:"taskID"                     binding:"required"`
}

type QueryTaskResp struct {
	Time      time.Time `json:"time"`
	Processed bool      `json:"processed"`
	Error     string    `json:"error,omitempty"`
}

type httpServer struct {
	engine *gin.Engine
	db     *db.DB
}

func (s *httpServer) queryTask(c *gin.Context) {
	req := &QueryTaskReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		slog.Error("failed to bind request", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(errors.Wrap(err, "invalid request payload")))
		return
	}
	taskID := common.HexToHash(req.TaskID)

	processed, errMsg, createdAt, err := s.db.ProcessedTask(req.ProjectID, taskID)
	if err != nil {
		slog.Error("failed to query processed task", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to query processed task")))
		return
	}

	c.JSON(http.StatusOK, &QueryTaskResp{
		Time:      createdAt,
		Processed: processed,
		Error:     errMsg,
	})
}

// this func will block caller
func Run(db *db.DB, address string) error {
	s := &httpServer{
		engine: gin.Default(),
		db:     db,
	}

	s.engine.GET("/task", s.queryTask)

	if err := s.engine.Run(address); err != nil {
		slog.Error("failed to start http server", "address", address, "error", err)
		return errors.Wrap(err, "could not start http server; check if the address is in use or network is accessible")
	}
	return nil
}
