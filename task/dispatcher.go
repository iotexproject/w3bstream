package task

import (
	"bytes"
	"log/slog"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/project/contracts"
	internaldispatcher "github.com/machinefi/sprout/task/internal/dispatcher"
	"github.com/machinefi/sprout/task/internal/handler"
	"github.com/machinefi/sprout/types"
)

type Persistence interface {
	Create(tl *types.TaskStateLog, t *types.Task) error
	FetchProjectProcessedTaskID(projectID uint64) (uint64, error)
	UpsertProjectProcessedTask(projectID, taskID uint64) error
}

type Datasource interface {
	Retrieve(projectID, nextTaskID uint64) (*types.Task, error)
}

type GetCachedProjectIDs func() []uint64

type dispatcher struct {
	projectDispatchers *sync.Map // projectID(uint64) -> *ProjectDispatcher
}

func (d *dispatcher) handleP2PData(data *p2p.Data, topic *pubsub.Topic) {
	if data.TaskStateLog == nil {
		return
	}
	s := data.TaskStateLog

	pd, ok := d.projectDispatchers.Load(s.ProjectID)
	if !ok {
		slog.Error("the project dispatcher not exist", "project_id", s.ProjectID)
		return
	}
	pd.(*internaldispatcher.ProjectDispatcher).Handle(s)
}

func RunDispatcher(persistence Persistence, newDatasource internaldispatcher.NewDatasource,
	getProjectIDs GetCachedProjectIDs, getProject handler.GetProject,
	bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519, chainEndpoint, projectContractAddress, projectFileDirectory string,
	iotexChainID int) error {
	projectDispatchers := &sync.Map{}
	d := &dispatcher{projectDispatchers: projectDispatchers}

	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return err
	}

	taskStateHandler := handler.NewTaskStateHandler(persistence.Create, getProject, operatorPrivateKey, operatorPrivateKeyED25519)

	if projectFileDirectory != "" {
		if err := dummyDispatch(persistence, newDatasource, getProjectIDs, getProject, projectDispatchers, ps, taskStateHandler); err != nil {
			return err
		}
		return nil
	}

	if err := dispatch(persistence, newDatasource, getProject, projectDispatchers, ps, taskStateHandler,
		chainEndpoint, projectContractAddress); err != nil {
		return err
	}
	return nil
}

func dummyDispatch(persistence Persistence, newDatasource internaldispatcher.NewDatasource,
	getProjectIDs GetCachedProjectIDs, getProject handler.GetProject,
	projectDispatchers *sync.Map, ps *p2p.PubSubs, handler *handler.TaskStateHandler) error {
	projectIDs := getProjectIDs()
	for _, id := range projectIDs {
		pm := &project.Meta{
			ProjectID: id,
		}
		_, ok := projectDispatchers.Load(id)
		if ok {
			continue
		}
		p, err := getProject(id)
		if err != nil {
			return errors.Wrapf(err, "failed to get project, project_id %v", id)
		}
		ps.Add(id)
		pd, err := internaldispatcher.NewProjectDispatcher(persistence.FetchProjectProcessedTaskID,
			persistence.UpsertProjectProcessedTask, p.DatasourceURI, newDatasource, pm, ps.Publish, handler)
		if err != nil {
			return errors.Wrapf(err, "failed to new project dispatcher, project_id %v", id)
		}
		projectDispatchers.Store(id, pd)
	}
	return nil
}

func dispatch(persistence Persistence, newDatasource internaldispatcher.NewDatasource, getProject handler.GetProject,
	projectDispatchers *sync.Map, ps *p2p.PubSubs, handler *handler.TaskStateHandler, chainEndpoint, projectContractAddress string) error {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return errors.Wrapf(err, "failed to dial chain endpoint %s", chainEndpoint)
	}
	instance, err := contracts.NewContracts(common.HexToAddress(projectContractAddress), client)
	if err != nil {
		return errors.Wrapf(err, "failed to new project contract instance")
	}

	emptyHash := [32]byte{}
	for projectID := uint64(1); ; projectID++ {
		mp, err := instance.Projects(nil, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get project meta from chain, project_id %v", projectID)
		}
		if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
			break
		}
		pm := &project.Meta{
			ProjectID: projectID,
			Uri:       mp.Uri,
			Hash:      mp.Hash,
		}
		// TODO support project update & watch project upsert
		_, ok := projectDispatchers.Load(projectID)
		if ok {
			continue
		}
		p, err := getProject(projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get project, project_id %v", projectID)
		}
		ps.Add(projectID)
		pd, err := internaldispatcher.NewProjectDispatcher(persistence.FetchProjectProcessedTaskID,
			persistence.UpsertProjectProcessedTask, p.DatasourceURI, newDatasource, pm, ps.Publish, handler)
		if err != nil {
			return errors.Wrapf(err, "failed to new project dispatcher, project_id %v", projectID)
		}
		projectDispatchers.Store(projectID, pd)
	}

	return nil
}
