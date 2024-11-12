package apinode

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/w3bstream/monitor"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/service/apinode/api"
	"github.com/iotexproject/w3bstream/service/apinode/config"
	"github.com/iotexproject/w3bstream/service/apinode/persistence"
	"github.com/pkg/errors"
)

type APINode struct {
	cfg *config.Config
	db  *persistence.Persistence
	prv *ecdsa.PrivateKey
}

func NewAPINode(cfg *config.Config, db *persistence.Persistence, prv *ecdsa.PrivateKey) *APINode {
	return &APINode{
		cfg: cfg,
		db:  db,
		prv: prv,
	}
}

func (n *APINode) Start() error {
	pubSub, err := p2p.NewPubSub(n.cfg.BootNodeMultiAddr, n.cfg.IoTeXChainID, nil)
	if err != nil {
		return errors.Wrap(err, "failed to new p2p pubsub")
	}

	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       n.db.ScannedBlockNumber,
			UpsertScannedBlockNumber: n.db.UpsertScannedBlockNumber,
			AssignTask:               n.db.UpsertAssignedTask,
			SettleTask:               n.db.UpsertSettledTask,
		},
		&monitor.ContractAddr{
			TaskManager: common.HexToAddress(n.cfg.TaskManagerContractAddr),
		},
		n.cfg.BeginningBlockNumber,
		n.cfg.ChainEndpoint,
	); err != nil {
		return errors.Wrap(err, "failed to run contract monitor")
	}

	go func() {
		if err := api.Run(n.db, n.prv, pubSub, n.cfg.ServiceEndpoint, n.cfg.ProverServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (n *APINode) Stop() error {
	return nil
}
