package api

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"log/slog"
	"net/http"

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
		slog.Error("Failed to bind request to HandleMessageReq struct", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(errors.Wrap(err, "invalid request payload")))
		return
	}

	sigStr := req.Signature
	req.Signature = ""

	reqJson, err := json.Marshal(req)
	if err != nil {
		slog.Error("Failed to marshal request into JSON format", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "failed to process request data")))
		return
	}

	sig, err := hexutil.Decode(sigStr)
	if err != nil {
                slog.Error("Failed to decode signature from hex format", "signature", sigStr, "error", err)
                c.JSON(http.StatusBadRequest, NewErrResp(errors.Wrap(err, "invalid signature format")))
		return
	}
	
	h := crypto.Keccak256Hash(reqJson)
	sigpk, err := crypto.SigToPub(h.Bytes(), sig)
	if err != nil {
		slog.Error("Failed to recover public key from signature", "error", err)
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
		slog.Error("Failed to save message to persistence layer", "error", err)
                c.JSON(http.StatusInternalServerError, NewErrResp(errors.Wrap(err, "internal server error; could not save task")))
		return
	}

	resp := &HandleMessageResp{}
	if !bytes.Equal(taskID[:], common.Hash{}.Bytes()) {
		resp.TaskID = taskID.String()
	}
	slog.Info("Successfully processed message", "taskID", resp.TaskID)
	c.JSON(http.StatusOK, resp)
}

func (s *httpServer) queryStateLogByID(c *gin.Context) {
	// messageID := c.Param("id")

	// ms, err := s.p.FetchMessage(messageID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
	// 	return
	// }
	// if len(ms) == 0 {
	// 	c.JSON(http.StatusOK, &apitypes.QueryMessageStateLogRsp{MessageID: messageID})
	// 	return
	// }
	// m := ms[0]

	// ss := []*apitypes.StateLog{
	// 	{
	// 		State: "received",
	// 		Time:  m.CreatedAt,
	// 	},
	// }

	// if m.InternalTaskID != "" {
	// 	ts, err := s.p.FetchTask(m.InternalTaskID)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
	// 		return
	// 	}
	// 	if len(ts) == 0 {
	// 		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(errors.New("cannot find task by internal task id")))
	// 		return
	// 	}
	// 	ss = append(ss, &apitypes.StateLog{
	// 		State: task.StatePacked.String(),
	// 		Time:  ts[0].CreatedAt,
	// 	})
	// 	resp, err := http.Get(fmt.Sprintf("http://%s/%s/%d/%d", "mock http endpoint", "task", m.ProjectID, ts[0].ID))
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
	// 		return
	// 	}
	// 	defer resp.Body.Close()

	// 	body, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
	// 		return
	// 	}
	// 	taskStateLog := &apitypes.QueryTaskStateLogRsp{}
	// 	if err := json.Unmarshal(body, &taskStateLog); err != nil {
	// 		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
	// 		return
	// 	}
	// 	ss = append(ss, taskStateLog.States...)
	// }

	// c.JSON(http.StatusOK, &apitypes.QueryMessageStateLogRsp{MessageID: messageID, States: ss})
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
