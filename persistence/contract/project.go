package contract

import (
	"bytes"
	"container/list"
	"math"
	"math/big"
	"sort"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/smartcontracts/go/blocknumber"
	"github.com/machinefi/sprout/smartcontracts/go/multicall"
	"github.com/machinefi/sprout/smartcontracts/go/project"
)

var (
	RequiredProverAmountHash            = crypto.Keccak256Hash([]byte("RequiredProverAmount"))
	VmTypeHash                          = crypto.Keccak256Hash([]byte("VmType"))
	ClientManagementContractAddressHash = crypto.Keccak256Hash([]byte("ClientManagementContractAddress"))

	attributeSetTopicHash         = crypto.Keccak256Hash([]byte("AttributeSet(uint256,bytes32,bytes)"))
	projectPausedTopicHash        = crypto.Keccak256Hash([]byte("ProjectPaused(uint256)"))
	projectResumedTopicHash       = crypto.Keccak256Hash([]byte("ProjectResumed(uint256)"))
	projectConfigUpdatedTopicHash = crypto.Keccak256Hash([]byte("ProjectConfigUpdated(uint256,string,bytes32)"))

	emptyHash = common.Hash{}
)

type Project struct {
	ID          uint64
	BlockNumber uint64
	Paused      *bool
	Uri         string
	Hash        common.Hash
	Attributes  map[common.Hash][]byte
}

type blockProject struct {
	BlockNumber uint64
	Projects    map[uint64]*Project
}

func (p *Project) merge(diff *Project) {
	if diff.ID != 0 {
		p.ID = diff.ID
	}
	if diff.BlockNumber != 0 {
		p.BlockNumber = diff.BlockNumber
	}
	if diff.Paused != nil {
		p.Paused = diff.Paused
	}
	if diff.Uri != "" {
		p.Uri = diff.Uri
	}
	if !bytes.Equal(diff.Hash[:], emptyHash[:]) {
		p.Hash = diff.Hash
	}
	for h, d := range diff.Attributes {
		p.Attributes[h] = d
	}
}

func (p *Project) isEmpty() bool {
	return p.ID == 0
}

func (ps *blockProject) merge(diff *blockProject) {
	ps.BlockNumber = diff.BlockNumber
	for id, p := range ps.Projects {
		diffP, ok := diff.Projects[id]
		if ok {
			p.merge(diffP)
		}
	}
	for id, p := range diff.Projects {
		if _, ok := ps.Projects[id]; !ok {
			np := &Project{Attributes: map[common.Hash][]byte{}}
			np.merge(p)
			ps.Projects[id] = np
		}
	}
}

type blockProjects struct {
	mu       sync.Mutex
	capacity uint64
	blocks   *list.List
}

func (c *blockProjects) project(projectID, blockNumber uint64) *Project {
	c.mu.Lock()
	defer c.mu.Unlock()

	if blockNumber == 0 {
		blockNumber = c.blocks.Back().Value.(*blockProject).BlockNumber
	}
	np := &Project{Attributes: map[common.Hash][]byte{}}

	for e := c.blocks.Front(); e != nil; e = e.Next() {
		ep := e.Value.(*blockProject)
		if blockNumber < ep.BlockNumber {
			break
		}
		p, ok := ep.Projects[projectID]
		if ok {
			np.merge(p)
		}
	}
	if np.isEmpty() {
		return nil
	}
	return np
}

func (c *blockProjects) projects() *blockProject {
	c.mu.Lock()
	defer c.mu.Unlock()

	np := &blockProject{Projects: map[uint64]*Project{}}

	for e := c.blocks.Front(); e != nil; e = e.Next() {
		ep := e.Value.(*blockProject)
		np.merge(ep)
	}
	return np
}

func (c *blockProjects) add(diff *blockProject) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.blocks.PushBack(diff)

	if uint64(c.blocks.Len()) > c.capacity {
		h := c.blocks.Front()
		np := &blockProject{Projects: map[uint64]*Project{}}
		np.merge(h.Value.(*blockProject))
		np.merge(h.Next().Value.(*blockProject))
		c.blocks.Remove(h.Next())
		c.blocks.Remove(h)
		c.blocks.PushFront(np)
	}
}

func listProject(client *ethclient.Client, projectContractAddress, blockNumberContractAddress, multiCallContractAddress common.Address) ([]*Project, uint64, uint64, error) {
	multiCallInstance, err := multicall.NewMulticall(multiCallContractAddress, client)
	if err != nil {
		return nil, 0, 0, errors.Wrap(err, "failed to new multi call contract instance")
	}
	blockNumberABI, err := abi.JSON(strings.NewReader(blocknumber.BlocknumberMetaData.ABI))
	if err != nil {
		return nil, 0, 0, errors.Wrap(err, "failed to decode block number contract abi")
	}
	projectABI, err := abi.JSON(strings.NewReader(project.ProjectMetaData.ABI))
	if err != nil {
		return nil, 0, 0, errors.Wrap(err, "failed to decode project contract abi")
	}
	blockNumberCallData, err := blockNumberABI.Pack("blockNumber")
	if err != nil {
		return nil, 0, 0, errors.Wrap(err, "failed to pack block number call data")
	}
	projectCountCallData, err := projectABI.Pack("count")
	if err != nil {
		return nil, 0, 0, errors.Wrap(err, "failed to pack project count call data")
	}
	ps := []*Project{}
	minBlockNumber := uint64(math.MaxUint64)
	maxBlockNumber := uint64(0)
	listedCount := uint64(0)
	for projectID := uint64(1); ; projectID++ {
		isValidProjectCallData, err := projectABI.Pack("isValidProject", new(big.Int).SetUint64(projectID))
		if err != nil {
			return nil, 0, 0, errors.Wrap(err, "failed to pack is valid project call data")
		}
		configCallData, err := projectABI.Pack("config", new(big.Int).SetUint64(projectID))
		if err != nil {
			return nil, 0, 0, errors.Wrap(err, "failed to pack project config call data")
		}
		isPausedCallData, err := projectABI.Pack("isPaused", new(big.Int).SetUint64(projectID))
		if err != nil {
			return nil, 0, 0, errors.Wrap(err, "failed to pack project is paused call data")
		}
		requiredProverAmountCallData, err := projectABI.Pack("attributes", new(big.Int).SetUint64(projectID), RequiredProverAmountHash)
		if err != nil {
			return nil, 0, 0, errors.Wrap(err, "failed to pack project attributes call data")
		}
		vmTypeCallData, err := projectABI.Pack("attributes", new(big.Int).SetUint64(projectID), VmTypeHash)
		if err != nil {
			return nil, 0, 0, errors.Wrap(err, "failed to pack project attributes call data")
		}
		managementAddressCallData, err := projectABI.Pack("attributes", new(big.Int).SetUint64(projectID), ClientManagementContractAddressHash)
		if err != nil {
			return nil, 0, 0, errors.Wrap(err, "failed to pack project attributes call data")
		}

		result, err := multiCallInstance.MultiCall(
			nil,
			[]common.Address{
				blockNumberContractAddress,
				projectContractAddress,
				projectContractAddress,
				projectContractAddress,
				projectContractAddress,
				projectContractAddress,
				projectContractAddress,
				projectContractAddress,
			},
			[][]byte{
				blockNumberCallData,
				projectCountCallData,
				isValidProjectCallData,
				configCallData,
				isPausedCallData,
				requiredProverAmountCallData,
				vmTypeCallData,
				managementAddressCallData,
			},
		)
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to multi call, project_id %v", projectID)
		}

		out, err := blockNumberABI.Unpack("blockNumber", result[0])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack block number result, project_id %v", projectID)
		}
		preBlockNumber := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		blockNumber := preBlockNumber.Uint64() - 1

		minBlockNumber = min(minBlockNumber, blockNumber)
		maxBlockNumber = max(maxBlockNumber, blockNumber)

		out, err = projectABI.Unpack("count", result[1])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack project count result, project_id %v", projectID)
		}
		projectCount := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		out, err = projectABI.Unpack("isValidProject", result[2])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack is valid project result, project_id %v", projectID)
		}
		isValidProject := *abi.ConvertType(out[0], new(bool)).(*bool)

		if !isValidProject {
			continue
		}

		out, err = projectABI.Unpack("config", result[3])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack project config result, project_id %v", projectID)
		}
		config := *abi.ConvertType(out[0], new(project.W3bstreamProjectProjectConfig)).(*project.W3bstreamProjectProjectConfig)

		out, err = projectABI.Unpack("isPaused", result[4])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack project is paused result, project_id %v", projectID)
		}
		isPaused := *abi.ConvertType(out[0], new(bool)).(*bool)

		out, err = projectABI.Unpack("attributes", result[5])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack project attributes result, project_id %v", projectID)
		}
		proverAmt := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

		out, err = projectABI.Unpack("attributes", result[6])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack project attributes result, project_id %v", projectID)
		}
		vmType := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

		out, err = projectABI.Unpack("attributes", result[7])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack project attributes result, project_id %v", projectID)
		}
		managementAddress := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

		attributes := map[common.Hash][]byte{}
		if len(proverAmt) != 0 {
			attributes[RequiredProverAmountHash] = proverAmt
		}
		if len(vmType) != 0 {
			attributes[VmTypeHash] = vmType
		}
		if len(managementAddress) != 0 {
			attributes[ClientManagementContractAddressHash] = managementAddress
		}

		ps = append(ps, &Project{
			ID:          projectID,
			BlockNumber: blockNumber,
			Paused:      &isPaused,
			Uri:         config.Uri,
			Hash:        config.Hash,
			Attributes:  attributes,
		})

		listedCount++
		if listedCount >= projectCount.Uint64() {
			break
		}
	}
	return ps, minBlockNumber, maxBlockNumber, nil
}

func processProjectLogs(add func(*blockProject), logs []types.Log, instance *project.Project) error {
	sort.Slice(logs, func(i, j int) bool {
		if logs[i].BlockNumber != logs[j].BlockNumber {
			return logs[i].BlockNumber < logs[j].BlockNumber
		}
		return logs[i].TxIndex < logs[j].TxIndex
	})
	psMap := map[uint64]*blockProject{}

	for _, l := range logs {
		ps, ok := psMap[l.BlockNumber]
		if !ok {
			ps = &blockProject{
				BlockNumber: l.BlockNumber,
				Projects:    map[uint64]*Project{},
			}
		}
		switch l.Topics[0] {
		case attributeSetTopicHash:
			e, err := instance.ParseAttributeSet(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse project attribute set event")
			}

			p, ok := ps.Projects[e.ProjectId.Uint64()]
			if !ok {
				p = &Project{
					ID:         e.ProjectId.Uint64(),
					Attributes: map[common.Hash][]byte{},
				}
			}
			p.Attributes[e.Key] = e.Value
			ps.Projects[e.ProjectId.Uint64()] = p

		case projectPausedTopicHash:
			e, err := instance.ParseProjectPaused(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse project paused event")
			}

			p, ok := ps.Projects[e.ProjectId.Uint64()]
			if !ok {
				p = &Project{
					ID:         e.ProjectId.Uint64(),
					Attributes: map[common.Hash][]byte{},
				}
			}
			paused := true
			p.Paused = &paused
			ps.Projects[e.ProjectId.Uint64()] = p

		case projectResumedTopicHash:
			e, err := instance.ParseProjectResumed(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse project resumed event")
			}

			p, ok := ps.Projects[e.ProjectId.Uint64()]
			if !ok {
				p = &Project{
					ID:         e.ProjectId.Uint64(),
					Attributes: map[common.Hash][]byte{},
				}
			}
			paused := false
			p.Paused = &paused
			ps.Projects[e.ProjectId.Uint64()] = p

		case projectConfigUpdatedTopicHash:
			e, err := instance.ParseProjectConfigUpdated(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse project config updated event")
			}

			p, ok := ps.Projects[e.ProjectId.Uint64()]
			if !ok {
				p = &Project{
					ID:         e.ProjectId.Uint64(),
					Attributes: map[common.Hash][]byte{},
				}
			}
			p.Uri = e.Uri
			p.Hash = e.Hash
			ps.Projects[e.ProjectId.Uint64()] = p
		}
		psMap[l.BlockNumber] = ps
	}

	psSlice := []*blockProject{}
	for _, p := range psMap {
		psSlice = append(psSlice, p)
	}
	sort.Slice(psSlice, func(i, j int) bool {
		return psSlice[i].BlockNumber < psSlice[j].BlockNumber
	})

	for _, p := range psSlice {
		add(p)
	}
	return nil
}
