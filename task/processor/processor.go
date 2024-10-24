package processor

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/smartcontracts/go/router"
	"github.com/iotexproject/w3bstream/task"
)

type HandleTask func(task *task.Task, vmTypeID uint64, code string, expParam string) ([]byte, error)
type Project func(projectID uint64) (*project.Project, error)
type RetrieveTask func(projectID uint64, taskID common.Hash) (*task.Task, error)

type DB interface {
	UnprocessedTask() (uint64, common.Hash, error)
	ProcessTask(uint64, common.Hash, error) error
}

type processor struct {
	db             DB
	retrieve       RetrieveTask
	handle         HandleTask
	project        Project
	prv            *ecdsa.PrivateKey
	waitingTime    time.Duration
	signer         types.Signer
	account        common.Address
	client         *ethclient.Client
	routerInstance *router.Router
}

func (r *processor) process(projectID uint64, taskID common.Hash) error {
	t, err := r.retrieve(projectID, taskID)
	if err != nil {
		return err
	}
	p, err := r.project(t.ProjectID)
	if err != nil {
		return err
	}
	c, err := p.Config(t.ProjectVersion)
	if err != nil {
		return err
	}
	slog.Debug("get a new task", "project_id", t.ProjectID, "task_id", t.ID)

	proof, err := r.handle(t, c.VMTypeID, c.Code, c.CodeExpParam)
	if err != nil {
		return err
	}

	tx, err := r.routerInstance.Route(
		&bind.TransactOpts{
			From: r.account,
			Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(t, r.signer, r.prv)
			},
		},
		new(big.Int).SetUint64(t.ProjectID),
		new(big.Int).SetUint64(1),
		t.DeviceID.String(),
		proof,
		taskID,
	)
	if err != nil {
		jsonErr := &struct {
			Code    int         `json:"code"`
			Message string      `json:"message"`
			Data    interface{} `json:"data,omitempty"`
		}{}
		je, nerr := json.Marshal(err)
		if nerr != nil {
			return errors.Wrap(err, "failed to marshal send tx error")
		}
		if err := json.Unmarshal(je, jsonErr); err != nil {
			return errors.Wrap(err, "failed to unmarshal send tx error")
		}
		return errors.Errorf("failed to send tx, error_code: %v, error_message: %v, error_data: %v", jsonErr.Code, jsonErr.Message, jsonErr.Data)
	}
	slog.Info("send tx to router contract success", "hash", tx.Hash().String())
	return nil
}

func (r *processor) run() {
	for {
		projectID, taskID, err := r.db.UnprocessedTask()
		if err != nil {
			slog.Error("failed to get unprocessed task", "error", err)
			time.Sleep(r.waitingTime)
			continue
		}
		if projectID == 0 {
			time.Sleep(r.waitingTime)
			continue
		}
		err = r.process(projectID, taskID)
		if err != nil {
			slog.Error("failed to process task", "error", err)
		}
		if err := r.db.ProcessTask(projectID, taskID, err); err != nil {
			slog.Error("failed to process db task", "error", err)
		}
	}
}

func Run(handle HandleTask, project Project, db DB, retrieve RetrieveTask, prv *ecdsa.PrivateKey, chainEndpoint string, routerAddr common.Address) error {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return errors.Wrap(err, "failed to dial chain endpoint")
	}
	routerInstance, err := router.NewRouter(routerAddr, client)
	if err != nil {
		return errors.Wrap(err, "failed to new router contract instance")
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get chain id")
	}
	p := &processor{
		db:             db,
		retrieve:       retrieve,
		handle:         handle,
		project:        project,
		prv:            prv,
		waitingTime:    3 * time.Second,
		signer:         types.NewLondonSigner(chainID),
		account:        crypto.PubkeyToAddress(prv.PublicKey),
		client:         client,
		routerInstance: routerInstance,
	}
	go p.run()
	return nil
}
