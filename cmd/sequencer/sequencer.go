package main

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/sequencer/config"
	"github.com/iotexproject/w3bstream/cmd/sequencer/db"
	"github.com/iotexproject/w3bstream/datasource"
	"github.com/iotexproject/w3bstream/monitor"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/task/assigner"
)

type Sequencer struct {
	cfg *config.Config
	db  *db.DB
	prv *ecdsa.PrivateKey
}

func NewSequencer(cfg *config.Config, db *db.DB, prv *ecdsa.PrivateKey) *Sequencer {
	return &Sequencer{
		cfg: cfg,
		db:  db,
		prv: prv,
	}
}

func (s *Sequencer) Start() error {
	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       s.db.ScannedBlockNumber,
			UpsertScannedBlockNumber: s.db.UpsertScannedBlockNumber,
			UpsertProver:             s.db.UpsertProver,
			SettleTask:               s.db.DeleteTask,
		},
		&monitor.ContractAddr{
			Prover:      common.HexToAddress(s.cfg.ProverContractAddr),
			Minter:      common.HexToAddress(s.cfg.MinterContractAddr),
			TaskManager: common.HexToAddress(s.cfg.TaskManagerContractAddr),
		},
		s.cfg.BeginningBlockNumber,
		s.cfg.ChainEndpoint,
	); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run contract monitor"))
	}

	if _, err := p2p.NewPubSub(s.cfg.BootNodeMultiAddr, s.cfg.IoTeXChainID, s.db.CreateTask); err != nil {
		log.Fatal(errors.Wrap(err, "failed to new p2p pubsub"))
	}

	datasource, err := datasource.NewPostgres(s.cfg.DatasourceDSN)
	if err != nil {
		return errors.Wrap(err, "failed to new datasource")
	}

	err = assigner.Run(s.db, s.prv, s.cfg.ChainEndpoint, datasource.Retrieve, common.HexToAddress(s.cfg.MinterContractAddr))
	return errors.Wrap(err, "failed to run task assigner")
}

func (s *Sequencer) Stop() error {
	return nil
}
