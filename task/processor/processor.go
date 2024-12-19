package processor

import (
	"bytes"
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

type HandleTask func(task *task.Task, projectConfig *project.Config) ([]byte, error)
type Project func(projectID *big.Int) (*project.Project, error)
type RetrieveTask func(taskIDs []common.Hash) ([]*task.Task, error)

type DB interface {
	UnprocessedTask() (common.Hash, error)
	ProcessTask(common.Hash, error) error
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

func (r *processor) process(taskID common.Hash) error {
	ts, err := r.retrieve([]common.Hash{taskID})
	if err != nil {
		return err
	}
	t := ts[0]
	p, err := r.project(t.ProjectID)
	if err != nil {
		return err
	}
	c, err := p.Config(t.ProjectVersion)
	if err != nil {
		return err
	}
	slog.Info("process task", "project_id", t.ProjectID.String(), "task_id", t.ID, "vm_type", c.VMTypeID)
	startTime := time.Now()
	proof, err := r.handle(t, c)
	if err != nil {
		metrics.FailedTaskNumMtc.WithLabelValues(t.ProjectID.String()).Inc()
		slog.Error("failed to handle task", "error", err)
		return err
	}
	processTime := time.Since(startTime)
	slog.Info("process task success", "project_id", t.ProjectID.String(), "task_id", t.ID, "process_time", processTime)
	metrics.TaskDurationMtc.WithLabelValues(t.ProjectID.String(), t.ProjectVersion, t.ID.String()).Set(processTime.Seconds())

	pubkey, err := crypto.UnmarshalPubkey(t.DevicePubKey)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal public key")
	}
	tx, err := r.routerInstance.Route(
		&bind.TransactOpts{
			From: r.account,
			Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(t, r.signer, r.prv)
			},
		},
		t.ProjectID,
		t.ID,
		r.account,
		crypto.PubkeyToAddress(*pubkey),
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
		taskID, err := r.db.UnprocessedTask()
		if err != nil {
			slog.Error("failed to get unprocessed task", "error", err)
			time.Sleep(r.waitingTime)
			continue
		}
		if bytes.Equal(taskID.Bytes(), common.Hash{}.Bytes()) {
			time.Sleep(r.waitingTime)
			continue
		}
		err = r.process(taskID)
		if err != nil {
			slog.Error("failed to process task", "error", err)
		}
		if err := r.db.ProcessTask(taskID, err); err != nil {
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
