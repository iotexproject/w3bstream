package api

import (
	"encoding/hex"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	solanatypes "github.com/blocto/solana-go-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/iotexproject/w3bstream/apitypes"
	"github.com/iotexproject/w3bstream/block"
	"github.com/iotexproject/w3bstream/cmd/sequencer/config"
	"github.com/iotexproject/w3bstream/persistence/postgres"
)

type HttpServer struct {
	engine          *gin.Engine
	persistence     *postgres.Postgres
	conf            *config.Config
	coordinatorConf *apitypes.CoordinatorConfigRsp
}

func NewHttpServer(persistence *postgres.Postgres, conf *config.Config) *HttpServer {
	s := &HttpServer{
		engine:      gin.Default(),
		persistence: persistence,
		conf:        conf,
	}

	s.coordinatorConf = &apitypes.CoordinatorConfigRsp{
		ProjectContractAddress: s.conf.ProjectContractAddr,
	}

	if len(s.conf.OperatorPriKey) > 0 {
		pk := crypto.ToECDSAUnsafe(common.FromHex(s.conf.OperatorPriKey))
		sender := crypto.PubkeyToAddress(pk.PublicKey)
		s.coordinatorConf.OperatorETHAddress = sender.String()
	}

	if len(s.conf.OperatorPriKeyED25519) > 0 {
		wallet, err := solanatypes.AccountFromHex(s.conf.OperatorPriKeyED25519)
		if err != nil {
			panic(errors.Wrapf(err, "invalid solana wallet address"))
		}
		s.coordinatorConf.OperatorSolanaAddress = wallet.PublicKey.String()
	}

	s.engine.GET("/", s.jsonRPC)
	s.engine.POST("/", s.jsonRPC)
	s.engine.GET("/live", s.liveness)
	s.engine.GET("/task/:project_id/:task_id", s.getTaskStateLog)
	s.engine.GET("/coordinator_config", s.getCoordinatorConfigInfo)
	s.engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return s
}

// this func will block caller
func (s *HttpServer) Run(endpoint string) error {
	if err := s.engine.Run(endpoint); err != nil {
		return errors.Wrap(err, "start http server failed")
	}
	return nil
}

func (s *HttpServer) jsonRPC(c *gin.Context) {
	req := &jsonRpcReq{}
	rsp := &jsonRpcRsp{
		ID:      1,
		Version: "1.0",
	}
	if err := c.ShouldBind(req); err != nil {
		slog.Error("failed to bind request param", "error", err)
		rsp.Error = err.Error()
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	switch req.Method {
	case "getblocktemplate":
		h := &block.Header{
			Meta:       [4]byte{},
			PrevHash:   common.Hash{},
			MerkleRoot: [32]byte{},
			Difficulty: 500,
			Nonce:      [8]byte{},
		}
		t := &blockTemplate{
			Meta:          hex.EncodeToString(h.Meta[:]),
			PrevBlockHash: hex.EncodeToString(h.PrevHash[:]),
			MerkleRoot:    hex.EncodeToString(h.MerkleRoot[:]),
			Difficulty:    h.Difficulty,
			Ts:            uint64(time.Time{}.Unix()),
			NonceRange:    hex.EncodeToString(h.Nonce[:]),
		}
		rsp.Result = t
		c.JSON(http.StatusOK, rsp)
	case "submitblock":
		params := []*submitBlockParam{}
		if err := json.Unmarshal(req.Params, &params); err != nil {
			slog.Error("failed to unmarshal submit block param", "error", err)
			rsp.Error = err.Error()
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
		if len(params) == 0 {
			slog.Error("empty submit block param")
			rsp.Error = "empty submit block param"
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
		param := params[0]
		h, err := param.toBlockHeader()
		if err != nil {
			slog.Error("failed to construct block header", "error", err)
			rsp.Error = err.Error()
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
		if !h.IsValid() {
			slog.Error("invalid nonce")
			rsp.Error = "invalid nonce"
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
		c.JSON(http.StatusOK, rsp)
	default:
		slog.Error("illegal method")
		rsp.Error = "illegal method"
		c.JSON(http.StatusBadRequest, rsp)
	}
}

func (s *HttpServer) liveness(c *gin.Context) {
	c.JSON(http.StatusOK, &apitypes.LivenessRsp{Status: "up"})
}

func (s *HttpServer) getTaskStateLog(c *gin.Context) {
	projectID, err := strconv.ParseUint(c.Param("project_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, apitypes.NewErrRsp(err))
		return
	}
	taskID, err := strconv.ParseUint(c.Param("task_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, apitypes.NewErrRsp(err))
		return
	}

	ls, err := s.persistence.Fetch(taskID, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apitypes.NewErrRsp(err))
		return
	}
	if len(ls) == 0 {
		c.JSON(http.StatusOK, &apitypes.QueryTaskStateLogRsp{
			TaskID:    taskID,
			ProjectID: projectID,
		})
		return
	}

	ss := []*apitypes.StateLog{}
	for _, l := range ls {
		ss = append(ss, &apitypes.StateLog{
			State:   l.State.String(),
			Time:    l.CreatedAt,
			Comment: l.Comment,
			Result:  string(l.Result),
		})
	}

	c.JSON(http.StatusOK, &apitypes.QueryTaskStateLogRsp{
		TaskID:    taskID,
		ProjectID: projectID,
		States:    ss,
	})
}

func (s *HttpServer) getCoordinatorConfigInfo(c *gin.Context) {
	c.JSON(http.StatusOK, s.coordinatorConf)
}
