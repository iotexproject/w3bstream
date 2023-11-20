package apis

import (
	"github.com/gin-gonic/gin"
	msghandler "github.com/machinefi/sprout/msg/handler"
	"github.com/machinefi/sprout/vm"
	"github.com/pkg/errors"
)

func NewServer(ep string, mh *msghandler.Handler) *Server {
	s := &Server{
		endpoint:   ep,
		engine:     gin.Default(),
		msgHandler: mh,
	}
	s.engine.POST("/message", s.handleRequest)
	s.engine.GET("/message/:messageID", s.queryByMessageID)
	return s
}

type Server struct {
	engine     *gin.Engine
	endpoint   string
	msgHandler *msghandler.Handler
	vmHandler  *vm.Handler
}

// this func will block caller
func (s *Server) Run() error {
	if err := s.engine.Run(s.endpoint); err != nil {
		return errors.Wrap(err, "start http server failed")
	}
	return nil
}
