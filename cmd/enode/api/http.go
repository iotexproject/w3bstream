package api

import (
	"log/slog"
	"net/http"
	"strings"

	solanatypes "github.com/blocto/solana-go-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/apitypes"
	"github.com/machinefi/sprout/auth/didvc"
	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/cmd/enode/config"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type HttpServer struct {
	engine         *gin.Engine
	pg             *persistence.Postgres
	projectManager *project.Manager
	conf           *config.Config
	enodeConf      *apitypes.ENodeConfigRsp
}

func NewHttpServer(pg *persistence.Postgres, projectManager *project.Manager, conf *config.Config) *HttpServer {
	s := &HttpServer{
		engine:         gin.Default(),
		pg:             pg,
		projectManager: projectManager,
		conf:           conf,
	}

	s.enodeConf = &apitypes.ENodeConfigRsp{
		ProjectContractAddress: s.conf.ProjectContractAddress,
	}

	if len(s.conf.OperatorPrivateKey) > 0 {
		pk := crypto.ToECDSAUnsafe(common.FromHex(s.conf.OperatorPrivateKey))
		sender := crypto.PubkeyToAddress(pk.PublicKey)
		s.enodeConf.OperatorETHAddress = sender.String()
	}

	if len(s.conf.OperatorPrivateKeyED25519) > 0 {
		wallet, err := solanatypes.AccountFromHex(s.conf.ServiceEndpoint)
		if err != nil {
			panic(errors.Wrapf(err, "invalid solana wallet address"))
		}
		s.enodeConf.OperatorSolanaAddress = wallet.PublicKey.String()
	}

	s.engine.GET("/live", s.liveness)
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

func (s *HttpServer) liveness(c *gin.Context) {
	c.JSON(http.StatusOK, &apitypes.LivenessRsp{Status: "up"})
}

func (s *HttpServer) handleMessage(c *gin.Context) {
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

	cliDID := ""
	if tok != "" {
		err := didvc.VerifyJWTCredential(s.conf.DIDAuthServerEndpoint, tok)
		if err != nil {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(err))
			return
		}
		cliDID, err = clients.VerifySessionAndProjectPermission(tok, req.ProjectID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(err))
		}
	}

	config, err := s.projectManager.Get(req.ProjectID, req.ProjectVersion)
	if err != nil {
		c.JSON(http.StatusBadRequest, apitypes.NewErrRsp(err))
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
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		return
	}

	slog.Debug("message was received", "messageID", id)
	c.JSON(http.StatusOK, &apitypes.HandleMessageRsp{MessageID: id})
}

func (s *HttpServer) queryStateLogByID(c *gin.Context) {
	tok := c.GetHeader("Authorization")
	if tok == "" {
		tok = c.Query("authorization")
	}
	tok = strings.TrimSpace(strings.Replace(tok, "Bearer", " ", 1))

	if tok != "" {
		err := didvc.VerifyJWTCredential(s.conf.DIDAuthServerEndpoint, tok)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}
	}

	messageID := c.Param("id")

	ms, err := s.pg.FetchMessage(messageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		return
	}
	if len(ms) == 0 {
		c.JSON(http.StatusOK, &apitypes.QueryMessageStateLogRsp{MessageID: messageID})
		return
	}

	ls, err := s.pg.FetchStateLog(messageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		return
	}

	ss := []*apitypes.StateLog{
		{
			State: "received",
			Time:  ms[0].CreatedAt,
		}}
	for _, l := range ls {
		ss = append(ss, &apitypes.StateLog{
			State:   l.State.String(),
			Time:    l.CreatedAt,
			Comment: string(l.Comment),
		})
	}

	c.JSON(http.StatusOK, &apitypes.QueryMessageStateLogRsp{MessageID: messageID, States: ss})
}

func (s *HttpServer) issueJWTCredential(c *gin.Context) {
	req := new(didvc.IssueCredentialReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, apitypes.NewErrRsp(err))
		return
	}

	rsp, err := didvc.IssueCredential(s.conf.DIDAuthServerEndpoint, req, true)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rsp)
	return
}

func (s *HttpServer) getENodeConfigInfo(c *gin.Context) {
	c.JSON(http.StatusOK, s.enodeConf)
}
