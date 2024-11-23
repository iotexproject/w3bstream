package prover

import (
	"crypto/ecdsa"
	"encoding/json"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/datasource"
	"github.com/iotexproject/w3bstream/monitor"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/service/prover/api"
	"github.com/iotexproject/w3bstream/service/prover/config"
	"github.com/iotexproject/w3bstream/service/prover/db"
	"github.com/iotexproject/w3bstream/task/processor"
	"github.com/iotexproject/w3bstream/vm"
)

type Prover struct {
	db        *db.DB
	cfg       *config.Config
	prv       *ecdsa.PrivateKey
	vmHandler *vm.Handler
}

func NewProver(cfg *config.Config, db *db.DB, privateKey *ecdsa.PrivateKey) *Prover {
	vmEndpoints := map[uint64]string{}
	if err := json.Unmarshal([]byte(cfg.VMEndpoints), &vmEndpoints); err != nil {
		log.Fatal(errors.Wrap(err, "failed to unmarshal vm endpoints"))
	}
	vmHandler, err := vm.NewHandler(vmEndpoints)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new vm handler"))
	}
	return &Prover{
		cfg:       cfg,
		vmHandler: vmHandler,
		prv:       privateKey,
		db:        db,
	}
}

func (p *Prover) Start() error {
	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       p.db.ScannedBlockNumber,
			UpsertScannedBlockNumber: p.db.UpsertScannedBlockNumber,
			AssignTask:               p.db.CreateTask,
			UpsertProject:            p.db.UpsertProject,
			SettleTask:               p.db.DeleteTask,
		},
		&monitor.ContractAddr{
			Project:     common.HexToAddress(p.cfg.ProjectContractAddr),
			TaskManager: common.HexToAddress(p.cfg.TaskManagerContractAddr),
		},
		p.cfg.BeginningBlockNumber,
		p.cfg.ChainEndpoint,
	); err != nil {
		return errors.Wrap(err, "failed to run monitor")
	}

	projectManager := project.NewManager(p.db.Project, p.db.ProjectFile, p.db.UpsertProjectFile)

	datasource, err := datasource.NewClickhouse(p.cfg.DatasourceDSN)
	if err != nil {
		return errors.Wrap(err, "failed to new datasource")
	}

	if err := processor.Run(p.vmHandler.Handle, projectManager.Project, p.db, datasource.Retrieve, p.prv, p.cfg.ChainEndpoint, common.HexToAddress(p.cfg.RouterContractAddr)); err != nil {
		return errors.Wrap(err, "failed to run task processor")
	}

	go func() {
		if err := api.Run(p.db, p.cfg.ServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (p *Prover) Stop() error {
	return nil
}
