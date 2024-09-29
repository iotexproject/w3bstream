package api

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"log/slog"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/block"
	"github.com/iotexproject/w3bstream/cmd/sequencer/config"
	"github.com/iotexproject/w3bstream/cmd/sequencer/db"
	"github.com/iotexproject/w3bstream/smartcontracts/go/minter"
)

type Sequencer struct {
	Addr        common.Address
	Operator    common.Address
	Beneficiary common.Address
}

type httpServer struct {
	engine         *gin.Engine
	db             *db.DB
	prv            *ecdsa.PrivateKey
	account        common.Address
	client         *ethclient.Client
	minterInstance *minter.Minter
	signer         types.Signer
}

func (s *httpServer) getTip(c *gin.Context, rsp *jsonRpcRsp) {
	head, hash, err := s.db.BlockHead()
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
}

func (s *httpServer) getBlockTemplate(c *gin.Context, rsp *jsonRpcRsp) {
	head, hash, err := s.db.BlockHead()
	if err != nil {
		slog.Error("failed to get prev hash", "error", err)
		rsp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}
	nbits, err := s.db.NBits()
	if err != nil {
		slog.Error("failed to get nbits", "error", err)
		rsp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}
	coinbase := Sequencer{
		Operator: s.account,
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
}

func (s *httpServer) submitBlock(c *gin.Context, req *jsonRpcReq, rsp *jsonRpcRsp) {
	data, err := json.Marshal(req.Params)
	if err != nil {
		slog.Error("failed to marshal submit block param", "error", err)
		rsp.Error = err.Error()
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
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

	nonce, err := s.client.PendingNonceAt(context.Background(), s.account)
	if err != nil {
		slog.Error("failed to get pending nonce", "error", err)
		rsp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}
	tx, err := s.minterInstance.Mint(
		&bind.TransactOpts{
			From: s.account,
			Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(t, s.signer, s.prv)
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
			Operator: s.account,
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
}

func (s *httpServer) jsonRPC(c *gin.Context) {
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

	reqJ, _ := json.Marshal(req)
	slog.Info("json rpc", "data", string(reqJ))

	switch req.Method {
	case "gettip":
		s.getTip(c, rsp)
	case "getblocktemplate":
		s.getBlockTemplate(c, rsp)
	case "submitblock":
		s.submitBlock(c, req, rsp)
	default:
		slog.Error("illegal method")
		rsp.Error = "illegal method"
		c.JSON(http.StatusBadRequest, rsp)
	}
}

// this func will block caller
func Run(db *db.DB, cfg *config.Config) error {
	client, err := ethclient.Dial(cfg.ChainEndpoint)
	if err != nil {
		return errors.Wrap(err, "failed to dial chain endpoint")
	}
	minterInstance, err := minter.NewMinter(common.HexToAddress(cfg.MinterContractAddr), client)
	if err != nil {
		return errors.Wrap(err, "failed to new minter contract instance")
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get chain id")
	}

	prv := crypto.ToECDSAUnsafe(common.FromHex(cfg.OperatorPrvKey))
	s := &httpServer{
		engine:         gin.Default(),
		db:             db,
		prv:            prv,
		account:        crypto.PubkeyToAddress(prv.PublicKey),
		client:         client,
		minterInstance: minterInstance,
		signer:         types.NewLondonSigner(chainID),
	}

	s.engine.GET("/", s.jsonRPC)
	s.engine.POST("/", s.jsonRPC)

	err = s.engine.Run(cfg.ServiceEndpoint)
	return errors.Wrap(err, "failed to start http server")
}
