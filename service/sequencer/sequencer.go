package sequencer

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/datasource"
	"github.com/iotexproject/w3bstream/monitor"
	"github.com/iotexproject/w3bstream/service/sequencer/api"
	"github.com/iotexproject/w3bstream/service/sequencer/config"
	"github.com/iotexproject/w3bstream/service/sequencer/db"
	"github.com/iotexproject/w3bstream/task/assigner"
)

type Sequencer struct {
	cfg        *config.Config
	db         *db.DB
	privateKey *ecdsa.PrivateKey
}

func NewSequencer(cfg *config.Config, db *db.DB, prv *ecdsa.PrivateKey) *Sequencer {
	return &Sequencer{
		cfg:        cfg,
		db:         db,
		privateKey: prv,
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
		return errors.Wrap(err, "failed to start monitor")
	}

	datasource, err := datasource.NewPostgres(s.cfg.DatasourceDSN)
	if err != nil {
		return errors.Wrap(err, "failed to new datasource")
	}

	err = assigner.Run(s.db, s.privateKey, s.cfg.ChainEndpoint, datasource.Retrieve, common.HexToAddress(s.cfg.MinterContractAddr), s.cfg.TaskProcessingBandwidth)

	go func() {
		if err := api.Run(s.cfg.ServiceEndpoint, s.db); err != nil {
			log.Fatal(err)
		}
	}()

	return errors.Wrap(err, "failed to run task assigner")
}

func (s *Sequencer) Stop() error {
	return nil
}

func (s *Sequencer) Address() common.Address {
	if s.privateKey == nil {
		log.Fatal("private key is nil")
	}
	return crypto.PubkeyToAddress(s.privateKey.PublicKey)
}
