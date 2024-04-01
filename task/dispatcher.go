package task

import (
	"log/slog"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type Datasource interface {
	Retrieve(nextTaskID uint64) (*types.Task, error)
}

type Persistence interface {
	Create(t *types.Task, tl *types.TaskStateLog) error
}

type ProjectConfigManager interface {
	Get(projectID uint64, version string) (*project.Config, error)
}

type Dispatcher struct {
	datasource                Datasource
	persistence               Persistence
	projectConfigManager      ProjectConfigManager
	pubSubs                   *p2p.PubSubs
	operatorPrivateKeyECDSA   string
	operatorPrivateKeyED25519 string
}

// will block caller
func (d *Dispatcher) Dispatch(nextTaskID uint64, pubkey []byte) {
	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		next, err := d.dispatchTask(nextTaskID, pubkey)
		if err != nil {
			slog.Error("failed to dispatch task", "error", err)
			continue
		}
		nextTaskID = next
	}
}

func (d *Dispatcher) dispatchTask(nextTaskID uint64, pubkey []byte) (uint64, error) {
	t, err := d.datasource.Retrieve(nextTaskID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to retrieve task from data source")
	}
	if t == nil {
		return nextTaskID, nil
	}
	if err := t.VerifySignature(pubkey); err != nil {
		return 0, errors.Wrap(err, "failed to verify task sign")
	}
	if err := d.pubSubs.Add(t.ProjectID); err != nil {
		return 0, errors.Wrapf(err, "failed to add project pubsub, project_id %v", t.ProjectID)
	}
	if err := d.pubSubs.Publish(t.ProjectID, &p2p.Data{Task: t}); err != nil {
		return 0, errors.Wrapf(err, "failed to publish data, project_id %v", t.ProjectID)
	}
	slog.Debug("dispatched a task", "project_id", t.ProjectID, "task_id", t.ID)
	return t.ID + 1, nil
}

func (d *Dispatcher) handleP2PData(data *p2p.Data, topic *pubsub.Topic) {
	if data.TaskStateLog == nil {
		return
	}
	l := data.TaskStateLog

	if err := d.persistence.Create(nil /*TODO*/, l); err != nil {
		slog.Error("failed to create task state log", "error", err, "task_id", l.TaskID)
		return
	}
	if l.State != types.TaskStateProved {
		return
	}

	if err := l.VerifySignature("", nil /*prover pubkey*/); err != nil {
		slog.Error("failed to verify proof sign", "error", err)
		return
	}

	// p, err := d.projectConfigManager.Get(l.Task.ProjectID, l.Task.ProjectVersion)
	// if err != nil {
	// 	//slog.Error("failed to get project", "error", err, "project_id", l.Task.ProjectID, "project_version", l.Task.ProjectVersion)
	// 	return
	// }
}

func NewDispatcher(persistence Persistence, projectConfigManager ProjectConfigManager, datasource Datasource, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string, iotexChainID int) (*Dispatcher, error) {
	d := &Dispatcher{
		datasource:                datasource,
		persistence:               persistence,
		projectConfigManager:      projectConfigManager,
		operatorPrivateKeyECDSA:   operatorPrivateKey,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
	}
	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}
	d.pubSubs = ps

	return d, nil
}
