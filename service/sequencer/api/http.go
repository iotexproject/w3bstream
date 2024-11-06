package api

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/metrics"
)

type httpServer struct {
	engine *gin.Engine
}

// this func will block caller
func Run(address string) error {
	if len(address) == 0 {
		slog.Warn("http server address is empty")
		return nil
	}
	s := &httpServer{
		engine: gin.Default(),
	}

	metrics.RegisterMetrics(s.engine)

	if err := s.engine.Run(address); err != nil {
		slog.Error("failed to start http server", "address", address, "error", err)
		return errors.Wrap(err, "could not start http server; check if the address is in use or network is accessible")
	}
	return nil
}
