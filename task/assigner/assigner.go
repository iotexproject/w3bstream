package assigner

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
	"golang.org/x/exp/rand"

	"github.com/iotexproject/w3bstream/smartcontracts/go/minter"
	"github.com/iotexproject/w3bstream/task"
)

type RetrieveTask func(projectID uint64, taskID common.Hash) (*task.Task, error)

type DB interface {
	UnassignedTask() (uint64, common.Hash, error)
	AssignTask(projectID uint64, taskID common.Hash, prover common.Address) error
	Provers() ([]common.Address, error)
}

type assigner struct {
	prv            *ecdsa.PrivateKey
	waitingTime    time.Duration
	signer         types.Signer
	account        common.Address
	client         *ethclient.Client
	db             DB
	retrieve       RetrieveTask
	minterInstance *minter.Minter
}

func (r *assigner) assign(projectID uint64, taskID common.Hash) error {
	t, err := r.retrieve(projectID, taskID)
	if err != nil {
		return err
	}
	provers, err := r.db.Provers()
	if err != nil {
		return errors.Wrap(err, "failed to get provers")
	}
	if len(provers) == 0 {
		return errors.New("no available prover")
	}
	prover := provers[rand.Intn(len(provers))]

	tx, err := r.minterInstance.Mint(
		&bind.TransactOpts{
			From: r.account,
			Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(t, r.signer, r.prv)
			},
		},
		minter.Sequencer{
			Addr:        r.account,
			Operator:    r.account,
			Beneficiary: r.account,
		},
		[]minter.TaskAssignment{
			{
				ProjectId: new(big.Int).SetUint64(projectID),
				TaskId:    taskID,
				Prover:    prover,
				Hash:      common.Hash{},
				Signature: t.Signature,
			},
		},
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
	slog.Info("send tx to minter contract success", "hash", tx.Hash().String())
	if err := r.db.AssignTask(projectID, taskID, prover); err != nil {
		return err
	}
	return nil
}

func (r *assigner) run() {
	for {
		projectID, taskID, err := r.db.UnassignedTask()
		if err != nil {
			slog.Error("failed to get unassigned task", "error", err)
			time.Sleep(r.waitingTime)
			continue
		}
		if projectID == 0 {
			time.Sleep(r.waitingTime)
			continue
		}
		if err := r.assign(projectID, taskID); err != nil {
			slog.Error("failed to assign task", "error", err)
			time.Sleep(r.waitingTime)
			continue
		}
		time.Sleep(1 * time.Second) // TODO after assign tx success, then get next task
	}
}

func Run(db DB, prv *ecdsa.PrivateKey, chainEndpoint string, retrieve RetrieveTask, minterAddr common.Address) error {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return errors.Wrap(err, "failed to dial chain endpoint")
	}
	minterInstance, err := minter.NewMinter(minterAddr, client)
	if err != nil {
		return errors.Wrap(err, "failed to new minter contract instance")
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get chain id")
	}
	p := &assigner{
		db:             db,
		prv:            prv,
		waitingTime:    3 * time.Second,
		signer:         types.NewLondonSigner(chainID),
		account:        crypto.PubkeyToAddress(prv.PublicKey),
		client:         client,
		retrieve:       retrieve,
		minterInstance: minterInstance,
	}
	go p.run()
	return nil
}
