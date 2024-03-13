package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/apitypes"
	"github.com/machinefi/sprout/types"
)

type httpServer struct {
	engine            *gin.Engine
	p                 *persistence
	enodeAddress      string
	aggregationAmount uint
}

func newHttpServer(p *persistence, aggregationAmount uint, enodeAddress string) *httpServer {
	s := &httpServer{
		engine:            gin.Default(),
		p:                 p,
		enodeAddress:      enodeAddress,
		aggregationAmount: aggregationAmount,
	}

	s.engine.POST("/message", s.handleMessage)
	s.engine.GET("/message/:id", s.queryStateLogByID)

	return s
}

// this func will block caller
func (s *httpServer) run(address string) error {
	if err := s.engine.Run(address); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}
	return nil
}

func (s *httpServer) handleMessage(c *gin.Context) {
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

func (s *httpServer) queryStateLogByID(c *gin.Context) {
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
	m := ms[0]
	ss := []*apitypes.StateLog{
		{
			State: "received",
			Time:  m.CreatedAt,
		},
	}

	if m.InternalTaskID != "" {
		ts, err := s.p.fetchTask(m.InternalTaskID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
			return
		}
		if len(ts) == 0 {
			c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(errors.New("cannot find task by internal task id")))
			return
		}
		resp, err := http.Get(fmt.Sprintf("http://%s/%s/%d/%d", s.enodeAddress, "task", m.ProjectID, ts[0].ID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
			return
		}
		taskStateLog := &apitypes.QueryTaskStateLogRsp{}
		if err := json.Unmarshal(body, &taskStateLog); err != nil {
			c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		}
		ss = append(ss, taskStateLog.States...)
	}

	c.JSON(http.StatusOK, &apitypes.QueryMessageStateLogRsp{MessageID: messageID, States: ss})
}
