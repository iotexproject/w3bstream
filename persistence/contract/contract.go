package contract

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"log/slog"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/smartcontracts/go/project"
	"github.com/machinefi/sprout/smartcontracts/go/prover"
)

const (
	chainHead   = "chain_head"
	blockPrefix = "block_"
)

var (
	allTopicHash = []common.Hash{
		attributeSetTopicHash,
		projectPausedTopicHash,
		projectResumedTopicHash,
		projectConfigUpdatedTopicHash,

		operatorSetTopicHash,
		nodeTypeUpdatedTopicHash,
		proverPausedTopicHash,
		proverResumedTopicHash,
	}
)

type Contract struct {
	db                     *pebble.DB
	size                   uint64
	beginningBlockNumber   uint64
	proverContractAddr     common.Address
	projectContractAddr    common.Address
	chainHeadNotifications []chan<- uint64
	projectNotifications   []chan<- uint64
	listStepSize           uint64
	watchInterval          time.Duration
	client                 *ethclient.Client
	proverInstance         *prover.Prover
	projectInstance        *project.Project
}

type block struct {
	blockProject
	blockProver
}

func (c *Contract) Project(projectID, blockNumber uint64) *Project {
	dataBytes, closer, err := c.db.Get(c.dbKey(blockNumber))
	if err != nil {
		if err != pebble.ErrNotFound {
			slog.Error("failed to get db block data", "block_number", blockNumber, "error", err)
		}
		return nil
	}
	dst := make([]byte, len(dataBytes))
	copy(dst, dataBytes)
	blockData := &block{}
	if err := json.Unmarshal(dst, blockData); err != nil {
		slog.Error("failed to unmarshal block data", "block_number", blockNumber, "error", err)
		return nil
	}
	if err := closer.Close(); err != nil {
		slog.Error("failed to close result of block data", "block_number", blockNumber, "error", err)
		return nil
	}
	return blockData.blockProject.Projects[projectID]
}

func (c *Contract) LatestProject(projectID uint64) *Project {
	bp := c.latestProjects()
	if bp == nil {
		return nil
	}
	return bp.Projects[projectID]
}

func (c *Contract) LatestProjects() []*Project {
	bp := c.latestProjects()
	if bp == nil {
		return nil
	}
	ps := []*Project{}
	for _, p := range bp.Projects {
		ps = append(ps, p)
	}
	return ps
}

func (c *Contract) latestProjects() *blockProject {
	batch := c.db.NewBatch()
	defer batch.Close()

	headBytes, closer, err := batch.Get(c.dbHead())
	if err != nil {
		slog.Error("failed to get chain head data", "error", err)
		return nil
	}
	dst := make([]byte, len(headBytes))
	copy(dst, headBytes)
	head := binary.LittleEndian.Uint64(dst)
	if err := closer.Close(); err != nil {
		slog.Error("failed to close result of chain head", "error", err)
		return nil
	}

	dataBytes, closer, err := batch.Get(c.dbKey(head))
	if err != nil {
		if err != pebble.ErrNotFound {
			slog.Error("failed to get db block data", "block_number", head, "error", err)
		}
		return nil
	}
	dst = make([]byte, len(dataBytes))
	copy(dst, dataBytes)
	blockData := &block{}
	if err := json.Unmarshal(dst, blockData); err != nil {
		slog.Error("failed to unmarshal block data", "block_number", head, "error", err)
		return nil
	}
	if err := closer.Close(); err != nil {
		slog.Error("failed to close result of block data", "block_number", head, "error", err)
		return nil
	}

	if err := batch.Commit(pebble.Sync); err != nil {
		slog.Error("failed to commit batch", "error", err)
		return nil
	}

	return &blockData.blockProject
}

func (c *Contract) Provers(blockNumber uint64) []*Prover {
	dataBytes, closer, err := c.db.Get(c.dbKey(blockNumber))
	if err != nil {
		if err != pebble.ErrNotFound {
			slog.Error("failed to get db block data", "block_number", blockNumber, "error", err)
		}
		return nil
	}
	dst := make([]byte, len(dataBytes))
	copy(dst, dataBytes)
	blockData := &block{}
	if err := json.Unmarshal(dst, blockData); err != nil {
		slog.Error("failed to unmarshal block data", "block_number", blockNumber, "error", err)
		return nil
	}
	if err := closer.Close(); err != nil {
		slog.Error("failed to close result of block data", "block_number", blockNumber, "error", err)
		return nil
	}

	ps := []*Prover{}
	for _, p := range blockData.blockProver.Provers {
		ps = append(ps, p)
	}
	return ps
}

func (c *Contract) LatestProvers() []*Prover {
	bp := c.latestProvers()
	if bp == nil {
		return nil
	}
	ps := []*Prover{}
	for _, p := range bp.Provers {
		ps = append(ps, p)
	}
	return ps
}

func (c *Contract) Prover(operator common.Address) *Prover {
	bp := c.latestProvers()
	if bp == nil {
		return nil
	}
	for _, p := range bp.Provers {
		if p.OperatorAddress == operator {
			return p
		}
	}
	return nil
}

func (c *Contract) latestProvers() *blockProver {
	batch := c.db.NewBatch()
	defer batch.Close()

	headBytes, closer, err := batch.Get(c.dbHead())
	if err != nil {
		slog.Error("failed to get chain head data", "error", err)
		return nil
	}
	dst := make([]byte, len(headBytes))
	copy(dst, headBytes)
	head := binary.LittleEndian.Uint64(dst)
	if err := closer.Close(); err != nil {
		slog.Error("failed to close result of chain head", "error", err)
		return nil
	}

	dataBytes, closer, err := batch.Get(c.dbKey(head))
	if err != nil {
		if err != pebble.ErrNotFound {
			slog.Error("failed to get db block data", "block_number", head, "error", err)
		}
		return nil
	}
	dst = make([]byte, len(dataBytes))
	copy(dst, dataBytes)
	blockData := &block{}
	if err := json.Unmarshal(dst, blockData); err != nil {
		slog.Error("failed to unmarshal block data", "block_number", head, "error", err)
		return nil
	}
	if err := closer.Close(); err != nil {
		slog.Error("failed to close result of block data", "block_number", head, "error", err)
		return nil
	}

	if err := batch.Commit(pebble.Sync); err != nil {
		slog.Error("failed to commit batch", "error", err)
		return nil
	}

	return &blockData.blockProver
}

func (c *Contract) notifyProject(bp *blockProjectDiff) {
	for _, p := range bp.diffs {
		for _, n := range c.projectNotifications {
			n <- p.id
		}
	}
}

func (c *Contract) notifyChainHead(chainHead uint64) {
	for _, n := range c.chainHeadNotifications {
		n <- chainHead
	}
}

func (c *Contract) dbKey(blockNumber uint64) []byte {
	return []byte(blockPrefix + strconv.FormatUint(blockNumber, 10))
}

func (c *Contract) dbHead() []byte {
	return []byte(chainHead)
}

func (c *Contract) updateDB(blockNumber uint64, projectDiff *blockProjectDiff, proverDiff *blockProverDiff) error {
	batch := c.db.NewBatch()
	defer batch.Close()

	preBlock := blockNumber - 1
	if blockNumber == 0 {
		preBlock = blockNumber
	}

	preBlockBytes, closer, err := batch.Get(c.dbKey(preBlock))
	if err != nil && err != pebble.ErrNotFound {
		return errors.Wrap(err, "failed to get pre block data")
	}
	var preBlockData *block
	if err == nil {
		dst := make([]byte, len(preBlockBytes))
		copy(dst, preBlockBytes)
		if err := json.Unmarshal(dst, preBlockData); err != nil {
			return errors.Wrap(err, "failed to unmarshal pre block data")
		}
		if err := closer.Close(); err != nil {
			return errors.Wrap(err, "failed to close result of pre block data")
		}
	} else {
		preBlockData = &block{
			blockProject: blockProject{
				Projects: map[uint64]*Project{},
			},
			blockProver: blockProver{
				Provers: map[uint64]*Prover{},
			},
		}
	}
	if projectDiff != nil {
		preBlockData.blockProject.merge(projectDiff)
	}
	if proverDiff != nil {
		preBlockData.blockProver.merge(proverDiff)
	}
	currBlockBytes, err := json.Marshal(preBlockData)
	if err != nil {
		return errors.Wrap(err, "failed to marshal block data")
	}

	numberBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(numberBytes, blockNumber)

	if err := batch.Set(c.dbHead(), numberBytes, nil); err != nil {
		return errors.Wrap(err, "failed to set chain head")
	}
	if err := batch.Set(c.dbKey(blockNumber), currBlockBytes, nil); err != nil {
		return errors.Wrap(err, "failed to set current block")
	}
	if blockNumber-c.beginningBlockNumber+1 > c.size {
		if err := batch.Delete(c.dbKey(blockNumber-c.size), nil); err != nil {
			return errors.Wrap(err, "failed to delete expired block")
		}
	}

	if err := batch.Commit(pebble.Sync); err != nil {
		return errors.Wrap(err, "failed to commit batch")
	}
	return nil
}

func (c *Contract) processLogs(from, to uint64, logs []types.Log, notify bool) error {
	sort.Slice(logs, func(i, j int) bool {
		if logs[i].BlockNumber != logs[j].BlockNumber {
			return logs[i].BlockNumber < logs[j].BlockNumber
		}
		return logs[i].TxIndex < logs[j].TxIndex
	})

	projectMap, err := c.processProjectLogs(logs)
	if err != nil {
		return err
	}
	proverMap, err := c.processProverLogs(logs)
	if err != nil {
		return err
	}
	for blockNumber := from; blockNumber <= to; blockNumber++ {
		projects, ok := projectMap[blockNumber]
		if ok && notify {
			c.notifyProject(projects)
		}
		if err := c.updateDB(blockNumber, projects, proverMap[blockNumber]); err != nil {
			return err
		}
	}
	return nil
}

func (c *Contract) list() (uint64, error) {
	headBytes, closer, err := c.db.Get(c.dbHead())
	if err != nil && err != pebble.ErrNotFound {
		return 0, errors.Wrap(err, "failed to get db chain head")
	}
	head := c.beginningBlockNumber
	if err == nil {
		dst := make([]byte, len(headBytes))
		copy(dst, headBytes)
		head = binary.LittleEndian.Uint64(dst)
		head++
		if err := closer.Close(); err != nil {
			return 0, errors.Wrap(err, "failed to close result of chain head")
		}
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.proverContractAddr, c.projectContractAddr},
		Topics:    [][]common.Hash{allTopicHash},
	}
	from := head
	to := from
	for {
		header, err := c.client.HeaderByNumber(context.Background(), nil)
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
		query.FromBlock = new(big.Int).SetUint64(from)
		query.ToBlock = new(big.Int).SetUint64(to)
		logs, err := c.client.FilterLogs(context.Background(), query)
		if err != nil {
			return 0, errors.Wrap(err, "failed to filter contract logs")
		}
		if err := c.processLogs(from, to, logs, false); err != nil {
			return 0, err
		}
		from = to + 1
	}
	return to, nil
}

func (c *Contract) watch(listedBlockNumber uint64) {
	queriedBlockNumber := listedBlockNumber
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.proverContractAddr, c.projectContractAddr},
		Topics:    [][]common.Hash{allTopicHash},
	}
	ticker := time.NewTicker(c.watchInterval)

	go func() {
		for range ticker.C {
			target := queriedBlockNumber + 1

			query.FromBlock = new(big.Int).SetUint64(target)
			query.ToBlock = new(big.Int).SetUint64(target)
			logs, err := c.client.FilterLogs(context.Background(), query)
			if err != nil {
				if !strings.Contains(err.Error(), "start block > tip height") {
					slog.Error("failed to filter contract logs", "error", err)
				}
				continue
			}

			if err := c.processLogs(target, target, logs, true); err != nil {
				slog.Error("failed to process logs", "error", err)
				continue
			}

			c.notifyChainHead(target)

			queriedBlockNumber = target
		}
	}()
}

func (c *Contract) Release() {
	if err := c.db.Close(); err != nil {
		slog.Error("failed to close pebble db", "error", err)
	}
}

func New(size, beginningBlockNumber uint64, localDBDir, chainEndpoint string, proverContractAddr, projectContractAddr common.Address, chainHeadNotifications []chan<- uint64, projectNotifications []chan<- uint64) (*Contract, error) {
	db, err := pebble.Open(localDBDir, &pebble.Options{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to open pebble db")
	}

	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial chain endpoint")
	}
	projectInstance, err := project.NewProject(projectContractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new project contract instance")
	}
	proverInstance, err := prover.NewProver(proverContractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new prover contract instance")
	}

	c := &Contract{
		db:                     db,
		size:                   size,
		beginningBlockNumber:   beginningBlockNumber,
		proverContractAddr:     proverContractAddr,
		projectContractAddr:    projectContractAddr,
		chainHeadNotifications: chainHeadNotifications,
		projectNotifications:   projectNotifications,
		listStepSize:           1000,
		watchInterval:          1 * time.Second,
		client:                 client,
		proverInstance:         proverInstance,
		projectInstance:        projectInstance,
	}

	listedBlockNumber, err := c.list()
	if err != nil {
		return nil, err
	}
	go c.watch(listedBlockNumber)

	return c, nil
}
