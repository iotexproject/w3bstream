package api

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/service/apinode/persistence"
	proverapi "github.com/iotexproject/w3bstream/service/prover/api"
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
	Error    string    `json:"error,omitempty"`
	Tx       string    `json:"transaction_hash,omitempty"`
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
	proverAddr        string
}

func (s *httpServer) handleMessage(c *gin.Context) {
	req := &HandleMessageReq{}
	if err := c.ShouldBindJSON(req); err != nil {
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

func (s *httpServer) queryTask(c *gin.Context) {
	req := &QueryTaskReq{}
	if err := c.ShouldBindJSON(req); err != nil {
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
		Time:     ta.CreatedAt,
		ProverID: "did:io:" + strings.TrimPrefix(ta.Prover.String(), "0x"),
	})

	ts, err := s.p.FetchSettledTask(req.ProjectID, taskID)
	if err != nil {
		slog.Error("failed to query settled task", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to query settled task")))
		return
	}
	if ts == nil {
		reqJ, err := json.Marshal(proverapi.QueryTaskReq{
			ProjectID: req.ProjectID,
			TaskID:    req.TaskID,
		})
		if err != nil {
			slog.Error("failed to marshal prover request", "error", err)
			c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to marshal prover request")))
			return
		}
		req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/task", s.proverAddr), bytes.NewBuffer(reqJ))
		if err != nil {
			slog.Error("failed to build http request", "error", err)
			c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to build http request")))
			return
		}
		req.Header.Set("Content-Type", "application/json")

		proverResp, err := http.DefaultClient.Do(req)
		if err != nil {
			slog.Error("failed to call prover http server", "error", err)
			c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to call prover http server")))
			return
		}
		defer proverResp.Body.Close()

		body, err := io.ReadAll(proverResp.Body)
		if err != nil {
			slog.Error("failed to read prover http server response", "error", err)
			c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to read prover http server response")))
			return
		}
		taskResp := &proverapi.QueryTaskResp{}
		if err := json.Unmarshal(body, &taskResp); err != nil {
			slog.Error("failed to unmarshal prover http server response", "error", err)
			c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to unmarshal prover http server response")))
			return
		}
		if taskResp.Processed && taskResp.Error != "" {
			resp.States = append(resp.States, &StateLog{
				State: "failed",
				Error: taskResp.Error,
				Time:  taskResp.Time,
			})
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.States = append(resp.States, &StateLog{
		State:   "settled",
		Time:    ts.CreatedAt,
		Comment: "The task has been completed. Please check the generated proof in the corresponding DApp contract.",
		Tx:      ts.Tx.String(),
	})
	c.JSON(http.StatusOK, resp)
}

// this func will block caller
func Run(p *persistence.Persistence, prv *ecdsa.PrivateKey, pubSub *p2p.PubSub, aggregationAmount int, addr, proverAddr string) error {
	s := &httpServer{
		engine:            gin.Default(),
		p:                 p,
		aggregationAmount: aggregationAmount,
		prv:               prv,
		pubSub:            pubSub,
		proverAddr:        proverAddr,
	}

	s.engine.POST("/message", s.handleMessage)
	s.engine.GET("/task", s.queryTask)

	if err := s.engine.Run(addr); err != nil {
		slog.Error("failed to start http server", "address", addr, "error", err)
		return errors.Wrap(err, "could not start http server; check if the address is in use or network is accessible")
	}
	return nil
}
