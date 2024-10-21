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

type (
	APINode struct {
		Config     *config.Config
		db         *persistence.Persistence
		privatekey *ecdsa.PrivateKey
	}
)

func NewAPINode(config *config.Config, db *persistence.Persistence, privatekey *ecdsa.PrivateKey) *APINode {
	return &APINode{
		Config:     config,
		db:         db,
		privatekey: privatekey,
	}
}

func (n *APINode) Start() error {
	pubSub, err := p2p.NewPubSub(n.Config.BootNodeMultiAddr, n.Config.IoTeXChainID, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create pubsub")
	}

	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       n.db.ScannedBlockNumber,
			UpsertScannedBlockNumber: n.db.UpsertScannedBlockNumber,
			AssignTask:               n.db.UpsertAssignedTask,
			SettleTask:               n.db.UpsertSettledTask,
		},
		&monitor.ContractAddr{
			TaskManager: common.HexToAddress(n.Config.TaskManagerContractAddr),
		},
		n.Config.BeginningBlockNumber,
		n.Config.ChainEndpoint,
	); err != nil {
		return errors.Wrap(err, "failed to start monitor")
	}

	go func() {
		if err := api.Run(n.db, n.privatekey, pubSub, n.Config.AggregationAmount, n.Config.ServiceEndpoint, n.Config.ProverServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (n *APINode) Stop() error {
	return nil
}
