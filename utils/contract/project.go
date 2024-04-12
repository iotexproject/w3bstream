package contract

import (
	"bytes"
	"context"
	"log/slog"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/smartcontracts/go/project"
)

var (
	RequiredProverAmountHash = crypto.Keccak256Hash([]byte("RequiredProverAmount"))
	VmTypeHash               = crypto.Keccak256Hash([]byte("VmType"))

	attributeSetTopicHash         = crypto.Keccak256Hash([]byte("AttributeSet(uint256,bytes32,bytes)"))
	projectPausedTopicHash        = crypto.Keccak256Hash([]byte("ProjectPaused(uint256)"))
	projectResumedTopicHash       = crypto.Keccak256Hash([]byte("ProjectResumed(uint256)"))
	projectConfigUpdatedTopicHash = crypto.Keccak256Hash([]byte("ProjectConfigUpdated(uint256,string,bytes32)"))

	emptyHash = common.Hash{}
)

type Projects struct {
	BlockNumber uint64
	Projects    map[uint64]*Project
}

type Project struct {
	ID          uint64
	BlockNumber uint64
	Paused      *bool
	Uri         string
	Hash        common.Hash
	Attributes  map[common.Hash][]byte
}

func (ps *Projects) Merge(diff *Projects) {
	ps.BlockNumber = diff.BlockNumber
	for id, p := range ps.Projects {
		diffP, ok := diff.Projects[id]
		if ok {
			p.Merge(diffP)
		}
	}
	for id, p := range diff.Projects {
		if _, ok := ps.Projects[id]; !ok {
			ps.Projects[id] = p
		}
	}
}

func (p *Project) Merge(diff *Project) {
	if diff.ID != 0 {
		p.ID = diff.ID
	}
	if diff.BlockNumber != 0 {
		p.BlockNumber = diff.BlockNumber
	}
	if diff.Paused != nil {
		paused := *diff.Paused
		p.Paused = &paused
	}
	if diff.Uri != "" {
		p.Uri = diff.Uri
	}
	if !bytes.Equal(diff.Hash[:], emptyHash[:]) {
		copy(p.Hash[:], diff.Hash[:])
	}
	for h, d := range diff.Attributes {
		p.Attributes[h] = d
	}
}

func ListAndWatchProject(chainEndpoint, contractAddress string) (<-chan *Projects, error) {
	ch := make(chan *Projects, 10)
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial chain endpoint")
	}

	instance, err := project.NewProject(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new project contract instance")
	}

	latestBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to query the latest block number")
	}
	if err := listProject(ch, instance, latestBlockNumber); err != nil {
		return nil, err
	}

	topics := []common.Hash{attributeSetTopicHash, projectPausedTopicHash, projectResumedTopicHash, projectConfigUpdatedTopicHash}
	watchProject(ch, client, instance, 3*time.Second, contractAddress, topics, 1000, latestBlockNumber)

	return ch, nil
}

// TOD list determinate block number
func listProject(ch chan<- *Projects, instance *project.Project, targetBlockNumber uint64) error {
	ps := &Projects{
		Projects: map[uint64]*Project{},
	}
	for projectID := uint64(1); ; projectID++ {
		mp, err := instance.Config(nil, new(big.Int).SetUint64(projectID))
		if err != nil {
			return errors.Wrapf(err, "failed to get project meta from chain, project_id %v", projectID)
		}
		if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
			break
		}

		isPaused, err := instance.IsPaused(nil, new(big.Int).SetUint64(projectID))
		if err != nil {
			return errors.Wrapf(err, "failed to get project pause status from chain, project_id %v", projectID)
		}

		proverAmt, err := instance.Attributes(nil, new(big.Int).SetUint64(projectID), RequiredProverAmountHash)
		if err != nil {
			return errors.Wrapf(err, "failed to get project attributes from chain, project_id %v, key %s", projectID, proverAmt)
		}
		vmType, err := instance.Attributes(nil, new(big.Int).SetUint64(projectID), VmTypeHash)
		if err != nil {
			return errors.Wrapf(err, "failed to get project attributes from chain, project_id %v, key %s", projectID, vmType)
		}
		attributes := make(map[common.Hash][]byte)
		attributes[RequiredProverAmountHash] = proverAmt
		attributes[VmTypeHash] = vmType

		ps.Projects[projectID] = &Project{
			ID:          projectID,
			BlockNumber: targetBlockNumber,
			Paused:      &isPaused,
			Uri:         mp.Uri,
			Hash:        mp.Hash,
			Attributes:  attributes,
		}
	}
	ch <- ps
	return nil
}

func watchProject(ch chan<- *Projects, client *ethclient.Client, instance *project.Project, interval time.Duration, contractAddress string, topics []common.Hash, step, startBlockNumber uint64) {
	queriedBlockNumber := startBlockNumber
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics: [][]common.Hash{{
			(topics[0]),
			(topics[1]),
			(topics[2]),
			(topics[3]),
		}},
	}
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			from := queriedBlockNumber + 1
			to := from + step

			latestBlockNumber, err := client.BlockNumber(context.Background())
			if err != nil {
				slog.Error("failed to query the latest block number", "error", err)
				continue
			}
			if to > latestBlockNumber {
				to = latestBlockNumber
			}
			if from > to {
				continue
			}
			query.FromBlock = new(big.Int).SetUint64(from)
			query.ToBlock = new(big.Int).SetUint64(to)
			logs, err := client.FilterLogs(context.Background(), query)
			if err != nil {
				slog.Error("failed to filter contract logs", "error", err)
				continue
			}
			if processProjectLogs(ch, logs, instance) {
				queriedBlockNumber = to
			}
		}
	}()
}

func processProjectLogs(ch chan<- *Projects, logs []types.Log, instance *project.Project) bool {
	if len(logs) == 0 {
		return true
	}
	sort.Slice(logs, func(i, j int) bool {
		if logs[i].BlockNumber != logs[j].BlockNumber {
			return logs[i].BlockNumber < logs[j].BlockNumber
		}
		return logs[i].TxIndex < logs[j].TxIndex
	})
	psMap := map[uint64]*Projects{}

	for _, l := range logs {
		ps, ok := psMap[l.BlockNumber]
		if !ok {
			ps = &Projects{
				BlockNumber: l.BlockNumber,
				Projects:    map[uint64]*Project{},
			}
		}
		switch l.Topics[0] {
		case attributeSetTopicHash:
			e, err := instance.ParseAttributeSet(l)
			if err != nil {
				slog.Error("failed to parse project attribute set event", "error", err)
				return false
			}

			p, ok := ps.Projects[e.ProjectId.Uint64()]
			if !ok {
				p = &Project{Attributes: map[common.Hash][]byte{}}
			}
			p.Attributes[e.Key] = e.Value
			ps.Projects[e.ProjectId.Uint64()] = p

		case projectPausedTopicHash:
			e, err := instance.ParseProjectPaused(l)
			if err != nil {
				slog.Error("failed to parse project paused event", "error", err)
				return false
			}

			p, ok := ps.Projects[e.ProjectId.Uint64()]
			if !ok {
				p = &Project{Attributes: map[common.Hash][]byte{}}
			}
			paused := true
			p.Paused = &paused
			ps.Projects[e.ProjectId.Uint64()] = p

		case projectResumedTopicHash:
			e, err := instance.ParseProjectResumed(l)
			if err != nil {
				slog.Error("failed to parse project resumed event", "error", err)
				return false
			}

			p, ok := ps.Projects[e.ProjectId.Uint64()]
			if !ok {
				p = &Project{Attributes: map[common.Hash][]byte{}}
			}
			paused := false
			p.Paused = &paused
			ps.Projects[e.ProjectId.Uint64()] = p

		case projectConfigUpdatedTopicHash:
			e, err := instance.ParseProjectConfigUpdated(l)
			if err != nil {
				slog.Error("failed to parse project config updated event", "error", err)
				return false
			}

			p, ok := ps.Projects[e.ProjectId.Uint64()]
			if !ok {
				p = &Project{Attributes: map[common.Hash][]byte{}}
			}
			p.Uri = e.Uri
			p.Hash = e.Hash
			ps.Projects[e.ProjectId.Uint64()] = p
		}
		psMap[l.BlockNumber] = ps
	}

	psSlice := []*Projects{}
	for _, p := range psMap {
		psSlice = append(psSlice, p)
	}
	sort.Slice(psSlice, func(i, j int) bool {
		return psSlice[i].BlockNumber < psSlice[j].BlockNumber
	})

	for _, p := range psSlice {
		ch <- p
	}
	return true
}
