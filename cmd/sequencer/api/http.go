package api

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/machinefi/ioconnect-go/pkg/ioconnect"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/apitypes"
	"github.com/iotexproject/w3bstream/clients"
	"github.com/iotexproject/w3bstream/cmd/sequencer/persistence"
	"github.com/iotexproject/w3bstream/task"
)

type httpServer struct {
	engine             *gin.Engine
	p                  *persistence.Persistence
	coordinatorAddress string
	aggregationAmount  uint
	privateKey         *ecdsa.PrivateKey
	jwk                *ioconnect.JWK
	clients            *clients.Manager
}

func NewHttpServer(p *persistence.Persistence, aggregationAmount uint, coordinatorAddress string, sk *ecdsa.PrivateKey, jwk *ioconnect.JWK, clientMgr *clients.Manager) *httpServer {
	s := &httpServer{
		engine:             gin.Default(),
		p:                  p,
		coordinatorAddress: coordinatorAddress,
		aggregationAmount:  aggregationAmount,
		privateKey:         sk,
		jwk:                jwk,
		clients:            clientMgr,
	}

	slog.Debug("jwk information",
		"did:io", jwk.DID(),
		"did:io#key", jwk.KID(),
		"ka did:io", jwk.KeyAgreementDID(),
		"ka did:io#key", jwk.KeyAgreementKID(),
		"doc", jwk.Doc(),
	)

	s.engine.POST("/issue_vc", s.issueJWTCredential)
	s.engine.POST("/message", s.verifyToken, s.handleMessage)
	s.engine.GET("/message/:id", s.verifyToken, s.queryStateLogByID)
	s.engine.GET("/didDoc", s.didDoc)

	return s
}

// this func will block caller
func (s *httpServer) Run(address string) error {
	if err := s.engine.Run(address); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}
	return nil
}

// verifyToken make sure the client token is issued by sequencer
func (s *httpServer) verifyToken(c *gin.Context) {
	tok := c.GetHeader("Authorization")
	if tok == "" {
		tok = c.Query("authorization")
	}

	if tok == "" {
		return
	}

	tok = strings.TrimSpace(strings.Replace(tok, "Bearer", " ", 1))

	clientID, err := s.jwk.VerifyToken(tok)
	if err != nil {
		c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(errors.Wrap(err, "invalid credential token")))
		return
	}
	client := s.clients.ClientByIoID(clientID)
	if client == nil {
		c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(errors.New("invalid credential token")))
		return
	}

	ctx := clients.WithClientID(c.Request.Context(), client)
	c.Request = c.Request.WithContext(ctx)
}

func (s *httpServer) handleMessage(c *gin.Context) {
	req := &apitypes.HandleMessageReq{}

	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(errors.Wrap(err, "failed to read request body")))
		return
	}
	defer c.Request.Body.Close()

	// decrypt did comm message
	client := clients.ClientIDFrom(c.Request.Context())
	if client != nil {
		payload, err = s.jwk.Decrypt(payload, client.DID())
		if err != nil {
			c.JSON(http.StatusBadRequest, apitypes.NewErrRsp(errors.Wrap(err, "failed to decrypt didcomm cipher data")))
			return
		}
	}

	// binding request
	if err := binding.JSON.BindBody(payload, req); err != nil {
		c.JSON(http.StatusBadRequest, apitypes.NewErrRsp(err))
		return
	}

	// validate project permission
	if client != nil {
		approved, err := s.clients.HasProjectPermission(client.DID(), req.ProjectID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(errors.Wrapf(err, "failed to check project %d permission for %s", req.ProjectID, client.DID())))
			return
		}
		if !approved {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(errors.Errorf("no permission project %d for %s", req.ProjectID, client.DID())))
			return
		}
	}

	clientDID := ""
	if client != nil {
		clientDID = client.DID()
	}

	// execute task committing
	id := uuid.NewString()
	if err := s.p.Save(&persistence.Message{
		MessageID:      id,
		ClientID:       clientDID,
		ProjectID:      req.ProjectID,
		ProjectVersion: req.ProjectVersion,
		Data:           []byte(req.Data),
	}, s.aggregationAmount, s.privateKey); err != nil {
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		return
	}

	response := &apitypes.HandleMessageRsp{MessageID: id}

	if client != nil {
		slog.Info("encrypt response task commit", "response", response)
		cipher, err := s.jwk.EncryptJSON(response, client.KeyAgreementKID())
		if err != nil {
			c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(errors.Wrap(err, "failed to encrypt response when commit task")))
			return
		}
		c.Data(http.StatusOK, "application/octet-stream", cipher)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (s *httpServer) queryStateLogByID(c *gin.Context) {
	messageID := c.Param("id")

	ms, err := s.p.FetchMessage(messageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		return
	}
	if len(ms) == 0 {
		c.JSON(http.StatusOK, &apitypes.QueryMessageStateLogRsp{MessageID: messageID})
		return
	}
	m := ms[0]

	client := clients.ClientIDFrom(c.Request.Context())
	if client != nil {
		if m.ClientID != client.DID() {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(errors.New("unmatched client DID")))
			return
		}
		approved, err := s.clients.HasProjectPermission(client.DID(), m.ProjectID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(errors.Wrapf(err, "failed to check project %d permission for %s", m.ProjectID, client.DID())))
			return
		}
		if !approved {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(errors.Errorf("no permission project %d for %s", m.ProjectID, client.DID())))
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
		ts, err := s.p.FetchTask(m.InternalTaskID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
			return
		}
		if len(ts) == 0 {
			c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(errors.New("cannot find task by internal task id")))
			return
		}
		ss = append(ss, &apitypes.StateLog{
			State: task.StatePacked.String(),
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
			return
		}
		ss = append(ss, taskStateLog.States...)
	}

	response := &apitypes.QueryMessageStateLogRsp{MessageID: messageID, States: ss}

	if client != nil {
		slog.Info("encrypt response task query", "response", response)
		cipher, err := s.jwk.EncryptJSON(response, client.KeyAgreementKID())
		if err != nil {
			c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(errors.Wrap(err, "failed to encrypt response when query task")))
			return
		}
		c.Data(http.StatusOK, "application/octet-stream", cipher)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (s *httpServer) didDoc(c *gin.Context) {
	if s.jwk == nil {
		c.JSON(http.StatusNotAcceptable, apitypes.NewErrRsp(errors.New("jwk is not config")))
		return
	}
	c.JSON(http.StatusOK, s.jwk.Doc())
}

func (s *httpServer) issueJWTCredential(c *gin.Context) {
	req := new(apitypes.IssueTokenReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, apitypes.NewErrRsp(err))
		return
	}

	client := s.clients.ClientByIoID(req.ClientID)
	if client == nil {
		c.String(http.StatusForbidden, errors.Errorf("client is not register to ioRegistry").Error())
		return
	}

	token, err := s.jwk.SignToken(req.ClientID)
	if err != nil {
		c.String(http.StatusInternalServerError, errors.Wrap(err, "failed to sign token").Error())
		return
	}
	slog.Info("token signed", "token", token)

	cipher, err := s.jwk.Encrypt([]byte(token), client.KeyAgreementKID())
	if err != nil {
		c.String(http.StatusInternalServerError, errors.Wrap(err, "failed to encrypt").Error())
		return
	}

	c.Data(http.StatusOK, "application/json", cipher)
}
