package api

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/machinefi/ioconnect-go/cmd/srv-did-vc/apis"
	"github.com/machinefi/ioconnect-go/pkg/ioconnect"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/apitypes"
	"github.com/machinefi/sprout/auth/didvc"
	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/cmd/sequencer/persistence"
	"github.com/machinefi/sprout/task"
)

type httpServer struct {
	engine                *gin.Engine
	p                     *persistence.Persistence
	coordinatorAddress    string
	aggregationAmount     uint
	didAuthServerEndpoint string
	privateKey            *ecdsa.PrivateKey
	didJWK                *ioconnect.JWK
}

func NewHttpServer(p *persistence.Persistence, aggregationAmount uint, coordinatorAddress, didAuthServerEndpoint string, sk *ecdsa.PrivateKey, did bool) *httpServer {
	s := &httpServer{
		engine:                gin.Default(),
		p:                     p,
		coordinatorAddress:    coordinatorAddress,
		aggregationAmount:     aggregationAmount,
		didAuthServerEndpoint: didAuthServerEndpoint,
		privateKey:            sk,
	}

	if did {
		didJWK, err := ioconnect.NewMasterJWK("io", "key")
		if err != nil {
			log.Fatal(err)
		}
		s.didJWK = didJWK

		slog.Info("server did:io:        " + s.didJWK.DID("io"))
		slog.Info("server did:io#key:    " + s.didJWK.KID("io"))
		slog.Info("server ka did:io:     " + s.didJWK.KeyAgreementDID("io"))
		slog.Info("server ka did:io#key: " + s.didJWK.KeyAgreementKID("io"))

		// generate did doc
		didDoc, err := s.didJWK.DIDDoc("io")
		if err != nil {
			log.Fatal(err)
		}
		docContent, _ := json.MarshalIndent(didDoc, "", "  ")
		slog.Info(string(docContent))
	}

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
	// TODO del no token
	if tok == "" {
		return
	}

	tok = strings.TrimSpace(strings.Replace(tok, "Bearer", " ", 1))
	clientID, err := didvc.VerifyJWTCredential(s.didAuthServerEndpoint, tok)
	if err != nil {
		c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(errors.Wrap(err, "invalid credential token")))
		return
	}
	ctx := didvc.WithClientID(c.Request.Context(), clientID)
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
	clientID, ok := didvc.ClientIDFrom(c.Request.Context())
	// TODO change ok
	if ok || s.didJWK != nil {
		payload, err = s.didJWK.DecryptBySenderDID("io", payload, clientID)
		//payload, err = s.didJWK.DecryptBySenderDID("io", payload, clientID)
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
	if clientID != "" {
		// TODO consider if project has public attribute
		if err = clients.VerifyProjectPermissionByClientDID(clientID, req.ProjectID); err != nil {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(err))
			return
		}
	}

	// execute task committing
	id := uuid.NewString()
	if err := s.p.Save(&persistence.Message{
		MessageID:      id,
		ClientID:       clientID,
		ProjectID:      req.ProjectID,
		ProjectVersion: req.ProjectVersion,
		Data:           []byte(req.Data),
	}, s.aggregationAmount, s.privateKey); err != nil {
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		return
	}

	response := &apitypes.HandleMessageRsp{MessageID: id}
	//TODO encrypt response and respond
	//if clientID != "" {
	//	//cipher, err := didcomm.EncryptJSON(response)
	//	cipher, err := s.didJWK.Encrypt("io", []byte("payload"), client.kid)
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(errors.Wrap(err, "failed to encrypt response when commit task")))
	//		return
	//	}
	//	c.Data(http.StatusOK, "application/octet-stream", cipher)
	//	return
	//}

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

	clientID, ok := didvc.ClientIDFrom(c.Request.Context())
	if ok {
		// TODO consider if project has public attribute
		if err = clients.VerifyProjectPermissionByClientDID(clientID, m.ProjectID); err != nil {
			c.JSON(http.StatusUnauthorized, apitypes.NewErrRsp(err))
			return
		}
		if m.ClientID != clientID {
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
		}
		ss = append(ss, taskStateLog.States...)
	}

	response := &apitypes.QueryMessageStateLogRsp{MessageID: messageID, States: ss}
	//TODO encrypt response and respond
	//if clientID != "" {
	//	cipher, err := didcomm.EncryptJSON(response)
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(errors.Wrap(err, "failed to encrypt response when query task")))
	//		return
	//	}
	//	c.Data(http.StatusOK, "application/octet-stream", cipher)
	//	return
	//}

	c.JSON(http.StatusOK, response)
}

func (s *httpServer) didDoc(c *gin.Context) {
	didDoc, err := s.didJWK.DIDDoc("io")
	if err != nil {
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, didDoc)
}

func (s *httpServer) issueJWTCredential(c *gin.Context) {
	req := new(apis.IssueTokenReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, apitypes.NewErrRsp(err))
		return
	}

	rsp, err := didvc.IssueCredential(s.didAuthServerEndpoint, req.ClientID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rsp)
}
