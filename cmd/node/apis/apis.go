package apis

import (
	"github.com/gin-gonic/gin"
	msghandler "github.com/machinefi/w3bstream-mainnet/msg/handler"
	"github.com/machinefi/w3bstream-mainnet/vm"
	"github.com/pkg/errors"
)

func NewServer(ep string, mh *msghandler.Handler) *Server {
	return &Server{
		endpoint:   ep,
		engine:     gin.Default(),
		msgHandler: mh,
	}
}

type Server struct {
	engine     *gin.Engine
	endpoint   string
	msgHandler *msghandler.Handler
	vmHandler  *vm.Handler
}

func (s *Server) Run() error {
	s.engine.POST("/message", s.handleRequest)

	if err := s.engine.Run(s.endpoint); err != nil {
		return errors.Wrap(err, "start http server failed")
	}
	return nil
}
