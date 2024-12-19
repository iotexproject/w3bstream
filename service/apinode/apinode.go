package apinode

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/w3bstream/monitor"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/service/apinode/aggregator"
	"github.com/iotexproject/w3bstream/service/apinode/api"
	"github.com/iotexproject/w3bstream/service/apinode/config"
	"github.com/iotexproject/w3bstream/service/apinode/db"
	"github.com/pkg/errors"
)

type APINode struct {
	cfg *config.Config
	db  *db.DB
}

func NewAPINode(cfg *config.Config, db *db.DB) *APINode {
	return &APINode{
		cfg: cfg,
		db:  db,
	}
}

func (n *APINode) Start() error {
	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       n.db.ScannedBlockNumber,
			UpsertScannedBlockNumber: n.db.UpsertScannedBlockNumber,
			AssignTask:               n.db.UpsertAssignedTask,
			SettleTask:               n.db.UpsertSettledTask,
			UpsertProjectDevice:      n.db.UpsertProjectDevice,
			UpsertProject:            n.db.UpsertProject,
		},
		&monitor.ContractAddr{
			Project:     common.HexToAddress(n.cfg.ProjectContractAddr),
			TaskManager: common.HexToAddress(n.cfg.TaskManagerContractAddr),
			IoID:        common.HexToAddress(n.cfg.IoIDContractAddr),
		},
		n.cfg.BeginningBlockNumber,
		n.cfg.ChainEndpoint,
	); err != nil {
		return errors.Wrap(err, "failed to run contract monitor")
	}

	projectManager := project.NewManager(n.db.Project, n.db.ProjectFile, n.db.UpsertProjectFile)

	go aggregator.Run(n.db, n.cfg.SequencerServiceEndpoint, time.Duration(n.cfg.TaskAggregatorIntervalSecond)*time.Second)

	go func() {
		if err := api.Run(n.db, projectManager, n.cfg.ServiceEndpoint, n.cfg.SequencerServiceEndpoint, n.cfg.ProverServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (n *APINode) Stop() error {
	return nil
}
