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

type ProjectManager interface {
	Get(projectID uint64) (*project.Project, error)
}

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

func RunDispatcher(persistence Persistence, newDatasource internaldispatcher.NewDatasource, getProject handler.GetProject, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519, chainEndpoint, projectContractAddress string, iotexChainID int) error {
	projectDispatchers := &sync.Map{}
	d := &dispatcher{projectDispatchers: projectDispatchers}

	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return err
	}

	taskStateHandler := handler.NewTaskStateHandler(persistence.Create, getProject, operatorPrivateKey, operatorPrivateKeyED25519)

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
		pd, err := internaldispatcher.NewProjectDispatcher(persistence.FetchProjectProcessedTaskID, persistence.UpsertProjectProcessedTask, p.DatasourceURI, newDatasource, pm, ps.Publish, taskStateHandler)
		if err != nil {
			return errors.Wrapf(err, "failed to new project dispatcher, project_id %v", projectID)
		}
		projectDispatchers.Store(projectID, pd)
	}

	return nil
}
