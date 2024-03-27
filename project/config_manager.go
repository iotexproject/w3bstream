package project

import (
	"bytes"
	"log/slog"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/project/contracts"
)

type ConfigManager struct {
	ipfsEndpoint string
	instance     *contracts.Contracts
	configs      sync.Map // projectID(uint64) -> []*Config
	cache        *cache   // optional
}

func (m *ConfigManager) Get(projectID uint64, version string) (*Config, error) {
	var configs []*Config
	var err error
	configsValue, ok := m.configs.Load(projectID)
	if !ok {
		configs, err = m.load(projectID)
		if err != nil {
			return nil, err
		}
	} else {
		configs = configsValue.([]*Config)
	}
	for _, c := range configs {
		if c.Version == version {
			return c, nil
		}
	}
	return nil, errors.Errorf("the version of the project config not exist, project_id %v, version %v", projectID, version)
}

func (m *ConfigManager) load(projectID uint64) ([]*Config, error) {
	emptyHash := [32]byte{}
	mp, err := m.instance.Projects(nil, projectID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get project meta from chain, project_id %v", projectID)
	}
	if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
		return nil, errors.Errorf("the project not exist, project_id %v", projectID)
	}

	pm := &ProjectMeta{
		ProjectID: projectID,
		Uri:       mp.Uri,
		Hash:      mp.Hash,
		Paused:    mp.Paused,
	}

	var data []byte
	cached := true
	if m.cache != nil {
		data = m.cache.get(projectID, mp.Hash[:])
	}
	if len(data) == 0 {
		cached = false
		data, err = pm.GetConfigData(m.ipfsEndpoint)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get project config, project_id %v", projectID)
		}
	}
	if !cached && m.cache != nil {
		m.cache.set(projectID, data)
	}

	cs, err := convertConfigs(data)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to convert project config, project_id %v", projectID)
	}
	m.configs.Store(projectID, cs)
	return cs, nil
}

func (m *ConfigManager) watchProjectContract() error {
	events := make(chan *contracts.ContractsProjectUpserted, 10)
	sub, err := m.instance.WatchProjectUpserted(nil, events, nil)
	if err != nil {
		return errors.Wrapf(err, "failed to watch project upserted event")
	}
	go func() {
		for {
			select {
			case err := <-sub.Err():
				slog.Error("got an error when watching project upserted event", "error", err)
			case e := <-events:
				m.configs.Delete(e.ProjectId)
			}
		}
	}()
	return nil
}

// TODO support local project config
func NewConfigManager(chainEndpoint, contractAddress, projectCacheDir, ipfsEndpoint string) (*ConfigManager, error) {
	var c *cache
	var err error
	if projectCacheDir != "" {
		c, err = newCache(projectCacheDir)
		if err != nil {
			return nil, errors.Wrap(err, "failed to new cache")
		}
	}

	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial chain, endpoint %s", chainEndpoint)
	}
	instance, err := contracts.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to new contract instance, endpoint %s, contractAddress %s", chainEndpoint, contractAddress)
	}

	m := &ConfigManager{
		ipfsEndpoint: ipfsEndpoint,
		instance:     instance,
		cache:        c,
	}
	if err := m.watchProjectContract(); err != nil {
		return nil, err
	}
	return m, nil
}
