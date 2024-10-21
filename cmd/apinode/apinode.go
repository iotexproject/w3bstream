package main

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/w3bstream/cmd/apinode/api"
	"github.com/iotexproject/w3bstream/cmd/apinode/config"
	"github.com/iotexproject/w3bstream/cmd/apinode/persistence"
	"github.com/iotexproject/w3bstream/monitor"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/pkg/errors"
)

type (
	APINode struct {
		config     *config.Config
		db         *persistence.Persistence
		privatekey *ecdsa.PrivateKey
	}
)

func NewAPINode(config *config.Config, db *persistence.Persistence, privatekey *ecdsa.PrivateKey) *APINode {
	return &APINode{
		config:     config,
		db:         db,
		privatekey: privatekey,
	}
}

func (n *APINode) Start() error {
	pubSub, err := p2p.NewPubSub(n.config.BootNodeMultiAddr, n.config.IoTeXChainID, nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new p2p pubsub"))
	}

	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       n.db.ScannedBlockNumber,
			UpsertScannedBlockNumber: n.db.UpsertScannedBlockNumber,
			AssignTask:               n.db.UpsertAssignedTask,
			SettleTask:               n.db.UpsertSettledTask,
		},
		&monitor.ContractAddr{
			TaskManager: common.HexToAddress(n.config.TaskManagerContractAddr),
		},
		n.config.BeginningBlockNumber,
		n.config.ChainEndpoint,
	); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run contract monitor"))
	}

	go func() {
		if err := api.Run(n.db, n.privatekey, pubSub, n.config.AggregationAmount, n.config.ServiceEndpoint, n.config.ProverServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (n *APINode) Stop() error {
	return nil
}
