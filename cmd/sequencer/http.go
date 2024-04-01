package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/apitypes"
	"github.com/machinefi/sprout/auth/didvc"
	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/types"
)

type httpServer struct {
	engine                *gin.Engine
	p                     *persistence
	coordinatorAddress    string
	aggregationAmount     uint
	didAuthServerEndpoint string
	privateKey            *ecdsa.PrivateKey
}

func newHttpServer(p *persistence, aggregationAmount uint, coordinatorAddress, didAuthServerEndpoint string, sk *ecdsa.PrivateKey) *httpServer {
	s := &httpServer{
		engine:                gin.Default(),
		p:                     p,
		coordinatorAddress:    coordinatorAddress,
		aggregationAmount:     aggregationAmount,
		didAuthServerEndpoint: didAuthServerEndpoint,
		privateKey:            sk,
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

	tok := c.GetHeader("Authorization")
	if tok == "" {
		tok = c.Query("authorization")
	}
	tok = strings.TrimSpace(strings.Replace(tok, "Bearer", " ", 1))

	clientDID := ""
	if tok != "" {
		err := didvc.VerifyJWTCredential(s.didAuthServerEndpoint, tok)
		if err != nil {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(err))
			return
		}
		if clientDID, err = clients.VerifySessionAndProjectPermission(tok, req.ProjectID); err != nil {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(err))
			return
		}
	}

	id := uuid.NewString()
	if err := s.p.save(&message{
		MessageID:      id,
		ClientDID:      clientDID,
		ProjectID:      req.ProjectID,
		ProjectVersion: req.ProjectVersion,
		Data:           []byte(req.Data),
	}, s.aggregationAmount, s.privateKey); err != nil {
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

	tok := c.GetHeader("Authorization")
	if tok == "" {
		tok = c.Query("authorization")
	}
	tok = strings.TrimSpace(strings.Replace(tok, "Bearer", " ", 1))

	if tok != "" {
		if err = didvc.VerifyJWTCredential(s.didAuthServerEndpoint, tok); err != nil {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(err))
			return
		}
		clientDID := ""
		if clientDID, err = clients.VerifySessionAndProjectPermission(tok, m.ProjectID); err != nil {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(err))
			return
		}
		if m.ClientDID != clientDID {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(errors.New("unmatched client DID")))
			return
		}
	}

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
		ss = append(ss, &apitypes.StateLog{
			State: types.TaskStatePacked.String(),
			Time:  ts[0].CreatedAt,
		})
		resp, err := http.Get(fmt.Sprintf("http://%s/%s/%d/%d", s.coordinatorAddress, "task", m.ProjectID, ts[0].ID))
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

func (s *httpServer) issueJWTCredential(c *gin.Context) {
	req := new(didvc.IssueCredentialReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, apitypes.NewErrRsp(err))
		return
	}

	rsp, err := didvc.IssueCredential(s.didAuthServerEndpoint, req, true)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rsp)
}
