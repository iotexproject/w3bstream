package persistence

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"log/slog"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/smartcontracts/go/ioidregistry"
	"github.com/machinefi/sprout/smartcontracts/go/projectclient"
)

var (
	approveTopic   = crypto.Keccak256Hash([]byte("Approve(uint256,address)"))
	unapproveTopic = crypto.Keccak256Hash([]byte("Unapprove(uint256,address)"))

	newDeviceTopic    = crypto.Keccak256Hash([]byte("NewDevice(address,address,bytes32,string)"))
	updateDeviceTopic = crypto.Keccak256Hash([]byte("UpdateDevice(address,address,bytes32,string)"))
)

type Contract struct {
	db                   *pebble.DB
	beginningBlockNumber uint64
	clientContractAddr   common.Address
	didContractAddr      common.Address
	listStepSize         uint64
	watchInterval        time.Duration
	client               *ethclient.Client
	clientInstance       *projectclient.Projectclient
	didInstance          *ioidregistry.Ioidregistry
}

type didDOC struct {
	URI  string
	Hash common.Hash
}

type contractData struct {
	ProjectClients map[uint64]map[common.Address]bool
	DIDDoc         map[common.Address]didDOC
}

func (c *Contract) IsApproved(projectID uint64, clientAddr common.Address) (bool, error) {
	dataBytes, closer, err := c.db.Get(c.dbKey())
	if err != nil {
		if err != pebble.ErrNotFound {
			return false, errors.Wrap(err, "failed to get db contract data")
		}
		return false, nil
	}
	dst := make([]byte, len(dataBytes))
	copy(dst, dataBytes)
	blockData := &contractData{ProjectClients: map[uint64]map[common.Address]bool{}}
	if err := json.Unmarshal(dst, blockData); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal contract data")
	}
	if err := closer.Close(); err != nil {
		return false, errors.Wrap(err, "failed to close result of contract data")
	}

	cs, ok := blockData.ProjectClients[projectID]
	if !ok {
		return false, nil
	}
	return cs[clientAddr], nil
}

func (c *Contract) dbHead() []byte {
	return []byte("chain_head")
}

func (c *Contract) dbKey() []byte {
	return []byte("contract_data")
}

func (c *Contract) processLogs(blockNumber uint64, logs []types.Log) error {
	sort.Slice(logs, func(i, j int) bool {
		if logs[i].BlockNumber != logs[j].BlockNumber {
			return logs[i].BlockNumber < logs[j].BlockNumber
		}
		return logs[i].TxIndex < logs[j].TxIndex
	})

	batch := c.db.NewIndexedBatch()
	defer batch.Close()

	dataBytes, closer, err := batch.Get(c.dbKey())
	if err != nil && err != pebble.ErrNotFound {
		return errors.Wrap(err, "failed to get local db contract data")
	}
	data := &contractData{
		ProjectClients: map[uint64]map[common.Address]bool{},
		DIDDoc:         map[common.Address]didDOC{},
	}
	if err == nil {
		dst := make([]byte, len(dataBytes))
		copy(dst, dataBytes)
		if err := json.Unmarshal(dst, data); err != nil {
			return errors.Wrap(err, "failed to unmarshal contract data")
		}
		if err := closer.Close(); err != nil {
			return errors.Wrap(err, "failed to close result of contract data")
		}
	}
	for _, l := range logs {
		switch l.Topics[0] {
		case approveTopic:
			e, err := c.clientInstance.ParseApprove(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse project client approve event")
			}

			cs, ok := data.ProjectClients[e.ProjectId.Uint64()]
			if !ok {
				cs = map[common.Address]bool{}
			}
			cs[e.Device] = true

		case unapproveTopic:
			e, err := c.clientInstance.ParseUnapprove(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse project client unapprove event")
			}

			cs, ok := data.ProjectClients[e.ProjectId.Uint64()]
			if !ok {
				continue
			}
			delete(cs, e.Device)

		case newDeviceTopic:
			e, err := c.didInstance.ParseNewDevice(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse new device event")
			}

			data.DIDDoc[e.Device] = didDOC{
				URI:  e.Uri,
				Hash: e.Hash,
			}

		case updateDeviceTopic:
			e, err := c.didInstance.ParseUpdateDevice(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse update device event")
			}

			data.DIDDoc[e.Device] = didDOC{
				URI:  e.Uri,
				Hash: e.Hash,
			}
		}
	}

	currData, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "failed to marshal contract data")
	}

	numberBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(numberBytes, blockNumber)

	if err := batch.Set(c.dbHead(), numberBytes, nil); err != nil {
		return errors.Wrap(err, "failed to set chain head")
	}
	if err := batch.Set(c.dbKey(), currData, nil); err != nil {
		return errors.Wrap(err, "failed to set contract data")
	}

	if err := batch.Commit(pebble.Sync); err != nil {
		return errors.Wrap(err, "failed to commit batch")
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
		Addresses: []common.Address{c.clientContractAddr, c.didContractAddr},
		Topics: [][]common.Hash{{
			approveTopic,
			unapproveTopic,

			newDeviceTopic,
			updateDeviceTopic,
		}},
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
		if err := c.processLogs(to, logs); err != nil {
			return 0, err
		}
		from = to + 1
	}
	return to, nil
}

func (c *Contract) watch(listedBlockNumber uint64) {
	queriedBlockNumber := listedBlockNumber
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.clientContractAddr, c.didContractAddr},
		Topics: [][]common.Hash{{
			approveTopic,
			unapproveTopic,

			newDeviceTopic,
			updateDeviceTopic,
		}},
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

			if err := c.processLogs(target, logs); err != nil {
				slog.Error("failed to process logs", "error", err)
				continue
			}

			queriedBlockNumber = target
		}
	}()
}

func New(db *pebble.DB, beginningBlockNumber uint64, chainEndpoint string, clientContractAddr, didContractAddr common.Address) (*Contract, error) {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial chain endpoint")
	}
	clientInstance, err := projectclient.NewProjectclient(clientContractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new project client contract instance")
	}
	didInstance, err := ioidregistry.NewIoidregistry(didContractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new ioid registry contract instance")
	}

	c := &Contract{
		db:                   db,
		beginningBlockNumber: beginningBlockNumber,
		clientContractAddr:   clientContractAddr,
		didContractAddr:      didContractAddr,
		listStepSize:         10000,
		watchInterval:        1 * time.Second,
		client:               client,
		clientInstance:       clientInstance,
		didInstance:          didInstance,
	}

	listedBlockNumber, err := c.list()
	if err != nil {
		return nil, err
	}
	go c.watch(listedBlockNumber)

	return c, nil
}
