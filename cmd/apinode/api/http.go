package api

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/apinode/persistence"
	"github.com/iotexproject/w3bstream/p2p"
)

type ErrResp struct {
	Error string `json:"error,omitempty"`
}

func NewErrResp(err error) *ErrResp {
	return &ErrResp{Error: err.Error()}
}

type HandleMessageReq struct {
	ProjectID      uint64 `json:"projectID"                  binding:"required"`
	ProjectVersion string `json:"projectVersion"             binding:"required"`
	Data           string `json:"data"                       binding:"required"`
	Signature      string `json:"signature,omitempty"        binding:"required"`
}

type HandleMessageResp struct {
	TaskID string `json:"taskID"`
}

type QueryTaskReq struct {
	ProjectID uint64 `json:"projectID"                  binding:"required"`
	TaskID    string `json:"taskID"                     binding:"required"`
}

type StateLog struct {
	State    string    `json:"state"`
	Time     time.Time `json:"time"`
	Comment  string    `json:"comment,omitempty"`
	ProverID string    `json:"prover_id,omitempty"`
}

type QueryTaskResp struct {
	ProjectID uint64      `json:"projectID"`
	TaskID    string      `json:"taskID"`
	States    []*StateLog `json:"states"`
}

type httpServer struct {
	engine            *gin.Engine
	p                 *persistence.Persistence
	aggregationAmount int
	prv               *ecdsa.PrivateKey
	pubSub            *p2p.PubSub
}

func (s *httpServer) handleMessage(c *gin.Context) {
	req := &HandleMessageReq{}
	if err := c.ShouldBind(req); err != nil {
		slog.Error("failed to bind request", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(errors.Wrap(err, "invalid request payload")))
		return
	}

	sigStr := req.Signature
	req.Signature = ""

	reqJson, err := json.Marshal(req)
	if err != nil {
		slog.Error("failed to marshal request into json format", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to process request data")))
		return
	}

	sig, err := hexutil.Decode(sigStr)
	if err != nil {
		slog.Error("failed to decode signature from hex format", "signature", sigStr, "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(errors.Wrap(err, "invalid signature format")))
		return
	}

	h := crypto.Keccak256Hash(reqJson)
	sigpk, err := crypto.SigToPub(h.Bytes(), sig)
	if err != nil {
		slog.Error("failed to recover public key from signature", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(errors.Wrap(err, "invalid signature; could not recover public key")))
		return
	}

	addr := crypto.PubkeyToAddress(*sigpk)

	taskID, err := s.p.Save(s.pubSub,
		&persistence.Message{
			DeviceID:       addr,
			ProjectID:      req.ProjectID,
			ProjectVersion: req.ProjectVersion,
			Data:           []byte(req.Data),
			TaskID:         common.Hash{},
		}, s.aggregationAmount, s.prv,
	)
	if err != nil {
		slog.Error("failed to save message to persistence layer", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "internal server error; could not save task")))
		return
	}

	resp := &HandleMessageResp{}
	if !bytes.Equal(taskID[:], common.Hash{}.Bytes()) {
		resp.TaskID = taskID.String()
	}
	slog.Info("successfully processed message", "taskID", resp.TaskID)
	c.JSON(http.StatusOK, resp)
}

func (s *httpServer) queryStateLogByID(c *gin.Context) {
	req := &QueryTaskReq{}
	if err := c.ShouldBind(req); err != nil {
		slog.Error("failed to bind request", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(errors.Wrap(err, "invalid request payload")))
		return
	}

	taskID := common.HexToHash(req.TaskID)
	resp := &QueryTaskResp{
		ProjectID: req.ProjectID,
		TaskID:    req.TaskID,
		States:    []*StateLog{},
	}

	t, err := s.p.FetchTask(req.ProjectID, taskID)
	if err != nil {
		slog.Error("failed to query task", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to query task")))
		return
	}
	if t == nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.States = append(resp.States, &StateLog{
		State: "packed",
		Time:  t.CreatedAt,
	})

	ta, err := s.p.FetchAssignedTask(req.ProjectID, taskID)
	if err != nil {
		slog.Error("failed to query assigned task", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to query assigned task")))
		return
	}
	if ta == nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.States = append(resp.States, &StateLog{
		State:    "assigned",
		Time:     t.CreatedAt,
		ProverID: "did:io:" + strings.TrimPrefix(ta.Prover.String(), "0x"),
	})

	ts, err := s.p.FetchSettledTask(req.ProjectID, taskID)
	if err != nil {
		slog.Error("failed to query settled task", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to query settled task")))
		return
	}
	if ts == nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.States = append(resp.States, &StateLog{
		State:   "settled",
		Time:    t.CreatedAt,
		Comment: "The task has been completed. Please check the generated proof in the corresponding DApp contract.",
	})
	c.JSON(http.StatusOK, resp)
}

// this func will block caller
func Run(p *persistence.Persistence, prv *ecdsa.PrivateKey, pubSub *p2p.PubSub, aggregationAmount int, address string) error {
	s := &httpServer{
		engine:            gin.Default(),
		p:                 p,
		aggregationAmount: aggregationAmount,
		prv:               prv,
		pubSub:            pubSub,
	}

	s.engine.POST("/message", s.handleMessage)
	s.engine.GET("/message/:id", s.queryStateLogByID)

	if err := s.engine.Run(address); err != nil {
		slog.Error("Failed to start HTTP server", "address", address, "error", err)
		return errors.Wrap(err, "could not start HTTP server; check if the address is in use or network is accessible")
	}
	return nil
}
