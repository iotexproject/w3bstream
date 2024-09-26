package api

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	. "github.com/agiledragon/gomonkey/v2"
	solanatypes "github.com/blocto/solana-go-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/apitypes"
	"github.com/iotexproject/w3bstream/cmd/sequencer/config"
	"github.com/iotexproject/w3bstream/persistence/postgres"
	"github.com/iotexproject/w3bstream/task"
)

func TestNewHttpServer(t *testing.T) {
	r := require.New(t)

	conf := &config.Config{
		OperatorPrvKey:        "privateKey",
		OperatorPriKeyED25519: "PrivateKeyED25519",
	}
	t.Run("InvalidSolanaAddress", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(crypto.ToECDSAUnsafe, &ecdsa.PrivateKey{})
		p.ApplyFuncReturn(crypto.PubkeyToAddress, common.Address{})
		defer func() {
			r.NotNil(recover())
		}()

		_ = NewHttpServer(nil, conf)
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(crypto.ToECDSAUnsafe, &ecdsa.PrivateKey{})
		p.ApplyFuncReturn(crypto.PubkeyToAddress, common.Address{})
		p.ApplyFuncReturn(solanatypes.AccountFromHex, solanatypes.Account{PublicKey: [32]byte{1}}, nil)

		s := NewHttpServer(nil, conf)
		r.NotNil(s)
	})
}

func TestHttpServer_Run(t *testing.T) {
	r := require.New(t)

	s := &httpServer{
		engine: gin.Default(),
	}

	t.Run("FailedToRun", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gin.Engine{}, "Run", errors.New(t.Name()))

		err := s.Run("")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&gin.Engine{}, "Run", nil)

		err := s.Run("")
		r.NoError(err)
	})
}

func TestHttpServer_liveness(t *testing.T) {
	r := require.New(t)

	s := &httpServer{
		engine: gin.Default(),
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	s.liveness(c)
	r.Equal(http.StatusOK, w.Code)

	actualResponse := &apitypes.LivenessRsp{}
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	r.NoError(err)
	r.Equal(&apitypes.LivenessRsp{Status: "up"}, actualResponse)
}

func TestHttpServer_getCoordinatorConfigInfo(t *testing.T) {
	r := require.New(t)

	s := &httpServer{
		engine: gin.Default(),
		coordinatorConf: &apitypes.CoordinatorConfigRsp{
			ProjectContractAddress: "projectContractAddress",
			OperatorETHAddress:     "operatorETHAddress",
			OperatorSolanaAddress:  "operatorSolanaAddress",
		},
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	s.getCoordinatorConfigInfo(c)
	r.Equal(http.StatusOK, w.Code)

	actualResponse := &apitypes.CoordinatorConfigRsp{}
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	r.NoError(err)
	r.Equal(&apitypes.CoordinatorConfigRsp{
		ProjectContractAddress: "projectContractAddress",
		OperatorETHAddress:     "operatorETHAddress",
		OperatorSolanaAddress:  "operatorSolanaAddress",
	}, actualResponse)
}

func TestHttpServer_getTaskStateLog(t *testing.T) {
	r := require.New(t)

	s := &httpServer{
		engine: gin.Default(),
		db:     &postgres.Postgres{},
	}

	t.Run("FailedToParseProjectID", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(strconv.ParseUint, uint64(0), errors.New(t.Name()))

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		s.getTaskStateLog(c)
		r.Equal(http.StatusBadRequest, w.Code)

		actualResponse := &apitypes.ErrRsp{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		r.NoError(err)
		r.Equal(&apitypes.ErrRsp{
			Error: t.Name(),
		}, actualResponse)
	})

	t.Run("FailedToParseTaskID", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		outputs := []OutputCell{
			{Values: Params{uint64(0), nil}},
			{Values: Params{uint64(0), errors.New(t.Name())}},
		}
		p.ApplyFuncSeq(strconv.ParseUint, outputs)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		s.getTaskStateLog(c)
		r.Equal(http.StatusBadRequest, w.Code)

		actualResponse := &apitypes.ErrRsp{}
		fmt.Println(string(w.Body.Bytes()))
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		r.NoError(err)
		r.Equal(&apitypes.ErrRsp{
			Error: t.Name(),
		}, actualResponse)
	})

	t.Run("FailedToFetch", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		outputs := []OutputCell{
			{Values: Params{uint64(0), nil}},
			{Values: Params{uint64(0), nil}},
		}
		p.ApplyFuncSeq(strconv.ParseUint, outputs)
		p.ApplyMethodReturn(&postgres.Postgres{}, "Fetch", nil, errors.New(t.Name()))

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		s.getTaskStateLog(c)
		r.Equal(http.StatusInternalServerError, w.Code)

		actualResponse := &apitypes.ErrRsp{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		r.NoError(err)
		r.Equal(&apitypes.ErrRsp{
			Error: t.Name(),
		}, actualResponse)
	})

	t.Run("FetchZero", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		outputs := []OutputCell{
			{Values: Params{uint64(0), nil}},
			{Values: Params{uint64(0), nil}},
			{Values: Params{uint64(0), nil}}, // need for json.Unmarshal
			{Values: Params{uint64(0), nil}},
		}
		p.ApplyFuncSeq(strconv.ParseUint, outputs)
		p.ApplyMethodReturn(&postgres.Postgres{}, "Fetch", []*task.StateLog{}, nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		s.getTaskStateLog(c)
		r.Equal(http.StatusOK, w.Code)

		actualResponse := &apitypes.QueryTaskStateLogRsp{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		r.NoError(err)
		r.Equal(&apitypes.QueryTaskStateLogRsp{
			TaskID:    0,
			ProjectID: 0,
		}, actualResponse)
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		ts := []*task.StateLog{
			{
				State:     task.StateOutputted,
				Comment:   "comment",
				Result:    []byte("result"),
				CreatedAt: time.Time{},
			},
		}

		outputs := []OutputCell{
			{Values: Params{uint64(0), nil}},
			{Values: Params{uint64(0), nil}},
			{Values: Params{uint64(0), nil}}, // need for json.Unmarshal
			{Values: Params{uint64(0), nil}},
		}
		p.ApplyFuncSeq(strconv.ParseUint, outputs)
		p.ApplyMethodReturn(&postgres.Postgres{}, "Fetch", ts, nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		s.getTaskStateLog(c)
		r.Equal(http.StatusOK, w.Code)

		actualResponse := &apitypes.QueryTaskStateLogRsp{}
		err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
		r.NoError(err)
		r.Equal(&apitypes.QueryTaskStateLogRsp{
			TaskID:    0,
			ProjectID: 0,
			States: []*apitypes.StateLog{
				{
					State:   ts[0].State.String(),
					Time:    ts[0].CreatedAt,
					Comment: ts[0].Comment,
					Result:  string(ts[0].Result),
				},
			},
		}, actualResponse)
	})
}
