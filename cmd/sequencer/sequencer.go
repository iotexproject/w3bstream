package main

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/sequencer/api"
	"github.com/iotexproject/w3bstream/cmd/sequencer/config"
	"github.com/iotexproject/w3bstream/cmd/sequencer/db"
	"github.com/iotexproject/w3bstream/monitor"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/task/assigner"
)

type (
	Sequencer struct {
		config     *config.Config
		database   *db.DB
		privateKey *ecdsa.PrivateKey
	}
)

func NewSequencer(config *config.Config, db *db.DB, privateKey *ecdsa.PrivateKey) *Sequencer {
	return &Sequencer{
		config:     config,
		database:   db,
		privateKey: privateKey,
	}
}

func (s *Sequencer) Start() error {
	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       s.database.ScannedBlockNumber,
			UpsertScannedBlockNumber: s.database.UpsertScannedBlockNumber,
			UpsertNBits:              s.database.UpsertNBits,
			UpsertBlockHead:          s.database.UpsertBlockHead,
			UpsertProver:             s.database.UpsertProver,
			SettleTask:               s.database.DeleteTask,
		},
		&monitor.ContractAddr{
			Prover:      common.HexToAddress(s.config.ProverContractAddr),
			Dao:         common.HexToAddress(s.config.DaoContractAddr),
			Minter:      common.HexToAddress(s.config.MinterContractAddr),
			TaskManager: common.HexToAddress(s.config.TaskManagerContractAddr),
		},
		s.config.BeginningBlockNumber,
		s.config.ChainEndpoint,
	); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run contract monitor"))
	}

	if _, err := p2p.NewPubSub(s.config.BootNodeMultiAddr, s.config.IoTeXChainID, s.database.CreateTask); err != nil {
		log.Fatal(errors.Wrap(err, "failed to new p2p pubsub"))
	}

	if err := assigner.Run(s.database, s.privateKey, s.config.ChainEndpoint, common.HexToAddress(s.config.MinterContractAddr)); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run task assigner"))
	}

	go func() {
		if err := api.Run(s.database, s.config, s.privateKey); err != nil {
			log.Fatal(errors.Wrap(err, "failed to run http server"))
		}
	}()

	return nil
}

func (s *Sequencer) Stop() error {
	return nil
}
