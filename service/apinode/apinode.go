package apinode

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/w3bstream/monitor"
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
			DeleteProjectDevice:      n.db.DeleteProjectDevice,
		},
		&monitor.ContractAddr{
			TaskManager:   common.HexToAddress(n.cfg.TaskManagerContractAddr),
			ProjectDevice: common.HexToAddress(n.cfg.ProjectDeviceContractAddr),
		},
		n.cfg.BeginningBlockNumber,
		n.cfg.ChainEndpoint,
	); err != nil {
		return errors.Wrap(err, "failed to run contract monitor")
	}

	go func() {
		if err := api.Run(n.db, n.cfg.ServiceEndpoint, n.cfg.SequencerServiceEndpoint, n.cfg.ProverServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (n *APINode) Stop() error {
	return nil
}
