package api

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
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
	"github.com/tidwall/gjson"

	"github.com/iotexproject/w3bstream/metrics"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/service/apinode/db"
	proverapi "github.com/iotexproject/w3bstream/service/prover/api"
)

type errResp struct {
	Error string `json:"error,omitempty"`
}

func newErrResp(err error) *errResp {
	return &errResp{Error: err.Error()}
}

// TODO move to project file
// mainnet 6
// testnet 923
var (
	defaultProject = project.Config{
		SignatureAlgorithm: "ecdsa",
		HashAlgorithm:      "sha256",
	}
	pebbleProject = project.Config{
		SignedKeys:         []project.SignedKey{{Name: "timestamp", Type: "uint64"}},
		SignatureAlgorithm: "ecdsa",
		HashAlgorithm:      "sha256",
	}
	geoProject = project.Config{
		SignedKeys: []project.SignedKey{
			{Name: "timestamp", Type: "uint64"},
			{Name: "latitude", Type: "uint64"},
			{Name: "longitude", Type: "uint64"},
		},
		SignatureAlgorithm: "ecdsa",
		HashAlgorithm:      "sha256",
	}
)

type CreateTaskReq struct {
	Nonce          uint64 `json:"nonce"                       binding:"required"`
	ProjectID      string `json:"projectID"                   binding:"required"`
	ProjectVersion string `json:"projectVersion,omitempty"`
	Payload        []byte `json:"payload"                     binding:"required"`
	Signature      string `json:"signature,omitempty"         binding:"required"`
}

func (ct CreateTaskReq) MarshalJSON() ([]byte, error) {
	if gjson.ValidBytes(ct.Payload) {
		return json.Marshal(struct {
			Nonce          uint64          `json:"nonce"`
			ProjectID      string          `json:"projectID"`
			ProjectVersion string          `json:"projectVersion,omitempty"`
			Payload        json.RawMessage `json:"payload"`
			Signature      string          `json:"signature,omitempty"`
		}{
			Nonce:          ct.Nonce,
			ProjectID:      ct.ProjectID,
			ProjectVersion: ct.ProjectVersion,
			Payload:        ct.Payload,
			Signature:      ct.Signature,
		})
	}
	payload := hex.EncodeToString(ct.Payload)
	return json.Marshal(struct {
		Nonce          uint64 `json:"nonce"`
		ProjectID      string `json:"projectID"`
		ProjectVersion string `json:"projectVersion,omitempty"`
		Payload        string `json:"payload"`
		Signature      string `json:"signature,omitempty"`
	}{
		Nonce:          ct.Nonce,
		ProjectID:      ct.ProjectID,
		ProjectVersion: ct.ProjectVersion,
		Payload:        payload,
		Signature:      ct.Signature,
	})
}

func (ct *CreateTaskReq) UnmarshalJSON(data []byte) error {
	if !gjson.ValidBytes(data) {
		return errors.New("invalid json format")
	}
	parsed := gjson.ParseBytes(data)
	ct.Nonce = parsed.Get("nonce").Uint()
	ct.ProjectID = parsed.Get("projectID").String()
	ct.ProjectVersion = parsed.Get("projectVersion").String()
	ct.Signature = parsed.Get("signature").String()

	switch parsed.Get("payload").Type {
	case gjson.JSON:
		ct.Payload = []byte(parsed.Get("payload").Raw)
		return nil
	case gjson.String:
		payload, err := hex.DecodeString(parsed.Get("payload").String())
		if err != nil {
			return errors.Wrap(err, "failed to decode payload from hex string")
		}
		ct.Payload = payload
		return nil
	default:
		return errors.New("invalid payload format")
	}
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
		slog.Error("failed to decode signature", "error", err)
		c.JSON(http.StatusBadRequest, newErrResp(errors.Wrap(err, "failed to decode signature")))
		return
	}

	// TODO move to project file
	cfg := &defaultProject
	if req.ProjectID == "923" {
		cfg = &pebbleProject
	}
	if req.ProjectID == "942" {
		cfg = &geoProject
	}

	recovered, sigAlg, hashAlg, err := Recover(*req, cfg, sig)
	if err != nil {
		slog.Error("failed to recover public key", "error", err)
		c.JSON(http.StatusBadRequest, newErrResp(errors.Wrap(err, "invalid signature; could not recover public key")))
		return
	}
	var matchedPubkey *ecdsa.PublicKey
	var approved bool
	for _, r := range recovered {
		addr := crypto.PubkeyToAddress(*r.pubkey)
		slog.Info("recovered address", "project_id", pid.String(), "address", addr.String())
		ok, err := s.db.IsDeviceApproved(pid, addr)
		if err != nil {
			slog.Error("failed to check device permission", "error", err)
			c.JSON(http.StatusInternalServerError, newErrResp(errors.Wrap(err, "failed to check device permission")))
			return
		}
		if ok {
			approved = true
			matchedPubkey = r.pubkey
			sig = r.sig
			break
		}
	}
	if !approved {
		slog.Error("device does not have permission", "project_id", pid.String())
		c.JSON(http.StatusForbidden, newErrResp(errors.New("device does not have permission")))
		return
	}

	//ioid := crypto.PubkeyToAddress(*matchedPubkey)
	taskID := crypto.Keccak256Hash(sig)

	if err := s.db.CreateTask(
		&db.Task{
			DevicePubKey:       hexutil.Encode(crypto.FromECDSAPub(matchedPubkey)),
			TaskID:             taskID.Hex(),
			Nonce:              req.Nonce,
			ProjectID:          pid.String(),
			ProjectVersion:     req.ProjectVersion,
			Payload:            string(req.Payload),
			Signature:          hexutil.Encode(sig),
			SignatureAlgorithm: sigAlg,
			HashAlgorithm:      hashAlg,
			CreatedAt:          time.Now(),
		},
	); err != nil {
		slog.Error("failed to create task to persistence layer", "error", err)
		c.JSON(http.StatusInternalServerError, newErrResp(errors.Wrap(err, "could not save task")))
		return
	}

	slog.Info("successfully processed message", "task_id", taskID.String())
	c.JSON(http.StatusOK, &CreateTaskResp{
		TaskID: taskID.String(),
	})
}

func Recover(req CreateTaskReq, cfg *project.Config, sig []byte) (res []*struct {
	pubkey *ecdsa.PublicKey
	sig    []byte
}, sigAlg, hashAlg string, err error) {
	slog.Debug("request json info", "signature", req.Signature)
	req.Signature = ""
	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "failed to marshal request into json format")
	}

	var hash [32]byte
	switch cfg.HashAlgorithm {
	default:
		hashAlg = "sha256"
		h1 := sha256.Sum256(reqJson)
		slog.Debug("request json info", "data", string(reqJson), "hash", hexutil.Encode(h1[:]))
		if len(cfg.SignedKeys) == 0 {
			hash = h1
			break
		} else {
			// concatenate payload hash with signed keys from payload json
			if ok := gjson.ValidBytes(req.Payload); !ok {
				slog.Error("failed to validate payload in json format")
				return nil, "", "", errors.Wrap(err, "failed to validate payload in json format")
			}
			buf := new(bytes.Buffer)
			buf.Write(h1[:])
			for _, k := range cfg.SignedKeys {
				value := gjson.GetBytes(req.Payload, k.Name)
				switch k.Type {
				case "uint64":
					slog.Debug("request json info", "json value uint64", value.Uint())
					if err := binary.Write(buf, binary.LittleEndian, value.Uint()); err != nil {
						return nil, "", "", errors.New("failed to convert uint64 to bytes array")
					}
				}
			}
			slog.Debug("request json info", "hash_d_final", hexutil.Encode(buf.Bytes()))
			hash = sha256.Sum256(buf.Bytes())
			slog.Debug("request json info", "hash_final", hexutil.Encode(hash[:]))
		}
	}

	switch cfg.SignatureAlgorithm {
	default:
		sigAlg = "ecdsa"
		rID := []uint8{0, 1}
		for _, id := range rID {
			ns := append(sig, byte(id))
			pk, err := crypto.SigToPub(hash[:], ns)
			if err != nil {
				return nil, "", "", errors.Wrapf(err, "failed to recover public key from signature, recover_id %d", id)
			}
			res = append(res, &struct {
				pubkey *ecdsa.PublicKey
				sig    []byte
			}{pubkey: pk, sig: ns})
		}
		return res, sigAlg, hashAlg, nil
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
		ProverID: "did:io:" + strings.TrimPrefix(assignedTask.Prover, "0x"),
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
			Tx:      settledTask.Tx,
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
