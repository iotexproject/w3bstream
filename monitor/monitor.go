package monitor

import (
	"bytes"
	"context"
	"log/slog"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/smartcontracts/go/dao"
	"github.com/iotexproject/w3bstream/smartcontracts/go/minter"
)

type (
	ScannedBlockNumber       func() (uint64, error)
	UpsertScannedBlockNumber func(uint64) error
	UpsertNBits              func(uint32) error
	UpsertBlockHead          func(uint64, common.Hash) error
	DeleteTask               func(common.Hash, uint64) error
)

type Handler struct {
	ScannedBlockNumber
	UpsertScannedBlockNumber
	UpsertNBits
	UpsertBlockHead
	DeleteTask
}

type ContractAddr struct {
	Prover  common.Address
	Project common.Address
	Dao     common.Address
	Minter  common.Address
}

type Contract struct {
	h                    *Handler
	addr                 *ContractAddr
	beginningBlockNumber uint64
	listStepSize         uint64
	watchInterval        time.Duration
	client               *ethclient.Client
	daoInstance          *dao.Dao
	minterInstance       *minter.Minter
}

var (
	blockAddedTopic   = crypto.Keccak256Hash([]byte("BlockAdded(uint256,bytes32,uint256)"))
	nbitsSetTopic     = crypto.Keccak256Hash([]byte("NBitsSet(uint32)"))
	taskAssignedTopic = crypto.Keccak256Hash([]byte("TaskAssigned(uint256,bytes32,address,uint256)"))
	taskSettledTopic  = crypto.Keccak256Hash([]byte("TaskSettled(uint256,bytes32,address)"))
)

var allTopic = []common.Hash{
	blockAddedTopic,
	nbitsSetTopic,
	taskAssignedTopic,
	taskSettledTopic,
}

func (a *ContractAddr) all() []common.Address {
	all := make([]common.Address, 0, 4)
	zero := common.Address{}
	if !bytes.Equal(a.Dao[:], zero[:]) {
		all = append(all, a.Dao)
	}
	if !bytes.Equal(a.Minter[:], zero[:]) {
		all = append(all, a.Minter)
	}
	if !bytes.Equal(a.Project[:], zero[:]) {
		all = append(all, a.Project)
	}
	if !bytes.Equal(a.Prover[:], zero[:]) {
		all = append(all, a.Prover)
	}
	return all
}

func (c *Contract) processLogs(logs []types.Log) error {
	sort.Slice(logs, func(i, j int) bool {
		if logs[i].BlockNumber != logs[j].BlockNumber {
			return logs[i].BlockNumber < logs[j].BlockNumber
		}
		return logs[i].TxIndex < logs[j].TxIndex
	})

	for _, l := range logs {
		switch l.Topics[0] {
		case blockAddedTopic:
			e, err := c.daoInstance.ParseBlockAdded(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse block added event")
			}
			if err := c.h.UpsertBlockHead(e.Num.Uint64(), e.Hash); err != nil {
				return err
			}
		case nbitsSetTopic:
			e, err := c.minterInstance.ParseNBitsSet(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse nbits set event")
			}
			if err := c.h.UpsertNBits(e.Nbits); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Contract) list() (uint64, error) {
	head := c.beginningBlockNumber
	h, err := c.h.ScannedBlockNumber()
	if err != nil {
		return 0, err
	}
	head = max(head, h)

	query := ethereum.FilterQuery{
		Addresses: c.addr.all(),
		Topics:    [][]common.Hash{allTopic},
	}
	ctx := context.Background()
	from := head + 1
	to := from
	for {
		header, err := c.client.HeaderByNumber(ctx, nil)
		if err != nil {
			return 0, errors.Wrap(err, "failed to retrieve latest block header")
		}
		currentHead := header.Number.Uint64()
		to = from + c.listStepSize
		if to > currentHead {
			to = currentHead
		}
		if from > to {
			break
		}
		slog.Debug("listing chain", "from", from, "to", to)
		query.FromBlock = new(big.Int).SetUint64(from)
		query.ToBlock = new(big.Int).SetUint64(to)
		logs, err := c.client.FilterLogs(ctx, query)
		if err != nil {
			return 0, errors.Wrap(err, "failed to filter contract logs")
		}
		if err := c.processLogs(logs); err != nil {
			return 0, err
		}
		if err := c.h.UpsertScannedBlockNumber(to); err != nil {
			return 0, err
		}
		from = to + 1
	}
	slog.Info("contract data synchronization completed", "current_height", to)
	return to, nil
}

func (c *Contract) watch(listedBlockNumber uint64) {
	scannedBlockNumber := listedBlockNumber
	query := ethereum.FilterQuery{
		Addresses: c.addr.all(),
		Topics:    [][]common.Hash{allTopic},
	}
	ticker := time.NewTicker(c.watchInterval)

	go func() {
		for range ticker.C {
			target := scannedBlockNumber + 1

			query.FromBlock = new(big.Int).SetUint64(target)
			query.ToBlock = new(big.Int).SetUint64(target)
			logs, err := c.client.FilterLogs(context.Background(), query)
			if err != nil {
				if !strings.Contains(err.Error(), "start block > tip height") {
					slog.Error("failed to filter contract logs", "error", err)
				}
				continue
			}
			if err := c.processLogs(logs); err != nil {
				slog.Error("failed to process logs", "error", err)
				continue
			}
			if err := c.h.UpsertScannedBlockNumber(target); err != nil {
				slog.Error("failed to upsert scanned block number", "error", err)
				continue
			}
			scannedBlockNumber = target
		}
	}()
}

func New(h *Handler, addr *ContractAddr, beginningBlockNumber uint64, chainEndpoint string) (*Contract, error) {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial chain endpoint")
	}
	daoInstance, err := dao.NewDao(addr.Dao, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new project contract instance")
	}
	minterInstance, err := minter.NewMinter(addr.Minter, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new prover contract instance")
	}

	c := &Contract{
		h:                    h,
		addr:                 addr,
		beginningBlockNumber: beginningBlockNumber,
		listStepSize:         500,
		watchInterval:        1 * time.Second,
		client:               client,
		daoInstance:          daoInstance,
		minterInstance:       minterInstance,
	}

	listedBlockNumber, err := c.list()
	if err != nil {
		return nil, err
	}
	go c.watch(listedBlockNumber)

	return c, nil
}
