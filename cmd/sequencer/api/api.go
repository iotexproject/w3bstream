package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/sequencer"
)

func NewServer(endpoint string, seq *sequencer.Sequencer) *Server {
	s := &Server{
		endpoint: endpoint,
		engine:   gin.Default(),
		seq:      seq,
	}

	s.engine.POST("/message", s.handleMessage)
	s.engine.GET("/message/:id", s.queryByID)

	return s
}

type Server struct {
	engine   *gin.Engine
	endpoint string
	seq      *sequencer.Sequencer
}

// this func will block caller
func (s *Server) Run() error {
	if err := s.engine.Run(s.endpoint); err != nil {
		return errors.Wrap(err, "start http server failed")
	}
	return nil
}
