package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/sequencer"
)

func NewHttpServer(seq *sequencer.Sequencer) *HttpServer {
	s := &HttpServer{
		engine: gin.Default(),
		seq:    seq,
	}

	s.engine.POST("/message", s.handleMessage)
	s.engine.GET("/message/:id", s.queryByID)

	return s
}

type HttpServer struct {
	engine *gin.Engine
	seq    *sequencer.Sequencer
}

// this func will block caller
func (s *HttpServer) Run(endpoint string) error {
	if err := s.engine.Run(endpoint); err != nil {
		return errors.Wrap(err, "start http server failed")
	}
	return nil
}
