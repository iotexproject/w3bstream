package project

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinefi/w3bstream-mainnet/project/contracts"
	"github.com/pkg/errors"
)

type Manager struct {
	mux             sync.Mutex
	pool            map[uint64]*Project
	chainEndpoint   string
	contractAddress string
}

func NewManager(chainEndpoint, contractAddress string) *Manager {
	return &Manager{
		pool:            make(map[uint64]*Project),
		chainEndpoint:   chainEndpoint,
		contractAddress: contractAddress,
	}
}

func (m *Manager) Get(projectID uint64) (*Project, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if p, ok := m.pool[projectID]; ok {
		return p, nil
	}

	client, err := ethclient.Dial(m.chainEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "dial chain endpoint failed")
	}
	address := common.HexToAddress(m.contractAddress)
	instance, err := contracts.NewContracts(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "new contracts instance failed")
	}
	p, err := instance.Projects(nil, new(big.Int).SetUint64(projectID))
	if err != nil {
		return nil, errors.Wrap(err, "get project from contracts failed")
	}

	resp, err := http.Get(p.Uri)
	if err != nil {
		return nil, errors.Wrap(err, "get project config file failed")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read project config file failed")
	}

	// TODO hash check

	np := Project{ID: projectID}
	if err := json.Unmarshal(data, &np.Config); err != nil {
		return nil, errors.Wrap(err, "unmarshal project config file failed")
	}
	m.pool[projectID] = &np
	return &np, nil
}
