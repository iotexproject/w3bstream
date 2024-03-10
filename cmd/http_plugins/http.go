package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/apitypes"
	"github.com/machinefi/sprout/types"
)

type HttpServer struct {
	engine            *gin.Engine
	p                 *persistence
	aggregationAmount uint
}

func NewHttpServer(p *persistence, aggregationAmount uint) *HttpServer {
	s := &HttpServer{
		engine:            gin.Default(),
		p:                 p,
		aggregationAmount: aggregationAmount,
	}

	s.engine.POST("/message", s.handleMessage)
	s.engine.GET("/message/:id", s.queryStateLogByID)

	return s
}

// this func will block caller
func (s *HttpServer) Run(address string) error {
	if err := s.engine.Run(address); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}
	return nil
}

func (s *HttpServer) handleMessage(c *gin.Context) {
	req := &apitypes.HandleMessageReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, apitypes.NewErrRsp(err))
		return
	}

	id := uuid.NewString()
	if err := s.p.save(&types.Message{
		ID:             id,
		ClientDID:      "",
		ProjectID:      req.ProjectID,
		ProjectVersion: req.ProjectVersion,
		Data:           req.Data,
	}, s.aggregationAmount); err != nil {
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		return
	}

	c.JSON(http.StatusOK, &apitypes.HandleMessageRsp{MessageID: id})
}

func (s *HttpServer) queryStateLogByID(c *gin.Context) {
	messageID := c.Param("id")

	ms, err := s.p.fetchMessage(messageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		return
	}
	if len(ms) == 0 {
		c.JSON(http.StatusOK, &apitypes.QueryMessageStateLogRsp{MessageID: messageID})
		return
	}

	// TODO query task

	c.JSON(http.StatusOK, &apitypes.QueryMessageStateLogRsp{MessageID: messageID})
}
