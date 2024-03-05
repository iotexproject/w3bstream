package api

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/auth/didvc"
	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type errResp struct {
	Error string `json:"error,omitempty"`
}

func newErrResp(err error) *errResp {
	return &errResp{Error: err.Error()}
}

type handleMessageReq struct {
	ProjectID      uint64 `json:"projectID"        binding:"required"`
	ProjectVersion string `json:"projectVersion"   binding:"required"`
	Data           string `json:"data"             binding:"required"`
}

type handleMessageResp struct {
	MessageID string `json:"messageID"`
}

type stateLog struct {
	State   string    `json:"state"`
	Time    time.Time `json:"time"`
	Comment string    `json:"comment"`
}

type queryMessageStateLogResp struct {
	MessageID string      `json:"messageID"`
	States    []*stateLog `json:"states"`
}

type ENodeConfigResp struct {
	ProjectContractAddress string `json:"projectContractAddress"`
	OperatorETHAddress     string `json:"OperatorETHAddress,omitempty"`
	OperatorSolanaAddress  string `json:"operatorSolanaAddress,omitempty"`
}

type eNodeConfig func() (*ENodeConfigResp, error)
type HttpServer struct {
	engine         *gin.Engine
	pg             *persistence.Postgres
	didAuthServer  string
	projectManager *project.Manager
	eNodeConfig    eNodeConfig
}

func NewHttpServer(pg *persistence.Postgres, didAuthServer string, projectManager *project.Manager, enodeConfig eNodeConfig) *HttpServer {
	s := &HttpServer{
		engine:         gin.Default(),
		pg:             pg,
		didAuthServer:  didAuthServer,
		projectManager: projectManager,
		eNodeConfig:    enodeConfig,
	}

	s.engine.POST("/message", s.handleMessage)
	s.engine.GET("/message/:id", s.queryStateLogByID)
	s.engine.POST("/sign_credential", s.issueJWTCredential)
	s.engine.GET("/enode_config", s.getENodeConfigInfo)

	return s
}

// this func will block caller
func (s *HttpServer) Run(endpoint string) error {
	if err := s.engine.Run(endpoint); err != nil {
		return errors.Wrap(err, "start http server failed")
	}
	return nil
}

func (s *HttpServer) handleMessage(c *gin.Context) {
	req := &handleMessageReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, newErrResp(err))
		return
	}

	tok := c.GetHeader("Authorization")
	if tok == "" {
		tok = c.Query("authorization")
	}
	tok = strings.TrimSpace(strings.Replace(tok, "Bearer", " ", 1))

	cliDID := ""
	if tok != "" {
		err := didvc.VerifyJWTCredential(s.didAuthServer, tok)
		if err != nil {
			c.JSON(http.StatusUnauthorized, newErrResp(err))
			return
		}
		cliDID, err = clients.VerifySessionAndProjectPermission(tok, req.ProjectID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, newErrResp(err))
		}
	}

	config, err := s.projectManager.Get(req.ProjectID, req.ProjectVersion)
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrResp(err))
		return
	}

	id := uuid.NewString()
	if err := s.pg.Save(&types.Message{
		ID:             id,
		ClientDID:      cliDID,
		ProjectID:      req.ProjectID,
		ProjectVersion: req.ProjectVersion,
		Data:           req.Data,
	}, config); err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	slog.Debug("message was received", "messageID", id)
	c.JSON(http.StatusOK, &handleMessageResp{MessageID: id})
}

func (s *HttpServer) queryStateLogByID(c *gin.Context) {
	tok := c.GetHeader("Authorization")
	if tok == "" {
		tok = c.Query("authorization")
	}
	tok = strings.TrimSpace(strings.Replace(tok, "Bearer", " ", 1))

	if tok != "" {
		err := didvc.VerifyJWTCredential(s.didAuthServer, tok)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}
	}

	messageID := c.Param("id")

	ms, err := s.pg.FetchMessage(messageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}
	if len(ms) == 0 {
		c.JSON(http.StatusOK, &queryMessageStateLogResp{MessageID: messageID})
		return
	}

	ls, err := s.pg.FetchStateLog(messageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	ss := []*stateLog{
		{
			State: "received",
			Time:  ms[0].CreatedAt,
		}}
	for _, l := range ls {
		ss = append(ss, &stateLog{
			State:   l.State.String(),
			Time:    l.CreatedAt,
			Comment: l.Comment,
		})
	}

	c.JSON(http.StatusOK, &queryMessageStateLogResp{MessageID: messageID, States: ss})
}

func (s *HttpServer) issueJWTCredential(c *gin.Context) {
	req := new(didvc.IssueCredentialReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, newErrResp(err))
		return
	}

	rsp, err := didvc.IssueCredential(s.didAuthServer, req, true)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rsp)
	return
}

func (s *HttpServer) getENodeConfigInfo(c *gin.Context) {
	configInfo, err := s.eNodeConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, newErrResp(err))
		return
	}

	c.JSON(http.StatusOK, configInfo)
}
