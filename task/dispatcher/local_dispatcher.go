package dispatcher

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/persistence/contract"
)

func NewLocal(persistence Persistence, newDatasource NewDatasource,
	projectManager ProjectManager, defaultDatasourceURI, operatorPrivateKey, operatorPrivateKeyED25519, bootNodeMultiaddr, contractWhitelist string,
	defaultDatasourcePubKey []byte, iotexChainID int) (*Dispatcher, error) {

	projectDispatchers := &sync.Map{}
	taskStateHandler := newTaskStateHandler(persistence, nil, projectManager, operatorPrivateKey, operatorPrivateKeyED25519, contractWhitelist)
	d := &Dispatcher{
		local:              true,
		projectDispatchers: projectDispatchers,
	}
	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}

	projectIDs, err := projectManager.ProjectIDs()
	if err != nil {
		return nil, err
	}
	for _, id := range projectIDs {
		_, ok := projectDispatchers.Load(id)
		if ok {
			continue
		}
		pf, err := projectManager.Project(id)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get project, project_id %v", id)
		}
		pfConf, err := pf.DefaultConfig()
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get project default config, project_id %v", id)
		}
		if err := ps.Add(id); err != nil {
			return nil, errors.Wrapf(err, "failed to add pubsubs, project_id %v", id)
		}
		cp := &contract.Project{
			ID:         id,
			Attributes: map[common.Hash][]byte{},
		}
		uri := pf.DatasourceURI
		if uri == "" {
			uri = defaultDatasourceURI
		}
		pubKey := defaultDatasourcePubKey
		if pf.DatasourcePubKey != "" {
			pubKey, err = hexutil.Decode(pf.DatasourcePubKey)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to decode datasource public key, project_id %v", id)
			}
		}
		pd, err := newProjectDispatcher(persistence, uri, newDatasource, cp, ps, taskStateHandler, pubKey, nil, pfConf.VMTypeID)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to new project dispatcher, project_id %v", id)
		}
		pd.window.setSize(pd.requiredProverAmount.Load())
		projectDispatchers.Store(id, pd)
	}
	return d, nil
}
