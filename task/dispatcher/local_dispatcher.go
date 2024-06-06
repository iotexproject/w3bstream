package dispatcher

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
)

func NewLocal(persistence Persistence, newDatasource NewDatasource,
	projectManager ProjectManager, defaultDatasourceURI, operatorPrivateKey, operatorPrivateKeyED25519, bootNodeMultiaddr string,
	sequencerPubKey []byte, iotexChainID int) (*Dispatcher, error) {

	projectDispatchers := &sync.Map{}
	taskStateHandler := newTaskStateHandler(persistence, nil, projectManager, operatorPrivateKey, operatorPrivateKeyED25519)
	d := &Dispatcher{
		local:              true,
		projectDispatchers: projectDispatchers,
	}
	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}

	projectIDs := projectManager.ProjectIDs()
	for _, id := range projectIDs {
		_, ok := projectDispatchers.Load(id)
		if ok {
			continue
		}
		p, err := projectManager.Project(id)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get project, project_id %v", id)
		}
		if err := ps.Add(id); err != nil {
			return nil, errors.Wrapf(err, "failed to add pubsubs, project_id %v", id)
		}
		cp := &contract.Project{
			ID:         id,
			Attributes: map[common.Hash][]byte{},
		}
		uri := p.DatasourceURI
		if uri == "" {
			uri = defaultDatasourceURI
		}
		pd, err := newProjectDispatcher(persistence, uri, newDatasource, cp, ps, taskStateHandler, sequencerPubKey)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to new project dispatcher, project_id %v", id)
		}
		pd.window.setSize(pd.requiredProverAmount.Load())
		projectDispatchers.Store(id, pd)
	}
	return d, nil
}
