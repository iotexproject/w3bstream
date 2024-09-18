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
	privateKey        *ecdsa.PrivateKey
	pubSub            *p2p.PubSub
}

func (s *httpServer) handleMessage(c *gin.Context) {
	req := &HandleMessageReq{}
	if err := c.ShouldBind(req); err != nil {
		slog.Error("failed to bind request", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(err))
		return
	}

	sigStr := req.Signature
	req.Signature = ""

	reqJson, err := json.Marshal(req)
	if err != nil {
		slog.Error("failed to marshal request", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrResp(err))
		return
	}

	sig, err := hexutil.Decode(sigStr)
	if err != nil {
		slog.Error("failed to decode signature", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(err))
		return
	}
	h := crypto.Keccak256Hash(reqJson)
	sigpk, err := crypto.SigToPub(h.Bytes(), sig)
	if err != nil {
		slog.Error("failed to recover public key", "error", err)
		c.JSON(http.StatusBadRequest, NewErrResp(err))
		return
	}
	addr := crypto.PubkeyToAddress(*sigpk)

	taskID, err := s.p.Save(s.pubSub,
		&persistence.Message{
			DeviceID:       addr,
			ProjectID:      req.ProjectID,
			ProjectVersion: req.ProjectVersion,
			Data:           []byte(req.Data),
		}, s.aggregationAmount, s.privateKey,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrResp(err))
		return
	}

	resp := &HandleMessageResp{}
	if !bytes.Equal(taskID[:], common.Hash{}.Bytes()) {
		resp.TaskID = taskID.String()
	}
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
		privateKey:        prv,
		pubSub:            pubSub,
	}

	s.engine.POST("/message", s.handleMessage)
	s.engine.GET("/message/:id", s.queryStateLogByID)

	if err := s.engine.Run(address); err != nil {
		return errors.Wrap(err, "failed to run http server")
	}
	return nil
}
