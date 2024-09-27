package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"math/big"
	"net/http"
	"strconv"
	"time"

	solanatypes "github.com/blocto/solana-go-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/iotexproject/w3bstream/apitypes"
	"github.com/iotexproject/w3bstream/block"
	"github.com/iotexproject/w3bstream/cmd/sequencer/config"
	"github.com/iotexproject/w3bstream/persistence/postgres"
	"github.com/iotexproject/w3bstream/smartcontracts/go/minter"
)

var prv = crypto.ToECDSAUnsafe(common.FromHex("33e6ba3e033131026903f34dfa208feb88c284880530cf76280b68d38041c67b"))

type Sequencer struct {
	Addr        common.Address
	Operator    common.Address
	Beneficiary common.Address
}

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
	if err := c.ShouldBindJSON(req); err != nil {
		slog.Error("failed to bind request param", "error", err)
		rsp.Error = err.Error()
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	switch req.Method {
	case "gettip":
		head, hash, err := s.persistence.ChainHead()
		if err != nil {
			slog.Error("failed to get prev hash", "error", err)
			rsp.Error = err.Error()
			c.JSON(http.StatusInternalServerError, rsp)
			return
		}
		t := &blockTip{
			BlockNumber: head,
			BlockHash:   hexutil.Encode(hash[:]),
		}
		rsp.Result = t
		c.JSON(http.StatusOK, rsp)
	case "getblocktemplate":
		head, hash, err := s.persistence.ChainHead()
		if err != nil {
			slog.Error("failed to get prev hash", "error", err)
			rsp.Error = err.Error()
			c.JSON(http.StatusInternalServerError, rsp)
			return
		}
		nbits, err := s.persistence.NBits()
		if err != nil {
			slog.Error("failed to get nbits", "error", err)
			rsp.Error = err.Error()
			c.JSON(http.StatusInternalServerError, rsp)
			return
		}
		coinbase := Sequencer{
			Operator: crypto.PubkeyToAddress(prv.PublicKey),
		}
		var rootData bytes.Buffer
		rootData.Write(coinbase.Addr[:])
		rootData.Write(coinbase.Operator[:])
		rootData.Write(coinbase.Beneficiary[:])

		h := &block.Header{
			Meta:       [4]byte{},
			PrevHash:   hash,
			MerkleRoot: [32]byte{},
			NBits:      nbits,
			Nonce:      [8]byte{},
		}
		t := &blockTemplate{
			PrevBlockNumber: head,
			Meta:            hexutil.Encode(h.Meta[:]),
			PrevBlockHash:   hexutil.Encode(h.PrevHash[:]),
			MerkleRoot:      hexutil.Encode(crypto.Keccak256Hash(rootData.Bytes()).Bytes()),
			NBits:           h.NBits,
			Ts:              uint64(time.Time{}.Unix()),
			NonceRange:      hexutil.Encode(h.Nonce[:]),
		}
		rsp.Result = t
		c.JSON(http.StatusOK, rsp)
	case "submitblock":
		data, err := json.Marshal(req.Params)
		if err != nil {
			slog.Error("failed to marshal submit block param", "error", err)
			rsp.Error = err.Error()
			c.JSON(http.StatusBadRequest, rsp)
			return
		}
		slog.Info("submitblock data", "data", string(data))
		params := []*submitBlockParam{}
		if err := json.Unmarshal(data, &params); err != nil {
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
		client, err := ethclient.Dial("https://babel-nightly.iotex.io")
		if err != nil {
			panic(err)
		}
		minterInstance, err := minter.NewMinter(common.HexToAddress("0xa77Be024413F955699E1eC3D0AdbbeAD8b11cFEE"), client)
		if err != nil {
			panic(err)
		}
		chainID, err := client.ChainID(context.Background())
		if err != nil {
			panic(err)
		}
		nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(prv.PublicKey))
		if err != nil {
			panic(err)
		}
		tx, err := minterInstance.Mint(&bind.TransactOpts{
			From: crypto.PubkeyToAddress(prv.PublicKey),
			Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(t, types.NewLondonSigner(chainID), prv)
			},
			Nonce: new(big.Int).SetUint64(nonce),
		},
			minter.BlockInfo{
				Meta:       h.Meta,
				Prevhash:   h.PrevHash,
				MerkleRoot: h.MerkleRoot,
				Nbits:      h.NBits,
				Nonce:      h.Nonce,
			},
			minter.Sequencer{
				Operator: crypto.PubkeyToAddress(prv.PublicKey),
			},
			nil,
		)
		if err != nil {
			slog.Error("failed to send tx", "error", err)
			rsp.Error = err.Error()
			c.JSON(http.StatusInternalServerError, rsp)
			return
		}
		slog.Info("mint block success", "hash", tx.Hash().Hex())
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
