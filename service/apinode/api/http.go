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
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/metrics"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/service/apinode/persistence"
	proverapi "github.com/iotexproject/w3bstream/service/prover/api"
	sequencerapi "github.com/iotexproject/w3bstream/service/sequencer/api"
)

type ErrResp struct {
	Error string `json:"error,omitempty"`
}

func NewErrResp(err error) *ErrResp {
	return &ErrResp{Error: err.Error()}
}

type CreateTaskReq struct {
	ProjectID      uint64   `json:"projectID"                  binding:"required"`
	ProjectVersion string   `json:"projectVersion"             binding:"required"`
	Payloads       []string `json:"payloads"                   binding:"required"`
	Signature      string   `json:"signature,omitempty"        binding:"required"`
}

type CreateTaskResp struct {
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
	engine        *gin.Engine
	p             *persistence.Persistence
	pubSub        *p2p.PubSub
	sequencerAddr string
	proverAddr    string
}

func (s *httpServer) createTask(c *gin.Context) {
	req := &CreateTaskReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		slog.Error("failed to bind request", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(errors.Wrap(err, "invalid request payload")))
		return
	}

	sig, err := hexutil.Decode(req.Signature)
	if err != nil {
		slog.Error("failed to decode signature from hex format", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(errors.Wrap(err, "failed to decode signature from hex format")))
		return
	}
	pubKey, err := s.recoverPubkey(*req, sig)
	if err != nil {
		slog.Error("failed to recover public key", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(errors.Wrap(err, "invalid signature; could not recover public key")))
		return
	}

	// TODO: Crosscheck pubkey with ioID

	addr := crypto.PubkeyToAddress(*pubKey)
	payloadsB := make([][]byte, 0, len(req.Payloads))
	for _, p := range req.Payloads {
		payloadsB = append(payloadsB, []byte(p))
	}
	payloadsJ, err := json.Marshal(payloadsB)
	if err != nil {
		slog.Error("failed to marshal payloads", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to marshal payloads")))
		return
	}
	taskID := crypto.Keccak256Hash([]byte(uuid.NewString()))

	if err := s.p.CreateTask(
		&persistence.Task{
			DeviceID:       addr,
			TaskID:         taskID,
			ProjectID:      req.ProjectID,
			ProjectVersion: req.ProjectVersion,
			Payloads:       payloadsJ,
			Signature:      sig,
		},
	); err != nil {
		slog.Error("failed to create task to persistence layer", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "could not save task")))
		return
	}

	reqSequencer := &sequencerapi.CreateTaskReq{ProjectID: req.ProjectID, TaskID: taskID}
	reqSequencerJ, err := json.Marshal(reqSequencer)
	if err != nil {
		slog.Error("failed to marshal sequencer request", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to marshal sequencer request")))
		return
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/task", s.sequencerAddr), "application/json", bytes.NewBuffer(reqSequencerJ))
	if err != nil {
		slog.Error("failed to call sequencer service", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to call sequencer service")))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			err = errors.New(string(body))
		}
		slog.Error("failed to call sequencer service", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to call sequencer service")))
		return
	}

	slog.Info("successfully processed message", "task_id", taskID.String())
	c.JSON(http.StatusOK, &CreateTaskResp{
		TaskID: taskID.String(),
	})
}

func (s *httpServer) recoverPubkey(req CreateTaskReq, sig []byte) (*ecdsa.PublicKey, error) {
	req.Signature = ""
	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request into json format")
	}

	h := crypto.Keccak256Hash(reqJson)
	sigpk, err := crypto.SigToPub(h.Bytes(), sig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to recover public key from signature")
	}
	return sigpk, nil
}

func (s *httpServer) getProverTaskState(projectID uint64, taskID string) (*StateLog, error) {
	reqBody, err := json.Marshal(proverapi.QueryTaskReq{
		ProjectID: projectID,
		TaskID:    taskID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal prover request")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/task", s.proverAddr), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, errors.Wrap(err, "failed to build http request")
	}
	req.Header.Set("Content-Type", "application/json")

	proverResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call prover http server")
	}
	defer proverResp.Body.Close()

	body, err := io.ReadAll(proverResp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read prover http server response")
	}

	taskResp := &proverapi.QueryTaskResp{}
	if err := json.Unmarshal(body, &taskResp); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal prover http server response")
	}

	if taskResp.Processed && taskResp.Error != "" {
		return &StateLog{
			State: "failed",
			Error: taskResp.Error,
			Time:  taskResp.Time,
		}, nil
	}
	return nil, nil
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

	// Get packed state
	task, err := s.p.FetchTask(req.ProjectID, taskID)
	if err != nil {
		slog.Error("failed to query task", "error", err)
		resp.States = append(resp.States, &StateLog{
			State: "error",
			Time:  time.Now(),
			Error: fmt.Sprintf("failed to query task: %v", err),
		})
		c.JSON(http.StatusOK, resp)
		return
	}
	if task == nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.States = append(resp.States, &StateLog{
		State: "packed",
		Time:  task.CreatedAt,
	})

	// Get assigned state
	assignedTask, err := s.p.FetchAssignedTask(req.ProjectID, taskID)
	if err != nil {
		slog.Error("failed to query assigned task", "error", err)
		resp.States = append(resp.States, &StateLog{
			State: "error",
			Time:  time.Now(),
			Error: fmt.Sprintf("failed to query assigned task: %v", err),
		})
		c.JSON(http.StatusOK, resp)
		return
	}
	if assignedTask == nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.States = append(resp.States, &StateLog{
		State:    "assigned",
		Time:     assignedTask.CreatedAt,
		ProverID: "did:io:" + strings.TrimPrefix(assignedTask.Prover.String(), "0x"),
	})

	// Get settled state
	settledTask, err := s.p.FetchSettledTask(req.ProjectID, taskID)
	if err != nil {
		slog.Error("failed to query settled task", "error", err)
		resp.States = append(resp.States, &StateLog{
			State: "error",
			Time:  time.Now(),
			Error: fmt.Sprintf("failed to query settled task: %v", err),
		})
		c.JSON(http.StatusOK, resp)
		return
	}

	if settledTask != nil {
		resp.States = append(resp.States, &StateLog{
			State:   "settled",
			Time:    settledTask.CreatedAt,
			Comment: "The task has been completed. Please check the generated proof in the corresponding DApp contract.",
			Tx:      settledTask.Tx.String(),
		})
		c.JSON(http.StatusOK, resp)
		return
	}

	// If not settled, check prover state
	proverState, err := s.getProverTaskState(req.ProjectID, req.TaskID)
	if err != nil {
		slog.Error("failed to get prover task state", "error", err)
		resp.States = append(resp.States, &StateLog{
			State: "error",
			Time:  time.Now(),
			Error: fmt.Sprintf("failed to get prover state: %v", err),
		})
		c.JSON(http.StatusOK, resp)
		return
	}
	if proverState != nil {
		resp.States = append(resp.States, proverState)
	}

	c.JSON(http.StatusOK, resp)
}

// this func will block caller
func Run(p *persistence.Persistence, pubSub *p2p.PubSub, addr, sequencerAddr, proverAddr string) error {
	s := &httpServer{
		engine:        gin.Default(),
		p:             p,
		pubSub:        pubSub,
		sequencerAddr: sequencerAddr,
		proverAddr:    proverAddr,
	}

	s.engine.POST("/task", s.createTask)
	s.engine.GET("/task", s.queryTask)
	metrics.RegisterMetrics(s.engine)

	if err := s.engine.Run(addr); err != nil {
		slog.Error("failed to start http server", "address", addr, "error", err)
		return errors.Wrap(err, "could not start http server; check if the address is in use or network is accessible")
	}
	return nil
}
