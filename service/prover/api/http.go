package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/metrics"
	"github.com/iotexproject/w3bstream/service/prover/db"
)

type errResp struct {
	Error string `json:"error,omitempty"`
}

func newErrResp(err error) *errResp {
	return &errResp{Error: err.Error()}
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
	taskIDStr := c.Param("id")
	taskID := common.HexToHash(taskIDStr)

	processed, errMsg, createdAt, err := s.db.ProcessedTask(taskID)
	if err != nil {
		slog.Error("failed to query processed task", "error", err)
		c.JSON(http.StatusInternalServerError, newErrResp(errors.Wrap(err, "failed to query processed task")))
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

	s.engine.GET("/task/:id", s.queryTask)
	metrics.RegisterMetrics(s.engine)

	if err := s.engine.Run(address); err != nil {
		slog.Error("failed to start http server", "address", address, "error", err)
		return errors.Wrap(err, "could not start http server; check if the address is in use or network is accessible")
	}
	return nil
}
