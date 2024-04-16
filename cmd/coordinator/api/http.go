package api

import (
	"net/http"
	"strconv"

	solanatypes "github.com/blocto/solana-go-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/apitypes"
	"github.com/machinefi/sprout/cmd/coordinator/config"
	"github.com/machinefi/sprout/persistence/postgres"
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
		ProjectContractAddress: s.conf.ProjectContractAddress,
	}

	if len(s.conf.OperatorPrivateKey) > 0 {
		pk := crypto.ToECDSAUnsafe(common.FromHex(s.conf.OperatorPrivateKey))
		sender := crypto.PubkeyToAddress(pk.PublicKey)
		s.coordinatorConf.OperatorETHAddress = sender.String()
	}

	if len(s.conf.OperatorPrivateKeyED25519) > 0 {
		wallet, err := solanatypes.AccountFromHex(s.conf.ServiceEndpoint)
		if err != nil {
			panic(errors.Wrapf(err, "invalid solana wallet address"))
		}
		s.coordinatorConf.OperatorSolanaAddress = wallet.PublicKey.String()
	}

	s.engine.GET("/live", s.liveness)
	s.engine.GET("/task/:project_id/:task_id", s.getTaskStateLog)
	s.engine.GET("/coordinator_config", s.getCoordinatorConfigInfo)

	return s
}

// this func will block caller
func (s *HttpServer) Run(endpoint string) error {
	if err := s.engine.Run(endpoint); err != nil {
		return errors.Wrap(err, "start http server failed")
	}
	return nil
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
