package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/metrics"
	"github.com/iotexproject/w3bstream/service/apinode/db"
	proverapi "github.com/iotexproject/w3bstream/service/prover/api"
	sequencerapi "github.com/iotexproject/w3bstream/service/sequencer/api"
)

type errResp struct {
	Error string `json:"error,omitempty"`
}

func newErrResp(err error) *errResp {
	return &errResp{Error: err.Error()}
}

type CreateTaskReq struct {
	Nonce          uint64 `json:"nonce"                        binding:"required"`
	ProjectID      string `json:"projectID"                    binding:"required"`
	ProjectVersion string `json:"projectVersion,omitempty"`
	Payload        string `json:"payload"                     binding:"required"`
	Algorithm      string `json:"algorithm,omitempty"` // Refer to the constants defined in JWT (JSON Web Token) https://jwt.io/
	Signature      string `json:"signature,omitempty"          binding:"required"`
}

type CreateTaskResp struct {
	TaskID string `json:"taskID"`
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
	ProjectID string      `json:"projectID"`
	TaskID    string      `json:"taskID"`
	States    []*StateLog `json:"states"`
}

type httpServer struct {
	engine        *gin.Engine
	db            *db.DB
	sequencerAddr string
	proverAddr    string
}

func (s *httpServer) createTask(c *gin.Context) {
	req := &CreateTaskReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		slog.Error("failed to bind request", "error", err)
		c.JSON(http.StatusBadRequest, newErrResp(errors.Wrap(err, "invalid request payload")))
		return
	}

	pid, ok := new(big.Int).SetString(req.ProjectID, 10)
	if !ok {
		slog.Error("failed to decode project id string", "project_id", req.ProjectID)
		c.JSON(http.StatusBadRequest, newErrResp(errors.New("failed to decode project id string")))
		return
	}
	sig, err := hexutil.Decode(req.Signature)
	if err != nil {
		slog.Error("failed to decode signature from hex format", "error", err)
		c.JSON(http.StatusBadRequest, newErrResp(errors.Wrap(err, "failed to decode signature from hex format")))
		return
	}

	recovered, alg, err := recover(*req, sig)
	if err != nil {
		slog.Error("failed to recover public key", "error", err)
		c.JSON(http.StatusBadRequest, newErrResp(errors.Wrap(err, "invalid signature; could not recover public key")))
		return
	}
	var addr common.Address
	var approved bool
	for _, r := range recovered {
		ok, err := s.db.IsDeviceApproved(pid, r.addr)
		if err != nil {
			slog.Error("failed to check device permission", "error", err)
			c.JSON(http.StatusInternalServerError, newErrResp(errors.Wrap(err, "failed to check device permission")))
			return
		}
		if ok {
			approved = true
			addr = r.addr
			sig = r.sig
			break
		}
	}
	if !approved {
		slog.Error("device does not have permission", "project_id", pid.String())
		c.JSON(http.StatusForbidden, newErrResp(errors.New("device does not have permission")))
		return
	}

	if _, err := hexutil.Decode(req.Payload); err != nil {
		slog.Error("failed to decode payload from hex format", "error", err)
		c.JSON(http.StatusBadRequest, newErrResp(errors.Wrap(err, "failed to decode payload from hex format")))
		return
	}

	taskID := crypto.Keccak256Hash(sig)

	if err := s.db.CreateTask(
		&db.Task{
			DeviceID:       addr.Hex(),
			TaskID:         taskID.Hex(),
			Nonce:          req.Nonce,
			ProjectID:      pid.String(),
			ProjectVersion: req.ProjectVersion,
			Payload:        req.Payload,
			Signature:      hexutil.Encode(sig),
			Algorithm:      alg,
			CreatedAt:      time.Now(),
		},
	); err != nil {
		slog.Error("failed to create task to persistence layer", "error", err)
		c.JSON(http.StatusInternalServerError, newErrResp(errors.Wrap(err, "could not save task")))
		return
	}

	reqSequencer := &sequencerapi.CreateTaskReq{TaskID: taskID}
	reqSequencerJ, err := json.Marshal(reqSequencer)
	if err != nil {
		slog.Error("failed to marshal sequencer request", "error", err)
		c.JSON(http.StatusInternalServerError, newErrResp(errors.Wrap(err, "failed to marshal sequencer request")))
		return
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/task", s.sequencerAddr), "application/json", bytes.NewBuffer(reqSequencerJ))
	if err != nil {
		slog.Error("failed to call sequencer service", "error", err)
		c.JSON(http.StatusInternalServerError, newErrResp(errors.Wrap(err, "failed to call sequencer service")))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			err = errors.New(string(body))
		}
		slog.Error("failed to call sequencer service", "error", err)
		c.JSON(http.StatusInternalServerError, newErrResp(errors.Wrap(err, "failed to call sequencer service")))
		return
	}

	slog.Info("successfully processed message", "task_id", taskID.String())
	c.JSON(http.StatusOK, &CreateTaskResp{
		TaskID: taskID.String(),
	})
}

func recover(req CreateTaskReq, sig []byte) (res []*struct {
	addr common.Address
	sig  []byte
}, alg string, err error) {
	req.Signature = ""
	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to marshal request into json format")
	}

	switch req.Algorithm {
	default:
		h := sha256.Sum256(reqJson)
		rID := []uint8{0, 1}
		for _, id := range rID {
			ns := append(sig, byte(id))
			pk, err := crypto.SigToPub(h[:], ns)
			if err != nil {
				return nil, "", errors.Wrapf(err, "failed to recover public key from signature, recover_id %d", id)
			}
			res = append(res, &struct {
				addr common.Address
				sig  []byte
			}{addr: crypto.PubkeyToAddress(*pk), sig: ns})
		}
		return res, "ES256", nil
	}
}

func (s *httpServer) getProverTaskState(taskID string) (*StateLog, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/task/%s", s.proverAddr, taskID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to call prover http server")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
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
	taskIDStr := c.Param("id")
	taskID := common.HexToHash(taskIDStr)

	resp := &QueryTaskResp{
		TaskID: taskIDStr,
		States: []*StateLog{},
	}

	// Get packed state
	task, err := s.db.FetchTask(taskID)
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
	resp.ProjectID = task.ProjectID
	resp.States = append(resp.States, &StateLog{
		State: "packed",
		Time:  task.CreatedAt,
	})

	// Get assigned state
	assignedTask, err := s.db.FetchAssignedTask(taskID)
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
	settledTask, err := s.db.FetchSettledTask(taskID)
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
	proverState, err := s.getProverTaskState(taskIDStr)
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
func Run(p *db.DB, addr, sequencerAddr, proverAddr string) error {
	s := &httpServer{
		engine:        gin.Default(),
		db:            p,
		sequencerAddr: sequencerAddr,
		proverAddr:    proverAddr,
	}

	s.engine.POST("/task", s.createTask)
	s.engine.GET("/task/:id", s.queryTask)
	metrics.RegisterMetrics(s.engine)

	if err := s.engine.Run(addr); err != nil {
		slog.Error("failed to start http server", "address", addr, "error", err)
		return errors.Wrap(err, "could not start http server; check if the address is in use or network is accessible")
	}
	return nil
}
