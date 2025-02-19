package processor

import (
	"context"
	"crypto/ecdsa"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/metrics"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/smartcontracts/go/router"
	"github.com/iotexproject/w3bstream/task"
)

type HandleTasks func(tasks []*task.Task, projectConfig *project.Config) ([]byte, error)
type Project func(projectID string) (*project.Project, error)
type RetrieveTask func(taskIDs []common.Hash) ([]*task.Task, error)

type DB interface {
	UnprocessedTasks(projectID string, limit uint64) ([]common.Hash, error)
	UnprocessedProjects() (map[string]uint64, error)
	ProcessTasks(taskIDs []common.Hash, err error) error
}

type processor struct {
	db             DB
	retrieve       RetrieveTask
	handle         HandleTasks
	project        Project
	prv            *ecdsa.PrivateKey
	waitingTime    time.Duration
	signer         types.Signer
	account        common.Address
	client         *ethclient.Client
	routerInstance *router.Router
}

func (r *processor) process(taskIDs []common.Hash, c *project.Config, pid string) error {
	ts, err := r.retrieve(taskIDs)
	if err != nil {
		return err
	}

	slog.Info("process tasks", "project_id", pid, "vm_type", c.VMTypeID)
	startTime := time.Now()
	proof, err := r.handle(ts, c)
	if err != nil {
		metrics.FailedTaskNumMtc.WithLabelValues(pid).Inc()
		slog.Error("failed to handle task", "error", err)
		return err
	}
	processTime := time.Since(startTime)
	slog.Info("process task success", "project_id", pid, "process_time", processTime)
	//metrics.TaskDurationMtc.WithLabelValues(pid, t.ProjectVersion, t.ID.String()).Set(processTime.Seconds())

	tids := [][32]byte{}
	for _, t := range ts {
		tids = append(tids, t.ID)
	}
	pidInt, ok := new(big.Int).SetString(pid, 10)
	if !ok {
		return errors.New("failed to decode project id string")
	}
	tx, err := r.routerInstance.Route(
		&bind.TransactOpts{
			From: r.account,
			Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(t, r.signer, r.prv)
			},
		},
		r.account,
		pidInt,
		tids,
		proof,
	)
	if err != nil {
		if jsonErr, ok := err.(rpc.DataError); ok {
			errData := jsonErr.ErrorData()
			errMsg := jsonErr.Error()
			errCode := err.(rpc.Error).ErrorCode()
			return errors.Wrapf(err, "failed to send tx to router contract, errData: %v, errMsg: %s, errCode: %d", errData, errMsg, errCode)
		}
		return errors.Wrap(err, "failed to send tx to router contract")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, r.client, tx)
	if err != nil {
		return errors.Wrap(err, "failed to wait tx mined")
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		slog.Error("tx failed", "tx_hash", tx.Hash().String())
		return errors.New("tx failed")
	}
	slog.Info("send tx to router contract success", "hash", tx.Hash().String())
	return nil
}

func (r *processor) run() {
	for {
		ps, err := r.db.UnprocessedProjects()
		if err != nil {
			slog.Error("failed to get unprocessed projects", "error", err)
			time.Sleep(r.waitingTime)
			continue
		}
		if len(ps) == 0 {
			time.Sleep(r.waitingTime)
			continue
		}
		for pid, n := range ps {
			p, err := r.project(pid)
			if err != nil {
				slog.Error("failed to get project file", "project_id", pid, "error", err)
				continue
			}
			c, err := p.DefaultConfig()
			if err != nil {
				slog.Error("failed to get project config", "project_id", pid, "error", err)
				continue
			}
			batch := uint64(1)
			if c.TaskProcessingBatch > 0 {
				batch = c.TaskProcessingBatch
			}
			if n < batch {
				slog.Info("the project currently doesn't have enough tasks", "project_id", pid, "task_processing_batch", batch, "current_number", n)
				time.Sleep(r.waitingTime)
				continue
			}
			taskIDs, err := r.db.UnprocessedTasks(pid, batch)
			if err != nil {
				slog.Error("failed to get tasks", "project_id", pid, "error", err)
				continue
			}
			err = r.process(taskIDs, c, pid)
			if err != nil {
				slog.Error("failed to process task", "error", err)
			}
			if err := r.db.ProcessTasks(taskIDs, err); err != nil {
				slog.Error("failed to process db tasks", "error", err)
			}
		}
	}
}

func Run(handle HandleTasks, project Project, db DB, retrieve RetrieveTask, prv *ecdsa.PrivateKey, chainEndpoint string, routerAddr common.Address) error {
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
